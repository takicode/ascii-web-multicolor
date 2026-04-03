package main

import (
  "html/template"
  "log"
  "net/http"
  "ascii-web-multicolor/core"
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
   


  if err := templ.Execute(w, pageData{}); 
  
  err != nil {
    log.Fatal("Error loading template:", err)
  }
}


func GenerateHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w,"405-Method not Allowed", http.StatusMethodNotAllowed)
    return
  }


  inputString := r.FormValue("text")
  banner      := r.FormValue("banner")
  inputWord   := r.FormValue("coloredWord")
  color       := r.FormValue("color")

<<<<<<< HEAD
  if strings.TrimSpace(inputString) == "" || banner == ""{
    http.Error(w, "400-Bad-Request", http.StatusBadRequest)
    return
  }


=======
>>>>>>> a0c3c2e7ace0cff288fb15478fbbb88725759db7
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
    http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
    return
  }
}


