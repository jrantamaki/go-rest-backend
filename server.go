package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"time"
	"math/rand"
)

type StaticRoute struct {
	Route            string  `json:"route"`
	ResponseFilePath *string `json:"responseFilePath"`
	HttpStatus       int     `json:"httpStatus"`
	Delay            *int    `json:"delay"`
}

type Routes struct {
	Routes []StaticRoute `json:"routes"`
}

var RoutingMap map[string]StaticRoute

func loadResponse(filePath string) interface{} {
	var resp interface{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to load the response file: ", err)
	}
	json.Unmarshal(data, &resp)
	return resp;
}

func LoadRoutes () {
	var configuredRoutes Routes

	file := "config.json"
	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("Could not load routes from configuration. ", err)
	}
	err = json.Unmarshal(data, &configuredRoutes)

	if err != nil {
		log.Fatal("Failed unmarshalling the routes ", err)
	}

	RoutingMap = make(map[string]StaticRoute)

	for index, route := range configuredRoutes.Routes {
		RoutingMap[route.Route] = configuredRoutes.Routes[index]
	}
}

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loading the url: ", r.URL.Path, " with request method: ", r.Method)

	route, exists := RoutingMap[r.URL.Path]

	if !exists {
		fmt.Println("Check your config.json. No route configured for ", r.URL.Path)
		w.WriteHeader(500)
		return
	}

	filePath := route.ResponseFilePath
	httpStatus := route.HttpStatus
	delay := route.Delay

	if (delay != nil && *delay > 0) {
		pause := rand.Intn(*delay)
		time.Sleep(time.Duration(pause) * time.Millisecond)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if filePath != nil {
		p := loadResponse(*filePath)

		pp, err := json.Marshal(p)

		if err != nil {
			log.Panic("Json marshalling failed with error: ", err)
		}
		w.Write(pp)
	}
}

func TrueMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	return true;
}

func main() {
	fmt.Println("Go rest backend, starting to listen to 8080!")
	LoadRoutes()

	r := mux.NewRouter()

	r.MatcherFunc(TrueMatcher).HandlerFunc(RouteHandler);

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
