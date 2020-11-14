package models

type ApiResponse struct {
	Code    int    `json:"status_code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
