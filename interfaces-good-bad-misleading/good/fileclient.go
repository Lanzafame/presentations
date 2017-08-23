package fix

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewFileClient() FileClient {
	return FileClient{}
}

type FileClient struct {
}

// errors ignored to fit on the slide
func (f FileClient) SendMsg(b []byte) error {
	dir := os.TempDir()
	fmt.Println(dir)
	msgfile, _ := os.Create(filepath.Join(dir, "msg"))
	msgfile.Write(b)
	return nil
}

func (f FileClient) GetMsg() ([]byte, error) {
	dir := os.TempDir()
	fmt.Println(dir)
	msgfile, _ := os.Open(filepath.Join(dir, "msg"))
	var msg []byte
	msgfile.Read(msg)
	return msg, nil
}
