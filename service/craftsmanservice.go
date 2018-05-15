package service

import "github.com/fernandoocampo/craftsman/data"

// CraftsmanService defines craftsman service behavior for management purpose.
type CraftsmanService interface {

	// Create inserts a new craftsman in the system. An error is returned if
	// something goes wrong.
	Create(craftsman data.Craftsman) error

	// FindByID search a craftsman with the given id
	// and return it.
	FindByID(id string) (*data.Craftsman, error)
}
