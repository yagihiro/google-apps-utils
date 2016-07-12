package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"google.golang.org/api/admin/directory/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// RandomString は初期パスワード等で利用することを想定した length 長の文字列を返す関数です
func RandomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// GetService は admin.Service をよしなに生成して返す関数です
func GetService() (*admin.Service, error) {
	b, err := getClientSecret()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	config, err := google.ConfigFromJSON(b,
		admin.AdminDirectoryUserScope,
		admin.AdminDirectoryGroupScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client := getClient(ctx, config)
	srv, err := admin.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve directory Client %v", err)
		return nil, err
	}

	return srv, err
}

// ResetToken はローカルにキャッシュ済みの oauth トークンを削除する関数です
func ResetToken() {
	path, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return
	}

	if err := os.Remove(path); err != nil {
		log.Fatalf("Unable to remove cached credential file. %v", err)
		return
	}
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

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then tyep the authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return tok
}

func tokenCacheFile() (string, error) {
	path, err := profilePath()
	if err != nil {
		log.Fatalf("Unable to get profile path: %v", err)
		return "", err
	}

	if !isExist(path) {
		os.MkdirAll(path, 0700)
	}

	return filepath.Join(path, url.QueryEscape("token.json")), err
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)

	defer f.Close()

	return t, err
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)

	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()

	json.NewEncoder(f).Encode(token)
}
