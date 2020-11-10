package model

type Message struct {
	Msg string
}

var messages []Message

func AddMessage(msg Message)  {
	messages = append(messages, msg)
}

func GetMessage() Message {
	return messages[len(messages)-1]
}
