package gojwt_test

import (
	"fmt"
	"github.com/prongbang/gojwt"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		payload map[string]any
		key     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return jwt token when generate token success",
			args: args{
				payload: map[string]any{
					"exp": 99999999999,
				},
				key: "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1",
			},
			want:    "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTk5fQ.rMKkGe6riuLZ3boYiMZsk5xrT7S-7VK6gZmFs1_7kKtVUkpvGatudYI5ZSkwIQ-iJKp2XskCxzn_6fVkCohtUQ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := gojwt.New()
			got, err := j.Generate(tt.args.payload, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		token string
		key   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true when verify success",
			args: args{
				key:   "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1",
				token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTk5fQ.rMKkGe6riuLZ3boYiMZsk5xrT7S-7VK6gZmFs1_7kKtVUkpvGatudYI5ZSkwIQ-iJKp2XskCxzn_6fVkCohtUQ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := gojwt.New()
			if got := j.Verify(tt.args.token, tt.args.key); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		token string
		key   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return exp payload when parse success",
			args: args{
				key:   "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1",
				token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTk5fQ.rMKkGe6riuLZ3boYiMZsk5xrT7S-7VK6gZmFs1_7kKtVUkpvGatudYI5ZSkwIQ-iJKp2XskCxzn_6fVkCohtUQ",
			},
			want:    "99999999999",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := gojwt.New()
			if got, err := j.Parse(tt.args.token, tt.args.key); fmt.Sprint(got["exp"]) != tt.want && err != nil != tt.wantErr {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	j := gojwt.New()
	key := "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1"
	payload := map[string]any{
		"exp": 999999999,
	}
	for i := 0; i < b.N; i++ {
		_, _ = j.Generate(payload, key)
	}
}

func BenchmarkParse(b *testing.B) {
	j := gojwt.New()
	key := "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1"
	jwe := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTk5fQ.rMKkGe6riuLZ3boYiMZsk5xrT7S-7VK6gZmFs1_7kKtVUkpvGatudYI5ZSkwIQ-iJKp2XskCxzn_6fVkCohtUQ"
	for i := 0; i < b.N; i++ {
		_, _ = j.Parse(jwe, key)
	}
}

func BenchmarkVerify(b *testing.B) {
	j := gojwt.New()
	key := "bdacaf398071931518f73917cb0c6f04b3a0ab45ee9cbedc258047a8c149a3e1"
	jwe := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTk5fQ.rMKkGe6riuLZ3boYiMZsk5xrT7S-7VK6gZmFs1_7kKtVUkpvGatudYI5ZSkwIQ-iJKp2XskCxzn_6fVkCohtUQ"
	for i := 0; i < b.N; i++ {
		_ = j.Verify(jwe, key)
	}
}
