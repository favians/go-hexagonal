package user

import (
	"chat-hex/business"
	"chat-hex/util/primitiveIDGenerator"
	"chat-hex/util/validator"
	"time"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//FindUserByID Get user by given ID, return nil if not exist
func (s *service) FindUserByID(id string) (*User, error) {
	return s.repository.FindUserByID(id)
}

//FindAllUserWithPagination Get all users , will be return empty array if no data or error occured
func (s *service) FindAllUserWithPagination(skip int, rowPerPage int) ([]User, error) {

	user, err := s.repository.FindAllUserWithPagination(skip, rowPerPage)
	if err != nil {
		return []User{}, err
	}

	return user, err
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	id := primitiveIDGenerator.GenerateID()
	user := NewUser(
		id,
		insertUserSpec.Name,
		insertUserSpec.Username,
		insertUserSpec.Password,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser will update found user, if not exists will be return error
func (s *service) UpdateUser(id string, name string, modifiedBy string, currentVersion int) error {

	user, err := s.repository.FindUserByID(id)
	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	} else if user.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	modifiedUser := user.ModifyUser(name, time.Now(), modifiedBy)

	return s.repository.UpdateUser(modifiedUser, currentVersion)
}
