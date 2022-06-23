package main
import (
    "fmt"
    "net/http"
func main(){
    http.HandleFunc("/ping",ping)
    http.ListenAndServe("8080",nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "pong\n")
}