package example_logic

import (
	"context"

	ccomand "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	data "github.com/pip-services3-gox/pip-services3-swagger-gox/example/data"
)

type DummyCommandSet struct {
	ccomand.CommandSet
	controller IDummyController
}

func NewDummyCommandSet(controller IDummyController) *DummyCommandSet {
	c := DummyCommandSet{}
	c.CommandSet = *ccomand.NewCommandSet()

	c.controller = controller

	c.AddCommand(c.makeGetPageByFilterCommand())
	c.AddCommand(c.makeGetOneByIdCommand())
	c.AddCommand(c.makeCreateCommand())
	c.AddCommand(c.makeUpdateCommand())
	c.AddCommand(c.makeDeleteByIdCommand())
	return &c
}

func (c *DummyCommandSet) makeGetPageByFilterCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_dummies",
		cvalid.NewObjectSchema().WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var filter *cdata.FilterParams
			var paging *cdata.PagingParams

			if data, contains := args.Get("filter"); contains {
				filter = cdata.NewFilterParamsFromValue(data)
			}

			if data, contains := args.Get("paging"); contains {
				paging = cdata.NewPagingParamsFromValue(data)
			}

			return c.controller.GetPageByFilter(ctx, correlationId, filter, paging)
		},
	)
}

func (c *DummyCommandSet) makeGetOneByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_dummy_by_id",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			id := args.GetAsString("dummy_id")
			return c.controller.GetOneById(ctx, correlationId, id)
		},
	)
}

func (c *DummyCommandSet) makeCreateCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"create_dummy",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy", data.NewDummySchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var entity data.Dummy

			if _val, ok := args.Get("dummy"); ok {
				val, _ := cconv.JsonConverter.ToJson(_val)
				obj, err := cconv.JsonConverter.FromJson(val)

				if err != nil {
					return nil, err
				}

				entity = obj.(data.Dummy)
			}

			return c.controller.Create(ctx, correlationId, entity)
		},
	)
}

func (c *DummyCommandSet) makeUpdateCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"update_dummy",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy", data.NewDummySchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var entity data.Dummy

			if _val, ok := args.Get("dummy"); ok {
				val, _ := cconv.JsonConverter.ToJson(_val)
				obj, err := cconv.JsonConverter.FromJson(val)

				if err != nil {
					return nil, err
				}

				entity = obj.(data.Dummy)
			}
			return c.controller.Update(ctx, correlationId, entity)
		},
	)
}

func (c *DummyCommandSet) makeDeleteByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"delete_dummy",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			id := args.GetAsString("dummy_id")
			return c.controller.DeleteById(ctx, correlationId, id)
		},
	)
}
