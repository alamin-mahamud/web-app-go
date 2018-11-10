package main

import (
	"html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var redisClient *redis.Client
var templates *template.Template

func main() {
	r := mux.NewRouter()

	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	templates = template.Must(
		template.ParseGlob("templates/*.html"),
	)

	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := redisClient.LRange("comments", 0, 10).Result()

	if err != nil {
		panic(err)
	}

	templates.ExecuteTemplate(w, "index.html", comments)
}
