package rest


// AuthService interface is used to convert a Bearer token
// into a valid user id.
// In an OAuth2 implementation would use the token to query 
// the OAuth2 service to obtain a user ID.
type AuthService interface {
	// authorizeUser returns a user id from an authToken
	AuthorizeUser(authToken string) (string, error)
}

// AllowAllAuthService is a dumb type to allow any Bearer token
// behave directly as user id
type AllowAllAuthService struct {
}

func (aas *AllowAllAuthService) AuthorizeUser(authToken string) (string, error) {
	return authToken, nil
}
