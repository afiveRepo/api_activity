package response

type BaseRespone struct {
	Status  string      `json:"status"`
	Message string      `json:"message`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}
