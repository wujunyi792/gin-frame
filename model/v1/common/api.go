package common

type CommonResponseApi struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Count   int           `json:"count"`
	Data    []interface{} `json:"data"`
}

type CommonRequestApi struct {
	Type string        `json:"type"`
	Data []interface{} `json:"data"`
}
