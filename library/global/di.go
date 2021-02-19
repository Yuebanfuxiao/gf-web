package global

import (
	"go.uber.org/dig"
	"log"
)

var (
	Di = &di{container: dig.New()}
)

type di struct {
	container *dig.Container
}

// 注入服务
func (o *di) Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := o.container.Provide(constructor, opts...)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// 调取服务
func (o *di) Invoke(function interface{}, opts ...dig.InvokeOption) {
	err := o.container.Invoke(function, opts...)

	if err != nil {
		log.Fatal(err.Error())
	}
}
