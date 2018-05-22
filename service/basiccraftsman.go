package service

import "github.com/fernandoocampo/craftsman/data"

// BasicCraftsman implements the behavior for craftsman functionalities.
type BasicCraftsman struct {
	data data.CraftsmanData
}

// Create a craftsman using a data function.
// This is an implementation of *CraftsmanService.Create
func (b *BasicCraftsman) Create(craftsman *data.Craftsman) error {
	return nil
}

// FindByID search a craftsman with the given id
// and return it.
func (b *BasicCraftsman) FindByID(id string) (*data.Craftsman, error) {
	return nil, nil
}

// ChangeState changes the state of a craftsman.
func (b *BasicCraftsman) ChangeState(id string, newstate int8) error {
	return nil
}
