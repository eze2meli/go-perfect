package solved

type Service struct {
	ServiceLogics func(input string) string
}

func (Service) serviceLogics(input string) string {
	panic("Not implement yet")
}

func NewService() *Service { // Constructor
	s := &Service{}
	s.ServiceLogics = s.serviceLogics
	return s
}
