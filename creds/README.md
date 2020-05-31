
一、证书认证
openssl genrsa -out server.key 2048
openssl req -new -x509 -days 3650  -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" -key server.key -out server.crt

openssl genrsa -out client.key 2048
openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" -key client.key -out client.crt

其中-subj中/CN=server.grpc.io 表示服务器的名称为server.grpc.io，在验证服务器时需要用到

当客户端没有证书时去请求服务器
报错：rpc error: code = Unavailable desc = connection closed
exit status

这种方式生成的证书，需要提前将服务器证书告知客户端，客户端才能服务器证书进行认证。在复复杂的网络环境下显得非常危险，如果在中间某个环节被监听或替换，那就玩大发了。



二、根证书生成方式
服务端
openssl genrsa -out server.key 2048
openssl req -new -x509 -days 3650  -subj "/C=GB/L=China/O=server/CN=server.io" -key server.key -out server.csr
openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.crt

客户端
openssl genrsa -out client.key 2048
openssl req -new -x509 -days 3650  -subj "/C=GB/L=China/O=client/CN=client.io" -key client.key -out client.csr
openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.crt

报错：PEM_read_bio:no start line:pem_lib.c:707:Expecting: CERTIFICATE REQUEST 
错误未解决