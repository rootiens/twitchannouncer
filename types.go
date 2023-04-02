package main

import "time"

type HttpHeaders struct {
	Key   string
	Value string
}

type HttpReq struct {
	Method  string
	Headers []HttpHeaders
	Url     string
}

type TokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type StreamerRequest struct{}

type StreamData struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserLogin    string    `json:"user_login"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	GameName     string    `json:"game_name"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
	TagIDs       []string  `json:"tag_ids"`
	Tags         []string  `json:"tags"`
	IsMature     bool      `json:"is_mature"`
}

type StreamResponse struct {
	Data       []StreamData `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type StreamerData struct {
	TwitchName     string `json:"twitch_login"`
	DiscordChannel int64  `json:"discord_channel"`
	State          string `json:"state"`
	OfflineAt      string `json:"offline_at"`
}

type Streamers struct {
	Streamers []StreamerData `json:"streamers"`
}
