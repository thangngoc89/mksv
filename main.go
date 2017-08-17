package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/qor"
	"github.com/qor/admin"
)

// User Model
type User struct { 
	gorm.Model
	Name string
}

// Product Model
type Product struct { 
	gorm.Model
	Name string
	Description string
}

func main() {
	DB, err := gorm.Open("sqlite3", "demo.db")
  if err != nil {
    panic("failed to connect database")
  }
  defer DB.Close()

	DB.AutoMigrate(&User{}, &Product{})

	Admin := admin.New(&qor.Config{DB: DB})

	// Create resources from GORM-backend model 
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// Register route
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe("0.0.0.0:9000", mux)

}
