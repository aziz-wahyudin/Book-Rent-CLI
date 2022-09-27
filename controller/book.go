package control

import (

)

type BookControll struct {
	Model model.Book
}

func (mc BookControll) GetAll() ([]model.Book, error) {
	res, err := mc.Model.GetAll ()
	if err!= nil {
		return nil, err
	}
	return res, nil
}

//mengambil data 
func (gc BookControll) Add(data model.Book) (model.Book, error) {
	res, err := gc.Model.Insert(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}