package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)
type Configuration struct{
	HealthCheckInterval string   `json:"healthCheckInterval"`
	Servers             []string `json:"servers"`
	Address          string   `json:"address"`
	Algorithm           string   `json:"algorithm"`
}

func (c *Configuration) CreateNewConfig(file string)(error){
	
    bytes , err := ioutil.ReadFile(file)
	fmt.Println("the file bytes" , string(bytes) , )
	if err !=nil{
		fmt.Println(err)
		return  err
	}
	
    if err = json.Unmarshal(bytes , c);err!=nil{
		fmt.Println(err)
		return  err
	}

   
    return nil
	}
