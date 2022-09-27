package controller

import (
	"Alterra/Project1-BE12-Book-Rent/model"
)

type UserControll struct {
	Model model.UserModel
}

func (uc UserControll) GetAll() ([]model.User, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc UserControll) Add(data model.User) (model.User, error) {
	res, err := uc.Model.Insert(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (uc UserControll) Find(Email, Password string) ([]model.User, error) {
	res, err := uc.Model.Find(Email, Password)
	if err != nil {
		return nil, err
	}
	return res, nil
}
