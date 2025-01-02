package postgres

import (
	"database/sql"
	"log"
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

func (a *ActivityRepository) GetTodayActivities() ([]models.Activity, error) {

	var search []models.Activity
	var description sql.NullString

	stmt := `SELECT id, start_time, end_time, description, AGE(activities.end_time, activities.start_time)
			FROM activities
			 WHERE activities.start_time::date = CURRENT_DATE
			 AND activities.end_time IS NOT NULL
			 ORDER BY activities.start_time DESC;`

	rows, err := a.Db.Query(stmt)

	if err != nil {
		return search, models.ErrDbOperation
	}

	for rows.Next() {
		var s models.Activity

		if err = rows.Scan(
			&s.ID,
			&s.StartTime,
			&s.EndTime, // Dont need to use sql.NullTime bc query checks IS NOT NULL
			&description,
			&s.Age,
		); err != nil {
			log.Println(err)
			return search, models.ErrDbOperation
		}

		if description.Valid {
			s.Description = description.String
		} else {
			s.Description = ""
		}

		search = append(search, s)
	}

	return search, nil
}
