package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lanzafame/presentations/interfaces-good-bad-misleading/misleading"
)

func main() {
	args := os.Args[1:]
	client := misleading.NewFileClient()
	if len(args) < 1 {
		log.Fatal("please provide command")
	}
	switch args[0] {
	case "get":
		msg, err := client.GetMsg()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", msg)

	case "send":
		msg := []byte(strings.Join(args[1:], ""))
		err := client.SendMsg(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
