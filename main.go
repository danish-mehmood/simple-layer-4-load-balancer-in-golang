package main

import (
	"fmt"

	configuration "github.com/danish-mehmood/simple-layer-4-load-balancer-in-golang/configuration"
)
func main (){
	c:=  configuration.Configuration{}
	c.CreateNewConfig("config.json")
	fmt.Println("done" , c)
}