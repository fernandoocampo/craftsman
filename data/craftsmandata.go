package data

// CraftsmanData defines behavior for data management related to Craftsman data.
type CraftsmanData interface {

	// Insert pushes a given craftsman data in a repository.
	// If something goes wrong an error must be returned
	// otherwise returns nil.
	Insert(craftsman *Craftsman) error

	// SearchById finds and returns a craftsman entity that has the given id.
	// If something goes wrong in the search, it returns an error otherwise returns nil.
	SearchById(id string) (*Craftsman, error)

	// UpdateState receives a craftsman id and a new state to update the craftsman's
	// state field. If something goes wrong it returns a error otherwise returns nil.
	UpdateState(id string, newstate int8) error
}
