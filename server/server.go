package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	log    *logrus.Logger
	engine *gin.Engine
}

func New() *Server {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	return &Server{
		log:    log,
		engine: gin.Default(),
	}
}

func (this *Server) Start() {
	this.engine.Use(gzip.Gzip(gzip.BestCompression))
	this.engine.Use(location.Default())
	this.engine.Use(cors.Default())

	this.engine.GET("/", func(ctx *gin.Context) {
		ctx.String(200, ctx.ClientIP())
	})

	if err := this.engine.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err)
	} else {
		this.log.Info("Successfully Started Server")
	}
}
