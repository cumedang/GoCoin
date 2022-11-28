package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/cumedang/GoCoin/blockchain"
)

type homedata struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

func home(rw http.ResponseWriter, r *http.Request) {
	data := homedata{"Home", blockchain.AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.FormValue("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}
func main() {
	templates = template.Must(template.ParseGlob(templateDir + "/pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partiais/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
