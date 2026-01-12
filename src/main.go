package main

import (
	"flag"
	_ "net/http/pprof" // 引入pprof用于性能分析

	"github.com/yaoxc/EasySwapBackend/src/api/router"
	"github.com/yaoxc/EasySwapBackend/src/app"
	"github.com/yaoxc/EasySwapBackend/src/config"
	"github.com/yaoxc/EasySwapBackend/src/service/svc"
)

const (
	port              = ":9000"
	repoRoot          = ""
	defaultConfigPath = "./config/config.toml"
)

// main 函数是程序的入口点
func main() {
	// 使用 flag 包解析命令行参数，获取配置文件路径
	// 默认值为 defaultConfigPath，参数名称为 "conf"
	conf := flag.String("conf", defaultConfigPath, "conf file path")
	flag.Parse()
	// 解析配置文件，将配置文件内容反序列化为配置结构体
	// Unmarshal = “反序列化”
	c, err := config.UnmarshalConfig(*conf)
	if err != nil {
		// 如果解析失败，程序终止并打印错误信息
		panic(err)
	}

	// 遍历配置中支持的链
	for _, chain := range c.ChainSupported {
		// 检查链ID和名称是否有效
		if chain.ChainID == 0 || chain.Name == "" {
			// 如果链配置无效，程序终止并打印错误信息
			panic("invalid chain_suffix config")
		}
	}

	// 创建服务上下文，初始化所需的服务组件
	serverCtx, err := svc.NewServiceContext(c)
	if err != nil {
		// 如果服务上下文创建失败，程序终止并打印错误信息
		panic(err)
	}
	// Initialize router
	r := router.NewRouter(serverCtx)
	app, err := app.NewPlatform(c, r, serverCtx)
	if err != nil {
		panic(err)
	}
	app.Start()
}
