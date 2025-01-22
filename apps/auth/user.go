package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shelllbyyyyy/belajar-api-go/internal/exception"
	"github.com/shelllbyyyyy/belajar-api-go/util"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       	string 		
	Username 	string		
	Email    	string		
	Password 	string		
	CreatedAt 	time.Time 	
	UpdatedAt 	time.Time 	
}

func newUser(r registerUserSchema) (*User, error) {
	if r.Email == "" || r.Password == "" || r.Username == "" {
		return nil, errors.New("input cannot be null")
	}

	return &User{
		Id: uuid.NewString(), 
		Email: r.Email, 
		Username: r.Username, 
		Password: r.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		}, nil
}

func (u *User) validate() (err error) {
	if err = u.validateEmail(); err != nil {
		return
	}

	if err = u.validatePassword(); err != nil {
		return
	}

	if err = u.validateUsername(); err != nil {
		return
	}

	return
}

func (u *User) validateEmail() (err error) {
	if u.Email == "" {
		return exception.ErrEmailRequired
	}

	emails := strings.Split(u.Email, "@")
	if len(emails) != 2 {
		return exception.ErrEmailInvalid
	}
	return
}

func (u *User) validatePassword() (err error) {
	if u.Password == "" {
		return exception.ErrPasswordRequired
	}

	if len(u.Password) < 6 {
		return exception.ErrPasswordInvalidLength
	}
	return
}

func (u *User) validateUsername() (err error) {
	if u.Username == "" {
		return exception.ErrUsernameRequired
	}

	if len(u.Username) < 6 {
		return exception.ErrUsernameInvalidLength
	}
	return
}

func (u *User) encryptPassword(salt int) (err error) {

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		return
	}
	u.Password = string(encryptedPass)
	return nil
}

func (u *User) comparePassword(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u *User) generateToken(exp float64) (tokenString string, err error) {
	return util.GenerateToken(u.Id, exp)
}