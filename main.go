package main

import (
 "fmt"

 "time"
)

func main() {
 //firstTest()
  secondTest()
}

// ---------------------------------------------------------------------------------------

// order count 1_000_000 marta qo'shib chiqqanda boshqa son chiqmoqda.
// shuni to'gri increment qiladigan qilish kerak.
// berilgan vaqt 5-10 daqiqa
func firstTest() {
 var orderCount int
 for i := 0; i < 1_000_000; i++ {
  go func() {
   orderCount++
  }()
 }

 fmt.Printf("order count: %d\n", orderCount)
}

// ---------------------------------------------------------------------------------------

// agar request 3 soniadan koproq vaqt oladigan bo'lsa, biz kutmay, panic qilishimiz kerak
// berilgan vaqt 5-10 daqiqa
func secondTest() {
 ctx, cancel := context.WithTimeOut(context.Backround(), 5 *time.Second)
 defer cancel()
 
 result := make(chan int)
 


 go func() {
	 result <- simulateRequest()
 }()
 
 select {
 case res := <-result:
 fmt.Println("malumot: ", res)
 case <-ctx.Done():
 fmt.Println("vaqt yetmadi")
 }
 
}

func simulateRequest() int {
 time.Sleep(time.Second * 5)
 return 1
}