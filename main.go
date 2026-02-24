package main

import (
	"fmt"
	"os/exec"
)

func main(){
	cmd := exec.Command("bash", "-ls", "ls", "la")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(out))
}