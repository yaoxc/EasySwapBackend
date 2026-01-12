package app

import (
	"context"
	"fmt"

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

	// 在初始化配置的地方添加打印
	fmt.Printf("[DEBUG] p.config.Api.Port 的值是: [%s]\n", p.config.Api.Port)
	fmt.Printf("[DEBUG] p.config 的完整内容: %+v\n", p.config)

	// 使用 xzap 包记录一条日志，表明 "EasySwap-End run" 正在运行，并记录监听的端口号
	xzap.WithContext(context.Background()).Info("EasySwap-End run", zap.String("port", p.config.Api.Port))

	// ！！！添加这一行，确认传入 Run 的参数！！！
	fmt.Printf("[DEBUG] 准备启动 HTTP 服务器，端口参数: [%s]\n", p.config.Api.Port)

	// 启动 HTTP 服务器，监听配置文件中指定的端口
	if err := p.router.Run(p.config.Api.Port); err != nil {
		panic(err)
	}
}
