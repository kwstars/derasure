package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/derasure/internal/app/middlewares"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var ProviderSet = wire.NewSet(NewRouter, New)

type InitControllers func(r *gin.Engine)

type Server struct {
	router     *gin.Engine
	httpServer http.Server
}

func NewRouter(initRouter InitControllers) *gin.Engine {
	r := gin.New()

	// 初始化静态文件目录
	r.Static("/public", "./public")

	// 初始化模板文件
	r.LoadHTMLGlob("templates/*")

	// 解决前端web与后端跨域调试问题
	r.Use(middlewares.Cors())

	// panic之后自动恢复
	r.Use(gin.Recovery())

	// 初始化路由
	initRouter(r)

	return r
}

func New(r *gin.Engine) (*Server, error) {
	var s = &Server{
		router: r,
	}

	return s, nil
}

func (s *Server) Start() error {
	s.httpServer = http.Server{Addr: ":8080", Handler: s.router}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Shutdown http server err")
	}

	return nil
}
