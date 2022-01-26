package users

import (
	"fmt"
	"foodHelper/data"
	"foodHelper/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/xsrftoken"
)

type UserCalls struct {
	store data.Store
}

type CreateUserRequest struct {
	username string
	userRole UserRole

	password string
}

func (c *UserCalls) CreateUser(request CreateUserRequest) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.password), 14)
	if err != nil {
		return User{}, err
	}

	return User{
		Id:                  "",
		Username:            "",
		EncodedUserPassword: hash,
		UserRole:            0,
		ApiKeys:             nil,
	}, nil
}

type UpdateUserRequest struct {
	userId   string
	username string
	userRole string

	password string
}

func (c *UserCalls) UpdateUser(request UpdateUserRequest) (*User, error) {
	return &User{}, nil
}

func (c *UserCalls) DeleteUser(userId string) error {
	return nil
}

type SessionToken string

func (c *UserCalls) LoginPassword(username string, password string) (SessionToken, error) {

	user, err := c.store.GetUserByName(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("user does not exist")
	}

	if !user.IsEnabled || user.IsDeleted {
		return "", fmt.Errorf("user not activ")
	}

	err = bcrypt.CompareHashAndPassword(user.EncodedUserPassword, []byte(password))
	if err != nil {
		return "", err
	}

	sessionToken := utils.RandStringBytes(64)
	
	xsrfToken := xsrftoken.Generate(utils.RandStringBytes(64), user.Id, "")

	session := data.Session{
		SessionId:        uuid.New().String(),
		SessionTokenHash: "",
		XsrfTokenHash:    "",
		UserId:           "",
	}

	return "", nil
}

func (c *UserCalls) LoginApiKey(username string, apiKey string) (SessionToken, error) {
	return "", nil
}

func createSession(username string) (SessionToken, error) {
	return "", nil
}
