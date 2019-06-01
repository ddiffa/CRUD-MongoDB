package main

import (
	"fmt"
	"mongoDB/config"
	"mongoDB/module/profile/model"
	"mongoDB/module/profile/repository"
	"time"
)

func main() {

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepository(db, "profile")

	//createProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile("1",profileRepository)
	//findById("1",profileRepository)
	getAllProfile(profileRepository)
}

func createProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "1"
	p.FirstName = "Diffa"
	p.LastName = "Dwi Desyawan"
	p.Email = "ddiffa2@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Create(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "1"
	p.FirstName = "Diffa"
	p.LastName = "Dwi Desyawan"
	p.Email = "diffadwi@gmail.com"
	p.Password = "12345678"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Update")
	}
}

func deleteProfile(id string, profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Deleted")
	}
}

func findById(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindById(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}

func getAllProfile(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profiles {
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}
