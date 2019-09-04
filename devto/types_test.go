package devto

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestWebURL_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
		err     error
	}{
		{
			"Unmarshal works fine",
			`{"visit": "https://dev.to/victoravelar"}`,
			false,
			nil,
		},
		{
			"Unmarshal fails when parsing url",
			`{"visit": " http://localhost"}`,
			false,
			errors.New("parse  http://localhost: first path segment in URL cannot contain colon"),
		},
	}

	type tester struct {
		Visit WebURL `json:"visit"`
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wu tester
			c, err := ioutil.ReadAll(strings.NewReader(tt.content))
			json.Unmarshal(c, &wu)
			if tt.wantErr && err != nil {
				if !reflect.DeepEqual(tt.err, err) {
					t.Errorf("want: %v, got: %v", tt.err, err)
				}
			}
		})
	}
}
