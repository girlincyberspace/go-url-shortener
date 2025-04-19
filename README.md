#  URL Shortener in Go

A simple URL shortener web application written in **Go** with an **HTML frontend**. Users can enter a long URL and receive a shortened version that redirects to the original.

---

##  Features

- Shorten any valid `http://` or `https://` URL
- Simple web interface using Go's `html/template`
- In-memory storage (no database needed)
- Instant redirection from short URL to long URL

---

##  Requirements

- [Go](https://golang.org/dl/) 1.18 or later

---

##  Project Structure

url-shortener/ 

├── main.go # Main Go server 

└── templates/ # HTML templates

    ├── index.html # Form to input long URLs

    └── result.html
    
---

##  Running the Project

1. **Clone the repo:**
   ```bash
   git clone https://github.com/girlincyberspace/url-shortener.git
   cd url-shortener

2. **Running the server:**
   ```bash
   go run main.go

3. **Visit in browser:**
   ```http://localhost:8080

