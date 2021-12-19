package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ListFiles() {
	res, _ := http.Get("http://82.165.184.35:8080/files")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body) + "\n")
}
