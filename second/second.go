package second

import (
	"log"

	protos "github.com/trystanj/proto"
)

func Help() string {
	demo := &protos.Sample{}
	log.Printf("demo.Info from helper: %v", demo.Info)

	return demo.Info
}
