package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mongoDB/module/profile/model"
	"time"
)

type profileRepository struct {
	db         *mgo.Database
	collection string
}

func NewProfileRepository(db *mgo.Database, collection string) *profileRepository {
	return &profileRepository{
		db:         db,
		collection: collection,
	}
}

func (r *profileRepository) Create(profile *model.Profile) error {
	err := r.db.C(r.collection).Insert(profile)
	return err
}

//selector menggunakan bson
func (r *profileRepository) Update(id string, profile *model.Profile) error {
	profile.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"id": id}, profile)
	return err
}

func (r *profileRepository) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"id": id})
	return err
}

func (r *profileRepository) FindById(id string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&profile)

	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) GetAll() (model.Profiles, error) {
	var profiles model.Profiles

	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
