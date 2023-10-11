package main

import (
	"testing"
)

func Test_hello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Case1",
			args:    args{name: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Case2",
			args:    args{name: "zheng"},
			want:    "Hello, zheng!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_introduce(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Case1",
			want: "Welcome to my Golang column.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := introduce(); got != tt.want {
				t.Errorf("introduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
