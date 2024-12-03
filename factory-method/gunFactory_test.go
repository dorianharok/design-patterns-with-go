package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGun(t *testing.T) {
	tests := []struct {
		name      string
		gunType   string
		wantName  string
		wantPower int
		wantErr   bool
	}{
		{
			name:      "should create ak47",
			gunType:   "ak47",
			wantName:  "AK47 gun",
			wantPower: 4,
			wantErr:   false,
		},
		{
			name:      "should create musket",
			gunType:   "musket",
			wantName:  "Musket gun",
			wantPower: 1,
			wantErr:   false,
		},
		{
			name:      "should return error for invalid gun type",
			gunType:   "invalid",
			wantName:  "",
			wantPower: 0,
			wantErr:   true,
		},
		{
			name:      "should return error for empty gun type",
			gunType:   "",
			wantName:  "",
			wantPower: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGun(tt.gunType)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, tt.wantName, got.getName())
			assert.Equal(t, tt.wantPower, got.getPower())
		})
	}
}
