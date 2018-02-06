package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("test")
	x := []byte("{}")
	r := json.Valid(x)
	fmt.Println(r)
}
