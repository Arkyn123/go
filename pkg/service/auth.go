package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Get() ([]byte, int, error) {
	return []byte{}, 500, nil
}

func (s *AuthService) Login(url string, contentType string, data any) ([]byte, int, error) {

	payload, err := json.Marshal(data)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, resp.StatusCode, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Request.URL)
	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}

	return body, resp.StatusCode, nil
}
