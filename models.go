package main

type Todo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
