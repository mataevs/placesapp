package users

import (
  "time"
  
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type User struct {
  Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
  Name string `json:"name,omitempty" bson:"name,omitempty"`
  User string `json:"user" bson:"user"`
  Pass string `json:"pass" bson:"pass"`
  JoinedDate time.Time `json:"joinedDate,omitempty" bson:"joineddate"`
  LastSeen time.Time `json:"lastSeen,omitempty" bson:"lastseen,omitempty"`
  Email string `json:"email" bson:"email"`
  IsAdmin bool `json:"isAdmin" bson:"isadmin"`
  Active bool `json:"active" bson:"active"`
}

type Repo struct {
  *mgo.Collection
}

func InstanceRepo(c *mgo.Collection) Repo {
  return Repo{c}
}

func (repo Repo) GetOneUserByField(field string, value string) (User, error) {
  var user User
  
  err := repo.Find(bson.M{field: value}).One(&user)
  return user, err
}

func (repo Repo) GetMultipleUsersByField(field string, value string) (users []User, err error) {
  err = repo.Find(bson.M{field: value}).All(&users)
  return
}

func (repo Repo) GetUserById(id bson.ObjectId) (User, error) {
  var user User
  
  err := repo.FindId(&id).One(&user)
  return user, err
}

func (repo Repo) GetUserByUsername(username string) (User, error) {
  return repo.GetOneUserByField("user", username)
}

func (repo Repo) AddUser(user User) error {
  user.JoinedDate = bson.Now()
  return repo.Insert(&user)
}

func (repo Repo) UpdateUserByUsername(username string, user User) error {
  return repo.Update(bson.M{"user": username}, &user)
}

func (repo Repo) UpdateUserById(id bson.ObjectId, user User) error {
  return repo.UpdateId(id, &user)
}

func (repo Repo) RemoveUserById(id bson.ObjectId) error {
  return repo.RemoveId(id)
}

func (repo Repo) RemoveUserByUsername(username string) error {
  return repo.Remove(bson.M{"user": username})
}