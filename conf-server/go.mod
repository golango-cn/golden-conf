module github.com/golango.cn/golden-conf/conf-server

go 1.14

replace github.com/golango.cn/golden-conf/conf-proto => ../conf-proto
replace github.com/golango.cn/golden-conf/conf-database => ../conf-database

require (
	github.com/golango.cn/golden-conf/conf-database v0.0.0-00010101000000-000000000000
	github.com/golango.cn/golden-conf/conf-proto v0.0.0-00010101000000-000000000000
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.0
)
