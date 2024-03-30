package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	config "github.com/danish-mehmood/simple-layer-4-load-balancer-in-golang/configuration"
)
func Index(writer http.ResponseWriter , request  *http.Request){
	target , err :=url.Parse("http://localhost:8081")
	fmt.Println("target and err" , target , err)
    proxy :=httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(writer , request)
   fmt.Fprintf(writer , "this is the index route , url is = %s , host is = %s" ,request.URL  , request.RemoteAddr)
}

func ChangeConfig(c *config.Configuration , key string  , value string ){
    // bytes , err := json.Marshal(c)
	
		fmt.Println("there is a field yeessss" , c)
		switch key{
		case "address": c.Address =value;
		case "algorithm" : c.Algorithm=value;
		// case "servers" : c.Servers=value;
		case "healthCheckInterval":c.HealthCheckInterval=value;
		default: fmt.Println("the key is wrong")
		}
		fmt.Println("!!!!!!!" , c)
		bytes , err := json.Marshal(c)
		if err == nil{
		
		fd , err := os.OpenFile("config.json" , os.O_WRONLY|os.O_CREATE|os.O_TRUNC , 0644)
		
		if err!=nil{
			fmt.Println("there is an error" , err)
		}else{
			defer fd.Close()
			_,err:=fd.Write(bytes)
			if err != nil {
				fmt.Println("error while writing file" , err)
			}


		}

	}
	

}


func hasField(bytes []byte, fieldName string) bool {
	var mapper map[string]any
    if err := json.Unmarshal(bytes , &mapper); err!=nil{
		log.Print(err)
		return false
	}

	if _  , ok:=mapper[fieldName];ok{

        return true
	}

    return false
}


func capitalize(word string) string {
	if word == "" {
		return ""
	}

	return strings.ToUpper(string(word[0])) + word[1:]
}


