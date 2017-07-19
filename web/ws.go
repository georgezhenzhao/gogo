package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CallGet(url string) string {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	respMsg := fmt.Sprintf("%s", robots)
	return respMsg
}
