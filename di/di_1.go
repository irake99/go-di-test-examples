package di

import "fmt"

type ServiceA struct{}

func (s *ServiceA) Process() {
	fmt.Println("Processing in the service.")
}

type ClientA struct{}

func (c *ClientA) CallServiceWithoutDI() {
	service := ServiceA{}
	service.Process()
}

// --------------------------------------------

type ServiceB interface {
	Process()
}

type ServiceBImpl struct{}

func (s *ServiceBImpl) Process() {
	fmt.Println("Processing in the service.")
}

type ClientB struct {
	service ServiceB
}

func NewClientB(service ServiceB) *ClientB {
	return &ClientB{
		service: service,
	}
}

func (c *ClientB) CallServiceWithDI() {
	c.service.Process()
}

// --------------------------------------------

func Test2() {
	c1 := ClientA{}
	c1.CallServiceWithoutDI()

	// ----------------------

	s := &ServiceBImpl{}
	c2 := NewClientB(s)
	c2.CallServiceWithDI()
}
