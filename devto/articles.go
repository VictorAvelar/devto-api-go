package devto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

// ArticlesResource implements the APIResource interface
// for devto articles.
type ArticlesResource struct {
	API *Client
}

// List will return the articles uploaded to devto, the result
// can be narrowed down, filtered or enhanced using query
// parameters as specified on the documentation.
// See: https://docs.dev.to/api/#tag/articles/paths/~1articles/get
func (ar *ArticlesResource) List(ctx context.Context, opt ArticleListOptions) ([]Article, error) {
	var l []Article
	q, err := query.Values(opt)
	if err != nil {
		return nil, err
	}
	req, _ := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles?%s", q.Encode()), nil)
	res, _ := ar.API.HTTPClient.Do(req)
	cont := decodeResponse(res)
	json.Unmarshal(cont, &l)
	return l, nil
}

// Find will retrieve an Article matching the ID passed.
func (ar *ArticlesResource) Find(ctx context.Context, id uint32) (Article, error) {
	var art Article
	req, _ := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles/%d", id), nil)
	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return art, err
	}
	cont := decodeResponse(res)
	json.Unmarshal(cont, &art)
	return art, nil
}

// New will create a new article on dev.to
func (ar *ArticlesResource) New(ctx context.Context, a Article) (Article, error) {
	if ar.API.Config.InsecureOnly {
		return a, ErrProtectedEndpoint
	}
	cont, err := json.Marshal(a)
	if err != nil {
		return a, err
	}
	req, err := ar.API.NewRequest(http.MethodPost, "api/articles", strings.NewReader(string(cont)))
	if err != nil {
		return a, err
	}
	req.Header.Add(APIKeyHeader, ar.API.Config.APIKey)
	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return a, err
	}
	content := decodeResponse(res)
	json.Unmarshal(content, &a)
	return a, nil
}

// Update will mutate the resource by id, and all the changes
// performed to the Article will be applied, thus validation
// on the API side.
func (ar *ArticlesResource) Update(ctx context.Context, a Article) (Article, error) {
	if ar.API.Config.InsecureOnly {
		return a, ErrProtectedEndpoint
	}
	cont, err := json.Marshal(a)
	if err != nil {
		return a, err
	}
	req, err := ar.API.NewRequest(http.MethodPut, fmt.Sprintf("api/articles/%d", a.ID), strings.NewReader(string(cont)))
	if err != nil {
		return a, err
	}
	req.Header.Add(APIKeyHeader, ar.API.Config.APIKey)
	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return a, err
	}
	content := decodeResponse(res)
	json.Unmarshal(content, &a)
	return a, nil
}
