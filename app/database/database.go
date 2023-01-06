package database

import (
	"fmt"
	"log"

	"github.com/FillanL/creatturlinks/app/config"
	"github.com/FillanL/creatturlinks/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func setupModels(db *gorm.DB)(error){
	var err error

	log.Println("creating models....")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("creating models....")
	err = db.AutoMigrate(&models.Link{})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("db migrated")

	return nil
}

func SetUpDatabase(){
	var db *gorm.DB
	var err error

	log.Println("Connecting to database...")
	connectionString := config.Env.DBbUrl
	log.Println("Connection string:", connectionString)
	
	if config.Env.DBbUrl == "" {
		log.Println("Connectingstring empty...")
		connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Env.DBHost, config.Env.DBPort, config.Env.DBUser, config.Env.DBName, config.Env.DBPass)
		log.Println("Connecting to database...")
	}

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil{
		log.Fatalln("could not connect to postgres")
	}
	log.Println("connected to postgres")
	// log.Println("creating models....")
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Println(err)
	// }
	err = setupModels(db)
	if err != nil{
		log.Fatalln("could not migrate db")
	}
	models.SetManager(db)
	NewStorage()
}