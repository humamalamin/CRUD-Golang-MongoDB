package main

import (
	"fmt"
	"go-mongo-tutorial/config"
	"go-mongo-tutorial/src/modules/profile/model"
	"go-mongo-tutorial/src/modules/profile/repository"
	"time"
)

func main(){
	fmt.Println("Go Mongo DB")

	db,err := config.GetMongoB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db,"profile")

	//saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)

	getProfile("U1",profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository){
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Aldo"
	p.LastName = "dino"
	p.Email = "aldo@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Profile saved...")
	}
}

func updateProfile(profileRepository repository.ProfileRepository){
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Aldo"
	p.LastName = "dino"
	p.Email = "aldo123@gmail.com"
	p.Password = "12345678"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1",&p)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Profile updated...")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository)  {
	err := profileRepository.Delete("U1")

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Profile deleted...")
	}

}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)

}