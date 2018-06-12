package data

// RepoSession defines behavior to connect to repository database.
type RepoSession interface {

	// InitSession creates and returns a new repository Session.
	InitSession() interface{}

	// newSession returns a copy of the repository session
	newSession() interface{}

	// CloseSession closes the root repository session.
	CloseSession()

	// SetAddresses set the addresses of the intances of the repository.
	SetAddresses(addrs []string)

	// SetTimeout set time out
	SetTimeout(ttl int)

	// DBname returns the name of the database.
	DBname() string

	// SetDBname set database name for the repository.
	SetDBname(name string)

	// SetUserdb set database user name
	SetUserdb(usr string)

	// SetPwddb set database password
	SetPwddb(pwd string)
}
