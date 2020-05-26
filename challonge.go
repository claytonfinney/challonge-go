package challonge

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const BASE_URL = "api.challonge.com/v1/"

type ChallongeClient struct {
	BaseUrl  string
	Username string
	ApiKey   string
	Client   *http.Client
}

func NewChallongeClient(u, k string) *ChallongeClient {
	return &ChallongeClient{
		BaseUrl:  BASE_URL,
		Username: u,
		ApiKey:   k,
		Client:   &http.Client{},
	}
}

func NewChallongeTestClient(b, u, k string) *ChallongeClient {
	return &ChallongeClient{
		BaseUrl:  b,
		Username: u,
		ApiKey:   k,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

// When unmarshalling the response body, if the supplied body is an error message,
// return either a helpful error message or just the entire raw body for better debugging
// instead of an unhelpful error returned when trying to unmarshal a bad response payload.
func checkForChallongeError(b string) error {
	if b == "HTTP Basic: Access denied." {
		return errors.New("challonge-go: invalid credentials")
	}
	if strings.Contains(b, "errors") {
		return errors.New(fmt.Sprintf("challonge-go: errors present: %s", b))
	}
	return nil
}

func (c *ChallongeClient) Get(uri string, p interface{}) error {
	url := fmt.Sprintf("https://%s:%s@%s%s.json", c.Username, c.ApiKey, c.BaseUrl, uri)
	resp, err := c.Client.Get(url)
	if err != nil {
		return err
	}
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(i))
	err = checkForChallongeError(string(i))
	if err != nil {
		return err
	}
	err = json.Unmarshal(i, p)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChallongeClient) Post(uri string, p interface{}, r interface{}) error {
	url := fmt.Sprintf("https://%s:%s@%s%s.json", c.Username, c.ApiKey, c.BaseUrl, uri)
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = checkForChallongeError(string(i))
	if err != nil {
		return err
	}
	err = json.Unmarshal(i, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChallongeClient) Put(uri string, p interface{}, r interface{}) error {
	url := fmt.Sprintf("https://%s:%s@%s%s.json", c.Username, c.ApiKey, c.BaseUrl, uri)
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(d))
	resp, err := c.Client.Do(req)
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = checkForChallongeError(string(i))
	if err != nil {
		return err
	}
	err = json.Unmarshal(i, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChallongeClient) Delete(uri string, r interface{}) error {
	url := fmt.Sprintf("https://%s:%s@%s%s.json", c.Username, c.ApiKey, c.BaseUrl, uri)
	req, err := http.NewRequest("DELETE", url, nil)
	resp, err := c.Client.Do(req)
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = checkForChallongeError(string(i))
	if err != nil {
		return err
	}
	err = json.Unmarshal(i, r)
	if err != nil {
		return err
	}
	return nil
}
