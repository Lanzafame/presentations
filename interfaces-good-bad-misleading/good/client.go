package good

type Client struct {
	Save(w io.Writer) error
	Load(r io.Reader) ([]byte, error)
}


