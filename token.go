package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

// Token is a token for Google Apps APIs
type Token struct {
	fileName string
	filePath string
}

// NewToken is the constructor
func NewToken(fileName string) (*Token, error) {
	if len(fileName) == 0 {
		err := fmt.Errorf("You must specify a fileName parameter")
		return nil, err
	}

	token := &Token{
		fileName: fileName,
	}

	return token, nil
}

func (t *Token) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
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

func (t *Token) tokenCacheFile() (string, error) {
	path, err := profilePath()
	if err != nil {
		log.Fatalf("Unable to get profile path: %v", err)
		return "", err
	}

	if !isExist(path) {
		os.MkdirAll(path, 0700)
	}

	t.filePath = filepath.Join(path, url.QueryEscape(t.fileName))

	return t.filePath, err
}

func (t *Token) tokenFromFile() (*oauth2.Token, error) {
	cacheFile, err := t.tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return nil, err
	}

	f, err := os.Open(cacheFile)
	if err != nil {
		log.Fatalf("Unable to open path to cached credential file. %v", err)
		return nil, err
	}

	ot := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(ot)

	defer f.Close()

	return ot, err
}

func (t *Token) saveToken(token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", t.filePath)

	f, err := os.Create(t.filePath)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

func (t *Token) resetToken() {
	path, err := t.tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return
	}

	os.Remove(path)
}
