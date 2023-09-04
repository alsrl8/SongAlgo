package selenium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isFilePathValid(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		expected bool
	}{
		{
			desc:     "Invalid file path",
			path:     "boo.boo",
			expected: false,
		},
		{
			desc:     "Valid file path",
			path:     "./../go.mod",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ok := isFilePathValid(tc.path)
			assert.Equal(t, tc.expected, ok, "Test case failed: %s", tc.desc)
		})
	}
}
