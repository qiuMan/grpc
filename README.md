# grpc
  gRPC是谷歌公司基于Protobuf开发的跨语言的开源RPC框架，gRPC是基于http/2协议设计，可以基于一个http/2链接提供多个服务。区别于RPC的是，RPC是远程函数调用，数据交互只能是函数参数和返回值，数据量不能太大，否则会影响响应时间，因此并不适合大量数据的上传和下载。而gRPC框架的流特性可以实现实时流数据处理，从而解决数据量问题。
 
 # protobuf
 ```json
 $ protoc --go_out=plugins=grpc:. hello.proto
 ```
