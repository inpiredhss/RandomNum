package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/rand", Rand)
	r.Run(":8080")
}

func Rand(c *gin.Context) {
	p := Rand0()
	c.JSON(200, p)
}

func Rand0() int {
	s2 := rand.NewSource(time.Now().Unix())
	r2 := rand.New(s2)
	return r2.Intn(4)
}
