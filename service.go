package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	ad "google.golang.org/api/admin/directory/v1"
	ar "google.golang.org/api/admin/reports/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Service is a wrapper struct for Google Apps Admin APIs
type Service struct {
	directory      *ad.Service
	directoryToken *Token
	reports        *ar.Service
	reportsToken   *Token
}

// NewService is the constructor
func NewService() (*Service, error) {
	b, err := getClientSecret()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	directoryToken, err := NewToken("directory-token.json")
	if err != nil {
		log.Fatalf("Unable to create token object: %v", err)
		return nil, err
	}

	directoryService, err := newDirectoryServiceWithSecret(b, directoryToken)
	if err != nil {
		return nil, err
	}

	reportsToken, err := NewToken("reports-token.json")
	if err != nil {
		log.Fatalf("Unable to create token object: %v", err)
		return nil, err
	}

	reportsService, err := newReportsServiceWithSecret(b, reportsToken)
	if err != nil {
		return nil, err
	}

	service := &Service{
		directory:      directoryService,
		directoryToken: directoryToken,
		reports:        reportsService,
		reportsToken:   reportsToken,
	}

	return service, err
}

func newDirectoryServiceWithSecret(bytes []byte, token *Token) (*ad.Service, error) {
	config, err := google.ConfigFromJSON(bytes,
		ad.AdminDirectoryUserScope,
		ad.AdminDirectoryGroupScope,
		ad.AdminDirectoryGroupMemberScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client := getClient(ctx, config, token)
	srv, err := ad.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory Client %v", err)
		return nil, err
	}

	return srv, err
}

func newReportsServiceWithSecret(bytes []byte, token *Token) (*ar.Service, error) {
	config, err := google.ConfigFromJSON(bytes,
		ar.AdminReportsAuditReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client := getClient(ctx, config, token)
	srv, err := ar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve reports client %v", err)
		return nil, err
	}

	return srv, err
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func profilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	path := filepath.Join(usr.HomeDir, ".google-apps-utils")
	if !isExist(path) {
		os.MkdirAll(path, 0700)
	}

	return path, nil
}

func getClientSecret() ([]byte, error) {
	path, err := profilePath()
	if err != nil {
		log.Fatalf("Unable to get profile path: %v", err)
		return nil, err
	}

	jsonPath := filepath.Join(path, "client_secret.json")

	b, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	return b, err
}

func getClient(ctx context.Context, config *oauth2.Config, token *Token) *http.Client {
	tok, err := token.tokenFromFile()
	if err != nil {
		tok = token.getTokenFromWeb(config)
		token.saveToken(tok)
	}

	return config.Client(ctx, tok)
}
