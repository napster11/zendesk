package zendeskUtil

// MetaData of HTTP API response
type MetaData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

// ConstructResponse - Creates HTTP Response in standard JSON format.
func ConstructResponse(statusCode int, msg string, data interface{}) Response {
	res := Response{}
	res.Meta.Code = statusCode
	res.Meta.Msg = msg
	res.Data = data
	return res
}