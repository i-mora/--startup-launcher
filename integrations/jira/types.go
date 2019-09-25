package jira

import (
	"context"
	"net/http"
)

type API struct {
	URL    string
	User   string
	Token  string
	Ctx    context.Context
	Client http.Client
}

type IssueResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type Issue struct {
	Fields struct {
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
		Summary     string `json:"summary"`
		Description struct {
			Type    string `json:"type"`
			Version int    `json:"version"`
			Content []struct {
				Type    string `json:"type"`
				Content []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"content"`
			} `json:"content"`
		} `json:"description"`
		Issuetype struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Reporter struct {
			ID string `json:"id"`
		} `json:"reporter"`
		Assignee struct {
			ID string `json:"id"`
		} `json:"assignee"`
	} `json:"fields"`
}

type GETprojectRes struct {
	Self       string    `json:"self"`
	MaxResults int64     `json:"maxResults"`
	StartAt    int64     `json:"startAt"`
	Total      int64     `json:"total"`
	IsLast     bool      `json:"isLast"`
	Values     []Project `json:"values"`
}

type Project struct {
	Expand         string `json:"expand"`
	Self           string `json:"self"`
	ID             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	ProjectTypeKey string `json:"projectTypeKey"`
	Simplified     bool   `json:"simplified"`
	Style          string `json:"style"`
	IsPrivate      bool   `json:"isPrivate"`
	EntityID       string `json:"entityId"`
	UUID           string `json:"uuid"`
}

type POSTprojectRes struct {
	Self        string  `json:"self"`
	Name        string  `json:"name"`
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Actors      []Actor `json:"actors"`
	Scope       Scope   `json:"scope"`
}

type Actor struct {
	ID          int64     `json:"id"`
	DisplayName string    `json:"displayName"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	ActorUser   ActorUser `json:"actorUser"`
}

type ActorUser struct {
	AccountID string `json:"accountId"`
}

type Scope struct {
	Type    string  `json:"type"`
	Project Project `json:"project"`
}

type UsersArray struct {
	User []string `json:"user"`
}

type Role struct {
	Self        string  `json:"self"`
	Name        string  `json:"name"`
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Scope       *Scope  `json:"scope,omitempty"`
	Actors      []Actor `json:"actors"`
}

type User struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	Locale       string `json:"locale"`
	Expand       string `json:"expand"`
}
