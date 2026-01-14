
- [一、代码解读](#一代码解读)
  - [一 、参数结构体验证](#一-参数结构体验证)
    - [1. 验证流程](#1-验证流程)
    - [2. 验证规则定义](#2-验证规则定义)
    - [3. 验证机制](#3-验证机制)
    - [4. 错误处理](#4-错误处理)
    - [5. 验证示例](#5-验证示例)
  - [二、生成token](#二生成token)
  - [三、返回](#三返回)
  - [四、 为什么 LoginReq 的ChainID接收不到值?](#四-为什么-loginreq-的chainid接收不到值)
    - [1. 问题原因](#1-问题原因)
    - [2. 解决方案](#2-解决方案)
    - [3. 验证规则](#3-验证规则)
    - [4. 注意事项](#4-注意事项)
    - [5. 支持的链ID](#5-支持的链id)
- [二、登录流程](#二登录流程)
    - [1. 登录流程的必要性](#1-登录流程的必要性)
    - [2. 安全机制](#2-安全机制)
    - [3. 登录流程](#3-登录流程)
    - [4. 为什么必须每次都获取](#4-为什么必须每次都获取)
- [三、如何理解“使用过的消息立即失效”？](#三如何理解使用过的消息立即失效)
    - [1. UUID的生成和验证机制](#1-uuid的生成和验证机制)
    - [2. 防止重复使用的机制](#2-防止重复使用的机制)
    - [3. 实际场景分析](#3-实际场景分析)
    - [4. 安全保障](#4-安全保障)
- [五、什么是“重放攻击”？](#五什么是重放攻击)
    - [1. 什么是重放攻击](#1-什么是重放攻击)
    - [2. 在NFT市场中的具体场景](#2-在nft市场中的具体场景)
    - [3. 防护机制](#3-防护机制)
    - [4. 实际防护效果](#4-实际防护效果)
    - [5. 其他防护措施](#5-其他防护措施)
- [六、前端登录交互流程](#六前端登录交互流程)
  - [1. 前端实现路径：](#1-前端实现路径)
  - [2. 详细流程：](#2-详细流程)
  - [3. 代码示例](#3-代码示例)
    - [1. 前端获取登录消息](#1-前端获取登录消息)
    - [2. 后端生成登录消息](#2-后端生成登录消息)
    - [3. 用户签名](#3-用户签名)
    - [4. 前端提交登录请求](#4-前端提交登录请求)
    - [5. 后端验证流程](#5-后端验证流程)
    - [6. 返回登录结果](#6-返回登录结果)
    - [7. 前端处理登录结果](#7-前端处理登录结果)
    - [8. 安全机制](#8-安全机制)
- [七、用户token生成过程](#七用户token生成过程)
  - [1. Token生成流程](#1-token生成流程)
    - [1.1 生成Token Key](#11-生成token-key)
    - [1.2 AES加密生成Token](#12-aes加密生成token)
    - [1.3 Token加密细节](#13-token加密细节)
  - [2. Token缓存](#2-token缓存)
    - [2.1 写入Redis](#21-写入redis)
    - [2.2 缓存配置](#22-缓存配置)
  - [3. Token返回](#3-token返回)
    - [3.1 返回格式](#31-返回格式)
    - [3.2 返回字段说明](#32-返回字段说明)
  - [安全机制](#安全机制)
- [八、Redis中缓存的用户token为什么是uuid.newstring() ？](#八redis中缓存的用户token为什么是uuidnewstring-)
  - [1. 安全性考虑](#1-安全性考虑)
    - [1.1 防止Token伪造](#11-防止token伪造)
    - [1.2 避免Token冲突](#12-避免token冲突)
  - [2. 实际用途](#2-实际用途)
    - [2.1 Token验证](#21-token验证)
    - [2.2 会话管理](#22-会话管理)
  - [3. 设计优势](#3-设计优势)
- [九、Token与UUID的匹配机制](#九token与uuid的匹配机制)
  - [1. 两者的不同用途](#1-两者的不同用途)
    - [1.1 UUID的作用](#11-uuid的作用)
    - [1.2 AES加密Token的作用](#12-aes加密token的作用)
  - [2. 匹配验证流程](#2-匹配验证流程)
    - [2.1 前端请求验证](#21-前端请求验证)
    - [2.2 验证逻辑](#22-验证逻辑)
  - [3. 安全机制](#3-安全机制)
    - [3.1 双重验证](#31-双重验证)
    - [3.2 时效性控制](#32-时效性控制)
- [十、支持多设备登录](#十支持多设备登录)
  - [1. 当前设计的问题](#1-当前设计的问题)
    - [1.1 覆盖机制](#11-覆盖机制)
    - [1.2 验证失败场景](#12-验证失败场景)
  - [2. 改进方案](#2-改进方案)
    - [2.1 多设备Token管理](#21-多设备token管理)
    - [2.2 设备标识](#22-设备标识)
    - [2.3 验证逻辑优化](#23-验证逻辑优化)
  - [3. 优势分析](#3-优势分析)
    - [3.1 多设备支持](#31-多设备支持)
    - [3.2 安全性](#32-安全性)
    - [3.3 灵活性](#33-灵活性)
- [十一、多设备登录用户识别机制](#十一多设备登录用户识别机制)
  - [1. 核心识别字段](#1-核心识别字段)
    - [1.1 用户地址（Address）](#11-用户地址address)
    - [1.2 Token Key生成](#12-token-key生成)
  - [2. 设备区分机制](#2-设备区分机制)
    - [2.1 设备ID（DeviceID）](#21-设备iddeviceid)
    - [2.2 Token存储结构](#22-token存储结构)
  - [3. 验证流程](#3-验证流程)
    - [3.1 登录验证](#31-登录验证)
    - [3.2 请求验证](#32-请求验证)
  - [4. 安全机制](#4-安全机制)
    - [4.1 地址验证](#41-地址验证)
    - [4.2 设备隔离](#42-设备隔离)
- [十二、前后端签名验证](#十二前后端签名验证)
  - [前端](#前端)
    - [1. API来源：](#1-api来源)
    - [2. 安全性](#2-安全性)
  - [后端](#后端)
  - [后端签名验证，前端是不是还要把公钥传递过来?](#后端签名验证前端是不是还要把公钥传递过来)
  - [后端验证流程说明](#后端验证流程说明)
    - [安全考虑](#安全考虑)
    - [实际应用](#实际应用)


# 一、代码解读
## 一 、参数结构体验证

```go
		if err := validator.Verify(&req); err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
```

这段代码通过 `validator.Verify` 函数进行参数验证，让我为你详细解析其验证机制：

### 1. 验证流程
```go
// 数据流向: HTTP Request Body → c.BindJSON → types.LoginReq结构体
req := types.LoginReq{}
if err := c.BindJSON(&req); err != nil {
    xhttp.Error(c, err)
    return
}

// 结构体验证
if err := validator.Verify(&req); err != nil {
    xhttp.Error(c, errcode.NewCustomErr(err.Error()))
    return
}
```

### 2. 验证规则定义
在 `types.LoginReq` 结构体中定义验证规则：
```go
type LoginReq struct {
    ChainID   int    `json:"chain_id" validate:"required"`     // 必填
    Message   string `json:"message" validate:"required"`     // 必填
    Signature string `json:"signature" validate:"required"`   // 必填
    Address   string `json:"address" validate:"required"`     // 必填
}
```

>   验证规则,支持多种验证标签：
> - required: 必填字段
> - min/max: 最小/最大值
> - len: 长度
> - email: 邮箱格式
> - oneof: 枚举值
> - httpmethod: HTTP方法验证
> - required_if: 条件必填

### 3. 验证机制
- `validator.Verify` 函数会检查结构体字段上的 `validate` 标签
- 每个字段都需要满足 `required` 约束
- 如果验证失败，返回自定义错误信息

### 4. 错误处理
```go
if err != nil {
    // 将验证错误转换为自定义错误
    xhttp.Error(c, errcode.NewCustomErr(err.Error()))
    return
}
```

### 5. 验证示例
```go
// 有效请求
{
    "chain_id": 1,
    "message": "xxx",
    "signature": "0x...",
    "address": "0x..."
}

// 无效请求 - 缺少必填字段
{
    "chain_id": 1,
    "message": "xxx"
    // 缺少 signature 和 address
}
```

## 二、生成token

```go
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": req.UserId,
		}).SignedString([]byte(config.JwtKey))
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
```

## 三、返回

```go
		xhttp.Success(c, &dto.LoginResp{
			Token: token,
}
```


## 四、 为什么 LoginReq 的ChainID接收不到值?

```go
// postman 发post请求，body传参：
 {
	"ChainID":115511,
	"Signature":"Nonce:ade2e741-2019-4d0a-a8af-43d156543f01",
	"address":"0x1D0ECb42d442baEb1A33aF74B45e0F148012941E",
	"message":"Welcome to EasySwap!\nNonce:ade2e741-2019-4d0a-a8af-43d156543f01"
}

```

根据代码分析，ChainID接收不到值的原因是JSON字段名不匹配。让我为你详细解析：

### 1. 问题原因
在 `types.LoginReq` 结构体中：
```go
type LoginReq struct {
    ChainID   int    `json:"chain_id"`  // 使用小写+下划线
    Message   string `json:"message"`
    Signature string `json:"signature"`
    Address   string `json:"address"`
}
```

而你的Postman请求中使用的是：
```json
{
    "ChainID": 115511  // 使用大写驼峰
}
```

### 2. 解决方案
修改Postman请求体为：
```json
{
    "chain_id": 115511,  // 改为小写+下划线
    "signature": "Nonce:ade2e741-2019-4d0a-a8af-43d156543f01",
    "address": "0x1D0ECb42d442baEb1A33aF74B45e0F148012941E",
    "message": "Welcome to EasySwap!\nNonce:ade2e741-2019-4d0a-a8af-43d156543f01"
}
```

### 3. 验证规则
```go
type LoginReq struct {
    ChainID   int    `json:"chain_id" validate:"required"`     // 必填
    Message   string `json:"message" validate:"required"`     // 必填
    Signature string `json:"signature" validate:"required"`   // 必填
    Address   string `json:"address" validate:"required"`     // 必填
}
```

### 4. 注意事项
1. JSON字段名必须与结构体标签完全匹配
2. 所有字段都是必填的（validate:"required"）
3. chain_id必须是支持的链ID（1: eth, 10: optimism, 11155111: sepolia）

### 5. 支持的链ID
```go
const (
    EthChainID      = 1
    OptimismChainID = 10
    SepoliaChainID  = 11155111
)
```

> 这就是为什么ChainID接收不到值的原因 - JSON字段名与结构体标签不匹配。


# 二、登录流程


每次登录都需要先调用获取登录消息的接口。让我通过代码分析来说明：

### 1. 登录流程的必要性

从 `EasySwapBackend/src/service/v1/user.go` 中的代码可以看出：

```go
func UserLogin(ctx context.Context, svcCtx *svc.ServerCtx, req types.LoginReq) (*types.UserLoginInfo, error) {
    // 从 redis 缓存中获取登录消息UUID
    cachedUUID, err := svcCtx.KvStore.Get(getUserLoginMsgCacheKey(req.Address))
    if cachedUUID == "" || err != nil {
        return nil, errcode.ErrTokenExpire
    }
    // 验证消息中的UUID
    loginUUID := strings.Trim(splits[1], "\n")
    if loginUUID != cachedUUID {
        return nil, errcode.ErrTokenExpire
    }
}
```

### 2. 安全机制

1. **一次性使用**
   ```go
   // GetUserLoginMsg 中每次生成新的UUID
   uuid := uuid.NewString()
   loginMsg := genLoginTemplate(uuid)
   ```

2. **时效性控制**
   ```go
   // 缓存72小时失效
   svcCtx.KvStore.Setex(getUserLoginMsgCacheKey(address), uuid, 72*60*60)
   ```

### 3. 登录流程

1. 前端先调用 `GetUserLoginMsg` 获取登录消息
2. 用户使用钱包签名消息
3. 前端调用 `UserLogin` 接口，提交签名和消息
4. 后端验证消息中的UUID是否匹配缓存中的UUID

### 4. 为什么必须每次都获取

1. **防重放攻击**
   - 每次生成唯一的UUID
   - 使用过的消息立即失效

2. **安全性考虑**
   - 即使消息未过期，也无法重复使用
   - 确保每次登录都是独立的会话

所以，是的，每次登录都必须先调用获取登录消息的接口，这是系统安全机制的重要组成部分。


# 三、如何理解“使用过的消息立即失效”？

> 我的疑问： 如果获取登录消息只调用一次，72小时之内再次调用登录接口，uuid不就可以重复使用么？


你的理解是正确的，但代码中有一个关键机制可以防止这种情况。让我详细解释：

### 1. UUID的生成和验证机制

在 `GetUserLoginMsg` 函数中：
```go
func GetUserLoginMsg(ctx context.Context, svcCtx *svc.ServerCtx, address string) (*types.UserLoginMsgResp, error) {
    // 每次都会生成新的uuid，旧的UUID即使未过期也无法再次使用
    uuid := uuid.NewString()
    loginMsg := genLoginTemplate(uuid)
    // 保存到redis中，72小时失效
    svcCtx.KvStore.Setex(getUserLoginMsgCacheKey(address), uuid, 72*60*60)
}
```

### 2. 防止重复使用的机制

1. **覆盖机制**
   - 每次调用 `GetUserLoginMsg` 都会生成新的UUID
   - 新的UUID会覆盖Redis中旧的UUID
   - 即使旧的UUID未过期，也会被替换

2. **验证逻辑**
   ```go
   func UserLogin(ctx context.Context, svcCtx *svc.ServerCtx, req types.LoginReq) (*types.UserLoginInfo, error) {
       // 获取当前缓存的UUID
       cachedUUID, err := svcCtx.KvStore.Get(getUserLoginMsgCacheKey(req.Address))
       // 从消息中提取UUID
       loginUUID := strings.Trim(splits[1], "\n")
       // 验证是否匹配
       if loginUUID != cachedUUID {
           return nil, errcode.ErrTokenExpire
       }

       // TODO：删除缓存的UUID
   }
   ```

   > uuid验证通过后，应该立即从redis中删除[代码中未给出，需要补充这个删除逻辑]。

### 3. 实际场景分析

1. **第一次登录**
   - 调用 `GetUserLoginMsg` 获取 UUID1
   - 使用 UUID1 登录成功（登陆成功后，立即使UUID1失效）

2. **72小时内再次登录**
   - 如果直接使用旧的 UUID1 登录，会失败
   - 因为Redis中已经被新的UUID2覆盖
   - 必须重新调用 `GetUserLoginMsg` 获取新的UUID

### 4. 安全保障

这种设计确保了：
- 每次登录都需要最新的UUID
- 防止了重放攻击
- 即使消息未过期，也无法重复使用
- 确保每次登录都是独立的会话

所以，虽然Redis中的UUID有72小时的过期时间，但每次调用 `GetUserLoginMsg` 都会生成新的UUID并覆盖旧的，这样就防止了UUID的重复使用。

# 五、什么是“重放攻击”？


让我为你详细解释重放攻击（Replay Attack）：

### 1. 什么是重放攻击

重放攻击是指攻击者截获一个合法的网络通信数据包，然后重新发送这个数据包，以达到欺骗系统的目的。

### 2. 在NFT市场中的具体场景

假设有一个恶意用户Bob想进行重放攻击：

```go
// 1. Alice正常登录流程
loginMsg := "Welcome to EasySwap!\nNonce:uuid-123"
signature := wallet.signMessage(loginMsg)

// 2. Bob截获了Alice的登录请求
{
    "chain_id": 1,
    "message": "Welcome to EasySwap!\nNonce:uuid-123",
    "signature": "0x...",
    "address": "Alice的地址"
}

// 3. Bob尝试重复发送这个请求
// 如果没有防护措施，系统可能会误认为是Alice的合法请求
```

### 3. 防护机制

在EasySwap中，通过以下机制防止重放攻击：

1. **UUID一次性使用**
```go
// 每次登录消息都有唯一UUID
uuid := uuid.NewString()
loginMsg := genLoginTemplate(uuid)

// 验证时检查UUID是否匹配
if loginUUID != cachedUUID {
    return nil, errcode.ErrTokenExpire
}
```

2. **时效性控制**
```go
// 登录消息72小时后自动失效
svcCtx.KvStore.Setex(getUserLoginMsgCacheKey(address), uuid, 72*60*60)
```

3. **覆盖机制**
```go
// 每次获取登录消息都会生成新的UUID
// 旧的UUID即使未过期也会被覆盖
uuid := uuid.NewString()
svcCtx.KvStore.Setex(getUserLoginMsgCacheKey(address), uuid, 72*60*60)
```

### 4. 实际防护效果

1. **防止消息重用**
   - 即使Bob截获了Alice的登录消息
   - 也无法重复使用，因为UUID已经被新的覆盖
   - 系统会拒绝旧的UUID

2. **防止会话劫持**
   - 每次登录都需要新的签名
   - 旧的签名即使被截获也无法使用
   - 确保每次会话都是独立的

### 5. 其他防护措施

1. **Token安全**
   - 使用AES加密
   - 30天有效期
   - 与用户地址绑定

2. **签名验证**
   - 验证签名者确实是地址所有者
   - 防止伪造身份

这种多重防护机制确保了系统的安全性，有效防止了重放攻击。



# 六、前端登录交互流程


根据代码库，我来补充完整的前端交互流程：


## 1. 前端实现路径：
- `nft-market-fe/components/wallet.tsx` - 钱包连接组件
- `nft-market-fe/api/user.ts` - 用户相关API调用

## 2. 详细流程：
1. **连接钱包**
   - 用户点击连接钱包按钮
   - 触发钱包连接，获取用户地址

2. **获取登录消息**
   ```typescript
   // 调用后端接口获取登录消息
   GET /api/v1/user/:address/login-message
   ```

3. **签名验证**
   - 用户使用钱包签名消息
   - 提交签名到后端验证
   ```typescript
   // 提交签名登录
   POST /api/v1/user/login
   ```

4. **状态管理**
   - 使用 `useGlobalState` hook 管理登录状态
   - 保存认证信息到全局状态

## 3. 代码示例

### 1. 前端获取登录消息
```typescript
// 调用获取登录消息接口
GET /user/{address}/login-message
```

### 2. 后端生成登录消息
```go
// 生成新的UUID
uuid := uuid.NewString()
// 构造登录消息模板
loginMsg := genLoginTemplate(uuid)
// 缓存UUID到Redis，72小时失效
svcCtx.KvStore.Setex(getUserLoginMsgCacheKey(address), uuid, 72*60*60)
```

### 3. 用户签名
- 前端展示登录消息给用户
- 用户使用钱包签名消息
- 获取签名结果

### 4. 前端提交登录请求
```typescript
// 提交登录参数
POST /user/login
{
    chain_id: number,
    message: string,
    signature: string,
    address: string
}
```

### 5. 后端验证流程
```go
// 1. 验证UUID匹配
cachedUUID, err := svcCtx.KvStore.Get(getUserLoginMsgCacheKey(req.Address))
loginUUID := strings.Trim(splits[1], "\n")
if loginUUID != cachedUUID {
    return nil, errcode.ErrTokenExpire
}

// 2. 验证成功后立即删除UUID
if err := svcCtx.KvStore.Del(getUserLoginMsgCacheKey(req.Address)); err != nil {
    return nil, errors.Wrap(err, "failed to delete used UUID")
}

// 3. 查询或创建用户
var user base.User
if user.Id == 0 {
    // 创建新用户
    user := &base.User{
        Address:    req.Address,
        IsAllowed:  false,
        IsSigned:   true,
        CreateTime: now,
        UpdateTime: now,
    }
}

// 4. 生成用户token
tokenKey := getUserLoginTokenCacheKey(req.Address)
userToken, err := AesEncryptOFB([]byte(tokenKey), []byte(middleware.CR_LOGIN_SALT))
```

### 6. 返回登录结果
```go
res := types.UserLoginInfo{
    Token:     hex.EncodeToString(userToken),
    IsAllowed: user.IsAllowed,
}
```

### 7. 前端处理登录结果
- 保存用户token
- 更新登录状态
- 根据IsAllowed状态决定后续操作

### 8. 安全机制
1. **一次性使用**
   - 每次生成唯一UUID
   - 使用后立即删除UUID

2. **时效性控制**
   - UUID 72小时自动失效
   - 防止长期未使用的消息被滥用

3. **防重放攻击**
   - UUID验证成功后立即删除
   - 确保消息只能使用一次


# 七、用户token生成过程


## 1. Token生成流程

### 1.1 生成Token Key
```go
tokenKey := getUserLoginTokenCacheKey(req.Address)
// 格式: CR_LOGIN_KEY:address
```

### 1.2 AES加密生成Token
```go
userToken, err := AesEncryptOFB([]byte(tokenKey), []byte(middleware.CR_LOGIN_SALT))
```

### 1.3 Token加密细节
1. **数据填充**
   ```go
   data = PKCS7Padding(data, aes.BlockSize)
   ```
   - 使用PKCS7填充确保数据块长度符合AES要求

2. **加密过程**
   ```go
   block, _ := aes.NewCipher([]byte(key))
   out := make([]byte, aes.BlockSize+len(data))
   iv := out[:aes.BlockSize]
   stream := cipher.NewOFB(block, iv)
   stream.XORKeyStream(out[aes.BlockSize:], data)
   ```
   - 使用OFB模式进行AES加密
   - 生成随机IV(初始化向量)

## 2. Token缓存

### 2.1 写入Redis
```go
if err := CacheUserToken(svcCtx, tokenKey, uuid.NewString()); err != nil {
    return nil, err
}
```

### 2.2 缓存配置
```go
func CacheUserToken(svcCtx *svc.ServerCtx, tokenKey, token string) error {
    if err := svcCtx.KvStore.Setex(tokenKey, token, 30*24*60*60); err != nil {
        return err
    }
    return nil
}
```
- 缓存时间：30天
- 使用新的UUID作为token值

## 3. Token返回

### 3.1 返回格式
```go
res := types.UserLoginInfo{
    Token:     hex.EncodeToString(userToken),
    IsAllowed: user.IsAllowed,
}
```

### 3.2 返回字段说明
- `Token`: 加密后的用户token
- `IsAllowed`: 用户是否被允许交易的标识

## 安全机制

1. **加密安全**
   - 使用AES-OFB模式加密
   - 随机IV确保每次加密结果不同
   - 固定盐值增加安全性

2. **时效性**
   - Token有效期30天
   - 过期后需要重新登录

3. **唯一性**
   - 每个地址生成唯一token
   - 使用地址作为token key的一部分


# 八、Redis中缓存的用户token为什么是uuid.newstring() ？

```go
// 使用随机UUID作为token值
if err := CacheUserToken(svcCtx, tokenKey, uuid.NewString()); err != nil {
    return nil, err
}
```

## 1. 安全性考虑

### 1.1 防止Token伪造

- UUID具有极高的随机性和唯一性
- 无法预测或伪造token值
- 增加了系统的安全性

### 1.2 避免Token冲突
- UUID保证全局唯一性
- 防止不同用户之间的token值冲突
- 确保每个用户都有唯一的token标识

## 2. 实际用途

### 2.1 Token验证
```go
// 在middleware/auth.go中
func AuthMiddleware() gin.HandlerFunc {
    // 解密前端传来的token
    decrptCode, err := AesDecryptOFB(data, []byte(CR_LOGIN_SALT))
    // 使用解密后的tokenKey从Redis获取UUID
    result, err := ctx.Get(string(decrptCode))
}

```
- UUID作为Redis中存储的token值
- 用于验证token的有效性
- 可以快速判断token是否存在

### 2.2 会话管理
- UUID作为会话的唯一标识
- 便于追踪和管理用户会话
- 支持多设备登录场景

## 3. 设计优势

1. **不可预测性**
   - UUID的随机性使得token无法被预测
   - 提高了系统的安全性

2. **全局唯一性**
   - 避免了token冲突的可能性
   - 确保每个token都是唯一的

3. **易于管理**
   - UUID格式统一，便于处理
   - 可以轻松实现token的创建和验证

4. **时效性**
   - 配合Redis的过期机制
   - 可以实现token的自动失效

这种设计既保证了安全性，又提供了良好的可管理性，是一个合理的token缓存方案。

# 九、Token与UUID的匹配机制

## 1. 两者的不同用途

### 1.1 UUID的作用
```go
// 缓存UUID作为token值
if err := CacheUserToken(svcCtx, tokenKey, uuid.NewString()); err != nil {
    return nil, err
}
```
- UUID作为Redis中存储的实际token值
- 用于验证token的有效性
- 30天后自动过期

### 1.2 AES加密Token的作用
```go
// 生成加密token返回给前端
userToken, err := AesEncryptOFB([]byte(tokenKey), []byte(middleware.CR_LOGIN_SALT))
res.Token = hex.EncodeToString(userToken)
```
- 加密token作为前端使用的凭证
- 包含了tokenKey的信息
- 用于后续请求的身份验证

## 2. 匹配验证流程

### 2.1 前端请求验证
```go
// 在middleware/auth.go中
func AuthMiddleware() gin.HandlerFunc {
    // 解密前端传来的token
    decrptCode, err := AesDecryptOFB(data, []byte(CR_LOGIN_SALT))
    // 使用解密后的tokenKey从Redis获取UUID
    result, err := ctx.Get(string(decrptCode))
}
```

### 2.2 验证逻辑
1. 前端携带加密token请求
2. 后端使用CR_LOGIN_SALT解密token
3. 解密后得到tokenKey
4. 用tokenKey从Redis获取UUID
5. 验证UUID是否存在且未过期

## 3. 安全机制

### 3.1 双重验证
- 加密token确保传输安全
- UUID验证确保会话有效

### 3.2 时效性控制
```go
// UUID有效期30天
svcCtx.KvStore.Setex(tokenKey, token, 30*24*60*60)
```

这种设计既保证了安全性（通过加密），又实现了会话管理（通过UUID），是一个合理的认证方案。


# 十、支持多设备登录

## 1. 当前设计的问题

### 1.1 覆盖机制
```go
// 每次登录都会覆盖旧的UUID
if err := CacheUserToken(svcCtx, tokenKey, uuid.NewString()); err != nil {
    return nil, err
}
```
- 新登录会立即覆盖Redis中的UUID
- 旧设备的token即使未过期也会失效
- 导致多设备登录体验不佳

### 1.2 验证失败场景
```go
// 在middleware/auth.go中
result, err := ctx.Get(string(decrptCode))
if result == "" || err != nil {
    xhttp.Error(c, errcode.ErrTokenExpire)
    c.Abort()
    return
}
```
- 旧设备请求时获取不到UUID
- 直接返回token过期错误
- 强制要求重新登录

## 2. 改进方案

### 2.1 多设备Token管理
```go
// 修改tokenKey生成策略
tokenKey := getUserLoginTokenCacheKey(req.Address, deviceID)
// 格式: CR_LOGIN_KEY:address:deviceID
```

### 2.2 设备标识
```go
type LoginReq struct {
    ChainID   int    `json:"chain_id" validate:"required"`
    Message   string `json:"message" validate:"required"`
    Signature string `json:"signature" validate:"required"`
    Address   string `json:"address" validate:"required"`
    DeviceID  string `json:"device_id"`  // 新增设备标识
}
```

### 2.3 验证逻辑优化
```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取设备ID
        deviceID := c.GetHeader("Device-ID")
        // 生成设备特定的tokenKey
        tokenKey := getUserLoginTokenCacheKey(address, deviceID)
        // 验证设备特定的token
        result, err := ctx.Get(tokenKey)
    }
}
```

## 3. 优势分析

### 3.1 多设备支持
- 每个设备独立的token
- 新设备登录不影响其他设备
- 提供更好的用户体验

### 3.2 安全性
- 设备隔离降低风险
- 单个设备泄露不影响其他设备
- 可以单独撤销特定设备权限

### 3.3 灵活性
- 支持设备管理
- 可以查看所有登录设备
- 支持远程登出特定设备

这种改进方案既保持了安全性，又提供了更好的多设备支持，是一个更合理的实现方式。

# 十一、多设备登录用户识别机制

## 1. 核心识别字段

### 1.1 用户地址（Address）
```go
type LoginReq struct {
    Address   string `json:"address" validate:"required"`
    // ...其他字段
}
```
- 用户地址是识别同一用户的主要标识
- 在区块链系统中，地址是唯一的身份标识
- 所有设备登录时必须使用相同的地址

### 1.2 Token Key生成
```go
tokenKey := getUserLoginTokenCacheKey(req.Address, deviceID)
// 格式: CR_LOGIN_KEY:address:deviceID
```
- 地址作为token key的第一部分
- 确保同一用户的不同设备共享相同的地址前缀

## 2. 设备区分机制

### 2.1 设备ID（DeviceID）
```go
type LoginReq struct {
    DeviceID  string `json:"device_id"`  // 设备标识
    // ...其他字段
}
```
- 每个设备有唯一的deviceID
- 用于区分同一用户的不同设备
- 通常由客户端生成或系统分配

### 2.2 Token存储结构
```
CR_LOGIN_KEY:address:deviceID -> UUID
```
- 同一地址可以有多个deviceID
- 每个deviceID对应独立的UUID
- 实现多设备独立管理

## 3. 验证流程

### 3.1 登录验证
```go
// 1. 验证地址
if !common.IsAddress(req.Address) {
    return nil, errcode.ErrInvalidAddress
}

// 2. 生成设备特定的token
tokenKey := getUserLoginTokenCacheKey(req.Address, deviceID)
```

### 3.2 请求验证
```go
func AuthMiddleware() gin.HandlerFunc {
    // 1. 从请求头获取设备ID
    deviceID := c.GetHeader("Device-ID")
    
    // 2. 生成设备特定的tokenKey
    tokenKey := getUserLoginTokenCacheKey(address, deviceID)
    
    // 3. 验证设备特定的token
    result, err := ctx.Get(tokenKey)
}
```

## 4. 安全机制

### 4.1 地址验证
- 验证地址格式是否合法
- 确保地址是有效的区块链地址
- 防止伪造地址

### 4.2 设备隔离
- 每个设备独立的token
- 单个设备泄露不影响其他设备
- 可以单独撤销特定设备权限

这种设计通过地址识别用户，通过设备ID区分设备，既保证了安全性，又提供了灵活的多设备支持。


# 十二、前后端签名验证

## 前端
### 1. API来源：
- 这个API来自ethers.js库，在代码中可以看到项目依赖了ethers包
- 它是Web3钱包交互的标准接口之一

```js
// 1. 获取登录消息
const { data: message } = await GetLoginMessage(address);

// 2. 用户使用钱包签名
const signature = await wallet.signMessage(message);

// 3. 提交签名登录
await Login({
    chain_id: chainId,
    message,
    signature,
    address
});

```
### 2. 安全性
- 签名过程在用户钱包中进行，私钥不会暴露
- 使用标准的ECDSA签名算法
- 签名结果会发送到后端进行验证

> 这是Web3应用中标准的用户认证流程，通过钱包签名来验证用户身份，确保安全性和去中心化特性。

## 后端

> go-ethereum包中提供了消息签名的方法。主要在github.com/ethereum/go-ethereum/crypto包中.

**核心签名方法：**

1. Sign方法:
```go
func Sign(hash []byte, prv *ecdsa.PrivateKey) ([]byte, error)
```
- 使用私钥对哈希值进行签名
- 返回65字节的签名数据（r,s,v）

2. SignHash方法
```go
func SignHash(data []byte) []byte

```
- 对数据进行Keccak256哈希
- 通常在签名前使用

3. 在项目中的典型使用场景：

```go   
import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

// 1. 准备消息
message := "Hello, world!"
hash := crypto.Keccak256Hash([]byte(message))

// 2. 使用私钥签名[r s v]
signature, err := crypto.Sign(hash.Bytes(), privateKey)
if err != nil {
    return err
}

// 3. 转换为十六进制字符串
sigHex := hexutil.Encode(signature)
```

注意事项:
- 签名前必须对消息进行Keccak256哈希
- 签名结果是65字节：32字节(r) + 32字节(s) + 1字节(v)
- v的值是27或28，用于恢复公钥

## 后端签名验证，前端是不是还要把公钥传递过来?

> 不需要前端传递公钥。在以太坊签名验证中，可以通过签名和消息恢复出公钥，然后推导出地址。

## 后端验证流程说明

1. **前端只需提供**：
   - 签名后的消息
   - 签名值
   - 签名地址

2. **后端验证过程**（见 `EasySwapBackend/src/common/utils/crypto.go`）：
```go
func VerifySig(addr, sigHex string, digest []byte) bool {
    // 1. 从签名恢复公钥
    publicKeyBytes, err := crypto.Ecrecover(digest, signature)
    
    // 2. 将公钥转换为地址
    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    
    // 3. 验证地址是否匹配
    if !strings.EqualFold(addr, address) {
        return false
    }
}
```

### 安全考虑

1. **不需要传递公钥的原因**：
   - 以太坊签名本身就包含了恢复公钥所需的信息
   - 可以通过签名和消息唯一确定签名者地址
   - 避免了公钥被篡改的风险

2. **验证机制**：
   - 使用ECDSA恢复算法从签名中提取公钥
   - 通过公钥计算地址
   - 验证计算出的地址与提供的地址是否匹配

### 实际应用

在 `EasySwapBackend/src/service/v1/user.go` 中的登录流程：
```go
// 验证签名 合法性
//ok := verifySignature(req.Message, req.Signature, req.Address)
//if !ok {
//    return nil, errors.New("invalid signature")
//}
```

这种设计更安全，因为：
1. 减少了前端需要传递的数据
2. 避免了公钥被伪造的风险
3. 利用了以太坊签名的内置验证机制

