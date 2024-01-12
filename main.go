package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "static/index.html")
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseMultipartForm(10 * 1024 * 1024) //size 10 MB
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, "Error Retrieving the File")
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	response := sendToOpenAI(fileBytes)
	w.Write([]byte(response))
}

func sendToOpenAI(file []byte) string {
	return "TODO"
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.POST("/upload", fileUploadHandler)
	fmt.Println("Running server on :8081")
	fmt.Println(http.ListenAndServe(":8081", router))
}
