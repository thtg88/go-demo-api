package types

// ErrorJSONResponse represents a JSON error response
type ErrorJSONResponse struct {
	Errors  map[string][]string
	Message string
}
