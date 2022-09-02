package example_services

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
)

type DummyCommandableHttpService struct {
	*cservices.CommandableHttpService
}

func NewDummyCommandableHttpService() *DummyCommandableHttpService {
	c := DummyCommandableHttpService{}
	c.CommandableHttpService = cservices.InheritCommandableHttpService(&c, "dummies2")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("pip-services-dummies", "controller", "default", "*", "*"))
	return &c
}
