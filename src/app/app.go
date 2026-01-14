package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yaoxc/EasySwapBase/logger/xzap"
	"go.uber.org/zap"

	"github.com/yaoxc/EasySwapBackend/src/config"
	"github.com/yaoxc/EasySwapBackend/src/service/svc"
)

type Platform struct {
	config    *config.Config
	router    *gin.Engine
	serverCtx *svc.ServerCtx
}

func NewPlatform(config *config.Config, router *gin.Engine, serverCtx *svc.ServerCtx) (*Platform, error) {
	return &Platform{
		config:    config,
		router:    router,
		serverCtx: serverCtx,
	}, nil
}

func (p *Platform) Start() {

	// Log组件使用
	// 使用 xzap 包记录一条日志，表明 "EasySwap-End run" 正在运行，并记录监听的端口号
	xzap.WithContext(context.Background()).Info("EasySwap-End run", zap.String("port", p.config.Api.Port))

	// 启动 HTTP 服务器，监听配置文件中指定的端口
	if err := p.router.Run(p.config.Api.Port); err != nil {
		panic(err)
	}
}
