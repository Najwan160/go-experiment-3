package base

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RespErrValidator struct {
	Message string            `json:"message"`
	ErrData map[string]string `json:"error_data"`
	Data    interface{}       `json:"data"`
}

type RespErr struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// for variant 2xx
type RespErrWithData struct {
	Message string      `json:"message"`
	ErrData interface{} `json:"error_data"`
	Data    interface{} `json:"data"`
}

type RespErrSendMessage struct {
	Message string      `json:"message"`
	ErrData string      `json:"error_data"`
	Data    interface{} `json:"data"`
}
