package domain

type Message struct {
	MsgType string `json:"msg_type"`
	Message string `json:"message"`
}
