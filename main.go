package main

import "net/http"

func main()  {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		log.Println("Hello World")
	})
	http.ListenAndServe(":9090", nil)
}