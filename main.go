package main

import "log"

// import "github.com/trystanj/external-proto/helper"
import "github.com/trystanj/external-proto/second"

func main() {
	second := second.Help()
	log.Printf("Main: second says %v", second)

	// help := helper.Help()
	// log.Printf("Main: helper says %v", help)
}
