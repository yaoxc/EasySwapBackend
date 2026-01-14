package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/yaoxc/EasySwapBackend/src/api/middleware"

	"github.com/yaoxc/EasySwapBackend/src/service/svc"
)

func NewRouter(svcCtx *svc.ServerCtx) *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() // 新建一个gin引擎实例

	// 捕获http请求处理过程中发生的panic错误，防止程序崩溃
	r.Use(middleware.RecoverMiddleware()) // 使用恢复中间件

	// RLog 请求响应日志打印处理器 【对每个请求 详细信息 打印】
	r.Use(middleware.RLog())

	// r.Use(middleware.AuthMiddleWare(svcCtx.KvStore)) // 使用认证中间件

	// r.Use(middleware.CacheApi(svcCtx.KvStore, 2)) // 使用API缓存中间件，缓存时间为2秒

	// 跨域资源共享（CORS）中间件配置
	r.Use(cors.New(cors.Config{ // 使用cors中间件
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-GW-Error-Code", "X-GW-Error-Message"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
	loadV1(r, svcCtx) // 加载v1路由

	return r
}
