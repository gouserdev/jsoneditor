# jsoneditor
JSON editor build with go, gin and vue.js.

## download bugs

if following bug occurs, please download source code and put it into goroot path.

```sh
package github.com/gouserdev/jsoneditor: directory "/root/go/GOCODE/src/github.com/gouserdev/jsoneditor" is not using a known version control system
```

## run or build jsoneditor

### install

```sh
make install
```

### run

```sh
make start
```

### build

```sh
make build
```

## use jsoneditor as module

```sh
go get -u "github.com/gouserdev/jsoneditor"
```

```go
package main

import (
	gin "github.com/gouserdev/jsoneditor/pkg/gin"
	"github.com/gouserdev/jsoneditor/pkg/util"
)

func main() {
	configpath := "./configs/config.json"
	util.CheckJsonFile(configpath)
	//util.Open("http://127.0.0.1:10889/j/index.html?pass=" + util.GetFromJson("password", configpath))
	gin.GinInit("0.0.0.0", "10889")
}
```


## thanks 

https://github.com/rakyll/statik
