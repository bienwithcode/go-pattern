package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

type Result struct {
	Id   int      `json:"id"`
	Name []string `json:"name"`
}

func main() {
	r := gin.Default()
	r.Use(recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/process-request", handleRequest())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, r)
			}
		}()
		c.Next()
	}
}

func handleRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []Student
		if err := c.ShouldBind(&students); err != nil {
			panic(err.Error())
		}

		// var result []Result
		result := make(map[int]Result)

		for _, v1 := range students {
			if _, ok := result[v1.Id]; !ok {
				var item Result
				item.Id = v1.Id
				for _, v2 := range students {
					if v1.Id == v2.Id {
						item.Name = append(item.Name, v2.Name)
					}
				}
				result[v1.Id] = item
			}

		}
		res := make([]Result, 0, len(result))

		for _, r := range result {
			res = append(res, r)
		}

		c.JSON(http.StatusOK, res)
	}
}
