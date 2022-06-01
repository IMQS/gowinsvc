package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IMQS/gowinsvc/service"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	handler := func() error {
		err := http.ListenAndServe(":8620", nil)
		fmt.Printf("Server error %v\n", err)
		return err
	}
	handlerNoRet := func() {
		handler()
	}
	var err error
	if !service.RunAsService(handlerNoRet) {
		fmt.Printf("Running not as service\n")
		err = handler()
	}
	fmt.Printf("Error %v\n", err)
	if err != nil {
		os.Exit(-1)
	} else {
		os.Exit(0)
	}
}
