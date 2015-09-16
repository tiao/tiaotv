package main

import (
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

var conf = Conf {
  directory: "/mnt/slideshow/SlideEldorado",
  http_port: 80,
  Timeout: 10000,
}

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/css/", fs)
  http.Handle("/js/", fs)
  fs_images := http.FileServer(http.Dir(conf.directory))
  // Remove prefix /images
  http.Handle("/images/", http.StripPrefix("/images/", fs_images))

  http.HandleFunc("/", serveTemplate)

  log.Println("Listening:8080...")
  http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  t := template.New("index") // Create a template.
  t, err := t.ParseFiles("templates/index.html")  // Parse template file.
  if err != nil { panic(err) }

  // Set refresh timeout
  conf.Files = getImages()
  conf.Refresh = (conf.Timeout * len(conf.Files)) / 1000
  conf.Refresh -= 1


  err = t.Execute(w, conf)  // merge.
  if err != nil { panic(err) }
}

func getImages() []string {
  files, _ := filepath.Glob(conf.directory + "/*.['j','J']['p','P']['g','G']")
  for i, p := range files {
        _, file := filepath.Split(p)
        // Add prefix /images
        files[i] = "images/" +file
  }
  log.Print(files) // contains a list of all iamges in the current directory

  return files
}
