package services

import (
	"goapi/app"
	"goapi/models"
)

// userDAO specifies the interface of the user DAO needed by UserService.
type userDAO interface {
	// Get returns the user with the specified user email.
	Get(rs app.RequestScope, email string) (*models.User, error)
	// Count returns the number of users.
	Count(rs app.RequestScope) (int, error)
	// Create saves a new user in the storage.
	Create(rs app.RequestScope, user *models.User) error
	// Update updates the user with given email in the storage.
	Update(rs app.RequestScope, email string, user *models.User) error
	// Delete removes the user with given email from the storage.
	Delete(rs app.RequestScope, email string) error
}

// UserService provides services related with users.
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get returns the user with the specified the user email.
func (s *UserService) Get(rs app.RequestScope, email string) (*models.User, error) {
	return s.dao.Get(rs, email)
}

// Create creates a new user.
func (s *UserService) Create(rs app.RequestScope, model *models.User) (*models.User, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Email)
}

// Update updates the user with the specified email.
func (s *UserService) Update(rs app.RequestScope, email string, model *models.User) (*models.User, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, email, model); err != nil {
		return nil, err
	}

	return s.dao.Get(rs, email)
}

// Delete deletes the user with the specified email.
func (s *UserService) Delete(rs app.RequestScope, email string) (*models.User, error) {
	user, err := s.dao.Get(rs, email)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, email)
	return user, err
}

// Count returns the number of users.
func (s *UserService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}
