package reviews

import (
  "time"
  
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Review struct {
  Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
  Title string `json:"title" bson:"title"`
  Body string `json:"content,omitempty" bson:"body,omitempty"`
  Created time.Time `json:"created,omitempty" bson:"created,omitempty"`
  Likes int `json:"likes,omitempty" bson:"likes,omitempty"`
  Dislikes int `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
  Author mgo.DBRef `json:"author" bson:"author"`
  Place mgo.DBRef `json:"place" bson:"place"`
}

type Repo struct {
  *mgo.Collection
}

func InstanceRepo(c *mgo.Collection) Repo {
  return Repo{c}
}

func (repo Repo) GetOneReviewByField(field string, value string) (Review, error) {
  var review Review
  
  err := repo.Find(bson.M{field: value}).One(&review)
  return review, err
}

func (repo Repo) GetMultipleReviewsByField(field string, value interface{}) (reviews []Review, err error) {
  err = repo.Find(bson.M{field: value}).All(&reviews)
  return
}

func (repo Repo) GetReviewById(id bson.ObjectId) (Review, error) {
  var review Review
  
  err := repo.FindId(&id).One(&review)
  return review, err
}

func (repo Repo) GetReviewsForPlace(place mgo.DBRef) ([]Review, error) {
  return repo.GetMultipleReviewsByField("place", place)
}

func (repo Repo) GetReviewsFromAuthor(author mgo.DBRef) ([]Review, error) {
  return repo.GetMultipleReviewsByField("author", author)
}