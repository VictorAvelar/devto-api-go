package devto

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *Config
		bc     httpClient
	}

	c, err := NewConfig(true, "dummy")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		want    Client
		wantErr bool
		err     error
		args    args
	}{
		{
			name: "client is build sucessfully",
			want: Client{
				Context:    context.Background(),
				HTTPClient: http.DefaultClient,
				Config:     c,
			},
			wantErr: false,
			err:     nil,
			args: args{
				ctx:    nil,
				config: c,
				bc:     nil,
			},
		},
		{
			name:    "client fails if config is nil",
			want:    Client{},
			wantErr: true,
			err:     ErrMissingConfig,
			args: args{
				ctx:    nil,
				config: nil,
				bc:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.ctx, tt.args.config, tt.args.bc)
			if tt.wantErr && err != nil {
				if !reflect.DeepEqual(err, tt.err) {
					t.Errorf("failed on error expectation, got: %v | want %v", err, tt.err)
				}
			} else {
				if !reflect.DeepEqual(tt.want, *got) {
					t.Errorf("missmatched expectation, got: %v, want: %v", got, tt.want)
				}
			}
		})
	}
}
