package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		pathItems := strings.Split(request.URL.Path,"/")
		var items []int64

		for i := range pathItems{
			out,err := strconv.ParseInt(pathItems[i],10,64)
			if err == nil{
				items = append(items,out)
			}
		}

		_,err := fmt.Fprint(writer,add(items...))
		if err != nil{
			log.Fatal(err)
		}
	})

	log.Println("Listen HTTP on port 8070")
	err := http.ListenAndServe(":8070",nil)
	if err != nil{
		log.Fatal(err)
	}
}

// Sum all items...
func add(items ...int64)int64  {
	var out int64
	for _,item := range items{
		out+=item
	}
	return out
}