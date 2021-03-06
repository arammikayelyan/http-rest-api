package store

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

// This helper returns test store and contains a function which is removing
// all the tables of the database.
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(
				fmt.Sprintf("TRANCATE %s CASCADE", strings.Join(tables, ", ")),
			); err != nil {
				log.Fatal(err)
			}
		}

		s.Close()
	}
}
