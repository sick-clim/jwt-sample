package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./auth"

	"github.com/gorilla/mux"
	"github.com/sick-clim/jwt-sample/auth"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	// public でハンドラーを実行
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handle(private))
	r.Handle("/auth", auth.GetTokenHandler)

	if err := http.ListenAndServe(":8070", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "VGolang to Google Cloud Vision",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NweEncoder(w).Encode(post)
})

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "VueCLI to Begin Vue.js",
		Tag:   "Vue.js",
		URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
})
