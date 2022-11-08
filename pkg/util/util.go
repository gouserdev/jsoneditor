package util

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func GetFromJson(gpath string, filePath string) string {
	if gpath == "" {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		return strutil.MustString(content)
	} else {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		value := gjson.Get(string(content), gpath).String()
		return value
	}
}

func SetJsonValue(gpath string, value string, filePath string) {
	value, _ = sjson.Set(string(value), gpath, value)
	//dump.P(value)
	value_byte := strutil.ToBytes(value)
	ioutil.WriteFile(filePath, value_byte, os.ModePerm)
}

func Forever() {
	for {
		//fmt.Printf("%v+\n", time.Now())
		//dump.P("t")
		time.Sleep(time.Second)
	}
}

func handle_util_Bug() {
	dump.P("handle_Bug")
}
