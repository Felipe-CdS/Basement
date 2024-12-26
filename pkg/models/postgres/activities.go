package postgres

import (
	"database/sql"
	"time"

	"nugu.dev/basement/pkg/models"
)

type ActivityRepository struct {
	Db *sql.DB
}

func (a *ActivityRepository) StartActivity() (int, error) {

	var newId, notFinishedCheck int

	stmt := "SELECT COUNT(*) FROM activities WHERE activities.end_time IS NULL;"
	if err := a.Db.QueryRow(stmt).Scan(&notFinishedCheck); err != nil {
		return 0, models.ErrDbOperation
	}
	if notFinishedCheck > 0 {
		return 0, models.ErrNotFinished
	}

	stmt = "INSERT INTO activities VALUES(DEFAULT) RETURNING activities.id;"
	if err := a.Db.QueryRow(stmt).Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}

func (a *ActivityRepository) EndActivity() error {

	var endTime sql.NullTime

	stmt := `SELECT end_time
			FROM activities
			WHERE id IN (SELECT MAX(id) FROM activities);`

	if err := a.Db.QueryRow(stmt).Scan(&endTime); err != nil {
		return models.ErrNotFound
	}

	// If last Activity is already Finished return nil
	if endTime.Valid {
		return nil
	}

	stmt = `UPDATE activities 
			SET end_time = NOW() 
			WHERE id IN (SELECT MAX(id) FROM activities);`

	_, err := a.Db.Exec(stmt)

	if err != nil {
		return models.ErrDbOperation
	}

	return nil
}

func (a *ActivityRepository) GetLastActivity() (models.Activity, error) {

	var search models.Activity
	var endTime sql.NullTime
	var description sql.NullString

	stmt := `SELECT id, start_time, end_time, description
			FROM activities
			WHERE id IN (SELECT MAX(id) FROM activities);`

	if err := a.Db.QueryRow(stmt).Scan(
		&search.ID,
		&search.StartTime,
		&endTime,
		&description,
	); err != nil {
		if err == sql.ErrNoRows {
			return search, models.ErrNotFound
		}

		return search, models.ErrDbOperation
	}

	if endTime.Valid {
		search.EndTime = endTime.Time
	} else {
		search.EndTime = time.Time{}
	}

	if description.Valid {
		search.Description = description.String
	} else {
		search.Description = ""
	}

	return search, nil
}
