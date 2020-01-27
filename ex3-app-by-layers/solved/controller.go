package solved

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func (ctrl *Controller) ControllerHandler(c *gin.Context) {
	input := c.Query("input")
	result := ctrl.service.ServiceLogics(input)
	c.JSON(200, fmt.Sprintf("The result is %v", result))
}

func NewController() *Controller { // Constructor
	c := &Controller{}
	c.service = NewService()
	return c
}
