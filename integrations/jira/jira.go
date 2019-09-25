package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// POSTissue ...
func (api *API) POSTissue(issue Issue) (IssueResponse, error) {
	resource := "issue"
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(&issue); err != nil {
		return IssueResponse{}, err
	}

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodPost, u, buffer)
	if err != nil {
		return IssueResponse{}, err
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return IssueResponse{}, err
	}
	defer resp.Body.Close()

	var response IssueResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// GETprojects ...
func (api *API) GETprojects() ([]Project, error) {
	resource := "project/search"
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response GETprojectRes
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Values, err
}

// POSTproject ...
func (api *API) POSTproject(projectName string, roleID string, users UsersArray) (POSTprojectRes, error) {
	resource := "project/" + projectName + "/role/" + roleID
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(&users); err != nil {
		return POSTprojectRes{}, err
	}

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodPost, u, buffer)
	if err != nil {
		return POSTprojectRes{}, err
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return POSTprojectRes{}, err
	}
	defer resp.Body.Close()

	var response POSTprojectRes
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// GETroles ...
func (api *API) GETroles() ([]Role, error) {
	resource := "role"
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("GETroles error making new requesr: %v", err)
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GETroles error doing requesr: %v", err)
	}
	defer resp.Body.Close()

	var response []Role
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, fmt.Errorf("GETroles error decoding response: %v", err)
}

// GETusers ...
func (api *API) GETusers() ([]User, error) {
	resource := "users/search"
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response []User
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// POSTuser ...
func (api *API) POSTuser(user User) (User, error) {
	resource := "user"
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(&user); err != nil {
		return User{}, err
	}

	req, err := http.NewRequestWithContext(api.Ctx, http.MethodPost, u, buffer)
	if err != nil {
		return User{}, err
	}
	req.SetBasicAuth(api.User, api.Token)

	resp, err := api.Client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	var response User
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
