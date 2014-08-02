package main

type User struct {
  Id  string
  Name string
  Avatar string
  Events []Event
}

type Event struct {
  Id  string
  Title string
  Participants []User
}