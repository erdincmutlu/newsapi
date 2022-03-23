package internal

import (
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/stretchr/testify/require"
)

func TestPrepareMessage(t *testing.T) {
	// The main is running in root folder and the tests are running in "internal" folder.
	// Make this test to run in root folder so it will be able to find the "templates" folder.
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	require.NoError(t, err)

	fillCache()

	tests := []struct {
		name     string
		id       string
		expected string
		err      error
	}{
		{
			name: "ok",
			id:   "1",
			expected: `<html>
<head>
    <title>News shared with you</title>
</head>
<body>
<h2>News shared with you</h2>
Your friend shared this news with you
<p>
<a href="">Title 1</a><br>
This is description 1
</body>
</html>`,
		},
		{
			name: "not found",
			id:   "3",
			err:  ttlcache.ErrNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := prepareMessage(test.id)
			require.Equal(t, test.err, err)
			// Get rid of carriage return char, windows is adding into
			require.Equal(t, test.expected, strings.ReplaceAll(string(message), "\r", ""))
		})
	}
}
