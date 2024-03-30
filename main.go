package main

import (
	"fmt"
	"net/http"

	configuration "github.com/danish-mehmood/simple-layer-4-load-balancer-in-golang/configuration"
	server "github.com/danish-mehmood/simple-layer-4-load-balancer-in-golang/server"
)


func main (){
	c:=  configuration.Configuration{}
	c.CreateNewConfig("config.json")
	// server.ChangeConfig(&c , "address" , ":7890")

	mux:= http.NewServeMux()
	mux.HandleFunc("/changeconfig" ,func (writer http.ResponseWriter , request *http.Request){
    					queryparams := request.URL.Query()
 						key:=queryparams.Get("key")
    					value:=queryparams.Get("value")
    					server.ChangeConfig(&c , key ,value )
    					c.CreateNewConfig("config.json")
						http.Redirect(writer , request , "http://localhost:7890/changeSuccessful" , http.StatusFound)
  						})

						mux.HandleFunc("/changeSuccessful" ,func (writer http.ResponseWriter , request *http.Request){
		                     fmt.Fprintf(writer , "success the config has changed !! %v" ,c)
	})

					
	mux.HandleFunc("/" , server.Index)
	node := http.Server{
		Addr: c.Address,
		Handler: mux,
	}
	node.ListenAndServe()

}