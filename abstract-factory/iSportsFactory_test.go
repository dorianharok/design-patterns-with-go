package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeShoe(t *testing.T) {
	tests := []struct {
		name     string
		brand    string
		wantLogo string
		wantSize int
		wantErr  bool
	}{
		{
			name:     "should create nike shoe",
			brand:    "nike",
			wantLogo: "nike",
			wantSize: 14,
			wantErr:  false,
		},
		{
			name:     "should create adidas shoe",
			brand:    "adidas",
			wantLogo: "adidas",
			wantSize: 14,
			wantErr:  false,
		},
		{
			name:     "should return error for invalid brand",
			brand:    "invalid",
			wantLogo: "",
			wantSize: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory, err := GetSportsFactory(tt.brand)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, factory)
				return
			}

			shoe := factory.makeShoe()
			assert.Equal(t, tt.wantLogo, shoe.getLogo())
			assert.Equal(t, tt.wantSize, shoe.getSize())
		})
	}
}

func TestMakeShirt(t *testing.T) {
	tests := []struct {
		name     string
		brand    string
		wantLogo string
		wantSize int
		wantErr  bool
	}{
		{
			name:     "should create nike shirt",
			brand:    "nike",
			wantLogo: "nike",
			wantSize: 14,
			wantErr:  false,
		},
		{
			name:     "should create adidas shirt",
			brand:    "adidas",
			wantLogo: "adidas",
			wantSize: 14,
			wantErr:  false,
		},
		{
			name:     "should return error for invalid brand",
			brand:    "invalid",
			wantLogo: "",
			wantSize: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory, err := GetSportsFactory(tt.brand)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, factory)
				return
			}

			shoe := factory.makeShoe()
			assert.Equal(t, tt.wantLogo, shoe.getLogo())
			assert.Equal(t, tt.wantSize, shoe.getSize())
		})
	}
}
