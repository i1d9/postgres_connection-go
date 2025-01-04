package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/i1d9/postgres_connection-go/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"strconv"
)

func createDatabaseConnectionPool() (*pgxpool.Pool, error) {
	// Create a new connection pool
	database_url := viper.GetString("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbpool, err

}

func loadConfig() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func getConsoleInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text1, _ := reader.ReadString('\n')
	text1 = strings.Replace(text1, "\n", "", -1)
	return text1

}

func renderUsers(users []models.User) {
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

		fmt.Println("  ---\n")
	}
}

func listUsers(dbpool *pgxpool.Pool) {

	users, err := models.GetAllUsers(dbpool)

	if err != nil {
		log.Fatalf("Error fetching users: %v\n", err)
	}

	renderUsers(users)

}

func main() {

	loadConfig()
	dbpool, _ := createDatabaseConnectionPool()
	defer dbpool.Close()
	for {
		fmt.Println("\nUser Management CLI")
		fmt.Println("---------------------")

		mainInput := getConsoleInput("1. List\n2. Create\n3. Delete\n4. Search\n5. Exit\n-> ")

		switch mainInput {
		case "1":
			listUsers(dbpool)
		case "2":
			

			firstName := getConsoleInput("Enter First Name: ")
			lastName := getConsoleInput("Enter Last Name: ")
			surname := getConsoleInput("Enter Surname: ")
			email := getConsoleInput("Enter Email: ")
			mobileNumber := getConsoleInput("Enter Mobile Number: ")
			gender := getConsoleInput("Enter Gender: ")

			user := models.User{
				FirstName:    &firstName,
				LastName:     &lastName,
				Surname:      &surname,
				Email:        &email,
				MobileNumber: &mobileNumber,
				Gender:       &gender,
			}

			err := models.CreateUser(dbpool, user)
			if err != nil {
				log.Fatalf("Error creating user: %v\n", err)
			} else {
				fmt.Println("User created successfully")
			}
		case "3":

			userID := getConsoleInput("\nEnter the User ID to be deleted\n-> ")
			user_id, err := strconv.Atoi(userID)
			if err != nil {
				log.Fatalf("Error searching for users: %v\n", err)
			}
			result := models.DeleteUser(dbpool, user_id)

			if result != nil {
				log.Fatalf("Error deleting the specified user: %v\n", result)
			}else
			{
				fmt.Println("User Deleted Successfully")
			}

		case "4":
			searchTerm := getConsoleInput("\nSearch.\nWho are you looking for?\n-> ")

			users, err := models.SearchUsers(dbpool, searchTerm)

			if err != nil {
				log.Fatalf("Error searching for users: %v\n", err)
			}
			fmt.Printf("\nFound %d users\n", len(users))
			renderUsers(users)

		case "5":
			os.Exit(0)
		case "q":
			os.Exit(0)
		}

	}

}
