package goecsclient

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/coredgeio/goecsclient/errors"
)

type ecsSession struct {
	Username string
	Password string
	Endpoint string
	Token    string
	c        *http.Client
}

const (
	TimeBufferInSeconds = int64(300)
)

func (s *ecsSession) Get(subUrl string, q url.Values, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", s.Endpoint+subUrl, nil)
	if q != nil {
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("X-SDS-AUTH-TOKEN", s.Token)
	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	
	resp, err := s.c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Println("failed to read Body", err)
			return nil, err
		}
	}
	if resp.StatusCode != http.StatusOK {
		if bodyBytes != nil {
			return nil, errors.ParseError(bodyBytes)
		}
		return nil, errors.Wrap(resp.Status)
	}
	return bodyBytes, nil
}

func (s *ecsSession) Post(subUrl string, d []byte, q url.Values, headers map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("POST", s.Endpoint+subUrl, bytes.NewReader(d))
	if q != nil {
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("X-SDS-AUTH-TOKEN", s.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := s.c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Println("failed to read Body", err)
			return nil, err
		}
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		if bodyBytes != nil {
			return nil, errors.ParseError(bodyBytes)
		}
		return nil, errors.Wrap(resp.Status)
	}
	return bodyBytes, nil
}

func (s *ecsSession) Put(subUrl string, d []byte, q url.Values) ([]byte, error) {
	req, _ := http.NewRequest("PUT", s.Endpoint+subUrl, bytes.NewReader(d))
	if q != nil {
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("X-SDS-AUTH-TOKEN", s.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := s.c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Println("failed to read Body", err)
			return nil, err
		}
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		if bodyBytes != nil {
			return nil, errors.ParseError(bodyBytes)
		}
		return nil, errors.Wrap(resp.Status)
	}
	return bodyBytes, nil
}

// internal function to perform login while client is created using user
// credentials. upon successful login attempt this updates the token that
// is used as part of various api triggers
func (s *ecsSession) performLogin() error {
	// token endpoint as of now is static and available at sub-path
	// /login
	req, err := http.NewRequest("GET", s.Endpoint+"/login", nil)
	req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.c.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap("login request failed, check endpoint or credentials")
	}
	token := ""
	if len(resp.Header) != 0 {
		token = resp.Header.Get("X-SDS-AUTH-TOKEN")
		maxAge := resp.Header.Get("X-SDS-AUTH-MAX-AGE")
		if maxAge != "" && token != "" {
			log.Println("got token age", maxAge)
			age, err := strconv.ParseInt(maxAge, 10, 64)
			if err != nil {
				log.Println("invalid age received", err)
			} else {
				go func() {
					// trigger token refresh upon approaching token age
					if age > TimeBufferInSeconds {
						age = age - TimeBufferInSeconds
					}
					time.Sleep(time.Duration(age) * time.Second)
					err := s.performLogin()
					if err != nil {
						// TODO(Prabhjot) need to evaluate if this situation
						// can be handled gracefully
						log.Fatalln("failed to refresh the session token")
					}
				}()
			}
		}
	}
	if token != "" {
		s.Token = token
		return nil
	}
	return errors.Wrap("Auth Token not available in response")
}

func createEcsSession(username, password, endpoint string) (*ecsSession, error) {
	// since certificate might be self signed, with mostly internal
	// communication with Dell ECS storage, it is safe to ignore
	// certificate validation
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	s := &ecsSession{
		Username: username,
		Password: password,
		Endpoint: endpoint,
		c:        &http.Client{Transport: tr},
	}
	err := s.performLogin()
	if err != nil {
		return nil, err
	}
	return s, nil
}
