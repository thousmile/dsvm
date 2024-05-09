# Database Schema Version Management

dsvm : "Database Schema Version Management"


```shell
proto 生成 go 代码

protoc --go_out=. --go-grpc_out=. proto/*.proto

golang 打包可执行文件
goreleaser --snapshot --skip-publish --clean

```
