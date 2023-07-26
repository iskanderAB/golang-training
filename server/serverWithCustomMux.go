package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute 🚀!")
}

func main() {
	var mux = new(MyMux)
	fooType := reflect.TypeOf(&MyMux{})
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
	http.ListenAndServe(":9090", mux)
}
