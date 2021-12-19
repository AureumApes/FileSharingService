package files

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type File struct {
	Id      string
	Name    string
	Content string
	Raw     bool
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	fileBytes, _ := ioutil.ReadFile("./upload/files/" + r.URL.Query().Get("id") + ".json")
	var file File
	json.Unmarshal(fileBytes, &file)
	if file.Raw {
		http.ServeFile(w, r, "./upload/raw/"+r.URL.Query().Get("id"))
	} else {
		w.Write([]byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>" + file.Name + "</title>\n    <style>\n        * {\n            background-color: #5d5d5d;\n            color: #d5d5d5;\n            text-align: center;\n        }\n\n        input {\n            font-size: 20px;\n            border: solid black 1px;\n        }\n    </style>\n</head>\n<body>\n<h1>" + file.Name + "</h1><div>" + strings.ReplaceAll(file.Content, "\n", "<br>") + "</div></body>\n</html>"))
	}
}

func TerminalFile(w http.ResponseWriter, r *http.Request) {
	fileBytes, _ := ioutil.ReadFile("./upload/files/" + r.URL.Query().Get("id") + ".json")
	var file File
	json.Unmarshal(fileBytes, &file)
	w.Write([]byte(file.Name + ":\n\n" + file.Content))
}
