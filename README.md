# Database Schema Version Management

dsvm : "Database Schema Version Management"


```shell

protoc dsvm.proto --go_out=. --go-grpc_out=.

goctl api go --api dsvm.api --dir .

```