package main

import (
	"flag"
	"fmt"

	"github.com/uudashr/goecho"
)

func main() {
	flagMsg := flag.String("m", "Hello", "Message")
	flag.Parse()

	fmt.Println(goecho.Echo(*flagMsg))
}
