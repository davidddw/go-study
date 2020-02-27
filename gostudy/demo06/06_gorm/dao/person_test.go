package dao

import (
	"fmt"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := GetAllPerson()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, user := range gotUsers {
				fmt.Printf("user: %v", user)
			}
		})
	}
}
