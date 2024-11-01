# maudit
微服务审计中心


### 架构说明
![arch.png](./docs/arch.png)
- 用户引入 maudit 包，会自动将Producer 中间件加载到 go-restful 框架
- 当请求到达时，Producer 会抓取相关数据写入到 Kafka 中
- Maudit 中的 Consumer 消费 Kafka 中的数据，并写入到 Mysql 数据库中
- 用户可通过 Maudit 提供的 API 查询审计数据

### 优先级问题
> 如果与 Auth 认证中间件开启的话，要通过优先级来调整顺序，确保请求先经过Auth、再经过Audit
> 否则Audit 中获取不到用户信息
- kafka：优先级 default/203，由 `ioc/default/kafka` 提供，是maudit Producer 的基础依赖
- producer：优先级 default/204，由 `maudit/provider/producer` 提供，给其他微服务使用

### 使用方式
- 配置说明
```yaml
# kafka 服务器地址、账号信息
kafka:
  username: "adminscram"
  password: "admin-secret-512"
  brokers: ["127.0.0.1:9093"]
  async: true
  batchTimeout: 10
# 审计中间件作为消费者，向Kafka 写入消息
mauditProducer:
  topic: "maudit"
```
- 加载审计中间件
```go
import (
    _ "github.com/qiaogy91/maudit/provider/producer"
)
```
- 指定资源元数据
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