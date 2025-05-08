package auth

import (
	"learn-fiber/core/config/database"
	"learn-fiber/core/exception"
	"learn-fiber/core/helper"
	"learn-fiber/model"
)

func findUserByEmailAndPassword(req *loginRequest) (*model.User, error) {
	var user model.User

	result := database.DB.Where(&model.User{
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
