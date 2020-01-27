package ex4_error_catalog_newrelic

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (ctrl *Controller) ControllerHandler(c *gin.Context) {
	input := c.Query("input")

	s := NewService()

	result := s.ServiceLogics(input)
	c.JSON(200, fmt.Sprintf("The result is %v", result))
}

func NewController() *Controller { // Constructor
	c := &Controller{}
	return c
}
