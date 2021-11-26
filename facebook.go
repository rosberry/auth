package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	Facebook struct{}
)

const (
	userDetailsTemplateFB = `https://graph.facebook.com/v7.0/me?fields=id,name,first_name,last_name,picture,email`
)

func (s *Facebook) auth(token string) (ud *UserDetails, err error) {
	// create request
	req, err := http.NewRequest("GET", userDetailsTemplateFB, nil)
	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Add("Authorization", "Bearer "+token)

	// send request with headers
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type userInfo struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Picture   struct {
			Data struct {
				URL string `json:"url"`
			}
		}
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
	}
	var info userInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if info.Error.Code != 0 {
		return nil, fmt.Errorf("[Request user details error] %v: %v", info.Error.Code, info.Error.Message)
	}

	ud = &UserDetails{
		ID:        info.ID,
		UserName:  info.Name,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
		Picture:   info.Picture.Data.URL,
	}

	return
}

func (s *Facebook) authWithCheckAUD(token, aud string) (ud *UserDetails, err error) {
	return s.auth(token)
}
