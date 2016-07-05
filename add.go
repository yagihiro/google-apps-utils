package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/api/admin/directory/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

func usage() {
	fmt.Println("")
	fmt.Println("add.go -- add user on Google Apps")
	fmt.Println("")
	fmt.Println("go run add.go GIVENNAME FAMILYNAME PRIMARYEMAIL")
	fmt.Println("")
	fmt.Println("    GIVENNAME    : Given Name")
	fmt.Println("    FAMILYNAME   : Family Name")
	fmt.Println("    PRIMARYEMAIL : email address, the format is given.family@recruit-rsc.io")
	fmt.Println("")
}

func randomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func main() {
	if (len(os.Args) - 1) < 3 {
		usage()
		os.Exit(0)
	}
	given := os.Args[1]
	family := os.Args[2]
	email := os.Args[3]
	password := randomString(8)
	log.Printf("Given:%v, Family:%v, Email:%v, Password:%v", given, family, email, password)

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, admin.AdminDirectoryUserScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	ctx := context.Background()
	client := GetClient(ctx, config)
	srv, err := admin.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory Client %v", err)
	}

	name := &admin.UserName{
		GivenName:  given,
		FamilyName: family,
	}
	user := &admin.User{
		Name:                      name,
		Password:                  password,
		PrimaryEmail:              email,
		ChangePasswordAtNextLogin: true,
	}

	user2, err := srv.Users.Insert(user).Do()
	if err != nil {
		log.Fatalf("Cannot create user in domain. %v", err)
	} else {
		log.Fatalf("Succeed to create user: %v", user2)
	}

}
