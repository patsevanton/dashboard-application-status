package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "os/exec"
  "time"
)

func main() {
  log.SetFlags(log.Lshortfile)
  // appVersion := []uint8{}
  appVersionChan := make(chan []uint8) 

  go func() {
    for {
      cmd := "/usr/bin/curl -I 'https://google.com/' 2>/dev/null | grep -i '^date:' | sed 's/^[Dd]ate: //g'"
      appVersion, err := exec.Command("bash","-c",cmd).Output()
      if err != nil {
          log.Fatal(err)
      }
      fmt.Printf("%s", appVersion)
      appVersionChan <- appVersion
      time.Sleep(5 * time.Second)
    }
  }()

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   tmpl, err := template.ParseFiles("index.html")
   if err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
   }
   appVersion := <-appVersionChan

   data := struct {
     AppVersion string
   }{
     AppVersion: string(appVersion[:]),
   }

   err = tmpl.Execute(w, data)
   if err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
   }
  })

  fmt.Println("http://127.0.0.1:3000")
  log.Fatalln(http.ListenAndServe("127.0.0.1:3000", nil))
}


