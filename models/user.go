package models

// User :
type User struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty" form:"firstName"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty" form:"lastName"`
	Email     string `json:"email,omitempty" bson:"email,omitempty" form:"email"`
	Password  string `json:"-" bson:"password"`
	Mobile    int64  `json:"mobile,omitempty" bson:"mobile,omitempty" form:"mobile"`
}

// Validate : validate function to validate the User docs
func (u User) Validate() error {
	// TODO: do validation here
	// Use ozzo-validation lib, or write custom validation here

	// INFO: return nil for now
	return nil
}
