package daos

import (
	"errors"
	"goapi/app"
	"goapi/models"

	"gopkg.in/mgo.v2/bson"
)

// UserDAO :
type UserDAO struct{}

// NewUserDao : create instance of userdao
func NewUserDao() *UserDAO {
	return &UserDAO{}
}

// Get reads the user with the specified email from the database.
func (dao *UserDAO) Get(rs app.RequestScope, email string) (*models.User, error) {
	var user models.User
	err := rs.DB().C("users").Find(bson.M{"email": email}).One(&user)
	return &user, err
}

// Create saves a new User record in the database.
// The User._id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(rs app.RequestScope, user *models.User) error {
	user.ID = bson.NewObjectId()
	return rs.DB().C("users").Insert(user)
}

// Update saves the changes to an user in the database.
func (dao *UserDAO) Update(rs app.RequestScope, email string, user *models.User) error {
	if _, err := dao.Get(rs, email); err != nil {
		return err
	}
	return rs.DB().C("users").Update(bson.M{"email": email}, bson.M{"$set": user})
}

// Delete deletes an user with the specified email from the database.
func (dao *UserDAO) Delete(rs app.RequestScope, email string) error {
	user, err := dao.Get(rs, email)
	if err != nil {
		return err
	}
	if user.Email == "" {
		return errors.New("User Not Found")
	}
	return rs.DB().C("users").Remove(bson.M{"email": email})
}

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count(rs app.RequestScope) (int, error) {
	return rs.DB().C("users").Count()
}
