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

## import jsoneditor into your project and then copy configs/config.json into your code

```sh
go get -u "github.com/gouserdev/jsoneditor"
```

```go
package main

import (
	je "github.com/gouserdev/jsoneditor"
)

func main(){
	je.Jsoneditor("0.0.0.0", "10889")
}
```


## thanks 

https://github.com/rakyll/statik
