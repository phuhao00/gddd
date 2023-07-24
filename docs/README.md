go build -ldflags "-X <import path>/config.Version=1.0.0" -o $(MY_BIN) $(MY_SRC)

go build -ldflags "-X my/package/config.Version=1.0.0" -o $(MY_BIN) $(MY_SRC)
ldflags 代表链接器标志


-X importpath.name=value
Set the value of the string variable in importpath named name to value.
Note that before Go 1.5 this option took two separate arguments.
Now it takes one argument split on the first = sign.


* 1. -w 去掉调试信息

* 2. -s 去掉符号表

* 3. -x 显示编译连接的过程

```go
go build -o=./main main.go && ls -lh main 

go build -o=./main -ldflags "-w -s" main.go && ls -lh main

```
