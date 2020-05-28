package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Sensor struct {
	gorm.Model
	Temp float32 `gorm:"type:real"`
}

var gdb = initDB()

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", env.PgHost, env.PgPort, env.PgUser, env.PgDBName, env.PgPassword))
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&Sensor{}) {
		db.Debug().AutoMigrate(&Sensor{})
	}
	return db
}

func addTemp(t float32) error {
	sensor := Sensor{Temp: t}

	if err := gdb.Create(&sensor).Error; err != nil {
		return err
	}
	return nil
}

func getLastTemp() (Sensor, error) {
	sensor := Sensor{}

	if err := gdb.Last(&sensor).Error; err != nil {
		return Sensor{}, err
	}
	return sensor, nil
}
