# 02_Web Handler

```go
package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 핸들러로 등록하는 방법 (어떤 일을 할 것인지) 여기시 '/' 표시는 경로를 의미한다. Index Page
		fmt.Fprint(w, "Hello World") 
	})
	/*
	w http.ResponseWriter 이쪽은 작성을 하는 부분 즉 화면에 써주는 부분이다.
	r *http.Request 이쪽은 무엇을 작성할지 읽는 부분이다.
	*/

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bar!")
	})

	http.Handle("/foo", &fooHandler{}) //fooHandler 라는 인스턴스 형태로 저장할때에는 이렇게 진행한다.

	http.ListenAndServe(":3000", nil)
}
```