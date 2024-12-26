package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"nugu.dev/basement/pkg/models"
	"nugu.dev/basement/pkg/models/postgres"
)

var store *postgres.ActivityRepository

func TestMain(m *testing.M) {

	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {

	log.Println("Setting up database..")

	db := postgres.NewPostgresDB(true)
	store = &postgres.ActivityRepository{Db: db}

	defer func() {
		postgres.DropTestMigrations(db)
		db.Close()
	}()

	return m.Run(), nil
}

func TestInsertStart(t *testing.T) {
	// No Error
	b, err := store.Insert()
	require.NoError(t, err)
	assert.Equal(t, 1, b, "Should be equal")

	// Trying to start and activity before finish other
	b, err = store.Insert()
	require.ErrorIs(t, err, models.ErrNotFinished)
}

func TestInsertEnd(t *testing.T) {
	// No Error
	err := store.EndActivity(1)
	require.NoError(t, err)

	// End inexistent activity
	err = store.EndActivity(2)
	require.ErrorIs(t, err, models.ErrNotFound)
}
