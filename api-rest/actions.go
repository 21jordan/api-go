package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var collection = getSession().DB("curso_go").C("movies")

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	log.Println("se conecto a la base de datos")
	return session
}

func responseMovie(res http.ResponseWriter, status int, results Movie) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(results)
}
func responseMovies(res http.ResponseWriter, status int, results []Movie) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(results)
}
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hola mundo desde un servidor web de go")
}
func MovieList(res http.ResponseWriter, req *http.Request) {
	var results []Movie
	err := collection.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("resultados", results)
	}
	responseMovies(res, 200, results)
}
func MovieShow(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	movieId := params["id"]
	results := Movie{}
	if !bson.IsObjectIdHex(movieId) {
		res.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movieId)
	fmt.Println(oid)
	err := collection.FindId(oid).One(&results)
	if err != nil {
		res.WriteHeader(404)
		return
	}
	responseMovie(res, 200, results)

}
func MovieAdd(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var movieData Movie
	//usando el "&"" se indica que una variable aun no tiene nada
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	err = collection.Insert(movieData)
	if err != nil {
		res.WriteHeader(500)
		return
	}
	responseMovie(res, 200, movieData)
}

func MovieUpdate(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	movieId := params["id"]
	if !bson.IsObjectIdHex(movieId) {
		res.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movieId)
	decoder := json.NewDecoder(req.Body)
	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
		res.WriteHeader(500)
		return
	}
	defer req.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movieData}
	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		res.WriteHeader(404)
		return
	}

	responseMovie(res, 200, movieData)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}
func (this *Message) setMessage(data string) {
	this.Message = data
}

func MovieRemove(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	movieId := params["id"]
	if !bson.IsObjectIdHex(movieId) {
		res.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movieId)
	fmt.Println(oid)
	err := collection.RemoveId(oid)
	if err != nil {
		res.WriteHeader(404)
		return
	}
	message := new(Message)
	message.setStatus("success")
	message.setMessage("La pelicula ha sido borrada correctamente")
	results := message

	// results := Message{"success", "La pelicula ha sido borrada correctamente"}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(results)
}
