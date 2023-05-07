package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

var postgres_url string

func init() {
	viper.SetConfigFile("configs/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig()failed,err:%v\n", err)
		return
	}

	postgres_url = viper.GetString("postgresql.url")
}

func main() {
	m, err := migrate.New(
		"file://deployment/migrations",
		postgres_url,
	)
	if err != nil {
		fmt.Printf("[error]: migrate.New - deployment/migration: %v", err)
	}

	if err := m.Down(); err != nil {
		fmt.Printf("[error]: m.Down() - deployment/migration: %v", err)
	}
}
