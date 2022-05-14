package temp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func patch(context *gin.Context) {
	uuid := context.Param("uuid")
	tempinfo, err := readFromFile(uuid)
	if err != nil {
		log.Println(err)
		context.String(http.StatusNotFound, "404 error")
		return
	}
	infoFile := os.Getenv("STORAGE_ROOT") + "/temp/" + uuid
	datFile := infoFile + ".dat"
	f, err := os.OpenFile(datFile, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500 error")
		return
	}
	defer f.Close()
	_, err = io.Copy(f, context.Request.Body)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500 error")
		return
	}
	info, err := f.Stat()
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500 error")
		return
	}

	actual := info.Size()
	if actual > tempinfo.Size {
		os.Remove(datFile)
		os.Remove(infoFile)
		log.Println("actual size", actual, "exceeds", tempinfo.Size)
		context.String(http.StatusInternalServerError, "500 error")
	}
}

func readFromFile(uuid string) (*tempInfo, error) {
	f, err := os.Open(os.Getenv("STORAGE_ROOT") + "/temp/" + uuid)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	var info tempInfo
	json.Unmarshal(b, &info)
	return &info, nil
}
