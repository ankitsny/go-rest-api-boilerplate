package services

import (
	"errors"
	"goapi/app"
	"goapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// Data store
func newMockUserDAO() userDAO {
	return &mockUserDAO{
		records: []models.User{
			{ID: "5a947f3a14032d3b384b0829", FirstName: "aaa", Email: "anks@anks.com"},
			{ID: "5a947f3a14032d3b384b0829", FirstName: "bbb", Email: "anso@ankso.com"},
			{ID: "5a947f3a14032d3b384b0829", FirstName: "ccc", Email: "yeah@yess.com"},
		},
	}
}

type mockUserDAO struct {
	records []models.User
}

// Implement all userDao methods

// Fake implementation goes here

// Implement Get Method of userDao interface
func (m *mockUserDAO) Get(rs app.RequestScope, email string) (*models.User, error) {
	for _, record := range m.records {
		if record.Email == email {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

// Implement Count func
func (m *mockUserDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

// Implement Create func
func (m *mockUserDAO) Create(rs app.RequestScope, user *models.User) error {
	if user.Email == "" {
		return errors.New("Email cannot be empty")
	}
	user.ID = bson.NewObjectId()
	m.records = append(m.records, *user)
	return nil
}

// Implement Update method
func (m *mockUserDAO) Update(rs app.RequestScope, email string, user *models.User) error {
	for i, record := range m.records {
		if record.Email == email {
			m.records[i] = *user
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockUserDAO) Delete(rs app.RequestScope, email string) error {
	for i, record := range m.records {
		if record.Email == email {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

// Test Code

func TestNewUserService(t *testing.T) {
	dao := newMockUserDAO()
	s := NewUserService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestNewUserService_Get(t *testing.T) {
	s := NewUserService(newMockUserDAO())

	// Valid User
	user, err := s.Get(nil, "anks@anks.com")
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, user.FirstName, "aaa")
	}

	user, err = s.Get(nil, "anks1@anks.com")
	assert.NotNil(t, err)
}

func TestUserService_Create(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user := models.User{
		Email:     "last@last.com",
		FirstName: "ddd",
		LastName:  "last",
	}
	err := s.Create(nil, &models.User{
		Email:     "last@last.com",
		FirstName: "ddd",
		LastName:  "last",
	})
	if assert.Nil(t, err) {
		assert.Equal(t, user.Email, "last@last.com")
		assert.Equal(t, "ddd", user.FirstName)
	}

	// validation error
	err = s.Create(nil, &models.User{
		Email: "",
	})
	assert.NotNil(t, err)
}

func TestUserService_Update(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	u, err := s.Update(nil, "anks@anks.com", &models.User{FirstName: "Hell", Email: "anks@anks.com"})
	assert.Nil(t, err)
	assert.NotNil(t, u)

	u1, err1 := s.Update(nil, "invalid_email@anks.com", &models.User{FirstName: "Hell", Email: "anks@anks.com"})
	assert.Nil(t, u1)
	assert.Equal(t, err1.Error(), "not found")

	u1, err1 = s.Update(nil, "anks@anks.com", &models.User{FirstName: "Hell", Email: ""})
	assert.Nil(t, u1)
	assert.Equal(t, err1.Error(), "Email can't be empty")

}

func TestNewUserService_Delete(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	err := s.Delete(nil, "anks@anks.com")
	assert.Nil(t, err)

	err = s.Delete(nil, "anks1111@anks.com")
	assert.NotNil(t, err)
}

func TestNewUserService_Count(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	count, err := s.Count(nil)
	assert.NotNil(t, count)
	assert.Nil(t, err)
}
