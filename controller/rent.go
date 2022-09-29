package controller

import "Alterra/Project1-BE12-Book-Rent/model"

type RentControll struct {
	Model model.RentModel
}

func (rc RentControll) Add(data model.Rent) (model.Rent, error) {
	res, err := rc.Model.Input(data)
	if err != nil {
		return model.Rent{}, err
	}
	return res, nil
}
