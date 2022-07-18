package auth

import "fmt"

var (
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = fmt.Errorf("secret key is required")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = fmt.Errorf("ginJWTMiddleware.Authenticator func is undefined")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = fmt.Errorf("failed to create JWT Token")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = fmt.Errorf("token is expired") // in practice, this is generated from the jwt library not by us

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = fmt.Errorf("auth header is empty")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = fmt.Errorf("missing exp field")

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = fmt.Errorf("exp must be float64 format")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = fmt.Errorf("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = fmt.Errorf("query token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = fmt.Errorf("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512
	ErrInvalidSigningAlgorithm = fmt.Errorf("invalid signing algorithm")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = fmt.Errorf("incorrect Username or Password")

	// ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrMissingLoginValues = fmt.Errorf("missing Username or Password")
)
