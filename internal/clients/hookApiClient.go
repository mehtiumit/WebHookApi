package clients

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"webhook/internal/domain/hook"
)

type HookApiClient interface {
	SendHookMessage(HookDto hook.HookDto) error
}

type hookApiClient struct {
	url    string
	token  string
	client *fasthttp.Client
}

func (h hookApiClient) SendHookMessage(HookDto hook.HookDto) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(h.url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.token)
	var jsonBody, _ = json.Marshal(HookDto)
	req.SetBodyRaw(jsonBody)
	resp := fasthttp.AcquireResponse()
	err := h.client.Do(req, resp)
	if err != nil {
		return err
	}
	return nil
}

func NewHookApiClient() HookApiClient {
	return &hookApiClient{
		url:    "https://webhook.site/c5008b2a-84bc-4adf-b16b-288da8cc4a90",
		client: new(fasthttp.Client),
	}
}
