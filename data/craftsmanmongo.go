package data

import (
	"errors"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CraftsmanMongo for mongo connection a data process.
type CraftsmanMongo struct {
	collection  string
	repoSession *MongoRepoSession
}

// Insert creates a new document for a craftsman.
func (m *CraftsmanMongo) Insert(craftsman *Craftsman) error {
	// make a connection to mongo database
	sessionCopy := m.repoSession.newMgoSession()
	defer sessionCopy.Close()

	// get the collection from mongo repository.
	c := sessionCopy.DB(m.repoSession.DBname()).C(m.collection)

	err := c.Insert(&craftsman)

	if err != nil {
		errmsg := fmt.Sprintf("Cannot create user: %v", err)
		log.Errorf("%s : %v\n", errmsg, err)
		return err
	}

	return nil
}

// SearchByID finds and returns a craftsman entity that has the given id.
// If something goes wrong in the search, it returns an error otherwise returns nil.
func (m *CraftsmanMongo) SearchByID(id string) (*Craftsman, error) {
	if id == "" || &id == nil {
		return nil, errors.New("Invalid craftsman id")
	}
	idval := bson.ObjectIdHex(id)
	// make a connection to mongo database
	sessionCopy := m.repoSession.newMgoSession()
	defer sessionCopy.Close()

	// References the mongo collection
	c := sessionCopy.DB(m.repoSession.DBname()).C(m.collection)

	result := Craftsman{}
	err := c.Find(bson.M{"_id": idval}).One(&result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		errmsg := "An error finding a craftsman by id - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return nil, fmt.Errorf(errmsg, err)
	}

	return &result, nil
}

// UpdateState receives a craftsman id and a new state to update the craftsman's
// state field. If something goes wrong it returns a error otherwise returns nil.
func (m *CraftsmanMongo) UpdateState(id string, newstate int8) error {
	if id == "" {
		return errors.New("Invalid craftsman id data")
	}
	// create update json map
	change := bson.M{"$set": bson.M{"state": newstate, "updated": time.Now()}}
	err := m.updateDataByID(id, change)

	if err != nil {
		errmsg := "An error updating a pack state - mongodao"
		log.Errorf("%s : %v\n", errmsg, err)
		return fmt.Errorf(errmsg, err)
	}
	return nil
}

// updateDataById update pack with the given parameter map
// that contains the data to update. Returns error if something
// goes wrong. id must be hex representation.
func (m *CraftsmanMongo) updateDataByID(id string, change interface{}) error {

	// make a connection to mongo database
	sessionCopy := m.repoSession.newMgoSession()
	defer sessionCopy.Close()

	// query the user in the database
	c := sessionCopy.DB(m.repoSession.DBname()).C(m.collection)

	//convert the ID to bson.ObjectID
	bsonid := bson.ObjectIdHex(id)

	//Here the filter is formed
	colQuerier := bson.M{"_id": bsonid}

	//Here the update is perform
	err := c.Update(colQuerier, change)

	return err

}

// SetMongoData set the data for database and collection
func (m *CraftsmanMongo) SetMongoData(repoSession RepoSession, collection string) {
	m.repoSession = repoSession.(*MongoRepoSession)
	m.collection = collection
}
