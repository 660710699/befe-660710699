package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student struct
type Toy struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var toy = []Toy{
	{ID: 1, Name: "Gundam", Category: "Model"},
	{ID: 2, Name: "Uno", Category: "BroadGame"},
	{ID: 3, Name: "PogDang", Category: "BroadGame"},
	{ID: 4, Name: "OnePiece Model", Category: "Model"},
	{ID: 5, Name: "Evagrlion", Category: "Model"},
}

func getToy(c *gin.Context) {
	toyCategory := c.Query("category")

	if toyCategory != "" {
		filter := []Toy{}
		for _, toy := range toy {
			if fmt.Sprint(toy.Category) == toyCategory {
				filter = append(filter, toy)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, toy)
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/toy", getToy)

	}

	r.Run(":8080")
}
