package src

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// RandomFox gets a picture of a fox
func RandomFox() (string, error) {
	var err error
	client := http.DefaultClient

	url := "https://randomfox.ca/floof/"

	resp, err := DoRequest(url, client)
	if err != nil {
		fmt.Println(err)
		return "", errors.Wrap(err, "RandomFoc() could DoRequest()")
	}

	fox := Fox{}

	err = json.NewDecoder(resp.Body).Decode(&fox)
	if err != nil {
		fmt.Println(err)
		return "", errors.Wrap(err, "RandomFoc() could Decode()")
	}

	defer resp.Body.Close()

	return fox.Image, err
}
