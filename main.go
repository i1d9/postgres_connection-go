package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/i1d9/postgres_connection-go/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {

	viper.AddConfigPath(".") // look for config in the working directory

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	database_url := viper.GetString("DATABASE_URL")

	// Create a new connection pool
	dbpool, err := pgxpool.New(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// models.InitiateDatabase(dbpool, context.Background())

	users, err := models.GetAllUsers(dbpool)

	if err != nil {
		log.Fatalf("Error fetching users: %v\n", err)
	}

	for i, user := range users {
		fmt.Printf("User %d:\n", i+1)
		fmt.Printf("  ID: %d\n", user.ID)

		fmt.Printf("  Email: %s\n", models.NullableString(user.Email))
		fmt.Printf("  Verified: %s\n", models.NullableBool(user.Verified))
		fmt.Printf("  Active: %s\n", models.NullableBool(user.Active))
		fmt.Printf("  Name: %s %s %s\n", models.NullableString(user.FirstName), models.NullableString(user.LastName), models.NullableString(user.Surname))
		fmt.Printf("  Mobile: %s (Verified: %s)\n", models.NullableString(user.MobileNumber), models.NullableBool(user.MobileVerified))
		fmt.Printf("  Gender: %s\n", models.NullableString(user.Gender))

		fmt.Printf("  Created At: %s\n", models.NullableTime(user.CreatedAt))
		fmt.Printf("  Updated At: %s\n", models.NullableTime(user.UpdatedAt))
		fmt.Printf("  Last Login At: %s\n", models.NullableTime(user.LastLoginAt))
		fmt.Printf("  Deleted At: %s\n", models.NullableTime(user.DeletedAt))

		fmt.Println("  ---")
	}

	

}
