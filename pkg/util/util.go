package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

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
