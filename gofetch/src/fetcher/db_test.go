package fetcher

import (
	"testing"
)

func TestDb(t *testing.T) {
	db := Db()
	defer db.Close()
}
