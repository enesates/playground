package main

import (
	"testing"
)

func TestUserTableCreated(t *testing.T) {
	t.Run("Check users table existence", func(t *testing.T) {
		var exists bool

		err := PostgresDB.QueryRow(`
            SELECT EXISTS (
               SELECT FROM information_schema.tables
				WHERE table_schema = 'public'
				AND table_name = 'users'
            )
        `).Scan(&exists)

		if err != nil {
			t.Fatalf("failed to check users table existence: %v", err)
		}

		if !exists {
			t.Fatal("users table does not exist")
		}
	})
}
