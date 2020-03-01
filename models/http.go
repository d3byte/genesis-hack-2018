package models

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPSuccess example
type HTTPSuccess struct {
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data" example:"object with information"`
}
