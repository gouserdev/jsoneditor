format:
	gofmt -w ./

build:
	export GOARCH=386
	export GOOS=linux
	go build -o jsoneditor ./cmd/main.go

install:
	go get -u "github.com/tidwall/sjson"
	go get -u "github.com/rakyll/statik"
	go get -u "github.com/Wissance/stringFormatter"
	go get -u "github.com/gouserdev/mongoescape"
	go get -u "github.com/tidwall/gjson"
	go get -u "go.mongodb.org/mongo-driver"
	go get -u "github.com/mongodb/mongo-go-driver"
	go get -u "github.com/dlclark/regexp2"
#===============================================
.DEFAULT_GOAL := start
.PHONY: start
start: ## Run USSD Telegram Bot
	go run ./cmd/main.go