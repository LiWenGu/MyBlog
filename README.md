https://github.com/jiujuan/go-collection/blob/master/README.md

golang语法： golang 引入包的地方 出现这个 _ "xxx"， 表示调用 xxx包下的init方法。
mian.go为入口函数，没有调用router.go的任意方法，但是注意引入包的地方 有 _ "github.com/jicg/liteblog/routers",就默认调用routers包下的init方法。

.zshrc:
export GOPROXY=https://goproxy.io
export PATH=$PATH:/Users/mac/go/bin
export GOPATH=/Users/mac/go

go get github.com/astaxie/beego
go get -u github.com/beego/bee
bee new liteblog
go mod vendor