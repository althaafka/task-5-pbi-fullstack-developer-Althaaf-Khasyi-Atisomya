package database
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/models"
)

var DB *gorm.DB;

func Connect() {
	database, err := gorm.Open(mysql.Open("root:qwerty@tcp(localhost:3306)/godb"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})

	DB = database
}