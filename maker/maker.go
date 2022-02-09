package apimaker

import (
	"errors"

	"github.com/gotbitoriginal/common"
)

var (
	ErrAlreadyAdded = errors.New("client already added")
	ErrNotFound     = errors.New("client not found")
)

type Maker interface {
	Add(factory common.APIClientFactory) (id string, err error)
	New(clientID string, options ...common.Option) (common.APIClientBase, error)
	GetIDs() []string
}

func NewMaker() Maker {
	return &maker{
		factories: make(map[string]common.APIClientFactory),
	}
}

var _ Maker = (*maker)(nil)

type maker struct {
	factories map[string]common.APIClientFactory
}

func (m *maker) Add(factory common.APIClientFactory) (id string, err error) {
	client := factory()

	id = client.GetClientID()
	if _, ok := m.factories[id]; ok {
		return id, ErrAlreadyAdded
	}

	m.factories[id] = factory
	return id, nil
}

func (m *maker) New(clientID string, options ...common.Option) (common.APIClientBase, error) {
	factory, ok := m.factories[clientID]
	if !ok {
		return nil, ErrNotFound
	}

	client := factory()
	if err := client.Init(options...); err != nil {
		return nil, err
	}
	return client, nil
}

func (m *maker) GetIDs() []string {
	keys := make([]string, len(m.factories))

	i := 0
	for k := range m.factories {
		keys[i] = k
		i++
	}
	return keys
}
