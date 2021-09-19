package main

import (
	"fmt"

	"github.com/istherepie/monorepo-test/one/helpers"
)

func main() {
	hello := helpers.Hello()

	fmt.Println(hello)
}
