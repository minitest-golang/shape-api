package utils

import "errors"

const (
	OK                    = "OK"
	UNKNOWN_ERROR         = "Unknown Error"
	BAD_REQUEST           = "Bad Request"
	BAD_JSON              = "Bad Json"
	OBJECT_EXISTED_ERROR  = "Object Existed"
	INTERNAL_SERVER_ERROR = "Internal Server Error"
	UNAUTHORIZED          = "Unauthorized: Please use login API to get access-token, then set 'Bearer <token>' to Authorization in the request header for other requests."
	SIGNUP_ERROR          = "Signup Failed"
	PASSWORD_TOO_SHORT    = "Password is too short!"
	USER_NAME_ERROR       = "Username must contains only lowercases uppercases, numbers, underscores and dashes!"
)

var (
	ErrBadInput         = errors.New("Bad Input!")
	ErrAtInternalServer = errors.New("Internal Server Error!")
	ErrDBUserExisted    = errors.New("sql: Username existed!")
	ErrDBObjectNotFound = errors.New("sql: Object not found!")
	ErrBadShape         = errors.New("shape: Invalid shape!")
	ErrBadEdge          = errors.New("shape: Invalid edge value!")
	ErrBadShapeId       = errors.New("shape: Bad shape ID!")
)
