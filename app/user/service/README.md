# User Service

## build
go build -o ./bin/user ./app/user/service/cmd/server/...


## run
bin/server  -conf app/user/service/configs 


## 打开grpc调试
grpcui -plaintext 127.0.0.1:8000
> 注意;需要给服务加个反射才可以运行
  `reflection.Register(svr)`