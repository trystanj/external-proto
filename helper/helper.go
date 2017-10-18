package helper

import (
	"log"

	protos "github.com/trystanj/proto/group"
)

func Help() string {
	demo := &protos.Outer{Counter: 1, Info: "hello"}
	log.Printf("demo.Info from helper: %v", demo.Info)

	return demo.Info
}
