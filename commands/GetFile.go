package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetFile() {
	var id string
	fmt.Print("ID >> ")
	fmt.Scanln(&id)
	res, _ := http.Get("http://82.165.184.35:8080/terminal?id=" + id)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body) + "\n")
}
