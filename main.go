package main

import (
    "fmt"
    "html"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "os"
)



func main() {

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))


    http.HandleFunc("/hallo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hallo von mir, Pfad: %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/", serveTemplate)

    log.Fatal(http.ListenAndServe(":8080", nil))

}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    layoutPath := filepath.Join("templates", "layout.html")
    fPath := filepath.Join("templates", filepath.Clean(r.URL.Path))

    info, err := os.Stat(fPath)
    // Return a 404 if the template doesn't exist
    if err != nil {
        if os.IsNotExist(err) {
        log.Println(err.Error())
        log.Println(fPath)
        http.NotFound(w, r)
        return
        }
    }

    // Return a 404 if the request is for a directory
    if info.IsDir() {
        
        http.NotFound(w, r)
        return
    }

  
    tmpl, err := template.ParseFiles(layoutPath, fPath)
    //Return 500 if Layout Template bad 
    if err != nil {
        // Log the detailed error
        log.Println(err.Error())
        // Return a generic "Internal Server Error" message
        http.Error(w, http.StatusText(500), 500)
        return
    }

    //Return 500 if Content bad 
    if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
        log.Println(err.Error())
        http.Error(w, http.StatusText(500), 500)
      }
}

