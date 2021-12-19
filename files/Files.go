package files

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Files(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("./upload/files")
	response := "ID\t\tName\n"
	for _, file := range files {
		fileBytes, _ := ioutil.ReadFile("./upload/files/" + file.Name())
		var fileInstance File
		json.Unmarshal(fileBytes, &fileInstance)
		response += fileInstance.Id + "\t" + fileInstance.Name + "\n"
	}
	w.Write([]byte(response))
}
