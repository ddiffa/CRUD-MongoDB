package repository

import (
	"gopkg.in/mgo.v2"
	"mongoDB/module/profile/model"
)

type profileRepository struct {
	db *mgo.Database
	collection string 
}

func NewProfileRepository(db *mgo.Database, collection string) *profileRepository  {
	return &profileRepository{
		db: db, 
		collection:collection,
	}
}

func (r *profileRepository) Create(profile *model.Profile) error{
	return nil
}

func (r *profileRepository) Update(id string, profile *model.Profile)error{
	return nil
}

func (r *profileRepository) Delete (id string) error{
	return nil
}

func (r *profileRepository) FindById(id string)(*model.Profile, error){
	return &model.Profile{}, nil
}

func (r *profileRepository) GetAll()(model.Profiles,error){
	return nil, nil
}