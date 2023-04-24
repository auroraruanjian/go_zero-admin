### 项目描述
项目采用go-zero和gorm开发，集成后端RBAC访问控制,登录日志，前端用户登录API。项目定位于大中型管理系统
> 待完成 数据校验，前端API流控，链路追踪，DB缓存，分布式事务，分布式定时任务

### 项目依赖
* mysql
* etcd
* go-zero
* gorm
* kafka（待加入）
* redis（待加入）
* Elasticsearch（待加入）

### 安装go和goctl工具和框架
1.
```
go mod tidy
etcd // 启动etcd
```

### 1.使用Gen工具创建Qury操作类和数据库结构
> 进到`models`目录操作
```
go run main.go
```

### 2.创建api
> 进到`api/doc/`目录执行
```
goctl api -o admin.api
goctl api go -api admin.api -dir ../
```

### 3.创建rpc
> 进到`rpc/sys/`目录操作
```
goctl rpc template -o sys.proto
goctl rpc protoc sys.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
```
### 4.运行测试
> 进到`api/`目录操作
```
go run admin.go -f etc/admin-api.yaml
```

> 进到`rpc/sys/`目录操作
```
go run sys.go -f etc/sys.yaml
```

### 5.测试运行
```bash
$ curl -i -X POST \
  http://127.0.0.1:8888/api/sys/user/login \
  -H 'content-type: application/json' \
  -d '{"userName":"admin", "password":"123456"}'
```
