package info

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arthurschuelter/go-git/models"
	"github.com/arthurschuelter/go-git/utils"
)

func GetUserData(url string, client *http.Client, token string) (models.User, error) {
	req, err := utils.ConfigRequest(url, token)
	if err != nil {
		return models.User{}, err
	}
	body, _, err := utils.MakeRequestAndRead(client, req)
	if err != nil {
		return models.User{}, fmt.Errorf("get user: %w", err)
	}
	user, err := ReadUser(body)
	if err != nil {
		return models.User{}, fmt.Errorf("decode user: %w", err)
	}
	return user, nil
}

func ReadUser(data []byte) (models.User, error) {
	var user models.User
	err := json.Unmarshal(data, &user)
	return user, err
}
