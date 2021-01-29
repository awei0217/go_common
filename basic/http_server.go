package basic

import (
	"fmt"
	"net/http"
)

//SET CGO_ENABLED=0
//SET GOOS=linux
//SET GOARCH=amd64
//go build cron_delete_compute_result_uat.go
func StartHttpServer() {

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello world:"+r.Form.Get("key"))
}
