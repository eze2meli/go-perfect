package solved

import "github.com/gin-gonic/gin"

// Info: Cant modify this file

var loggedNewRelicErr error

type Transaction struct { }

func (*Transaction) LogToNewRelic(err error) {
	loggedNewRelicErr = err
}

func CreateTransaction(c *gin.Context) *Transaction {
	if c == nil {
		panic("gin Context cant be nil. Its is required for magic")
	}

	return &Transaction{}
}