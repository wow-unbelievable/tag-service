package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wow-unbelievable/tag-service/pkg/errcode"
	"io"
	"net/http"
	"net/url"
)

const (
	APP_KEY    = "abv"
	APP_SECRET = "asdz"
)

type API struct {
	URL string
}

type AccessToekn struct {
	Token string `json:"token"`
}

func NewApi(url string) *API {
	return &API{URL: url}
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func (a *API) httpPost(ctx context.Context, path string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/%s", a.URL, path), form)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTokenFail)
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	form := url.Values{}
	form.Set("app_key", APP_KEY)
	form.Set("app_secret", APP_SECRET)
	body, err := a.httpPost(ctx, fmt.Sprintf("%s", "auth"), form)
	if err != nil {
		return "", err
	}

	var accessToken AccessToekn
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}
