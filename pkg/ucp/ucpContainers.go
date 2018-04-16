package ucp

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListContainers() error {
	// Add the /auth/log to the URL
	url := fmt.Sprintf("%s/containers/json", c.UCPURL)

	data := map[string]string{
		"username": c.Username,
		"password": c.Password,
	}
	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	response, err := c.getRequest(url, b)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", response)
	return nil
}