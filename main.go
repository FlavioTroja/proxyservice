package main
import (
    "bytes"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

type AddParams struct {
    Url string `json:"url" binding:"required,gte=1"`
    Method string `json:"method"`
    Data string `json:"data"`
}

func call(c *gin.Context) {
    var ap AddParams
    if err := c.ShouldBindJSON(&ap); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input!"})
        return
    }

    // TODO: add fetch
    jsonBody := []byte(ap.Data)
    bodyReader := bytes.NewReader(jsonBody)
    res, err := http.NewRequest(ap.Method, ap.Url, bodyReader)

    log.Printf("url: %a, method: %b, data: %c", ap.Url, ap.Method, ap.Data)

    if err != nil {
        c.JSON(400, err)
        return
    }

    c.JSON(200,  gin.H{"response": res})
}


func main() {
    router := gin.Default()
    router.Use(cors.Default())
    router.POST("/", call)

    router.Run(":5000")
}
