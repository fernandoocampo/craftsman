package data

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// mongo addresses
var mgoaddrs []string

// mgo conn in seconds
var timeout int

// mongo database name
var dbname string

// mongo db user name
var usrdb string

// mongo database password
var pwddb string

// mongo session
var mgoSession *mgo.Session

// InitMgoSession creates and returns a new Mongo Session.
func InitMgoSession() *mgo.Session {
	mgodialinfo := &mgo.DialInfo{
		Addrs:    mgoaddrs,
		Timeout:  time.Duration(timeout) * time.Second,
		Database: dbname,
		Username: usrdb,
		Password: pwddb,
	}

	fmt.Println("Connecting to db...")
	session, err := mgo.DialWithInfo(mgodialinfo)

	if err != nil {
		panic(err)
	}
	session.SetSafe(&mgo.Safe{})
	session.SetSyncTimeout(3 * time.Second)
	session.SetSocketTimeout(3 * time.Second)

	mgoSession = session

	return session
}

// CloseMgoSession closes the root mongo session.
func CloseMgoSession() {
	mgoSession.Close()
}

// newMgoSession returns a copy of the mongo session
func newMgoSession() *mgo.Session {
	return mgoSession.Copy()
}

// SetMongoAddrs set *mgoaddrs value
func SetMongoAddrs(addrs []string) {
	mgoaddrs = addrs
}

// SetTimeout set time out
func SetTimeout(ttl int) {
	timeout = ttl
}

// SetDBname set mongo database name
func SetDBname(name string) {
	dbname = name
}

// SetUserdb set mongo db user name
func SetUserdb(usr string) {
	usrdb = usr
}

// SetPwddb set mongo database password
func SetPwddb(pwd string) {
	pwddb = pwd
}
