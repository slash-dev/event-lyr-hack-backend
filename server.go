package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings"
)

func me(w http.ResponseWriter, r *http.Request) {
  log.Print("access token: " + r.FormValue("access_token"))
  resp, err := http.Get(
    "https://www.googleapis.com/plus/v1/people/me?access_token=" +
    r.FormValue("access_token"))
  if err != nil {
    return
  }
  user := new(googlePlusUser)
  decoder := json.NewDecoder(resp.Body)
  decoder.Decode(&user)
  resp.Body.Close()

  user.Image.Avatar = strings.Replace(user.Image.Avatar, "sz=50", "sz=200", 1)

  serialized_user, err := json.Marshal(user)

  if err != nil {
    return
  }

  fmt.Fprintf(w, "%s", serialized_user)
}

func create_event(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "create_event!\n")
}

func main() {
  http.HandleFunc("/me", me)
  http.HandleFunc("/me/create_event", create_event)
  http.ListenAndServe(":3000", nil)
}
