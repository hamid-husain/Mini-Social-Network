package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"project/config"

	_ "project/db/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(len(args))
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
		config.AppConfig.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	ctx := context.Background()
	migrationsDir := "/home/ubuntu/Desktop/project/db/migrations"
	if err := goose.RunContext(ctx, command, db, migrationsDir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
