# go-quickstart

结合**DDD&整洁架构&责任链**构建的go项目开发基础框架

### 涉及到：
- Kafka
- RabbitMQ
- Redis (封装LRU)
- MySQL
- Mongo
- Elastic Search
- Hive (TBD)
- gRpc (TBD)

### 责任链：
核心代码于`api/middleware/taskchain/core.go`下
使用方法：继承此目录下`TaskContextData`接口，自定义链子业务流程中所需数据，task包内定义责任链节点（具体实现的任务），实现应用即可

PS:所有配置目前仅适用于单机，集群相关TBD。
