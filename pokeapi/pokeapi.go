package pokeapi

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/google/go-querystring/query"
)

const (
    BaseUrl   = "https://pokeapi.co/api/v2/"
    userAgent = "gopoke/v1.0.0"
)

type PokeApi struct {
    client *http.Client
}

func NewPokeApi() *PokeApi {
    c := &http.Client{Timeout: 10 * time.Second}

    return &PokeApi{client: c}
}

func (p *PokeApi) NewRequest(endpoint string, queryParams interface{}) (*http.Request, error) {
    reqUrl := fmt.Sprintf("%s/%s/", BaseUrl, endpoint)

    if queryParams != nil {
        v, _ := query.Values(queryParams)
        reqUrl = fmt.Sprintf("%s?%s", reqUrl, v.Encode())
    }

    req, err := http.NewRequest(http.MethodGet, reqUrl, nil)

    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("User-Agent", userAgent)

    return req, nil
}

func (p *PokeApi) Do(req *http.Request, resource interface{}) error {
    resp, err := p.client.Do(req)

    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resource != nil {
        err := json.NewDecoder(resp.Body).Decode(&resource)

        if err != nil {
            return err
        }
    }

    return nil
}
