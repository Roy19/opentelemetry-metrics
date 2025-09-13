package dto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToCart_Validate(t *testing.T) {
	cartName := "cart1"
	itemName := "item1"

	tests := []struct {
		name      string
		input     AddToCart
		expectErr error
	}{
		{
			name:      "valid input",
			input:     AddToCart{CartName: &cartName, ItemName: &itemName},
			expectErr: nil,
		},
		{
			name:      "missing cart name",
			input:     AddToCart{CartName: nil, ItemName: &itemName},
			expectErr: errors.New("cart_name cannot be empty"),
		},
		{
			name:      "missing item name",
			input:     AddToCart{CartName: &cartName, ItemName: nil},
			expectErr: errors.New("item_name cannot be empty"),
		},
		{
			name:      "missing both",
			input:     AddToCart{CartName: nil, ItemName: nil},
			expectErr: errors.New("cart_name cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.expectErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectErr.Error())
			}
		})
	}
}
