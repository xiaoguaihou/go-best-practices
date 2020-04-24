package core

// the api implement by this project

type Status int

const (
	StatusOK = iota
	StatusFailure
	StatusInvalidParam
	StatusEmptyQuery
	StatusUserNotExist
	StatusInterError
)

func (s Status) Message() string {
	return [...]string{"ok",
		"request failure",
		"invalid parameter",
		"empty query condition",
		"user not exist",
		"internal error"}[s]
}

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Status     Status    `json:"status"`
	Message    string `json:"message"`
	OpenID     string `json:"openId"`
	SessionKey string `json:"sessionKey"`
	UnionID    string `json:"unionId"`
	Mobile     string `json:"mobile"`
}