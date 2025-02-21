// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type CreateUserError interface {
	IsCreateUserError()
}

type Author struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}

type ErrorInvalidEmail struct {
	Message string `json:"message"`
}

func (ErrorInvalidEmail) IsCreateUserError() {}

type ErrorInvalidPassword struct {
	Message string `json:"message"`
}

type ErrorInvalidUserName struct {
	Message string `json:"message"`
}

func (ErrorInvalidUserName) IsCreateUserError() {}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterResponse struct {
	ID   string `json:"Id"`
	Name string `json:"name"`
}

type UserIn struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ConfirmAuthor struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

type NewAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
