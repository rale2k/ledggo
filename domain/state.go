package domain

type MessageType int32

const (
	NODE        MessageType = 0
	COORDINATOR MessageType = 1
)

type State struct {
	Role        MessageType
	Coordinator Node
	Nodes       []Node
}
