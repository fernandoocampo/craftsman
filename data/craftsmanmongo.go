package data

// CraftsmanMongo for mongo connection a data process.
type CraftsmanMongo struct {
	collection string
}

// Insert creates a new document for a craftsman.
func (m *CraftsmanMongo) Insert(craftsman *Craftsman) error {
	return nil
}

// SearchByID finds and returns a craftsman entity that has the given id.
// If something goes wrong in the search, it returns an error otherwise returns nil.
func (m *CraftsmanMongo) SearchByID(id string) (*Craftsman, error) {
	return nil, nil
}

// UpdateState receives a craftsman id and a new state to update the craftsman's
// state field. If something goes wrong it returns a error otherwise returns nil.
func (m *CraftsmanMongo) UpdateState(id string, newstate int8) error {
	return nil
}

// SetMongoData set the data for database and collection
func (m *CraftsmanMongo) SetMongoData(collection string) {
	m.collection = collection
}
