package example_logic

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	data "github.com/pip-services3-gox/pip-services3-swagger-gox/example/data"
)

type IDummyController interface {
	GetPageByFilter(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result *data.DummyDataPage, err error)
	GetOneById(ctx context.Context, correlationId string, id string) (result *data.Dummy, err error)
	Create(ctx context.Context, correlationId string, entity data.Dummy) (result *data.Dummy, err error)
	Update(ctx context.Context, correlationId string, entity data.Dummy) (result *data.Dummy, err error)
	DeleteById(ctx context.Context, correlationId string, id string) (result *data.Dummy, err error)
}
