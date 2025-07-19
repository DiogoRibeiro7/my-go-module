package main

import (
	"flag"
	"fmt"

	"github.com/DiogoRibeiro7/my-go-module/greetings"
)

func main() {
	name := flag.String("name", "", "name to greet")
	flag.Parse()

	fmt.Println(greetings.Hello(*name))
}
