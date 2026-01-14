package types

type LoginReq struct {
	ChainID   int    `json:"chain_id"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
	// 可选字段，视具体需求而定【在以太坊签名验证中，可以通过签名和消息恢复出公钥，然后推导出地址】
	PublicKey string `json:"public_key"`
	Address   string `json:"address"`
}

type UserLoginInfo struct {
	Token     string `json:"token"`
	IsAllowed bool   `json:"is_allowed"`
}

type UserLoginResp struct {
	Result interface{} `json:"result"`
}

type UserLoginMsgResp struct {
	Address string `json:"address"`
	Message string `json:"message"`
}

type UserSignStatusResp struct {
	IsSigned bool `json:"is_signed"`
}
