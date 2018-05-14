package data

// Skill contain data for craftsman's skills
type Skill struct {
	Name  string `json:"name" bson:"name"`
	Stats Stats  `json:"stats,omitempty" bson:"stats,omitempty"` // statistics of skill
}

// Stats contains the data of different statistics for skills.
type Stats struct {
	Feedback  int `json:"feedback" bson:"feedback"` // How many users has given a feedback
	Validated int `json:"orders" bson:"orders"`     // How many users has given a credit
	Rating    int `json:"rating" bson:"rating"`     // The score average given by other users
}

// Craftsman defines state
type Craftsman struct {
	Username  string  `json:"username" bson:"username"` // The name of the user
	Firstname string  `json:"firstname" bson:"firstname"`
	Lastname  string  `json:"lastname" bson:"lastname"`
	Email     string  `json:"email" bson:"email"`
	Email2    string  `json:"email2" bson:"email2"`                     // Second email
	State     int8    `json:"state" bson:"cstate"`                      // The craftsman current state
	Kwds      string  `json:"kwds" bson:"kwds"`                         // keywords for searching
	Skills    []Skill `json:"skills,omitempty" bson:"skills,omitempty"` // Craftsman's skills
}
