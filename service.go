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
	directory *ad.Service
	reports   *ar.Service
}

func NewService() (*Service, error) {
	b, err := getClientSecret()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	directoryService, err := newDirectoryServiceWithSecret(b)
	if err != nil {
		return nil, err
	}

	reportsService, err := newReportsServiceWithSecret(b)
	if err != nil {
		return nil, err
	}

	service := &Service{
		directory: directoryService,
		reports:   reportsService,
	}

	return service, err
}

func newDirectoryServiceWithSecret(bytes []byte) (*ad.Service, error) {
	config, err := google.ConfigFromJSON(bytes,
		ad.AdminDirectoryUserScope,
		ad.AdminDirectoryGroupScope,
		ad.AdminDirectoryGroupMemberScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client := getClient(ctx, config)
	srv, err := ad.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory Client %v", err)
		return nil, err
	}

	return srv, err
}

func newReportsServiceWithSecret(bytes []byte) (*ar.Service, error) {
	config, err := google.ConfigFromJSON(bytes,
		ar.AdminReportsAuditReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client := getClient(ctx, config)
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

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return nil
	}

	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}

	return config.Client(ctx, tok)
}
