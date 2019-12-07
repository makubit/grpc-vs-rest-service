package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", test)
}

func test(c *gin.Context) {

}
