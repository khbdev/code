package main

import "fmt"


func sayHi(n int){
for i := 0; i < 10; i++ {
		if n == 0 {
		return
	}
	fmt.Println("Hello 1")
	sayHi(n -1)
}
}

func main(){

for i := 0; i < 10; i++ {
	sayHi(3)
}
}