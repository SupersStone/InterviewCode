package api

import (
	"dao-exchange/config"
	"net/http"

	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func InitEngine(conf *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	// pprof
	pprof.Register(engine)

	// prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(engine)

	initRouter(engine)
	return engine
}

func initRouter(engine *gin.Engine) {
	// health check
	engine.GET("/info", health)

	// expvars
	engine.GET("/debug/vars", expvar.Handler())
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    "",
		"success": true,
	})
}
