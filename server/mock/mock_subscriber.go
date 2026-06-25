package mock

import (
	"micro.labqa.pp.ua/v6/registry"
	"micro.labqa.pp.ua/v6/server"
)

type MockSubscriber struct {
	Opts server.SubscriberOptions
	Sub  interface{}
	Id   string
}

func (m *MockSubscriber) Topic() string {
	return m.Id
}

func (m *MockSubscriber) Subscriber() interface{} {
	return m.Sub
}

func (m *MockSubscriber) Endpoints() []*registry.Endpoint {
	return []*registry.Endpoint{}
}

func (m *MockSubscriber) Options() server.SubscriberOptions {
	return m.Opts
}
