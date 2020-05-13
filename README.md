
### 启动micro网关

```
micro --api_namespace=go.micro.conf.web  api --handler=web
```

### 运行服务
```
go run conf-server/main.go
```

### 运行api接口
```
go run conf-api/main.go
```

### protobuf生成方式
```
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. *.proto
```