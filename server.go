package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

func loadResponse(filePath string) interface{} {
	var resp interface{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic("Could not read the file: %s", filePath)
	}
	json.Unmarshal(data, &resp)
	return resp;
}


func getFilePath(path string) string {
	// TODO make dynamic
	var m = make(map[string]string)
	m["/todo/1"] = "todo1.json"
	m["/todo/2"] = "todo2.json"
	m["/todo/3"] = "todo3.json"

	return m[path];
}

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loading the url: ", r.URL.Path)

	filePath := getFilePath(r.URL.Path)
	p := loadResponse(filePath)

	pp, err := json.Marshal(p)

	if err != nil {
		log.Panic("Could not marshall to json the file ", filePath)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pp)
}

func TrueMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return true;
}

func main() {
	fmt.Println("Go rest backend, starting to listen to 8080!")
	r := mux.NewRouter()

	r.MatcherFunc(TrueMatcher).HandlerFunc(RouteHandler);

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

