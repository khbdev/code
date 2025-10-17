package main

import (
	"encoding/json"
	"fmt"
)


type  User struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func main(){
data := `{"name": "Azizbek", "age": 17}`
var u User

json.Unmarshal([]byte(data), &u)
fmt.Println("Name", u.Name)
fmt.Println("Age", u.Age)
}