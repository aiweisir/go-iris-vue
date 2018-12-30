package dispatch

import (
	"casbin-demo/routes/dispatch/di"

	"github.com/kataras/golog"
	"github.com/kataras/iris/context"
)

// def is the default herp value which can be used for dependencies share.
var def = New()

// Dispater contains the Dependencies which will be binded
// to the controller(s) or handler(s) that can be created
// using the Dispater's `Handler` and `Controller` methods.
//
// This is not exported for being used by everyone, use it only when you want
// to share heroes between multi mvc.go#Application
// or make custom hero handlers that can be used on the standard
// iris' APIBuilder. The last one reason is the most useful here,
// although end-devs can use the `MakeHandler` as well.
//
// For a more high-level structure please take a look at the "mvc.go#Application".
type Dispater struct {
	values di.Values
}

// New returns a new Dispater, a container for dependencies and a factory
// for handlers and controllers, this is used internally by the `mvc#Application` structure.
// Please take a look at the structure's documentation for more information.
func New() *Dispater {
	return &Dispater{
		values: di.NewValues(),
	}
}

// Dependencies returns the dependencies collection if the default hero,
// those can be modified at any way but before the consumer `Handler`.
func Dependencies() *di.Values {
	return def.Dependencies()
}

// Dependencies returns the dependencies collection of this hero,
// those can be modified at any way but before the consumer `Handler`.
func (h *Dispater) Dependencies() *di.Values {
	return &h.values
}

// Register adds one or more values as dependencies.
// The value can be a single struct value-instance or a function
// which has one input and one output, the input should be
// an `iris.Context` and the output can be any type, that output type
// will be binded to the handler's input argument, if matching.
//
// Example: `.Register(loggerService{prefix: "dev"}, func(ctx iris.Context) User {...})`.
func Register(values ...interface{}) *Dispater {
	return def.Register(values...)
}

// Register adds one or more values as dependencies.
// The value can be a single struct value-instance or a function
// which has one input and one output, the input should be
// an `iris.Context` and the output can be any type, that output type
// will be binded to the handler's input argument, if matching.
//
// Example: `.Register(loggerService{prefix: "dev"}, func(ctx iris.Context) User {...})`.
func (h *Dispater) Register(values ...interface{}) *Dispater {
	h.values.Add(values...)
	return h
}

// Clone creates and returns a new hero with the default Dependencies.
// It copies the default's dependencies and returns a new hero.
func Clone() *Dispater {
	return def.Clone()
}

// Clone creates and returns a new hero with the parent's(current) Dependencies.
// It copies the current "h" dependencies and returns a new hero.
func (h *Dispater) Clone() *Dispater {
	child := New()
	child.values = h.values.Clone()
	return child
}

// Handler accepts a "handler" function which can accept any input arguments that match
// with the Dispater's `Dependencies` and any output result; like string, int (string,int),
// custom structs, Result(View | Response) and anything you can imagine.
// It returns a standard `iris/context.Handler` which can be used anywhere in an Iris Application,
// as middleware or as simple route handler or subdomain's handler.
func Handler(handler interface{}) context.Handler {
	return def.Handler(handler)
}

// Handler accepts a handler "fn" function which can accept any input arguments that match
// with the Dispater's `Dependencies` and any output result; like string, int (string,int),
// custom structs, Result(View | Response) and anything you can imagine.
// It returns a standard `iris/context.Handler` which can be used anywhere in an Iris Application,
// as middleware or as simple route handler or subdomain's handler.
func (h *Dispater) Handler(fn interface{}) context.Handler {
	handler, err := makeHandler(fn, h.values.Clone()...)
	if err != nil {
		golog.Errorf("Dispatch Handler: %v", err)
	}
	return handler
}
