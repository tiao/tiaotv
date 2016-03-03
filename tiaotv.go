package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Conf struct {
	directory string
	http_port int
	Timeout   int
	Refresh   int
	Files     []string
}

var conf = Conf{
	directory: "images",
	http_port: 8080,
	Timeout:   10000,
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	t, err := t.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	// Set Files list and refresh timeout
	conf.Files = getImages()
	conf.Refresh = (conf.Timeout * len(conf.Files)) / 1000
	conf.Refresh -= 1

	err = t.Execute(w, conf)
	if err != nil {
		panic(err)
	}
}

func getImages() []string {
	files, _ := filepath.Glob(conf.directory + "/*.jpg")
	for i, p := range files {
		_, file := filepath.Split(p)
		// Add prefix /images
		files[i] = "images/" + file
	}

	log.Print(files)
	return files
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)

	fs_images := http.FileServer(http.Dir(conf.directory))
	http.Handle("/images/", http.StripPrefix("/images/", fs_images))

	http.HandleFunc("/", serveTemplate)
	port := fmt.Sprintf(":%d", conf.http_port)
	log.Println("Listening " + port)
	http.ListenAndServe(port, nil)
}
