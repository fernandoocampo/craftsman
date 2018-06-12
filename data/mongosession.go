package data

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// MongoRepoSession contains all data to connect to mongo repository.
type MongoRepoSession struct {
	mgoaddrs   []string     // mongo addresses
	timeout    int          // mgo conn in seconds
	dbname     string       // mongo database name
	usrdb      string       // mongo db user name
	pwddb      string       // mongo database password
	mgoSession *mgo.Session // mongo session
}

// InitSession creates and returns a new Mongo Session.
func (m *MongoRepoSession) InitSession() interface{} {
	mgodialinfo := &mgo.DialInfo{
		Addrs:    m.mgoaddrs,
		Timeout:  time.Duration(m.timeout) * time.Second,
		Database: m.dbname,
		Username: m.usrdb,
		Password: m.pwddb,
	}

	fmt.Println("Connecting to db...")
	session, err := mgo.DialWithInfo(mgodialinfo)

	if err != nil {
		panic(err)
	}
	session.SetSafe(&mgo.Safe{})
	session.SetSyncTimeout(3 * time.Second)
	session.SetSocketTimeout(3 * time.Second)

	m.mgoSession = session

	return session
}

// CloseSession closes the root mongo session.
func (m *MongoRepoSession) CloseSession() {
	m.mgoSession.Close()
}

// newSession returns a copy of the mongo session
func (m *MongoRepoSession) newSession() interface{} {
	return m.newMgoSession()
}

// newMgoSession returns a copy of the mongo session
func (m *MongoRepoSession) newMgoSession() *mgo.Session {
	return m.mgoSession.Copy()
}

// SetAddresses set *mgoaddrs value
func (m *MongoRepoSession) SetAddresses(addrs []string) {
	m.mgoaddrs = addrs
}

// SetTimeout set time out
func (m *MongoRepoSession) SetTimeout(ttl int) {
	m.timeout = ttl
}

// SetDBname set mongo database name
func (m *MongoRepoSession) SetDBname(name string) {
	m.dbname = name
}

// DBname returns the name of the database.
func (m *MongoRepoSession) DBname() string {
	return m.dbname
}

// SetUserdb set mongo db user name
func (m *MongoRepoSession) SetUserdb(usr string) {
	m.usrdb = usr
}

// SetPwddb set mongo database password
func (m *MongoRepoSession) SetPwddb(pwd string) {
	m.pwddb = pwd
}

// SetMgoSession set the mongo session
func (m *MongoRepoSession) SetMgoSession(mgosession *mgo.Session) {
	m.mgoSession = mgosession
}
