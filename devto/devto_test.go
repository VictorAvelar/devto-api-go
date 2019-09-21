package devto

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// ----- Testing utilities -----

var (
	testMux       *http.ServeMux
	testClientPro *Client
	testClientPub *Client
	testServer    *httptest.Server
	testConfigPro *Config
	testConfigPub *Config
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	testConfigPro, _ = NewConfig(true, "demo-token")
	testConfigPub, _ = NewConfig(false, "")
	testClientPro, _ = NewClient(nil, testConfigPro, nil, testServer.URL)
	testClientPub, _ = NewClient(nil, testConfigPub, nil, testServer.URL)
}

func teardown() {
	testServer.Close()
}

func TestNewClient(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *Config
		bc     httpClient
		bu     string
	}

	c, err := NewConfig(true, "dummy")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		wantErr bool
		err     error
		args    args
	}{
		{
			name:    "client is build successfully",
			wantErr: false,
			err:     nil,
			args: args{
				ctx:    nil,
				config: c,
				bc:     nil,
				bu:     "",
			},
		},
		{
			name:    "client fails if config is nil",
			wantErr: true,
			err:     ErrMissingConfig,
			args: args{
				ctx:    nil,
				config: nil,
				bc:     nil,
				bu:     "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.ctx, tt.args.config, tt.args.bc, tt.args.bu)
			if tt.wantErr && err != nil {
				if !reflect.DeepEqual(err, tt.err) {
					t.Errorf("failed on error expectation, got: %v | want %v", err, tt.err)
				}
			} else {
				if !reflect.DeepEqual(tt.err, nil) || got == nil {
					t.Errorf("missmatched expectation, got: %v, want: %v", err, tt.err)
				}
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	setup()
	defer teardown()
	type args struct {
		m    string
		uri  string
		body io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			"new request is created",
			args{
				m:    "",
				uri:  "",
				body: nil,
			},
			false,
			nil,
		},
		{
			"new fails for invalid url",
			args{
				m:    "",
				uri:  " http://localhost",
				body: nil,
			},
			true,
			errors.New("parse  http://localhost: first path segment in URL cannot contain colon"),
		},
		{
			"new fails for invalid url",
			args{
				m:    "\\\\\\\\\\\\\\",
				uri:  "",
				body: nil,
			},
			true,
			fmt.Errorf("net/http: invalid method %q", "\\\\\\\\\\\\\\"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClientPro.NewRequest(tt.args.m, tt.args.uri, tt.args.body)
			if tt.wantErr && err != nil {
				if !reflect.DeepEqual(err.Error(), tt.err.Error()) {
					t.Errorf("failed on error expectation, got: %v | want %v", err, tt.err)
				}
			} else {
				if !reflect.DeepEqual(tt.err, nil) || got == nil {
					t.Errorf("missmatched expectation, got: %v, want: %v", err, tt.err)
				}
			}
		})
	}
}
