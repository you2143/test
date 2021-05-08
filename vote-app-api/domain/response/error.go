package response

type ResponseError struct {
	Message    string `json:"message"`
	StatisCode int    `json:"status_code"`
}
