package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kgermando/lobi/config"
	"github.com/kgermando/lobi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Silent, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  ParameterizedQueries:      true,           // Don't include params in the SQL log
		  Colorful:                  false,          // Disable color
		},
	  )
	  
	var err error
	p := config.ConfigEnv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.ConfigEnv("DB_HOST"), port, config.ConfigEnv("DB_USER"), config.ConfigEnv("DB_PASSWORD"), config.ConfigEnv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")




	DB.AutoMigrate(
		&models.User{},
	)

}
