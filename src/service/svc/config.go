package svc

import (
	"github.com/yaoxc/EasySwapBase/evm/erc"
	//"github.com/yaoxc/EasySwapBase/image"
	"github.com/yaoxc/EasySwapBase/stores/xkv"
	"gorm.io/gorm"

	"github.com/yaoxc/EasySwapBackend/src/dao"
)

type CtxConfig struct {
	db *gorm.DB
	//imageMgr image.ImageManager
	dao     *dao.Dao
	KvStore *xkv.Store
	Evm     erc.Erc
}

type CtxOption func(conf *CtxConfig)

// NewServerCtx 创建一个新的ServerCtx实例
// 参数:
//
//	options: 可变参数，用于配置CtxConfig的结构体选项
//
// 返回值:
//
//	*ServerCtx: 返回一个配置好的ServerCtx指针
func NewServerCtx(options ...CtxOption) *ServerCtx {
	// 初始化一个空的CtxConfig结构体
	c := &CtxConfig{}
	// 遍历所有传入的选项函数，并应用到CtxConfig结构体上
	// for index, value := range collection { } ,  第一个返回值（索引）被 _ 接收了，表示忽略它
	for _, opt := range options {
		opt(c)
	}
	// 返回一个ServerCtx实例，使用配置好的CtxConfig进行初始化
	return &ServerCtx{
		DB: c.db,
		//ImageMgr: c.imageMgr, // 此行被注释掉，表示暂时不使用ImageMgr
		KvStore: c.KvStore, // 使用配置的KvStore
		Dao:     c.dao,     // 使用配置的Dao
	}
}

func WithKv(kv *xkv.Store) CtxOption {
	return func(conf *CtxConfig) {
		conf.KvStore = kv
	}
}

func WithDB(db *gorm.DB) CtxOption {
	return func(conf *CtxConfig) {
		conf.db = db
	}
}

func WithDao(dao *dao.Dao) CtxOption {
	return func(conf *CtxConfig) {
		conf.dao = dao
	}
}
