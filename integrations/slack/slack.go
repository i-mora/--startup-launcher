package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "https://slack.com/api"

	// KeyToken ...
	KeyToken = "token"
)

// GetUsers ...
func GetUsers(ctx context.Context) ([]Member, error) {
	fmt.Println("getting users...")

	var userListResponse UserListResponse
	resource := "users.list"
	token, ok := ctx.Value(KeyToken).(string)
	if !ok {
		return nil, fmt.Errorf("Key: %v was not found", KeyToken)
	}
	u := fmt.Sprintf("%v/%v?token=%v", apiURL, resource, token)

	fmt.Println("fetching users...")
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("unmarshalling users...")
	if err := json.Unmarshal(body, &userListResponse); err != nil {
		return nil, err
	}
	return userListResponse.Members, nil
}
