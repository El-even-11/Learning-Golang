package main

import (
	"os"
	"topicpage/controller"
	"topicpage/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repo.Init(filePath); err != nil {
		return err
	}
	return nil
}