package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
)

func GetHashFromHeader(context *gin.Context) string {
	digest := context.GetHeader("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256" {
		return ""
	}
	return digest[8:]
}

func GetSizeFromHeader(context *gin.Context) int64 {
	return context.Request.ContentLength
}

//func GetSizeFromHeader(h http.Header) int64 {
//	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
//	return size
//}

func CalculateHash(r io.Reader) string {
	h := sha256.New()
	io.Copy(h, r)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
