# http-shell

提供通过http接口执行shell脚本的功能

## 编译执行

编译适合当前系统的可执行文件：
```
make http-shell
```

编译全平台的可执行文件：
```
make all
```

生成的可执行文件在bin目录下

默认端口为8080，也可以通过`-p <port>`指定端口启动项目


## Docker运行

执行下面这行命令可以得到一个编译好的镜像
```
docker build --no-cache -t http-shell:latest .
```

编译好镜像以后，执行下面的命令，可以后台启动项目
```
docker run --name http-shell -p 8080:8080 -d http-shell:latest
```

## 接口介绍

### /sh

`/sh`接口支持使用sh命令运行指定的脚本

使用示例
```
curl --request POST 'http://127.0.0.1:8080/sh' \
--header 'Content-Type: application/json' \
--data '{
    "shellPath":"/Users/martin/Downloads/test.sh"
}'
```

### /execShell

`/execShell`接口支持运行指定的脚本

使用示例
```
curl --request POST 'http://127.0.0.1:8080/execShell' \
--header 'Content-Type: application/json' \
--data '{
    "shellPath":"/Users/martin/Downloads/test.sh"
}'
```


