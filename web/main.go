package main

import (
	"net/http"

	"web/myapp"
)

func main() {

	//리스닝하는 포트 지정.
	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}
