package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	VK struct{}
)

const (
	apiVersion            = "5.110"
	userDetailsTemplateVK = `https://api.vk.com/method/users.get?access_token=%s&v=%s&fields=photo_max_orig`
)

func (s *VK) auth(token string) (ud *UserDetails, err error) {
	reqUserLink := fmt.Sprintf(userDetailsTemplateVK, token, apiVersion)
	resp, err := http.Get(reqUserLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type userInfo struct {
		Response []struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			AvatarURL string `json:"photo_max_orig"`
		}
		Error struct {
			Code    int    `json:"error_code"`
			Message string `json:"error_msg"`
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

	if len(info.Response) == 0 {
		return nil, err
	}

	ud = &UserDetails{
		ID:        fmt.Sprint(info.Response[0].ID),
		FirstName: info.Response[0].FirstName,
		LastName:  info.Response[0].LastName,
		Picture:   info.Response[0].AvatarURL,
	}

	return
}

func (s *VK) authWithCheckAUD(token, aud string) (ud *UserDetails, err error) {
	return s.auth(token)
}