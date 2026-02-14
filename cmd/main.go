package main

import "code/config"


func main(){
	db, err := config.ConfigPostgress(config.Postgress{
		Host: ,
	})
}