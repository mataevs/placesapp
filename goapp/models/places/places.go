package places

import (
  "time"
  
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Place struct {
  Id bson.ObjectId `json:"id,omitempty"  bson:"_id,omitempty"`
  Name string `json:"name,omitempty"  bson:"name"`
  Address string `json:"address,omitempty"  bson:"address"`
  DateAdded time.Time `json:"DateAdded,omitempty"  bson:"dateadded"`
  Phone string `json:"phone,omitempty"  bson:"phone"`
  Email string `json:"email,omitempty"  bson:"email"`
  Website string `json:"website,omitempty"  bson:"website"`
  Type string `json:"type,omitempty"  bson:"type"`
  votes int `json:"votes,omitempty"  bson:"votes"`
}

type Repo struct {
  *mgo.Collection
}


func (p Place) SetName(n string) Place {
  p.Name = n
  return p
}

func (p Place) SetAddress(a string) Place {
  p.Address = a
  return p
}

func (p Place) SetDate(d time.Time) Place {
  p.DateAdded = d
  return p
}

func (p Place) SetPhone(ph string) Place {
  p.Phone = ph
  return p
}

func (p Place) SetEmail(e string) Place {
  p.Email = e
  return p
}

func (p Place) SetWebsite(w string) Place {
  p.Website = w
  return p
}

func (p Place) SetType(t string) Place {
  p.Type = t
  return p
}

func (p Place) IncVotes() Place {
  p.votes++
  return p
}

func (p Place) ResetVotes() Place {
  p.votes = 0
  return p
}

func (p Place) Votes() int {
  return p.votes
}


func InstanceRepo(c *mgo.Collection) Repo {
  return Repo{c}
}

func (repo Repo) GetPlaceByName(name string) (Place, error) {
  var place Place
  
  err := repo.Find(bson.M{"name": name}).One(&place)
  return place, err
}

func (repo Repo) GetPlaceById(id bson.ObjectId) (Place, error) {
  var place Place
  
  err := repo.FindId(&id).One(&place)
  return place, err
}

func (repo Repo) All() (places []Place, err error) {
  err = repo.Find(bson.M{}).All(&places)
  return
}


func (repo Repo) AddPlace(place Place) error {
  place.SetDate(bson.Now())
  return repo.Insert(&place)
}


func (repo Repo) UpdatePlaceByName(name string, place Place) error {
  return repo.Update(bson.M{"name": name}, &place)
}

func (repo Repo) UpdatePlaceById(id bson.ObjectId, place Place) error {
  return repo.UpdateId(id, &place)
}


func (repo Repo) RemovePlaceById(id bson.ObjectId) error {
  return repo.RemoveId(id)
}

func (repo Repo) RemovePlaceByName(name string) error {
  return repo.Remove(bson.M{"name" : name})
}