package config

import (
	"fmt"
	"os"

	"dao-exchange/internal/mqs"
	"dao-exchange/pkg/orm"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/shima-park/agollo"
	remote "github.com/shima-park/agollo/viper-remote"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config config
type Config struct {
	Port          int           `json:"port"`
	Redis         RedisConf     `json:"redis"`
	Db            orm.DBCfg     `json:"db"`
	ReaderKafka   mqs.KafkaCfg  `json:"reader_kafka"`
	ProducerKafka mqs.KafkaCfg  `json:"producer_kafka"`
	Scanner       []*Scanner    `json:"scanner"`
	Credential    AWSCredential `json:"credential"`
}

// AWSCredential secret
type AWSCredential struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Region          string `json:"region"`
}

// Scanner scan block
type Scanner struct {
	NodeURL      string `json:"node_url"`
	Chain        string `json:"chain"`
	ChainID      int    `json:"chain_id"`
	DelayBlock   uint64 `json:"delay_block"`
	StartBlock   uint64 `json:"start_block"`
	ScanInterval int    `json:"scan_interval"`
	ScanAmount   int    `json:"scan_amount"`
}

// RedisConf config of redis
type RedisConf struct {
	Addrs  []string `json:"addrs"`
	Passwd string   `json:"passwd"`
}

// KafkaConf config of kafka
type KafkaConf struct {
	Brokers        []string `json:"brokers"`  // Kafka bootstrap brokers to connect to, as a comma separated list
	Group          string   `json:"group"`    // Kafka consumer group definition
	Topics         []string `json:"topics"`   //"Kafka topics to be consumed, as a comma separated list"
	Version        string   `json:"version"`  // "2.1.1", "Kafka cluster version"
	Assignor       string   `json:"assignor"` // "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)"
	Oldest         bool     `json:"oldest"`   // true, "Kafka consumer consume initial offset from oldest"
	Verbose        bool     `json:"verbose"`  // false, "Sarama logging"
	FlushFrequency int      `json:"flush_frequency"`
	FlushMessages  int      `json:"flush_messages"`
}

// LoadConf load config from apollo
func LoadConf(fpath, appID, clusterName, namespace, tagName string) (*Config, error) {
	conf := &Config{}
	vip := viper.New()
	vip.SetConfigType("yaml")

	if fpath != "" {
		logrus.Info("read configuration from local yaml file : ", fpath)
		err := localConfig(fpath, vip)
		if err != nil {
			return nil, err
		}
	} else {
		confServerURL := os.Getenv("CONFIG_SERVER_URL")
		if confServerURL == "" {
			return nil, errors.New("must set env CONFIG_SERVER_URL or add local config path")
		}

		logrus.Info("read configuration from remote conf server :", confServerURL)

		err := remoteConfig(namespace, appID, clusterName, confServerURL, vip)
		if err != nil {
			return nil, err
		}
	}

	withTagName := func(dc *mapstructure.DecoderConfig) {
		dc.TagName = tagName
		dc.Squash = true
	}

	err := vip.Unmarshal(conf, withTagName)

	if err != nil {
		return nil, err
	}

	return conf, nil
}

func localConfig(filename string, v *viper.Viper) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("err : %v", err)
	}

	// TODO 这里是相对路径了
	v.AddConfigPath(path) //设置读取的文件路径

	v.SetConfigName(filename) //设置读取的文件名

	err = v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read conf file err : %v", err)
	}

	return err
}

func remoteConfig(namespace, appID, clusterName, confServerURL string, v *viper.Viper) error {
	remote.SetAppID(appID)
	remote.SetAgolloOptions(
		agollo.Cluster(clusterName),
		agollo.DefaultNamespace("config.yaml"),
		agollo.AutoFetchOnCacheMiss(),       // namespace不存在或者未初始化时，从apollo接口拉取配置
		agollo.FailTolerantOnBackupExists(), // 容灾在apollo连接不上时
	)

	err := v.AddRemoteProvider("apollo", confServerURL, namespace)
	if err != nil {
		return fmt.Errorf("add remote provider error : %v", err)
	}

	err = v.ReadRemoteConfig()
	if err != nil {
		return fmt.Errorf("read remote config error : %v", err)
	}

	err = v.WatchRemoteConfigOnChannel() // 启动一个goroutine来同步配置更改

	return err
}
