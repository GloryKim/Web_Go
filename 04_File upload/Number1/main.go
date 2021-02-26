package main

import (
	"net/http"

	"github.com/GloryKim/Web_Go_Private/04_File upload/Number1/myapp"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir("public"))) //이 경로의 파일을 오픈
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
