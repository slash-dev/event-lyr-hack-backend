package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

var db Database

func GetUser(access_token string) User {
  resp, err := http.Get(
    "https://www.googleapis.com/plus/v1/people/me?access_token=" +
    access_token)
  if err != nil {
    log.Print("Error getting google user")
    var user User
    return user
  }
  user := new(googlePlusUser)
  decoder := json.NewDecoder(resp.Body)
  decoder.Decode(&user)
  resp.Body.Close()
  return db.GetUser(user)
}

func me(w http.ResponseWriter, r *http.Request) {
  log.Print("access token: " + r.FormValue("access_token"))
  user := GetUser(r.FormValue("access_token"))
  serialized_user, err := json.Marshal(user)

  if err != nil {
    log.Print("Error serializing user: ")
    log.Print(user)
    return
  }

  fmt.Fprintf(w, "%s", serialized_user)
}

func create_event(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "create_event!\n")
}

func main() {
  db.Init()
  http.HandleFunc("/me", me)
  http.HandleFunc("/me/create_event", create_event)
  http.ListenAndServe(":3000", nil)
}
