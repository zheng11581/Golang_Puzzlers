package main

import (
	"testing"
)

func Test_introduce(t *testing.T) {
	tests := []struct {
		name          string
		wantIntroduce string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIntroduce := introduce(); gotIntroduce != tt.wantIntroduce {
				t.Errorf("introduce() = %v, want %v", gotIntroduce, tt.wantIntroduce)
			}
		})
	}
}

func Test_hello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		wantGreeting string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGreeting, err := hello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotGreeting != tt.wantGreeting {
				t.Errorf("hello() = %v, want %v", gotGreeting, tt.wantGreeting)
			}
		})
	}
}
