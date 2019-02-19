package basic

import (
	"fmt"
	"net/http"
)

func StartHttpServer(){


	http.HandleFunc("/",IndexHandler)
	http.ListenAndServe(":8080",nil)
	
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintln(w,"hello world")
}
