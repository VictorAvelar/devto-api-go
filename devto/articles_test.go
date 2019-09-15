package devto

import (
	"context"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/VictorAvelar/devto-api-go/testdata"
)

var ctx = context.Background()

func TestArticlesResource_List(t *testing.T) {
	setup()
	defer teardown()
	cont, err := ioutil.ReadAll(strings.NewReader(testdata.ListResponse))
	if err != nil {
		t.Error(err)
	}
	testMux.HandleFunc("/api/articles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(cont)
	})

	var ctx = context.Background()
	list, err := testClientPub.Articles.List(ctx, ArticleListOptions{})
	if err != nil {
		t.Error(err)
	}

	if len(list) != 3 {
		t.Errorf("not all articles where parsed")
	}

	for _, a := range list {
		if a.Title == "" {
			t.Error("parsing failed / empty titles")
		}
	}
}

func TestArticlesResource_ListWithQueryParams(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/articles?page=1&state=fresh&tag=go&top=1&username=victoravelar", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix("username=victoravelar", r.URL.String()) {
			t.Error("url mismatch")
		}
		w.WriteHeader(http.StatusOK)
	})

	q := ArticleListOptions{
		Tags:     "go",
		Username: "victoravelar",
		State:    "fresh",
		Top:      "1",
		Page:     1,
	}
	list, err := testClientPub.Articles.List(ctx, q)
	if err != nil {
		t.Error(err)
	}
	if len(list) != 0 {
		t.Error("response is unexpected")
	}
}

func TestArticlesResource_Find(t *testing.T) {
	setup()
	defer teardown()
	cont, err := ioutil.ReadAll(strings.NewReader(testdata.FindResponse))
	if err != nil {
		t.Error(err)
	}
	testMux.HandleFunc("/api/articles/164198", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(cont)
	})

	article, err := testClientPub.Articles.Find(ctx, 164198)
	if err != nil {
		t.Error(err)
	}

	if article.ID != 164198 {
		t.Error("article returned is not the one requested")
	}
}

func TestArticlesResource_New(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("invalid method for request")
		}
		cont, _ := ioutil.ReadAll(r.Body)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(cont)
	})

	res, err := testClientPro.Articles.New(ctx, ArticleUpdate{
		Title:     "Demo article",
		Published: false,
	})
	if err != nil {
		t.Fatalf("unexpected error publishing article: %v", err)
	}

	if res.Title != "Demo article" {
		t.Error("article parsing failed")
	}
}

func TestArticlesResource_NewFailsWhenInsecure(t *testing.T) {
	setup()
	defer teardown()

	_, err := testClientPub.Articles.New(ctx, ArticleUpdate{
		Title:     "Demo article",
		Published: false,
	})

	if !reflect.DeepEqual(err, ErrProtectedEndpoint) {
		t.Error("auth check failed")
	}
}

func TestArticlesResource_Update(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/articles/164198", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Error("invalid method for request")
		}
		cont, _ := ioutil.ReadAll(r.Body)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(cont)
	})

	res, err := testClientPro.Articles.Update(ctx, ArticleUpdate{
		Title:     "Demo article",
		Published: false,
	}, 164198)
	if err != nil {
		t.Fatalf("unexpected error publishing article: %v", err)
	}

	if res.Title != "Demo article" {
		t.Error("article parsing failed")
	}
}

func TestArticlesResource_UpdateFailsWhenInsecure(t *testing.T) {
	setup()
	defer teardown()

	_, err := testClientPub.Articles.Update(ctx, ArticleUpdate{
		Title:     "Demo article",
		Published: false,
	}, 164198)

	if !reflect.DeepEqual(err, ErrProtectedEndpoint) {
		t.Error("auth check failed")
	}
}
