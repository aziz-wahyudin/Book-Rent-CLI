package main

import (
	"Alterra/Project1-BE12-Book-Rent/controller"
	"Alterra/Project1-BE12-Book-Rent/model"
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/Book_Rent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Rent{})
}

func main() {
	gconn, err := connectGorm()
	migrate(gconn)
	if err != nil {
		fmt.Println("cannot connect to DB", err.Error())
	}
	var isRunning bool = true
	var inputMenu int
	var loginSession int
	userMdl := model.UserModel{DB: gconn}
	userCtl := controller.UserControll{Model: userMdl}
	bookMdl := model.BookModel{DB: gconn}
	bookCtl := controller.BookControll{Model: bookMdl}
	rentMdl := model.RentModel{DB: gconn}
	rentCtl := controller.RentControll{Model: rentMdl}

	for isRunning {
		fmt.Println("\t--Menu--")
		fmt.Println("1. Tampilkan Semua Buku")
		fmt.Println("2. Login Atau Registrasi")
		fmt.Println("3. Exit")
		fmt.Println("Masukkan Input: ")
		fmt.Scanln(&inputMenu)
		callClear()

		switch inputMenu {
		case 1:
			var isRunning2 bool = true

			for isRunning2 {
				var subMenu int
				fmt.Println("--Submenu Tampilkan Semua Buku--")
				fmt.Println("============================")
				fmt.Println("1. Pencarian Buku")
				fmt.Println("2. Daftar Buku")
				fmt.Println("3. Exit")
				fmt.Println("Masukkan Input: ")
				fmt.Scanln(&subMenu)
				callClear()

				switch subMenu {
				case 1:
					var searchingBuku model.Book
					fmt.Println("Pencarian Buku")
					fmt.Println("============================")
					fmt.Println("Ketik judul buku yang ingin dicari: ")
					fmt.Scanln(&searchingBuku.Name)
					result, _ := bookCtl.Searching(searchingBuku.Name)
					if len(result) == 0 {
						fmt.Println("Buku tidak ditemukan")
					} else {
						for i := 0; i < len(result); i++ {
							fmt.Println(result[i].IdBook, result[i].Name, result[i].Status)
						}
					}
				case 2:
					fmt.Println("Daftar Buku")
					fmt.Println("============================")
					fmt.Println("")
					res, _ := bookCtl.ShowBook()
					for i := 0; i < len(res); i++ {
						fmt.Println(res[i].IdBook, res[i].Name, res[i].Status)
					}
				case 3:
					isRunning2 = false
					callClear()
				}

			}

		case 2:
			var menuDua bool = true
			for menuDua {
				fmt.Println("Menu Login Atau Registrasi")
				fmt.Println("============================")
				fmt.Println("1. Registrasi")
				fmt.Println("2. Login")
				fmt.Println("3. Exit")
				fmt.Scanln(&inputMenu)
				callClear()
				switch inputMenu {
				case 1:
					fmt.Println("Menu Registrasi")
					fmt.Println("============================")
					var userBaru model.User
					fmt.Println("Masukkan Nama: ")
					fmt.Scanln(&userBaru.Name)
					fmt.Println("Masukkan Email: ")
					fmt.Scanln(&userBaru.Email)
					fmt.Println("Masukkan Password: ")
					fmt.Scanln(&userBaru.Password)

					res, err := userCtl.Add(userBaru)
					if err != nil {
						fmt.Println("some error on add", err.Error())
					}
					fmt.Println("sukses membuat akun", res)
					fmt.Println("silakan melakukan login")
				case 2:
					fmt.Println("Menu Login")
					fmt.Println("============================")
					var Email string
					var Password string
					fmt.Println("Masukkan Email: ")
					fmt.Scanln(&Email)
					fmt.Println("Masukkan Password: ")
					fmt.Scanln(&Password)

					res, _ := userCtl.Find(Email, Password)
					if len(res) == 0 {
						fmt.Println("Gagal login. Email atau Password salah.")
						fmt.Println("Atau pastikan anda sudah punya akun")
					} else {
						loginSession = res[0].User_Id
						fmt.Println("login sukses")
					}
					if len(res) != 0 {
						var menuLogin bool = true
						for menuLogin {
							fmt.Println("==============")
							fmt.Println("1. Tampilkan Buku Saya")
							fmt.Println("2. Tambah Koleksi Buku")
							fmt.Println("3. Ubah Informasi Buku")
							fmt.Println("4. Hapus Buku")
							fmt.Println("5. Pinjam Buku")
							fmt.Println("6. Buku Pinjaman Saya")
							fmt.Println("7. Kembalikan Buku")
							fmt.Println("8. Update Akun")
							fmt.Println("9. Hapus Akun")
							fmt.Println("10. Exit")
							fmt.Scanln(&inputMenu)
							callClear()
							switch inputMenu {
							case 1:
								mybook, err := bookCtl.Show(loginSession)
								if err != nil {
									fmt.Println("some error on show", err.Error())
								}
								if len(mybook) == 0 {
									fmt.Println("Anda tidak punya buku")
								}
								if len(mybook) != 0 {
									for i := 0; i < len(mybook); i++ {
										fmt.Println(mybook[i].IdBook, mybook[i].Name, mybook[i].Status)
									}
								}
							case 2:
								fmt.Println("Tambah koleksi buku")
								var bukuBaru model.Book
								fmt.Println("Judul buku: ")
								fmt.Scanln(&bukuBaru.Name)
								bukuBaru.Owner = loginSession
								bukuBaru.User_Id = loginSession
								bukuBaru.Status = "tersedia"

								addedbook, err := bookCtl.Add_New(bukuBaru)
								if err != nil {
									fmt.Println("some error on add", err.Error())
								} else {
									fmt.Println("Sukses menambahkan buku", addedbook)
								}
							case 3:
								fmt.Println("Update info buku")
								var updateBuku model.Book
								var inputkode int
								fmt.Println("Pilih kode buku yang ingin anda ubah")
								fmt.Scanln(&inputkode)
								fmt.Println("Judul buku baru: ")
								fmt.Scanln(&updateBuku.Name)
								updateBuku.IdBook = inputkode
								updateBuku.User_Id = loginSession
								updateBuku.Owner = loginSession
								updateBuku.Status = "tersedia"

								update, err := bookCtl.Update(updateBuku)
								if err != nil {
									fmt.Println("some error on update", err.Error())
								} else {
									fmt.Println("Sukses mengubah buku", update)
								}
							case 4:
								var kodeBuku int
								fmt.Println("Pilih kode buku yang ingin dihapus: ")
								fmt.Scanln(&kodeBuku)
								deletebook, err := bookCtl.Delete(kodeBuku, loginSession)
								if err != nil {
									fmt.Println("some error on delete", err.Error())
								} else {
									fmt.Println("Sukses menghapus buku", deletebook)
								}
							case 5:
								//menampilkan daftar semua buku di sistem
								fmt.Println("Daftar buku yang dapat Anda pinjam")
								availBook, _ := bookCtl.AvailableBook(loginSession)
								if len(availBook) == 0 {
									fmt.Println("Tidak ada buku yang dapat Anda pinjam")
								} else {
									for i := 0; i < len(availBook); i++ {
										fmt.Println(availBook[i].IdBook, availBook[i].Name)
									}
								}
								//meminjam buku
								var bukuIncaran model.Rent
								fmt.Println("Masukkan kode buku yang ingin dipinjam: ")
								fmt.Scanln(&bukuIncaran.IdBook)
								bukuIncaran.User_Id = loginSession

								borrowingBook, err := rentCtl.Add(bukuIncaran)
								if err != nil {
									fmt.Println("some error on borrowing a book", err.Error())
								} else {
									fmt.Println("Sukses meminjam buku", borrowingBook)
								}
								//update status buku
								var updatePinjam model.Book
								updatePinjam.IdBook = bukuIncaran.IdBook
								updatePinjam.Status = "tidak tersedia"
								updatePinjam.Rent_By = loginSession
								borrowedBook, err := bookCtl.UpdateBorrowed(updatePinjam)
								if err != nil {
									fmt.Println("some error on update", err.Error())
								} else {
									fmt.Println("Sukses mengubah buku", borrowedBook)
								}
							case 6:
								fmt.Println("Buku yang saya pinjam:")
								myrent, err := bookCtl.RentedBook(loginSession)
								if err != nil {
									fmt.Println("some error on rentedBook", err.Error())
								}
								if len(myrent) == 0 {
									fmt.Println("Anda tidak meminjam buku")
								}
								if len(myrent) != 0 {
									for i := 0; i < len(myrent); i++ {
										fmt.Println(myrent[i].IdBook, myrent[i].Name)
									}
								}
							case 7:
								// buku pinjaman
								fmt.Println("Buku yang saya pinjam:")
								myrent, err := bookCtl.RentedBook(loginSession)
								if err != nil {
									fmt.Println("some error on rentedBook", err.Error())
								}
								if len(myrent) == 0 {
									fmt.Println("Anda tidak meminjam buku")
								}
								if len(myrent) != 0 {
									for i := 0; i < len(myrent); i++ {
										fmt.Println(myrent[i].IdBook, myrent[i].Name)
									}
								}
								// mengembalikan buku
								fmt.Println("Pengembalian buku pinjaman")
								fmt.Println("Pilih kode buku yang ingin dikembalikan:")
								var kembalikanBuku model.Rent
								fmt.Scanln(&kembalikanBuku.IdBook)
								kembalikanBuku.Return_date = time.Now()
								// update di rents
								returnBook, err := rentCtl.ReturnBook(kembalikanBuku)
								if err != nil {
									fmt.Println("some error on update", err.Error())
								} else {
									fmt.Println("Sukses mengembalikan buku 1/2", returnBook)
								}
								// update di books
								var updateKembali model.Book
								updateKembali.IdBook = kembalikanBuku.IdBook
								updateKembali.Status = "tersedia"
								updateKembali.Rent_By = 0
								returnedBook, err := bookCtl.UpdateBorrowed(updateKembali)
								if err != nil {
									fmt.Println("some error on update", err.Error())
								} else {
									fmt.Println("Sukses mengembalikan buku 2/2", returnedBook)
								}
							case 8:
								fmt.Println("Update info akun")
								var updateAkun model.User
								fmt.Println("Nama baru:")
								fmt.Scanln(&updateAkun.Name)
								fmt.Println("Email baru:")
								fmt.Scanln(&updateAkun.Email)
								fmt.Println("Password baru:")
								fmt.Scanln(&updateAkun.Password)
								updateAkun.User_Id = loginSession

								newAccount, err := userCtl.Update(updateAkun)
								if err != nil {
									fmt.Println("some error on update", err.Error())
								} else {
									fmt.Println("Sukses mengubah info akun", newAccount)
								}
							case 9:
								fmt.Println("Apakah anda yakin untuk menghapus akun?")
								fmt.Println("jika ya, ketik y")
								fmt.Println("jika tidak, ketik n")
								var delConfirm string
								fmt.Scanln(&delConfirm)
								switch delConfirm {
								case "y":
									// menghapus rent data
									delRentAccount, err := rentCtl.DeleteAccount(loginSession)
									if err != nil {
										fmt.Println("some error on delete", err.Error())
									} else {
										fmt.Println("Sukses menghapus 1/3", delRentAccount)
									}
									// menghapus book data
									delBookAccount, err := bookCtl.DeleteAccount(loginSession)
									if err != nil {
										fmt.Println("some error on delete", err.Error())
									} else {
										fmt.Println("Sukses menghapus 2/3", delBookAccount)
									}
									// menghapus user data
									delUserAccount, err := userCtl.DeleteAccount(loginSession)
									if err != nil {
										fmt.Println("some error on delete", err.Error())
									} else {
										fmt.Println("Sukses menghapus 3/3", delUserAccount)
									}
								case "n":
									break
								}
								menuLogin = false
							case 10:
								callClear()
								menuLogin = false
							}
						}
					}
				case 3:
					callClear()
					menuDua = false
				}
			}
		case 3:
			callClear()
			isRunning = false
		}
	}
}
