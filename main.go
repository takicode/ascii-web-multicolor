package main

import (
  "html/template"
  "log"
  "net/http"
  "ascii-web-multicolor/core"
  "fmt"
)

type pageData struct {
  Result string
  Message string
  WordString string
  Substring  string
  Coloring   string
}

var templ = template.Must(
  template.New("index.html").Funcs(template.FuncMap{
    "safe": func(s string) template.HTML {
      return template.HTML(s)
    },
  }).ParseFiles("./templates/index.html"),
)



func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  http.HandleFunc("/", HomeHandler)
  http.HandleFunc("/ascii-art", GenerateHandler)

  log.Println("Server running on port 8085")
  log.Fatal(http.ListenAndServe(":8085", nil))
}



func HomeHandler(w http.ResponseWriter, r *http.Request) {
  if err := templ.Execute(w, pageData{}); err != nil {
    log.Fatal("Error loading template:", err)
  }
}


func GenerateHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }


  inputString := r.FormValue("text")
  banner      := r.FormValue("banner")
  inputWord   := r.FormValue("coloredWord")
  color       := r.FormValue("color")

  fmt.Println(inputString)
  fmt.Println(inputWord)
  fmt.Println(color)
  fmt.Println(banner)
  result, message := core.ColorLogic(inputString, banner, inputWord, color)

  data := pageData{
    Result: result,
    Message: message,
    WordString:inputString,
    Substring:inputWord,
    Coloring:color,
}

  if err := templ.Execute(w, data); err != nil {
    log.Println("Template error:", err)
  }
}