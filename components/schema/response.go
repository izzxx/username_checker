package schema

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Url        string `json:"url"`
}
