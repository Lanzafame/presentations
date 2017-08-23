package misleading

type Client interface {
	SendMsg(b []byte) error
	GetMsg() ([]byte, error)
}
