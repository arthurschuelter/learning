package info

import (
	"encoding/json"
	"net/http"

	"github.com/arthurschuelter/go-git/models"
	"github.com/arthurschuelter/go-git/utils"
)

func GetUserData(url string, client *http.Client, token string) models.User {
	req := utils.ConfigRequest(url, token)
	body, err := utils.MakeRequestAndRead(client, req)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", body)
	user, err := ReadUser(body)

	if err != nil {
		panic(err)
	}
	return user
}

func ReadUser(data []byte) (models.User, error) {
	var user models.User
	err := json.Unmarshal(data, &user)
	return user, err
}
