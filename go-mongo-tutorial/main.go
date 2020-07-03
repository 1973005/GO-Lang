package main

import (
	"fmt"
	"time"

	"github.com/apiang/go-mongo-tutorial/config"
	"github.com/apiang/go-mongo-tutorial/src/modules/profile/model"
	"github.com/apiang/go-mongo-tutorial/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Mongo DB")

	_, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfile("U1", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U2"
	p.FirstName = "jeremi"
	p.LastName = "Purnawan"
	p.Email = "ferly.apiang9@gmail.com"
	p.Password = "12345"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Ferly"
	p.LastName = "Apiang"
	p.Email = "1973005@maranatha.ac.id"
	p.Password = "12345"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Update..")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile delete..")
	}
}

func getProfile(id string.profileRepository repository.ProfileRepository){
	profile, err := profileRepository.FindByID(id)

	if err != nil{
		fmt.Println()
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)
}


func getProfiles(profileRepository repository.ProfileRepository){
	profiles, err := profileRepository.FindAll()

	if err != nil{
		fmt.Println(err)
	}
	for _, profile := range profile{
	fmt.Println("----------------")
	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)
	}
}