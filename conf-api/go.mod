module github.com/golango.cn/golden-conf/conf-api

go 1.14

replace github.com/golango.cn/golden-conf/conf-proto => ../conf-proto

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/golango.cn/golden-conf/conf-proto v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.5.0
)
