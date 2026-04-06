package main


import (
  "html/template"
  "net/http"
  "log"
  "ascii-web-multicolor/core"
)

type pageData struct{
  Result     string
  WordString string
  Coloring   string
  Substring  string
  Message    string
  Error404   string
  Error      string
}


var templ = template.Must(template.New("index.html").Funcs(template.FuncMap{
  "safe" : func(s string) template.HTML{
    return template.HTML(s)
  }}).ParseGlob("./templates/*.html"))

func main(){
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   http.HandleFunc("/", HomeHandler)
   http.HandleFunc("/ascii-art", GenerateHandler)

   log.Println("Server listening on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil)) 
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/" {
      w.WriteHeader(http.StatusNotFound)
      error404, _ := core.ColorLogic(`404`, "standard", "", "")
      errorText, _ := core.ColorLogic(`Page Not Found`, "standard", "", "")
      err:= templ.ExecuteTemplate(w, "notFound.html", pageData{
        Error : errorText,
        Error404:error404,
      })
      if err != nil{
      log.Println("error loading 404 page", err)
      http.Error(w, "404-Not found", http.StatusNotFound)
    }
      return
    }

    err := templ.ExecuteTemplate(w, "index.html", pageData{})
    if err != nil{
      log.Println("error loading template file", err)
      http.Error(w, "500-internal server error", http.StatusInternalServerError)
    }
}

func GenerateHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost{
      log.Println("Method not allowed",)
      http.Error(w, "405-Not allowed", http.StatusMethodNotAllowed)
      return
    }
    err := r.ParseForm()
     if err != nil {
    http.Error(w, "400 - Bad Request", http.StatusBadRequest)
    return
    }

    text := r.FormValue("text")
    banner := r.FormValue("banner")
    color := r.FormValue("color")
    coloredWord := r.FormValue("coloredWord")

    if text == ""{
      http.Error(w, "400-Missing required field", http.StatusBadRequest)
      return
    }
  
  result, message :=  core.ColorLogic(text, banner, coloredWord, color)

  errs := templ.ExecuteTemplate(w, "index.html", pageData{
    Result:result,
    Message:message,
    WordString: text,
    Coloring :color,
    Substring: coloredWord,
  } )

  if errs != nil{
    log.Println("error parsing content", errs)
    http.Error(w, "500-internal server error", http.StatusInternalServerError)
    return
  }

}



