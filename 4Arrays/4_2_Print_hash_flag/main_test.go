package main

import (
	"fmt"
	"testing"
)

func TestPrintHashFlag(t *testing.T) {
	tests := []struct {
		name      string
		algorithm string
		input     string
		want      string
		wantErr   bool
	}{
		{"SHA512",
			"512",
			"Hello, World!",
			"374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeHash(tt.algorithm, []byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("computeHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && fmt.Sprintf("%x", got) != tt.want {
			}
		})

	}

}
