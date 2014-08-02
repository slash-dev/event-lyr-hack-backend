package main

import (
  "log"
)

type Database struct {
  users    map[string]User
  // Key user id, value, events ids.
  user_events    map[string][]string
  events   map[string]Event
  // Key event id, value, users ids.
  event_users    map[string][]string
}

func (db *Database) Init() {
  db.users = make(map[string]User)
  db.events = make(map[string]Event)
}

func (db *Database) GetUser(google_user *googlePlusUser) User {
  user, user_exists := db.users[google_user.Id]
  if !user_exists {
    user.Id = google_user.Id
    user.Name = google_user.Name
    user.Avatar = google_user.Image.Avatar
    db.users[user.Id] = user
    log.Print("New user:")
    log.Print(user)
  }
  for _, event_id := range db.user_events[user.Id] {
    user.Events = append(user.Events, db.events[event_id])
  }
  return user
}
