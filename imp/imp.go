package main

import (
	"fmt"
	"github.com/distributed/exportunexported/exp"
)

func main() {
	i, err := exp.GetLeStuff()
	fmt.Println(i, err)
}
