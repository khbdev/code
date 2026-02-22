package main

import (
	"fmt"
	"strings"
)

func main(){
var u string
fmt.Print("URl: ")
fmt.Scanln(&u)

if !strings.HasPrefix(u, "http") {
	u = "https://" + u
}

}