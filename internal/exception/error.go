package exception

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound        						= errors.New("not found")
	ErrUnauthorized    						= errors.New("unauthorized")
	ErrForbiddenAccess 						= errors.New("forbidden access")
	ErrBadRequest 	   						= errors.New("bad request")
	ErrInternalServer  						= errors.New("internal server error")
)

var (
	ErrTokenRequired 		 				= errors.New("token is required")
	ErrTokenInvalid 		 				= errors.New("token is invalid")
	ErrTokenExpired 		 				= errors.New("token is expired")
	ErrIsNotAdmin 		 	 				= errors.New("only admin can access this endpoint")

	ErrEmailRequired         				= errors.New("email is required")
	ErrEmailInvalid          				= errors.New("email is invalid")
	ErrPasswordRequired      				= errors.New("password is required")
	ErrPasswordInvalidLength 				= errors.New("password must have minimimum 6 character")
	ErrUsernameRequired      				= errors.New("username is required")
	ErrUsernameInvalidLength 				= errors.New("username must have minimimum 4 character")
	ErrAuthIsNotExists       				= errors.New("auth is not exists")
	ErrEmailAlreadyUsed      				= errors.New("email already used")
	ErrPasswordNotMatch      				= errors.New("password not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorInternalServer         	   		= NewError(ErrInternalServer.Error(), 		"500", 	   http.StatusInternalServerError)
	ErrorBadRequest     	   				= NewError(ErrBadRequest.Error(), 			"400", 	   http.StatusBadRequest)
	ErrorNotFound        	   				= NewError(ErrNotFound.Error(), 			"404", 	   http.StatusNotFound)
	ErrorUnauthorized    	   				= NewError(ErrUnauthorized.Error(), 		"401", 	   http.StatusUnauthorized)
	ErrorForbiddenAccess 	   				= NewError(ErrForbiddenAccess.Error(), 		"403", 	   http.StatusForbidden)
)

var (
	ErrorTokenRequired 		   				= NewError(ErrTokenRequired.Error(), 		"401-000", http.StatusUnauthorized)
	ErrorTokenInvalid 		   				= NewError(ErrTokenInvalid.Error(), 		"401-001", http.StatusUnauthorized)
	ErrorTokenExpired 		   				= NewError(ErrTokenExpired.Error(), 		"401-002", http.StatusUnauthorized)
	ErrorIsNotAdmin 		   				= NewError(ErrIsNotAdmin.Error(), 			"403-000", http.StatusForbidden)
)

var (
	ErrorEmailRequired         				= NewError(ErrEmailRequired.Error(), 		"400-001", http.StatusBadRequest)
	ErrorEmailInvalid          				= NewError(ErrEmailInvalid.Error(), 		"400-002", http.StatusBadRequest)
	ErrorPasswordRequired      				= NewError(ErrPasswordRequired.Error(), 	"400-003", http.StatusBadRequest)
	ErrorPasswordInvalidLength 				= NewError(ErrPasswordInvalidLength.Error(),"400-004", http.StatusBadRequest)
	ErrorUsernameRequired      				= NewError(ErrUsernameRequired.Error(), 	"400-005", http.StatusBadRequest)
	ErrorUsernameInvalidLength 				= NewError(ErrUsernameInvalidLength.Error(),"400-006", http.StatusBadRequest)

	ErrorAuthIsNotExists  	   				= NewError(ErrAuthIsNotExists.Error(), 		"404-001", http.StatusNotFound)
	ErrorEmailAlreadyUsed 	   				= NewError(ErrEmailAlreadyUsed.Error(), 	"409-001", http.StatusConflict)
	ErrorPasswordNotMatch 	   				= NewError(ErrPasswordNotMatch.Error(), 	"401-003", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              				ErrorNotFound,
		ErrUnauthorized.Error():          				ErrorUnauthorized,
		ErrForbiddenAccess.Error():       				ErrorForbiddenAccess,
		ErrBadRequest.Error():       	  				ErrorForbiddenAccess,
		ErrInternalServer.Error():        				ErrorInternalServer,

		ErrTokenRequired.Error():         				ErrorTokenRequired,
		ErrTokenInvalid.Error():          				ErrorTokenInvalid,
		ErrTokenExpired.Error():          				ErrorTokenExpired,
		ErrIsNotAdmin.Error():            				ErrorIsNotAdmin,

		ErrEmailRequired.Error():         				ErrorEmailRequired,
		ErrEmailInvalid.Error():          				ErrorEmailInvalid,
		ErrUsernameRequired.Error():      				ErrorUsernameRequired,
		ErrUsernameInvalidLength.Error(): 				ErrorUsernameInvalidLength,
		ErrPasswordRequired.Error():      				ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): 				ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error():       				ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error():      				ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error():      				ErrorPasswordNotMatch,
	}
)