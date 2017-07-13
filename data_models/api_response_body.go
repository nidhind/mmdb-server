package data_models

// Models for Errors:
// * Invalid input
type APIError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewAPIError(name string, cause string) APIError {
	switch name {
	case "INVALID_INPUT":
		return APIError{
			Code: 2000,
			Message:"Invalid Input",
			Cause: cause}
	default:
		return APIError{
			Code: 500,
			Message:"Internal Server Error",
			Cause: "Unknown"}
	}
}