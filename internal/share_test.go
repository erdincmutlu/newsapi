package internal

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/erdincmutlu/newsapi/types"
)

func TestShare(t *testing.T) {
	fillCache()

	tests := []struct {
		name     string
		request  types.ShareRequest
		expected string
		err      error
	}{
		{
			name: "ok",
			request: types.ShareRequest{
				ID:        "2",
				Action:    "email",
				Recipient: "test@example.com",
			},
		},
		{
			name: "invalid params",
			request: types.ShareRequest{
				ID:        "123",
				Action:    "invalid",
				Recipient: "test@example.com",
			},
			err: types.ErrInvalidInparams,
		},
		{
			name:    "empty request",
			request: types.ShareRequest{},
			err:     types.ErrInvalidInparams,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := share(test.request)
			require.Equal(t, test.err, err)
		})
	}
}
