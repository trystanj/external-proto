package main

import "log"
import "github.com/trystanj/external-proto/helper"

func main() {
	help := helper.Help()
	log.Printf("Main: helper says %v", help)
}
