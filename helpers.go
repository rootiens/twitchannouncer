package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	GetTokenURL        = "https://id.twitch.tv/oauth2/token"
	GetStreamerDetails = "https://api.twitch.tv/helix/streams?user_login="
)

var AuthToken string
var TokenExpireAt time.Time

func CheckToken() {
	isPassed := time.Now().After(TokenExpireAt)
	if AuthToken == "" || isPassed {
		if err := GetBearerToken(); err != nil {
			panic(err)
		}
		fmt.Println("obtained/refreshed token at: ", time.Now())
	}
}

func GetBearerToken() error {
	Client_ID := os.Getenv("TWITCH_CLIENT_ID")
	Client_Secret := os.Getenv("TWITCH_SECRET")
	reqBody := TokenRequest{
		ClientId:     Client_ID,
		ClientSecret: Client_Secret,
		GrantType:    "client_credentials",
	}

	headers := HttpReq{
		Url:    GetTokenURL,
		Method: "POST",
		Headers: []HttpHeaders{
			HttpHeaders{
				Key:   "Content-Type",
				Value: "application/json",
			},
		},
	}

	resp, err := httpRequest(reqBody, headers)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	var response TokenResponse
	json.Unmarshal(body, &response)

	duration := time.Second * time.Duration(response.ExpiresIn)
	expires := time.Now().Add(duration)

	AuthToken = response.AccessToken
	TokenExpireAt = expires

	return nil
}

func IsStreamerOnline(streamer string) (bool, error) {
	reqBody := StreamerRequest{}

	Client_ID := os.Getenv("TWITCH_CLIENT_ID")

	headers := HttpReq{
		Url:    GetStreamerDetails + streamer,
		Method: "GET",
		Headers: []HttpHeaders{
			HttpHeaders{Key: "Content-Type", Value: "application/json"},
			HttpHeaders{Key: "Authorization", Value: "Bearer " + AuthToken},
			HttpHeaders{Key: "Client-Id", Value: Client_ID},
		},
	}

	resp, err := httpRequest(reqBody, headers)

	defer resp.Body.Close()

	if err != nil {
		return false, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, errors.New(string(body))
	}

	var response StreamResponse
	json.Unmarshal(body, &response)

	if len(response.Data) > 0 {
		if response.Data[0].Type == "live" {
			return true, nil
		}
	}

	return false, nil
}

func httpRequest(body interface{}, headers HttpReq) (*http.Response, error) {
	parsedBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(headers.Method, headers.Url, bytes.NewBuffer(parsedBody))

	for _, header := range headers.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetStreamers() (Streamers, error) {
	file, err := os.ReadFile("data.json")
	if err != nil {
		return Streamers{}, err
	}
    
    data, _ := io.ReadAll(bytes.NewBuffer(file))

	var streamers Streamers
    err = json.Unmarshal(data, &streamers)
    if err != nil{
        return Streamers{}, err
    }

	return streamers, nil
}
