package main

import (
	"clockify/config"
	"clockify/project/application"
	entity2 "clockify/project/infrastructure/entity"
	mysql2 "clockify/project/infrastructure/mysql"
	"clockify/project/presentation/projecthttp"
	"clockify/project/presentation/projectrouter"
	application3 "clockify/timeEntry/application"
	entity3 "clockify/timeEntry/infrastructure/entity"
	mysql3 "clockify/timeEntry/infrastructure/mysql"
	"clockify/timeEntry/presentation/timehttp"
	"clockify/timeEntry/presentation/timerouter"
	application2 "clockify/users/application"
	"clockify/users/infrastructure/entity"
	"clockify/users/infrastructure/mysql"
	"clockify/users/presentation/router"
	"clockify/users/presentation/userhttp"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

// @title clockify project
// @version 1.0
// @description clockify project for time-entry

// @host localhost:1212
// @basePath /api

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Connect to the database
	db := config.Connect()
	fmt.Println(db)
	// Migrate User table
	err = migrateUserTable(db)
	if err != nil {
		log.Fatalf("Error migrating User table: %v", err)
	}

	// Migrate Projects table
	err = migrateProjectsTable(db)
	if err != nil {
		log.Fatalf("Error migrating Projects table: %v", err)
	}
	// Migrate TimeEntry Table
	err = migrateTimeEntryTable(db)
	if err != nil {
		log.Fatalf("Error migrating TimeEntry table: %v", err)
	}

	// Initialize User repository
	userRepo := mysql.NewUserEpoImpl(db)

	// Initialize User service with the JWT secret key
	jwtSecretKey := []byte("your_secret_key")
	userService := application2.NewUserServiceImp(userRepo, jwtSecretKey)

	// Initialize User controller
	userController := userhttp.NewUserController(userService)

	// Initialize Project repository
	projectRepo := mysql2.NewProjectRepoImpl(db)

	// Initialize Project service
	projectService := application.NewProjectServiceImpl(projectRepo, userRepo)

	// Initialize Project controller
	projectController := projecthttp.NewProjectController(projectService)

	// Initialize TimeEntry repository
	timeRepo := mysql3.NewTimeEntryRepoImp(db)

	// Initialize TimeEntry service
	timeService := application3.NewTimeEntryServiceImpl(timeRepo, userRepo)

	// Initialize TimeEntry controller
	timeController := timehttp.NewTimeEntryController(timeService)

	// Setup routes for user, project, and TimeEntry
	userRoutes := router.NewRouter(*userController)
	fmt.Println(userRoutes)
	authRouter := router.NewRouter(*userController)
	fmt.Println(authRouter)
	projectRoutes := projectrouter.ProjectRouter(*projectController)
	timeRoutes := timerouter.TimeEntryRouter(*timeController)

	// Run User HTTP server in a Goroutine
	go func() {
		if err := userRoutes.Run(":1212"); err != nil {
			log.Fatalf("Failed to start User server: %v", err)
		}
	}()

	go func() {
		if err := projectRoutes.Run(":1213"); err != nil {
			log.Fatalf("Failed to start Project server: %v", err)
		}
	}()

	go func() {
		if err := timeRoutes.Run(":1214"); err != nil {
			log.Fatalf("Failed to start TimeEntry server: %v", err)
		}
	}()

	select {}
}

func migrateUserTable(db *gorm.DB) error {
	// Check if the table already exists before attempting to create it
	if !db.Migrator().HasTable(&entity.User{}) {
		err := db.AutoMigrate(&entity.User{})
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateProjectsTable(db *gorm.DB) error {
	// Check if the table already exists before attempting to create it
	if !db.Migrator().HasTable(&entity2.Projects{}) {
		err := db.AutoMigrate(&entity2.Projects{})
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateTimeEntryTable(db *gorm.DB) error {
	if !db.Migrator().HasTable(&entity3.TimeEntry{}) {
		err := db.AutoMigrate(&entity3.TimeEntry{})
		if err != nil {
			return err
		}
	}
	return nil
}
