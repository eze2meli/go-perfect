package solved

import "github.com/gin-gonic/gin"

type AppContext struct {
	c *gin.Context
	transaction *Transaction // Request-scoped transaction
}

func (a *AppContext) LogToNR(err error) {
	if a.transaction == nil {
		a.transaction = CreateTransaction(a.c)
	}
	a.transaction.LogToNewRelic(err)
}

func NewAppContext(c *gin.Context) *AppContext { // Constructor
	return &AppContext{c: c}
}