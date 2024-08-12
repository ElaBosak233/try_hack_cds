package db

import (
	"fmt"
	"github.com/elabosak233/cloudsdale/internal/app/config"
	"github.com/elabosak233/cloudsdale/internal/app/logger/adapter"
	"github.com/elabosak233/cloudsdale/internal/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbInfo string

func Db() *gorm.DB {
	return db
}

// InitDatabase initializes the database connection and performs the necessary migrations.
func InitDatabase() {
	initDatabaseEngine()
	zap.L().Info(fmt.Sprintf("Database Connect Information: %s", dbInfo))
	db.Logger = adapter.NewGORMAdapter(zap.L())
	migrate()
	initAdmin()
	initDefaultCategories()
	selfCheck()
}

// Debug enables the debug mode of the database connection.
func Debug() {
	db = db.Debug()
}

// initDatabaseEngine initializes the database connection engine.
// It supports PostgreSQL, MySQL, and SQLite.
// The connection information is read from the configuration file.
// The connection information is formatted according to the database type.
// The connection is established using the GORM library.
// The database connection is stored in the global variable db.
// If an error occurs during the connection, the program will exit.
func initDatabaseEngine() {
	var err error
	switch config.AppCfg().DB.Provider {
	case "postgres":
		dbInfo = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.AppCfg().DB.Postgres.Host,
			config.AppCfg().DB.Postgres.Port,
			config.AppCfg().DB.Postgres.Username,
			config.AppCfg().DB.Postgres.Password,
			config.AppCfg().DB.Postgres.Dbname,
			config.AppCfg().DB.Postgres.Sslmode,
		)
		db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	case "mysql":
		dbInfo = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.AppCfg().DB.MySQL.Username,
			config.AppCfg().DB.MySQL.Password,
			config.AppCfg().DB.MySQL.Host,
			config.AppCfg().DB.MySQL.Port,
			config.AppCfg().DB.MySQL.Dbname,
		)
		db, err = gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	case "sqlite":
		dbInfo = config.AppCfg().DB.SQLite.Path
		db, err = gorm.Open(sqlite.Open(dbInfo), &gorm.Config{})
	}
	if err != nil {
		zap.L().Fatal("Database connection failed.", zap.Error(err))
	}
}

// migrate performs the necessary migrations.
// It creates the tables if they do not exist.
func migrate() {
	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Challenge{},
		&model.Team{},
		&model.UserTeam{},
		&model.Submission{},
		&model.Nat{},
		&model.Pod{},
		&model.Game{},
		&model.GameChallenge{},
		&model.GameTeam{},
		&model.Flag{},
		&model.Port{},
		&model.Nat{},
		&model.Env{},
		&model.Notice{},
		&model.Webhook{},
	)
	if err != nil {
		zap.L().Fatal("Database sync failed.", zap.Error(err))
	}
}

// selfCheck performs a self-check.
func selfCheck() {
	db.Exec("DELETE FROM pods")
}

func initAdmin() {
	var count int64
	db.Model(&model.User{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		zap.L().Warn("Administrator account does not exist, will be created soon.")
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		admin := model.User{
			Username: "admin",
			Nickname: "Administrator",
			Group:    "admin",
			Password: string(hashedPassword),
			Email:    "admin@admin.com",
		}
		err := db.Create(&admin).Error
		if err != nil {
			zap.L().Fatal("Super administrator account creation failed.", zap.Error(err))
			return
		}
		zap.L().Info("Super administrator account created successfully.")
	}
}

// initDefaultCategories initializes the default categories.
// If the categories do not exist, they will be created.
func initDefaultCategories() {
	var count int64
	db.Model(&model.Category{}).Count(&count)
	if count == 0 {
		zap.L().Warn("Categories do not exist, will be created soon.")
		defaultCategories := []model.Category{
			{
				Name:        "misc",
				Description: "misc",
				Color:       "#3F51B5",
				Icon:        "fingerprint",
			},
			{
				Name:        "web",
				Description: "web",
				Color:       "#009688",
				Icon:        "language",
			},
			{
				Name:        "pwn",
				Description: "pwn",
				Color:       "#673AB7",
				Icon:        "function",
			},
			{
				Name:        "crypto",
				Description: "crypto",
				Color:       "#607D8B",
				Icon:        "tag",
			},
			{
				Name:        "reverse",
				Description: "reverse",
				Color:       "#6D4C41",
				Icon:        "keyboard_double_arrow_left",
			},
		}
		err := db.Create(&defaultCategories).Error
		if err != nil {
			zap.L().Fatal("Category initialization failed.", zap.Error(err))
			return
		}
	}
}
