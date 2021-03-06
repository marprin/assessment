package constant

import "errors"

var (
	ErrTokenIsNotValid   = errors.New("Token is not valid")
	ErrClaimsIsNotValid  = errors.New("Claim is not valid")
	ErrIssuerIsNotValid  = errors.New("Issuer is not valid")
	ErrTokenIsExpired    = errors.New("Token is expired")
	ErrInternalServer    = errors.New("Internal server error")
	ErrStartDateNotValid = errors.New("Start date is not valid")
	ErrEndDateNotValid   = errors.New("End date is not valid")
)
