package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	version = "v2"
)

type CodeMessage struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
}

func (c *CodeMessage) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}
func main() {
	router := gin.Default()
	msg := &CodeMessage{
		Code:    200,
		Version: version,
	}
	router.GET("/", func(c *gin.Context) {
		c.String(200, msg.ToJsonString())
	})
	http.ListenAndServe(":9000", router)
}
