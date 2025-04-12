package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	urlStore  = make(map[string]string)
	mutex     = &sync.RWMutex{}
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

const (
	shortURLPrefix  = "http://localhost:8080/"
	shortCodeLength = 6
	charset         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateShortCode() string {
	b := make([]byte, shortCodeLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	url := r.FormValue("url")
	if url == "" || (!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://")) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()

	mutex.Lock()
	urlStore[shortCode] = url
	mutex.Unlock()

	shortURL := shortURLPrefix + shortCode
	templates.ExecuteTemplate(w, "result.html", shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	mutex.RLock()
	originalURL, found := urlStore[code]
	mutex.RUnlock()

	if !found {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {}) // avoid 404
	http.HandleFunc("/redirect/", redirectHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
