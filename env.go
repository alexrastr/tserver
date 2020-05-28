package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
)

var env = loadEnv()

//системные переменные
type Env struct {
	GinMode string

	PgUser     string
	PgPassword string
	PgHost     string
	PgPort     string
	PgDBName   string

	MqttUser     string
	MqttPassword string
	MqttHost     string
	MqttPort     string
}

func loadEnv() *Env {
	e := &Env{
		//gin
		GinMode: os.Getenv("GIN_MODE"),
		//postgres
		PgUser:     os.Getenv("PG_USER"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgHost:     os.Getenv("PG_HOST"),
		PgPort:     os.Getenv("PG_PORT"),
		PgDBName:   os.Getenv("PG_DBNAME"),
		//MQTT
		MqttUser:     os.Getenv("MQTT_USER"),
		MqttPassword: os.Getenv("MQTT_PASSWORD"),
		MqttHost:     os.Getenv("MQTT_HOST"),
		MqttPort:     os.Getenv("MQTT_PORT"),
	}

	fmt.Println("=== LOAD ENVIRONMENT VARIABLES ===")
	e.print()
	fmt.Println("==================================")

	return e
}

//вывод всех переменных из структуры
func (e *Env) print() {
	v := reflect.ValueOf(e).Elem()
	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varValue := v.Field(i).Interface()
		//сокрытие паролей
		//если в названии переменной есть пароль то скрыть значение
		matched, _ := regexp.Match("(Password|Token)$", []byte(varName))
		if matched && varValue != "" {
			varValue = "************"
		}
		fmt.Printf("%v %v %v\n", varName, "=", varValue)
	}
}
