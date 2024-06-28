package goecsclient

import (
	"log"
	"net/url"
)

type EcsClient interface {
	Get(subUrl string, query url.Values, h map[string]string) ([]byte, error)
	Post(subUrl string, data []byte, query url.Values, h map[string]string) ([]byte, error)
	Put(subUrl string, data []byte, query url.Values) ([]byte, error)
}

type ecsClient struct {
	Username string
	Password string
	Endpoint string
	Session  *ecsSession
}

func (c *ecsClient) Get(subUrl string, query url.Values, h map[string]string) ([]byte, error) {
	return c.Session.Get(subUrl, query, h)
}

func (c *ecsClient) Post(subUrl string, data []byte, query url.Values, h map[string]string) ([]byte, error) {
	return c.Session.Post(subUrl, data, query, h)
}

func (c *ecsClient) Put(subUrl string, data []byte, query url.Values) ([]byte, error) {
	return c.Session.Put(subUrl, data, query)
}

// creates Ecs management API client using username and password of provided
// management api user.
//
// Additionally it also requires api endpoint for the dell ecs management API
// server
//
// Upon successful creation of client it returns the client handle.
// whereas, if the creation fails the handle will be nil and corresponding
// error is returned
func CreateEcsClientWithUserCred(username, password, endpoint string) (EcsClient, error) {
	session, err := createEcsSession(username, password, endpoint)
	if err != nil {
		log.Println("ecs client failed to create ecs session", err)
		return nil, err
	}
	cl := &ecsClient{
		Username: username,
		Password: password,
		Endpoint: endpoint,
		Session:  session,
	}

	return cl, nil
}
