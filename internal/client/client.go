package client

import (
	"FT_ServerClient/pkg/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const address = "localhost:8080"

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Start() {
	err := c.helloRequest()
	if tools.MinorError(err, "Failed helloRequest") {
		fmt.Println(err)
	}

	err = c.helloUsernameRequest()
	if tools.MinorError(err, "Failed helloUsernameRequest") {
		fmt.Println(err)
	}
}

func (c *Client) helloRequest() error {
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:8000/hello", nil)
	if tools.MinorError(err, "Failed http.NewRequest()") {
		return err
	}
	fmt.Println("Send request: GET /hello")
	resp, err := client.Do(req)
	if tools.MinorError(err, "Failed http.NewRequest()") {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if tools.MinorError(err, "Failed to read response body") {
		return err
	}
	fmt.Printf("Received response: %d %s \n", resp.StatusCode, respBody)
	return err
}

func (c *Client) helloUsernameRequest() error {
	client := http.Client{Timeout: time.Duration(1) * time.Second}

	data, err := json.Marshal(map[string]string{
		"username": "Name",
	})
	body := bytes.NewReader(data)
	if tools.MinorError(err, "Failed to marshal response") {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/hello_username", body)
	if tools.MinorError(err, "Failed http.NewRequest()") {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if tools.MinorError(err, "Failed to send request") {
		return err
	}
	if tools.MinorError(err, "Failed http.NewRequest()") {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if tools.MinorError(err, "Failed to read response body") {
		return err
	}
	fmt.Printf("Received response: %d %s \n", resp.StatusCode, respBody)
	return err
}
