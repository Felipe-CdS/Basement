package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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
	b, err := store.StartActivity()
	require.NoError(t, err)
	assert.Equal(t, 1, b, "Should be equal")

	// Trying to start and activity before finish other
	_, err = store.StartActivity()
	require.ErrorIs(t, err, models.ErrNotFinished)
}

func TestInsertEnd(t *testing.T) {
	// No Error
	err := store.EndActivity()
	require.NoError(t, err)

	// End inexistent activity
	err = store.EndActivity()
	require.ErrorIs(t, err, models.ErrNotFound)
}

func TestInsertComplete(t *testing.T) {

	activity := models.Activity{
		Title:       "aaaa",
		Description: "bbb",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(time.Hour),
	}

	// No Error
	b, err := store.NewCompleteActivity(activity)
	require.NoError(t, err)
	assert.Equal(t, 2, b, "Should be equal")

	// Trying to insert an activity without StartTime
	activity.StartTime = time.Time{}
	_, err = store.NewCompleteActivity(activity)
	require.ErrorIs(t, err, models.ErrInvalidInsert)

	// Trying to insert an activity without StartTime
	activity.StartTime = time.Now()
	activity.EndTime = time.Time{}
	_, err = store.NewCompleteActivity(activity)
	require.ErrorIs(t, err, models.ErrInvalidInsert)
}
