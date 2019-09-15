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
	req, err := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles?%s", q.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	cont := decodeResponse(res)
	if err := json.Unmarshal(cont, &l); err != nil {
		return nil, err
	}
	return l, nil
}

// ListForTag is a convenience method for retrieving articles
// for a particular tag, calling the base List method.
func (ar *ArticlesResource) ListForTag(ctx context.Context, tag string, page int) ([]Article, error) {
	return ar.List(ctx, ArticleListOptions{Tags: tag, Page: page})
}

// ListForUser is a convenience method for retrieving articles
// by a particular user, calling the base List method.
func (ar *ArticlesResource) ListForUser(ctx context.Context, username string, page int) ([]Article, error) {
	return ar.List(ctx, ArticleListOptions{Username: username, Page: page})
}

// Find will retrieve an Article matching the ID passed.
func (ar *ArticlesResource) Find(ctx context.Context, id uint32) (Article, error) {
	var art Article
	req, err := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles/%d", id), nil)
	if err != nil {
		return art, err
	}

	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return art, err
	}
	cont := decodeResponse(res)
	if err := json.Unmarshal(cont, &art); err != nil {
		return Article{}, err
	}
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
	if err := json.Unmarshal(content, &a); err != nil {
		return Article{}, err
	}
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
	if err := json.Unmarshal(content, &a); err != nil {
		return Article{}, err
	}
	return a, nil
}
