package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"dao-exchange/apps/fsm/consumer"
	"dao-exchange/apps/fsm/handler"
	"dao-exchange/apps/scanner/api"
	"dao-exchange/config"
	"dao-exchange/internal/cache"
	"dao-exchange/pkg/orm"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var (
	appID, configPath, namespace, clusterName, tagName string
)

func init() {
	flag.StringVar(&appID, "appid", "chain-data-scanner", "app id use in read remote config")
	flag.StringVar(&configPath, "config", "", "local config file path")
	flag.StringVar(&namespace, "namespace", "config.yaml", "apollo config namespace")
	flag.StringVar(&clusterName, "cluster", "default", "apollo cluster name")
	flag.StringVar(&tagName, "tag", "json", "config tag, default json")
	flag.Parse()
}

func main() {
	conf, err := config.LoadConf(configPath, appID, clusterName, namespace, tagName)
	if err != nil {
		logrus.Error("read config failed, err : ", err)
		return
	}

	logrus.SetLevel(logrus.DebugLevel)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	engine := api.InitEngine(conf)
	srv := http.Server{
		Addr:    ":" + strconv.Itoa(conf.Port),
		Handler: engine,
	}

	db, err := orm.NewGorm(conf.Db)
	if err != nil {
		log.Fatalf("connect db err %s", err)
	}

	// 缓存数据
	go cacheContracts(ctx, db)

	// terminal signal
	g.Go(func() error {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		for {
			select {
			case sig := <-sigCh:
				logrus.Infof("got terminal signal : %s", sig.String())
				cancel()

				ctx1, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
				defer cancelFunc()

				if err := srv.Shutdown(ctx1); err != nil {
					return err
				}
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	g.Go(func() error {
		logrus.Infof("listen on: %d", conf.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	reader := consumer.NewReader(conf.ReaderKafka, conf.Credential, db, handler.InitAction(db))
	go reader.Start(ctx)

	if err := g.Wait(); err != nil {
		logrus.Error(err)
	}
}

func cacheContracts(ctx context.Context, db *gorm.DB) {
	cache.NewLocalEventCache().UpdateCache(db)
	cache.NewPayTokenCache().UpdateCache(db)

	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			cache.NewLocalEventCache().UpdateCache(db)
			cache.NewPayTokenCache().UpdateCache(db)
		case <-ctx.Done():
			ticker.Stop()
		}
	}
}
