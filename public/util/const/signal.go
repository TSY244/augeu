package _const

const (
	MsgTypeSuccess = "success"
	MsgTypeError   = "error"
	MsgTypeWarning = "warning"
)

type MessageSignal struct {
	Message string
	MsgType string
}
