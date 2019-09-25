package slack

import (
	"context"
	"net/http"
)

type API struct {
	URL    string
	Token  string
	Ctx    context.Context
	Client http.Client
}

// UserListResponse ...
type UserListResponse struct {
	Ok      bool     `json:"ok"`
	Members []Member `json:"members"`
	CacheTs int64    `json:"cache_ts"`
	Error   string   `json:"error"`
}

// Member ...
type Member struct {
	ID       string  `json:"id"`
	TeamID   string  `json:"team_id"`
	Name     string  `json:"name"`
	RealName string  `json:"real_name"`
	Profile  Profile `json:"profile"`
	Updated  int64   `json:"updated"`
	Has2Fa   *bool   `json:"has_2fa,omitempty"`
}

// Profile ...
type Profile struct {
	Title                 string  `json:"title"`
	RealName              string  `json:"real_name"`
	RealNameNormalized    string  `json:"real_name_normalized"`
	DisplayName           string  `json:"display_name"`
	DisplayNameNormalized string  `json:"display_name_normalized"`
	Email                 string  `json:"email,omitempty"`
	Team                  string  `json:"team"`
	FirstName             *string `json:"first_name,omitempty"`
	LastName              *string `json:"last_name,omitempty"`
}
