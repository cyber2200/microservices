package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BASE_URL = "http://auth_service:3002"

func Request(c *gin.Context, reqPath string) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic(err)
	}
	resp, err := http.Post(BASE_URL+reqPath, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Panic(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	c.JSON(http.StatusOK, res)
}
