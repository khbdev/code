package main

import (
	"fmt"
	"sync"

	"time"
)
 var mu sync.Mutex
 var wy sync.WaitGroup

func main() {
 firstTest()
 // secondTest()
 
}

// ---------------------------------------------------------------------------------------

// order count 1_000_000 marta qo'shib chiqqanda boshqa son chiqmoqda.
// shuni to'gri increment qiladigan qilish kerak.
// berilgan vaqt 5-10 daqiqa
func firstTest() {
 var orderCount int

 for i := 0; i < 1_000_000; i++ {
	wy.Add(1)
  go func() {
	mu.Lock()
   orderCount++
   mu.Unlock()
   wy.Done()
  }()
 }
  wy.Wait()

 fmt.Printf("order count: %d\n", orderCount)

}

// ---------------------------------------------------------------------------------------

// agar request 3 soniadan koproq vaqt oladigan bo'lsa, biz kutmay, panic qilishimiz kerak
// berilgan vaqt 5-10 daqiqa
func secondTest() {
 response := simulateRequest()
 fmt.Printf("received response: %d\n", response)
}

func simulateRequest() int {
 time.Sleep(time.Second * 5)
 return 1
}