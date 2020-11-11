package models

// Message ...
type Message struct {
	Code    int
	Message string
}

// MessageWithData ...
type MessageWithData struct {
	Code    int
	Message string
	Data    interface{}
}
