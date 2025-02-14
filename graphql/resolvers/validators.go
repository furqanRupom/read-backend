package graphqlresolvers

import (
   "read-backend/graphql"
	"net/mail"
)


func validateEmail(email string) *graphql.ErrorInvalidEmail {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return &graphql.ErrorInvalidEmail{}
	}
	if len(email) > 320 {
		return &graphql.ErrorInvalidEmail{}
	}
	return nil
}

func isPasswordValid(password string) bool {
	return len(password) > 10
}

func validateUserIn(userIn graphql.UserIn) graphql.CreateUserError {
	emailErr := validateEmail(userIn.Email)
	if emailErr != nil {
		return *emailErr
	}
	if len(userIn.Name) > 100 {
		return graphql.ErrorInvalidUserName{}
	}
	return nil
}
