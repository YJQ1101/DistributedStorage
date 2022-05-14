package temp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type tempInfo struct {
	Uuid string
	Name string
	Size int64
}

func post(context *gin.Context) {
	output, _ := exec.Command("uuidgen").Output()
	uuid := strings.TrimSuffix(string(output), "\n")
	name := context.Param("uuid")
	size, err := strconv.ParseInt(context.Request.Header.Get("size"), 0, 64)
	if err != nil {
		log.Println(err)
		return
	}
	t := tempInfo{uuid, name, size}
	err = t.writeToFile()
	if err != nil {
		log.Println(err)
		return
	}
	os.Create(os.Getenv("STORAGE_ROOT") + "/temp/" + t.Uuid + ".dat")
	context.String(http.StatusOK, uuid)
}

func (t *tempInfo) writeToFile() error {
	f, err := os.Create(os.Getenv("STORAGE_ROOT") + "/temp/" + t.Uuid)
	if err != nil {
		return err
	}
	defer f.Close()
	b, _ := json.Marshal(t)
	f.Write(b)
	return nil
}
