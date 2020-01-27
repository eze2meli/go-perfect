package solved

import (
	"errors"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"strconv"
)

type Service struct {
	*AppContext
	ServiceLogics func(input string) int
}

var ErrServiceLogics = struct {
	InputEmpty string
	InputInvalid string
}{
	InputEmpty: "Input must not be empty ",
	InputInvalid: "Cant convert input to int. Error: ",
}
func (s *Service) serviceLogics(input string) int {
	if input == "" {
		// Non-returnable err
		err := errors.New(ErrServiceLogics.InputEmpty)
		logger.Error(err.Error(), err)
		s.LogToNR(err)
		return -1
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		// Non-returnable err
		err := errors.New(ErrServiceLogics.InputInvalid + err.Error())
		logger.Error(err.Error(), err)
		s.LogToNR(err)
		return -1
	}

	return output
}

func NewService() *Service { // Constructor
	s := &Service{}
	s.ServiceLogics = s.serviceLogics
	return s
}
