package orm

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google-service/models"
	"google-service/repository"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"
)

func GoogleRepo() repository.GoogleRepo {
	return &ormGoogleRepo{}
}

type ormGoogleRepo struct{}

func (o ormGoogleRepo) CopyTable(ctx context.Context, s *models.Spreadsheet) (*models.Spreadsheet, error) {
	srv := SetupGoogle()

	rb := &sheets.CopySheetToAnotherSpreadsheetRequest{
		DestinationSpreadsheetId: "167wX9gl7QWHJCQMi8oPCCU7t8v14y6OPZEye0y3YDbk",
	}

	res, err := srv.Spreadsheets.Sheets.CopyTo("16vNIioTnNfmdd0clMvgjzIkpGbABtqclnHB4Acd3_08", int64(0), rb).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	fmt.Println(res)

	return nil, nil
}

func (o ormGoogleRepo) CreateTable(ctx context.Context) (*models.Spreadsheet, error) {
	srv := SetupGoogle()

	rb := &sheets.Spreadsheet{
		// TODO: Add desired fields of the request body.
	}

	res, err := srv.Spreadsheets.Create(rb).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return &models.Spreadsheet{URL: res.SpreadsheetUrl, Token: res.SpreadsheetId}, nil
}

func (o ormGoogleRepo) ClearData(ctx context.Context, s *models.SpreadsheetClear) error {
	srv := SetupGoogle()

	spreadsheetId := s.Token

	clearRange := "A1:Z100000000"

	rb := &sheets.ClearValuesRequest{
		// TODO: Add desired fields of the request body.
	}

	_, err := srv.Spreadsheets.Values.Clear(spreadsheetId, clearRange, rb).Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (o ormGoogleRepo) AppendData(ctx context.Context, s *models.Spreadsheet) error {
	srv := SetupGoogle()

	spreadsheetId := s.Token

	writeRange := "A2"

	var vr sheets.ValueRange

	for _, v := range s.Data {
		vr.Values = append(vr.Values, v)
	}

	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		panic(err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		panic(err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("%v", err)
	}
	json.NewEncoder(f).Encode(token)
}
