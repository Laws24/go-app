package user

import (
	"github.com/BoyYangZai/go-server-lib/pkg/database"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	//Birthday time.Time `json:"birthday"`
	Email string
}

func Run() {
	db := database.Db
	customer := Customer{
		Name: "ydh", Age: 201, Email: "666@66.com",
	}
	db.Create(&customer)

	println("scucess")
}
