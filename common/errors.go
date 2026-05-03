package common

import "fmt"

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
}

//// IsAPIError check if e is an API error
//func IsAPIError(e error) bool {
//	_, ok := e.(*APIError)
//	return ok
//}
