package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Invoke[T any](url string, method string, payload any) (resp *T, err error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	log.Println(res)

	if err != nil {
		return nil, errors.Wrap(err, "making request to endpoint")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading ADB response body")
	}

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "decoding JSON response body")
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("response not OK from ADB, with body: %+v", resp)
	}

	return resp, nil
}

type AddParams struct {
	Url     string `json:"url" binding:"required,gte=1"`
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

func call(c *gin.Context) {
	var ap AddParams
	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}

	res, err := Invoke[string](ap.Url, ap.Method, ap.Payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	log.Println("Response from endpoint: " + *res)

	c.JSON(200, gin.H{"response": *res})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/", call)

	router.Run(":5000")
}
