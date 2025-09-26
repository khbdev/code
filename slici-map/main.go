package main

import "fmt"

func main(){
	numbers := []int{1,2,2,3,3,3,3,3,3,3,4,4,4,4,}

	counMap := make(map[int]int)

	for _, v := range numbers {
		counMap[v]++
	}
	for s, v := range counMap {
		fmt.Printf("%d -> %d marta\n", s, v)
	}
}