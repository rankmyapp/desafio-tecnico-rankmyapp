package util_test

import (
	"database/sql"
	"testing"

	"github.com/otaviomart1ns/backend-challenge/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestToNullString(t *testing.T) {
	t.Run("Valid string input", func(t *testing.T) {
		val := "RankMyApp"
		expected := sql.NullString{String: val, Valid: true}

		result := util.ToNullString(val)

		assert.Equal(t, expected, result)
	})

	t.Run("Empty string input", func(t *testing.T) {
		val := ""
		expected := sql.NullString{String: "", Valid: false}

		result := util.ToNullString(val)

		assert.Equal(t, expected, result)
	})
}
