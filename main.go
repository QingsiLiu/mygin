package main

import (
	"fmt"
	"mygin"
	"net/http"
)

func main() {
	r := mygin.New()
	r.Get("/", test)

	r.Run(":8001")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello: lhq")
}
