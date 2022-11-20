package domain

type HttpResponse interface {
	Status() int
	Response() interface{}
}

type ApiResponse struct {
	StatusCode      int
	ServiceResponse interface{}
}

func (s *ApiResponse) Status() int {
	return s.StatusCode
}

func (s *ApiResponse) Response() interface{} {
	return s.ServiceResponse
}

