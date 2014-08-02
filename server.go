package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

var db Database

var debug bool

func GetUser(access_token string) User {
  google_user := new(googlePlusUser)
  if !debug {
    resp, err := http.Get(
      "https://www.googleapis.com/plus/v1/people/me?access_token=" +
      access_token)
    if err != nil {
      log.Print("Error getting google user")
      return User{}
    }
    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&google_user)
    resp.Body.Close()
  } else {
    google_user.Id = access_token;
    google_user.Name = "Name" + access_token;
    google_user.Image.Avatar = "http://avatar.com/" + access_token;
  }
  user := db.GetUser(google_user)
  log.Print(user)
  return user
}

func me(w http.ResponseWriter, r *http.Request) {
  access_token := r.FormValue("access_token")
  log.Print("access token: " + access_token)
  user := GetUser(access_token)
  serialized_user, err := json.Marshal(user)

  if err != nil {
    log.Print("Error serializing user: ")
    log.Print(user)
    return
  }

  fmt.Fprintf(w, "%s", serialized_user)
}

func create_event(w http.ResponseWriter, r *http.Request) {
  access_token := r.FormValue("access_token")
  log.Print("access token: " + access_token)
  user := GetUser(access_token)
  user.Events = []Event{}
  title := r.FormValue("title")
  var participants []User
  err := json.Unmarshal([]byte(r.FormValue("participants")), &participants)
  if err != nil {
    log.Print(err)
    return
  }
  participants = append(participants, user)
  event := db.CreateEvent(title, participants)
  log.Print(event)
  serialized_event, err := json.Marshal(event)

  if err != nil {
    log.Print("Error serializing event...")
    return
  }
  fmt.Fprintf(w, "%s", serialized_event)
}

func main() {
  debug = false
  db.Init()
  http.HandleFunc("/me", me)
  http.HandleFunc("/me/create_event", create_event)
  http.ListenAndServe(":3000", nil)
}
