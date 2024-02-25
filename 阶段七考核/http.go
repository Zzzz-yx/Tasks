package main

import(
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.FPrint(w,"hello world")
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":8000",nil)
}