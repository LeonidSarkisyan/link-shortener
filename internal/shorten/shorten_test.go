package shorten_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"url-shotener-api/internal/shorten"
)

func TestShorten(t *testing.T) {
	t.Run("returns an alphameric short identifier", func(t *testing.T) {
		type testCase struct {
			id       uint32
			excepted string
		}

		testCases := []testCase{
			{
				id:       1024,
				excepted: "Mv",
			},
			{
				id:       0,
				excepted: "",
			},
		}

		for _, tc := range testCases {
			actual := shorten.Shorten(tc.id)
			assert.Equal(t, tc.excepted, actual)
		}
	})
	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, "Mv", shorten.Shorten(1024))
		}
	})
}
