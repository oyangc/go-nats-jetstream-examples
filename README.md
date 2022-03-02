
# go-nats-jetstream-examples

# 1、nats安装
服务器下载安装
```shell
wget https://github.com/nats-io/nats-server/releases/download/v2.7.2/nats-server-v2.7.2-linux-amd64.zip --no-check-certificate
#直接解压即可
```
服务器运行
```shell
./nats-server &
```

客户端下载安装
```sh
#1、安装go环境，直接下载安装 
安装包下载地址为：https://golang.org/dl/。
如果打不开可以使用这个地址：https://golang.google.cn/dl/。

#2、go设置代理（用来下载github.com里的组件）
$ go env -w GOPROXY=https://goproxy.cn

#3、在要运行的go脚本目录下，安装nats客户端
$ go get github.com/nats-io/nats.go/
```

# 2、代码
jetstream \
|-- jspub.go  #jetstream pub消息 \
|-- jssub.go  #jetstream sub消息 \
|-- ncsub.go  #connect sub消息

## 运行pub/sub
1、运行之前

    go get github.com/nats-io/nats.go/

2、先运行ncsub.go和jssub.go，再运行jspub.go



# 3、nats监控natsboard

    natsboard --nats-mon-url http://127.0.0.1:8222 & 