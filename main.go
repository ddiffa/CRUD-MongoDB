package main

import (
	"fmt"
	"mongoDB/config"
)

func main(){
	_, err:=config.GetMongoDB()

	if err!=nil{
		fmt.Println(err)
	}
}