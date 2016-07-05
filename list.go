package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/api/admin/directory/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

func main() {
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

	r, err := srv.Users.List().Customer("my_customer").MaxResults(10).OrderBy("email").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain. %v", err)
	}

	count := len(r.Users)
	if count == 0 {
		fmt.Print("No users found.\n")
	} else {
		// fmt.Printf("Users[%v]:\n", count)
		for _, u := range r.Users {
			fmt.Printf("%s (%s)\n", u.PrimaryEmail, u.Name.FullName)
		}
	}
}
