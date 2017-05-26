package main

import (
  "fmt"
  "log"
  "net/http"
  "runtime"
)

func main() {
  proc := runtime.GOMAXPROCS(1)
  log.Print("GOMAXPROC = ", proc)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    fmt.Fprintf(w, "{\"hello\":\"world\"}")
  })

  log.Print("server running ...")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
