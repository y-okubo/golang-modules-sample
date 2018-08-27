## 使ってみる

```
env GO111MODULE=on go mod init
env GO111MODULE=on go build -a
```

/go/pkg/mod になにやらできる

## GOPATH 外でやる

以下の記述が必要

```
package main // import "github.com/y-okubo/golang-modules-sample"
```