package fix

type Client interface {
	SendMsg(b []byte) error
	GetMsg() ([]byte, error)
}
