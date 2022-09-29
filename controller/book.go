package controller

import "Alterra/Project1-BE12-Book-Rent/model"

type BookControll struct {
	Model model.BookModel
}

func (bc BookControll) GetAll() ([]model.Book, error) {
	res, err := bc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BookControll) Add(data model.Book) (model.Book, error) {
	res, err := bc.Model.Insert(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) ShowBook() ([]model.Book, error) {
	res, err := bc.Model.ShowBook()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BookControll) Show(User_Id int) ([]model.Book, error) {
	res, err := bc.Model.Show(User_Id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BookControll) Add_New(data model.Book) (model.Book, error) {
	res, err := bc.Model.Input(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) Update(data model.Book) (model.Book, error) {
	res, err := bc.Model.Update(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) Delete(IdBook int, User_Id int) (model.Book, error) {
	res, err := bc.Model.Delete(IdBook, User_Id)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) UpdateBorrowed(data model.Book) (model.Book, error) {
	res, err := bc.Model.UpdateBorrowed(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) DeleteAccount(User_Id int) (model.Book, error) {
	res, err := bc.Model.DeleteAccount(User_Id)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}

func (bc BookControll) RentedBook(Rent_By int) ([]model.Book, error) {
	res, err := bc.Model.RentedBook(Rent_By)
	if err != nil {
		return nil, err
	}
	return res, nil
}
