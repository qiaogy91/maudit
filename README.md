# maudit
微服务审计中心


### 架构说明
![arch.png](./docs/arch.png)
1. 用户引入 Maudit 包，会自为框架加载 Maudit 中间件
2. 用户的所有路由都会被识别并记录，而后在中间件中通过 gRPC 方式发送到 Maudit 服务端
3. Maudit 服务端内部会将用户的审计数据写入到 Kafka 中（异步）
4. Maudit 服务端内部会启动一个 Kafka Consumer，负责从 Kafka 中读取数据，并调用 Event 模块写入到 Mysql 中
5. Maudit 同时提供 Rest API 供用户查询审计数据

### 优先级问题
> 如果与 Auth 认证中间件开启的话，要通过优先级来调整顺序，确保请求先经过Auth、再经过Audit
> 否则Audit 中获取不到用户信息
1. 在 `Project/provider/` 目录下，确保 `Client` 对象在 `Middleware` 对象之前加载
2. 如果开启了 Auth 认证中间件，请确保 Auth 在 Maudit 之前加载（否则 Maudit 获取不到用户信息）


### 配置说明
#### 服务端配置
```yaml
application:
  appName: "maudit"
  description: "微服务审计中心"
  domain: "example.com"

log:
  otlp: true

http:
  enabled: true
  host: "0.0.0.0"
  port: 8081
  otlp: true

grpc:
  enabled: true
  host: "0.0.0.0"
  port: 18081
  otlp: true
  token: "maudit-token"

otlp:
  httpEndpoint: "10.54.44.78:4318"
  grpcEndpoint: "10.54.44.78:4317"
  enableTLS: false

# 内部使用Kafka 客户端
kafka:
  username: "adminscram"
  password: "admin-secret-512"
  brokers: [ "127.0.0.1:9093" ]
  async: true
  batchTimeout: 10

# bridge 模块负责从Kafka 读取数据，而后调用Event 模块写入数据库
bridge:
  topic: "maudit"
  duration: 5
  
# Event 模块写入Mysql 数据库，也可替换为其他实现
datasource:
  otlp: true
  host: "127.0.0.1"
  port: 3306
  database: "CloudManager"
  username: "root"
  password: "redhat"
  debug: true
```
#### 客户端配置
- 配置 mauditClient
```yaml
# 指向审计中心的gRPC服务
mauditClient:
  host: "127.0.0.1"
  port: 18081
  token: "maudit-token"
```
- 加载审计中间件
```go
import (
    _ "github.com/qiaogy91/maudit/provider/authentication"
)
```
- 指定路由元数据
```go
func (h *Handler) Registry() {
    tag := []string{"凭证管理"}
    
    ws := gorestful.ModuleWebservice(h)
    ws.Route(ws.GET("/").To(h.Test).
    Doc("测试接口").
    Metadata(restfulspec.KeyOpenAPITags, tag).
    Metadata("auth", false).
    Metadata("audit", true).
    Metadata("resource", secret.AppName).
    Metadata("action", "get"),
    )
}
```