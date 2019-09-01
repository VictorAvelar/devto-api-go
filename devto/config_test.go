package devto

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		p bool
		k string
	}
	tests := []struct {
		name    string
		a       args
		want    Config
		wantErr bool
		err     error
	}{
		{
			name: "new config call is successful for public endpoints",
			a: args{
				p: false,
				k: "",
			},
			want: Config{
				InsecureOnly: true,
				APIKey:       "",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "new config call is successful for protected endpoints",
			a: args{
				p: true,
				k: "dummy-api-key",
			},
			want: Config{
				InsecureOnly: false,
				APIKey:       "dummy-api-key",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "new config call fails for protected access without a key",
			a: args{
				p: true,
				k: "",
			},
			want:    Config{},
			wantErr: true,
			err:     ErrMissingRequiredParameter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.a.p, tt.a.k)
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
