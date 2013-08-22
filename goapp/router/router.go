package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  
  "github.com/gorilla/mux"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "places/models/places"
  "places/models/users"
  
  "encoding/json"
  "strconv"
)


var (
  mongoSession *mgo.Session
  database *mgo.Database
  repo places.Repo
  usersRepo users.Repo
  
  router = new(mux.Router)
)

func main() {
  var err error
  
  if mongoSession, err = mgo.Dial("mongodb://radubanciu:dromader@localhost/placesdb"); err != nil {
    panic(err)
  }
  log.Println("Connected to mongodb")
  
  database = mongoSession.DB("placesdb")
  repo = places.InstanceRepo(database.C("places"))
  usersRepo = users.InstanceRepo(database.C("users"))
  
  router.HandleFunc("/places/{id}", handlePlaceRequest).Methods("GET").Name("get_place")
  router.HandleFunc("/places/{id}", handlePlaceDelete).Methods("DELETE").Name("del_place")
  router.HandleFunc("/places/{id}", handlePlaceUpdate).Methods("PUT").Name("update_place")
  router.HandleFunc("/places", handlePlaceCreate).Methods("POST").Name("add_place")
  router.HandleFunc("/places", handlePlaces).Methods("GET").Name("get_places")
  //router.HandleFunc("/", handleMainPage).Name("main_page")
  
  
  router.HandleFunc("/users/{id}", handleUserRequest).Methods("GET").Name("get_user")
  router.HandleFunc("/users", handleUserCreate).Methods("POST").Name("add_user")

  http.Handle("/", router)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

  err = http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal(err)
  }
}


func handlePlaceRequest(w http.ResponseWriter, r *http.Request) {
  var place places.Place
  var err error
  
  vars := mux.Vars(r)
  id := vars["id"]
  
  if place, err = repo.GetPlaceById(bson.ObjectIdHex(id)); err != nil {
    serveError(w, err)
    return
  }
  
  writeJson(w, place)
}

func handlePlaces(w http.ResponseWriter, r *http.Request) {
  var places []places.Place
  var err error
  
  if places, err = repo.All(); err != nil {
    serveError(w, err)
    return
  }
  
  writeJson(w, places)
}


func handlePlaceDelete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  
  if err := repo.RemovePlaceById(bson.ObjectIdHex(id)); err != nil {
    serveError(w, err)
    return
  }
  fmt.Fprintf(w, "Place deleted.")
}


func handlePlaceCreate(w http.ResponseWriter, r *http.Request) {
  var place places.Place
  var err error
  
  if readJson(r, &place) {
    if err = repo.AddPlace(place); err != nil {
      log.Printf("%v", err)
      serveError(w, err)
      return
    }
  }
  
  fmt.Fprintf(w, "Place added.")
}


func handlePlaceUpdate(w http.ResponseWriter, r *http.Request) {
  var place places.Place
  var err error
  
  vars := mux.Vars(r)
  id := vars["id"]
  
  if readJson(r, &place) {
    if err = repo.UpdatePlaceById(bson.ObjectIdHex(id), place); err != nil {
      log.Printf("%v", err)
      serveError(w, err)
      return
    }
  }
  
  fmt.Fprintf(w, "Place updated.")
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
  var user users.User
  var err error
  
  vars := mux.Vars(r)
  id := vars["id"]
  
  if user, err = usersRepo.GetUserById(bson.ObjectIdHex(id)); err != nil {
    serveError(w, err)
    return
  }
  
  writeJson(w, user)
}

func handleUserCreate(w http.ResponseWriter, r *http.Request) {
  var user users.User
  var err error
  
  if readJson(r, &user) {
    if err = usersRepo.AddUser(user); err != nil {
      log.Printf("%v", err)
      serveError(w, err)
      return
    }
  }
  
  fmt.Fprintf(w, "User added.")
}


func readJson(r *http.Request, v interface{}) bool {
  defer r.Body.Close()
  
  var body []byte
  var err error
  
  body, err = ioutil.ReadAll(r.Body)
  
  if err != nil {
    log.Printf("Couldn't read request body %v", err)
    return false
  }
  
  if err = json.Unmarshal(body, v); err != nil {
    log.Printf("Couldn't parse request body %v", err)
    return false
  }
  
  return true
}


func writeJson(w http.ResponseWriter, v interface{}) {
  
  if data, err := json.Marshal(v); err != nil {
    log.Printf("Error marshalling json: %v", err)
  } else {
    w.Header().Set("Content-Length", strconv.Itoa(len(data)))
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
  }
}


func serveError(w http.ResponseWriter, err error) {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}