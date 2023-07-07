## 编译项目
* `./build.sh  helloworld/v1`
* 如果重新生成 batch 下的 protoc，可以 `./build.sh  batch/v1`

## 注意事项
* 如果缺少依赖，可以执行以下命令安装

```shell
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
```
