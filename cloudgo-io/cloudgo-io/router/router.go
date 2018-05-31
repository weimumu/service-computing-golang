package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.New()

	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println(root)
		}
	}
	mx.HandleFunc("/unknown", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "unknown path", 503)
	})
	mx.HandleFunc("/form", func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "form", struct{}{})
	})
	mx.HandleFunc("/login", formmerForLogin(formatter)).Methods("POST")
	mx.HandleFunc("/", apiTestHandler(formatter)).Methods("GET")
	mx.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(webRoot+"/public/"))))
}
