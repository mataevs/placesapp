package main

import (
    "fmt"
    
    "places/router/places"
    "labix.org/v2/mgo"
)

func main() {
  session, err := mgo.Dial("mongodb://radubanciu:dromader@localhost/placesdb")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  
  session.SetMode(mgo.Monotonic, true)
  
  c := session.DB("placesdb").C("places")
  
  if err := places.RemovePlace("Gigi's Place", c); err != nil && err != mgo.ErrNotFound {
    panic(err)
  }
  
  var place places.Place
  place = place.SetName("Gigi's Place").SetPhone("0724-111-222").SetEmail("contact@gigi.ro").SetType("Fusion")
  
  if err := places.AddPlace(place, c); err != nil {
    panic(err)
  }
  
  place, err = places.FindPlace("Gigi's Place", c)
  if err != nil {
    panic(err)
  }
  
  fmt.Println(place)
}