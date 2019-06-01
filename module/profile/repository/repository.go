package repository

import "mongoDB/module/profile/model"

type ProfileRepository interface {
	Create(profile *model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindById(string) (*model.Profile, error)
	GetAll() (model.Profiles, error)
}
