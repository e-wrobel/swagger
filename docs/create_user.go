package docs

import "github.com/e-wrobel/swagger/internal/handlers"
import _ "github.com/pdrum/swagger-automation/api"

// swagger:route POST /users users-tag idCreateUser
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
	// Create user request.
	// in: body
	Body handlers.CreateUserRequest
}
