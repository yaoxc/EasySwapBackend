package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoxc/EasySwapBase/errcode"
	"github.com/yaoxc/EasySwapBase/kit/validator"
	"github.com/yaoxc/EasySwapBase/xhttp"

	"github.com/yaoxc/EasySwapBackend/src/service/svc"
	"github.com/yaoxc/EasySwapBackend/src/service/v1"
	"github.com/yaoxc/EasySwapBackend/src/types/v1"
)

func UserLoginHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ===============================================
		// c *gin.Context 参数的值是在 Gin 框架的路由匹配过程中自动设置
		// ===============================================

		// 数据流向: HTTP Request Body → c.BindJSON → types.LoginReq结构体
		// 初始化请求结构体
		req := types.LoginReq{}
		// 将JSON数据绑定到结构体指针
		if err := c.BindJSON(&req); err != nil {
			xhttp.Error(c, err)
			return
		}

		// 结构体验证
		if err := validator.Verify(&req); err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		// controller 调用 service 层处理业务逻辑
		res, err := service.UserLogin(c.Request.Context(), svcCtx, req)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, types.UserLoginResp{
			Result: res,
		})
	}
}

func GetLoginMessageHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Params.ByName("address")
		if address == "" {
			xhttp.Error(c, errcode.NewCustomErr("user addr is null"))
			return
		}

		// controller 调用 service 层, 生成登录消息
		res, err := service.GetUserLoginMsg(c.Request.Context(), svcCtx, address)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, res)
	}
}

func GetSigStatusHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Params.ByName：获取路径参数
		userAddr := c.Params.ByName("address")
		if userAddr == "" {
			xhttp.Error(c, errcode.NewCustomErr("user addr is null"))
			return
		}

		res, err := service.GetSigStatusMsg(c.Request.Context(), svcCtx, userAddr)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, res)
	}
}
