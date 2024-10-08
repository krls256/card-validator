package http

type ValidationResponse struct {
	Valid bool             `json:"valid"`
	Error *ValidationError `json:"error,omitempty"`
}

type ValidationError struct {
	Code    int    `json:"int"`
	Message string `json:"message"`
}
