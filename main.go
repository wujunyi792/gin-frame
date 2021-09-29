package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/wujunyi792/ginFrame/router/v1"
)

func main() {
	r := gin.Default()
	v1.LoadApi(r)
	if err := r.Run("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
}
