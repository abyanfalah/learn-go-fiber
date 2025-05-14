package auth

import (
	"fmt"
	"learn-fiber/core/config/db"
	"learn-fiber/core/exception"
	"learn-fiber/core/helper"
	"learn-fiber/core/helper/generator"
	"learn-fiber/model"
)

func findUserByEmailAndPassword(req *loginRequest) (*model.User, error) {
	var user model.User

	result := db.Use().Where(&model.User{
		Email: req.Email,
	}).First(&user)

	if result.Error != nil {
		return nil, exception.BadCredentials(result.Error)
	}

	isCorrect, err := helper.IsCorrectPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}

	if !isCorrect {
		return nil, exception.ErrBadCredentials
	}

	return &user, nil
}

func saveUser(req *registerRequest) (*model.User, error) {
	// err := user.IsUsedEmail(req.Email)
	// if err != nil {
	// 	return nil, err
	// }

	u := new(model.User)
	u.ID = generator.GenerateId()
	u.Name = req.Name
	u.Email = req.Email
	u.Password = req.Password

	r := db.Use().Create(u)
	if r.Error != nil {
		return nil, r.Error
	}

	return u, nil
}

func Hehe() {
	fmt.Println("hehe")
}
