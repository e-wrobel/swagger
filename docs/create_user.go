package docs

import "github.com/e-wrobel/swagger/internal/handlers"
import _ "github.com/pdrum/swagger-automation/api"

// swagger:route POST /users users idCreateUser
// Create new user with address.
// responses:
//   200: CreateUserResponse

// swagger:response CreateUserResponse
type CreateUserResponseWrapper struct {
	// in: body
	Body handlers.CreateUserResponse
}

// swagger:parameters CreateUserRequest idCreateUser
type CreateUserRequestWrapper struct {
	// in: header
	Token string
	// Create user request.
	// in: body
	Body handlers.CreateUserRequest
}
