package common

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

const (
	LoginMessageType 				= "LoginMessageType"
	RegisterMessageType				= "RegisterMessageType"
	LoginResponseMessageType		= "LoginResponseMessageType"
	RegisterResponseMessageType		= "ResponseMessageType"
)