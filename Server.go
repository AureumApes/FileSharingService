package main

import (
	"FileSharing/files"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", files.Upload)
	http.HandleFunc("/files", files.Files)
	http.HandleFunc("/file", files.FileHandler)
	http.HandleFunc("/terminal", files.TerminalFile)
	http.ListenAndServe(":7623", nil)
}
