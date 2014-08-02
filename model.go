package main

type User struct {
  Id  string `json:"id"`
  Name string `json:"name"`
  Avatar string `json:"avatar"`
  Events []Event `json:"events"`
}

type Event struct {
  Id  string `json:"id"`
  Title string `json:"title"`
  Participants []User `json:"participants"`
}