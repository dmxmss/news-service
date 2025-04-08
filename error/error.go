package error

type Error int

const (
	ErrDbTransactionFailed Error = iota
	ErrDbInitError
	ErrDbNewsNotFound
	ErrInvalidRequestData
	ErrNotAuthorized
	ErrTokenClaimsInvalid
	ErrUserIsNotAuthor
)

func (e Error) Error() string {
	var err string
	switch e {
	case ErrDbTransactionFailed:
		err = "Database: transaction failed"
	case ErrDbInitError:
		err = "Database: init error"
	case ErrDbNewsNotFound:
		err = "Database: news not found"
	case ErrInvalidRequestData:
		err = "Invalid request data"
	case ErrNotAuthorized:
		err = "Not authorized"
	case ErrTokenClaimsInvalid:
		err = "Token claims are invalid"
	case ErrUserIsNotAuthor:
		err = "User is not an author"
	}

	return err
}
