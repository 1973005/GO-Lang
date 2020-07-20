package repository

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/wuryscamp/go-mongo-tutorial/src/modules/profile/model"
)

//profile repository mongo
type profileRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

//save
func (r *profileRepositoryMongo) Save(profile *model.Profile) error {
	err := r.db.C(r.collection).Insert(profile)
	return err
}

//update
func (r *profileRepositoryMongo) Update(id string, profile *model.Profile) error {
	profile.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"id": id}, profile)
	return err
}

//delete
func (r *profileRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"id": id})
	return err
}

//findbyid
func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	var profile model.Profile

	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

//findall
func (r *profileRepositoryMongo) FindAll() (model.Profile, error) {
	var profile model.Profiles

	err := r.db.C(r.collection).Find(bson.M{}).All(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
