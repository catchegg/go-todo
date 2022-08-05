# go-todo
todo app with golang



开发环境: go version go1.16.1 windows/amd64



TODO

- [ ] 用户注册、登录
- [ ] TODO分类管理
- [ ] TODO事项管理







LOG

2022-08-05 15:53:03: 找一个命令行交互框架 cobra,最新版1.5.0 

下载

```
go install github.com/spf13/cobra-cli@latest
```

$GOPATH/bin 下面有了`cobra-cli.exe`, idea里面Terminal 执行 `cobra-cli` 提示找不到命令。

要把$GOPATH/bin添加到环境变量 $PATH

初始化项目

```
E:\go\go-todo>cobra-cli init go-todo
Your Cobra application is ready at
E:\go\go-todo/go-todo

E:\go\go-todo>

```

在main.go 所在的目录，执行 `cobra-cli add user`， 添加user命令,会在cmd目录生成user.go 文件。

```
E:\go\go-todo>cobra-cli add user
user created at E:\go\go-todo

E:\go\go-todo>go run main.go user
user called

E:\go\go-todo>
```



注册账号

```
go run main.go user add hello 123456
```

