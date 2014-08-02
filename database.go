package main

import (
  "log"
  "strconv"
)

type Database struct {
  users    map[string]User
  // Key user id, value, events ids.
  user_events    map[string][]string
  events   map[string]Event
  // Key event id, value, users ids.
  event_users    map[string][]string

  last_user_id int64
  last_event_id int64
}

func (db *Database) Init() {
  db.users = make(map[string]User)
  db.events = make(map[string]Event)
  db.event_users = make(map[string][]string)
  db.user_events = make(map[string][]string)
  db.last_user_id = 0
  db.last_event_id = 0
}

func (db *Database) GetUser(google_user *googlePlusUser) User {
  user, user_exists := db.users[google_user.Id]
  if !user_exists {
    user.Id = google_user.Id
    user.Name = google_user.Name
    user.Avatar = google_user.Image.Avatar
    db.users[user.Id] = user
    log.Print("New user:")
  }
  for _, event_id := range db.user_events[user.Id] {
    new_event := db.events[event_id]
    for _, participant_id := range db.event_users[new_event.Id] {
      new_event.Participants =
          append(new_event.Participants, db.users[participant_id])
    }
    user.Events = append(user.Events, new_event)
  }
  return user
}

func (db *Database) CreateEvent(title string, participants []User) Event {
  db.last_event_id++
  event := Event{ Id: strconv.FormatInt(db.last_event_id, 10), Title: title}
  db.events[event.Id] = event;
  event.Participants = participants
  for _, participant := range participants {
    if _, ok := db.users[participant.Id]; !ok {
      log.Print("ERROR: Create event with invalid Id:", participant.Id)
      continue;
    }
    db.user_events[participant.Id] =
        append(db.user_events[participant.Id], event.Id);
    db.event_users[event.Id] =
        append(db.event_users[event.Id], participant.Id);
  }
  return event
}
