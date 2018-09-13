package main

import (
	"fmt"
	"os"
)

func main() {
    c, err := os.Executable() 
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(c)
}
