package ex4_error_catalog_newrelic

import (
	"errors"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"strconv"
)

type Service struct {
	ServiceLogics func(input string) int
}

func (Service) serviceLogics(input string) int {
	if input == "" {
		// Non-returnable err
		err := errors.New("Input must not be empty ")
		logger.Error(err.Error(), err)
		return -1
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		// Non-returnable err
		err := errors.New("Cant convert input to int. Error: " + err.Error())
		logger.Error(err.Error(), err)
		return -1
	}

	return output
}

func NewService() *Service { // Constructor
	s := &Service{}
	s.ServiceLogics = s.serviceLogics
	return s
}
