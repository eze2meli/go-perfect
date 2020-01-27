package controller

import (
	"fmt"
	"github.com/eze2meli/go-perfect/ex3-app-by-layers/service"
	"github.com/gin-gonic/gin"
)

func ControllerHandler(c *gin.Context) {
	input := c.Query("input")

	result := service.ServiceLogics(input)

	c.JSON(200, fmt.Sprintf("The result is %v", result))
}
