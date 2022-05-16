package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for i := 1; i <= 52; i++ {
		rs, err := http.Get("http://www.heidelberg-catechism.com/en/lords-days/" + fmt.Sprint(i) + ".html")
		if err != nil {
			log.Fatalln(err)
		}

		bytes, err := ioutil.ReadAll(rs.Body)
		if err != nil {
			log.Fatalln(err)
		}

		rs.Body.Close()

		dst := "LD-" + fmt.Sprint(i) + ".html"
		ioutil.WriteFile(dst, bytes, os.ModePerm)
		fmt.Println("wrote: ", dst)
	}
}
