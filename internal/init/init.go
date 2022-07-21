package init

import (
	"flag"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	initGlog()
}

func initGlog() {
	var logDirStr = "./log"
	//config for gin request log
	createFolder("log")
	{
		f, _ := os.Create(path.Join(logDirStr, "gin.log"))
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
	//config for glog
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", logDirStr)
	flag.Set("v", strconv.Itoa(1))

	if _, err := os.Stat(logDirStr); os.IsNotExist(err) {
		os.Mkdir(logDirStr, os.ModeAppend)
	}
	flag.Parse()
}

func createFolder(dir string) error {
	if len(dir) == 0 {
		return nil
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		return err
	} else if err != nil {
		return err
	}
	return nil
}
