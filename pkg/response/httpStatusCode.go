package response

const (
	ErrCodeSuccess      = 20001
	ErrCodePramsInvalid = 20003

	ErrInvalidToken = 30001
)

// message

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodePramsInvalid: "Email invalid",
	ErrInvalidToken:     "Invalid Token",
}
