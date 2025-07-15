package testhelpers

import (
	"database/sql"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/otaviomart1ns/backend-challenge/internal/infra/db/models"
)

func NewTestDB(t *testing.T) (*sql.DB, *models.Queries) {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	_, b, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to get runtime caller")
	}
	basePath := filepath.Dir(b)
	schemaPath := filepath.Join(basePath, "schema.sql")

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("failed to read schema: %v", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		t.Fatalf("failed to execute schema: %v", err)
	}

	return db, models.New(db)
}
