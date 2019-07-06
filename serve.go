package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

func serveListDns(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	d, err := listDns("dns.toml") // todo
	if err != nil{
		http.Error(w, "InternalServerError", 500)
		return
	}

	b, err := json.Marshal(d)
	if err != nil{
		http.Error(w, "InternalServerError", 500)
		return
	}
	fmt.Fprint(w, string(b))
}

func serveInterface(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	i, err := getInterfaces()
	if err != nil{
		http.Error(w, "InternalServerError", 500)
		return
	}

	b, err := json.Marshal(i)
	fmt.Fprint(w, string(b))
}

func serveEnableDns(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// validation

	if err := enableDns("Wi-Fi","8.8.8.8"); err != nil{
		http.Error(w, "InternalServerError", 500)
		return
	}
}

func serveDisableDns(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if err := disableDns("Wi-Fi"); err != nil{
		http.Error(w, "InternalServerError", 500)
		return
	}
}

func CmdServe(c *cli.Context){
	router := httprouter.New()
	router.GET("/interface", serveInterface)
	router.GET("/dnsList", serveListDns)
	router.POST("/enable", serveEnableDns)
	router.POST("/disable", serveDisableDns)
	log.Fatal(http.ListenAndServe(":32819", router))
}
