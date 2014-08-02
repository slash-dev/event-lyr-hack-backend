package main

type googlePlusUser struct {
  Id    string `json:"id"`
  Name  string `json:"displayName"`
  Image struct {
    Avatar string `json:"url"`
  } `json:"image"`
}
