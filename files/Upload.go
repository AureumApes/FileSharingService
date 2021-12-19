package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	MB = 1 << 20
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		uploadedFile, fileHeader, _ := r.FormFile("file")
		fileName := r.FormValue("fileName")
		id := generateId()
		if fileHeader.Size > MB*200 {
			w.WriteHeader(http.StatusConflict)
			logFile, _ := os.OpenFile("./logs/upload.txt", os.O_APPEND|os.O_WRONLY, 0644)
			defer logFile.Close()
			logFile.WriteString("[" + r.RemoteAddr + "] >> Failed upload" + fileName + " with the ID " + id + " because it was too big. Size: " + strconv.Itoa(int(fileHeader.Size/MB)) + "MB\n")
			return
		} else {
			defer uploadedFile.Close()
			fileBytes, _ := ioutil.ReadAll(uploadedFile)
			if r.FormValue("raw") == "on" {
				os.WriteFile("./upload/raw/"+id, fileBytes, os.ModePerm)
			}
			fmt.Println()
			file := File{
				Name:    fileName,
				Id:      id,
				Content: string(fileBytes),
				Raw:     r.FormValue("raw") == "on",
			}
			fileBytes, _ = json.Marshal(file)
			os.WriteFile("./upload/files/"+id+".json", fileBytes, os.ModePerm)
			logFile, _ := os.OpenFile("./logs/upload.txt", os.O_APPEND|os.O_WRONLY, 0644)
			defer logFile.Close()
			logFile.WriteString("[" + r.RemoteAddr + "] >> Uploaded " + fileName + " with the ID " + id + "\n")
			w.Write([]byte("<script>location.replace(\"/upload\")</script>"))
		}
	} else {
		http.ServeFile(w, r, "./frontend/Upload.html")
	}
}

func generateId() string {
	index := 11
	returned := ""
	for index >= 0 {
		rand.Seed(time.Now().UnixNano())
		returned += strconv.Itoa(rand.Intn(9))
		index--
	}
	return returned
}
