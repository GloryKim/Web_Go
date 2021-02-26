# 03_JSON Transfer

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {//JSON 예시
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Email     string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`//이렇게 ``안에 내용을 써서 하면 json에서는 저렇게 쓴다는걸 의미한다. 추가로 디코더와 마샬링할때 자동으로 변환해준다.
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)//JSON 형태로 파싱을 해야한다. 데이터는 리퀘스트 바디 안에 있다. 안의 자료를 읽어서 디코딩 하겠다는 뜻
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)// 자료를 받아서 바이트 배열 형대로 받는다.
	w.Header().Add("content-type", "application/json")// 여기서 이걸 안써주면 text로 인식해서 모양이 안이쁘게 나온다.
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
	
	//27~33줄까지 지우고 주소창에 http://localhost:3002/foo 작성하면 Bad Request: EOF가 나올것이다.
	//이유는 데이터가 없기 때문이다.

	//가장 바람직한 방법은 구글크롬 확장 프로그램에서 ARC 프로그램을 설치해서 바디에 직접 값을 넣어주는 방법이다.
	//참고자료 https://www.youtube.com/watch?v=vOW0j6hd-Rg&list=PLy-g2fnSzUTDALoERcKDniql16SAaQYHF&index=2
	//12분 부터
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") //name이라는 인스턴스를 만들어서
	if name == "" {//name이 없으면
		name = "World" //name은 World이다.
	}
	fmt.Fprintf(w, "Hello %s!", name)
}//컴파일할 때에 주소창에 'localhost:3002/bar?name=glory' 라고 검색하면 Hello glory가 나온다.

func main() {
	mux := http.NewServeMux() //mux라는 인스턴스를 만들어서 사용할때에는 아래의 코드를 바꿔주면서 써야한다.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3002", mux)//만약에 컴파일 했는데 바로 컴파일이 꺼지게 된다면 포트 번호를 수정하자
}
```

- Client가 Server에 Req를 요청할때에 Input 정보를 넣을수 있다.
- 방식으로는 URL에 직접 넣는 방식이 있고
- Body에 직접 넣는 방식이 있다.
- 이러한 규칙에는 여러가지가 존재하는데 Binary가 예시이다.
- 하지만 Binary는 양쪽에 규약을 맞춰야하기 때문에 문제가 발생할 수 있다.
- 그래서 대표적으로 String 형태로 보낸다.

# JSON

- XML 방식보다는 요즘은 JSON 방식을 많이 쓴다.
- Java Script Object Notation
- 자바스크립트에서 오브젝트를 표현하는 방식인데 워낙 심플해서 주로 많이 쓰인다.

```go
{
	"Key" : "value"
	"Key2" : 3
	"Key3" : [3,4,5]
}
```