package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GETusers ...
func (api *API) GETusers() ([]Member, error) {
	fmt.Println("getting users...")

	resource := "users.list?token=" + api.Token
	u := fmt.Sprintf("%v/%v", api.URL, resource)

	fmt.Println("fetching users...")
	req, err := http.NewRequestWithContext(api.Ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("unmarshalling users...")
	var response UserListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("GETusers error: %v", response.Error)
	}
	return response.Members, nil
}
