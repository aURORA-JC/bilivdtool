package core

import (
	"testing"
)

func TestDoFileOperations(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "1",
			filePath: "./v.m4s",
			wantErr:  false,
		},
		{
			name:     "2",
			filePath: "./a.m4s",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DoFileOperations(tt.filePath); (err != nil) != tt.wantErr {
				t.Errorf("DoFileOperations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
