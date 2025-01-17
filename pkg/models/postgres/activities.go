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
		return models.ErrNotFound
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

func (a *ActivityRepository) NewCompleteActivity(x models.Activity) (int, error) {

	var newId int

	tx, err := a.Db.Begin()

	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO activities(start_time, end_time, title, description)
			VALUES($1, $2, $3, $4) 
			RETURNING activities.id;`

	if x.StartTime.IsZero() || x.EndTime.IsZero() || x.StartTime.After(x.EndTime) {
		return 0, models.ErrInvalidInsert
	}

	if err = tx.QueryRow(stmt, x.StartTime, x.EndTime, x.Title, x.Description).Scan(&newId); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, tag := range x.Tags {
		stmt = `INSERT INTO activities_tags(fk_activity_id, fk_tag_id) VALUES ($1, $2);`

		_, err = tx.Exec(stmt, newId, tag.ID)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return newId, nil
}

func (a *ActivityRepository) GetLastActivity() (models.Activity, error) {

	var search models.Activity
	var endTime sql.NullTime
	var title, description sql.NullString

	stmt := `SELECT id, start_time, end_time, title, description
			FROM activities
			WHERE id IN (SELECT MAX(id) FROM activities);`

	if err := a.Db.QueryRow(stmt).Scan(
		&search.ID,
		&search.StartTime,
		&endTime,
		&title,
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

	if title.Valid {
		search.Title = title.String
	} else {
		search.Title = ""
	}

	return search, nil
}

func (a *ActivityRepository) GetDailyLog(date time.Time) ([]models.Activity, error) {

	var search []models.Activity
	var title, description sql.NullString

	stmt := `SELECT id, start_time, end_time, title, description, AGE(activities.end_time, activities.start_time)
			FROM activities
			 WHERE activities.start_time::date = $1
			 AND activities.end_time IS NOT NULL
			 ORDER BY activities.start_time DESC;`

	rows, err := a.Db.Query(stmt, date.Format(time.DateOnly))

	if err != nil {
		return search, models.ErrDbOperation
	}

	for rows.Next() {
		var s models.Activity
		var activityTags []models.Tag

		if err = rows.Scan(
			&s.ID,
			&s.StartTime,
			&s.EndTime, // Dont need to use sql.NullTime bc query checks IS NOT NULL
			&title,
			&description,
			&s.Age,
		); err != nil {
			return search, models.ErrDbOperation
		}

		if title.Valid {
			s.Title = title.String
		} else {
			s.Title = ""
		}

		if description.Valid {
			s.Description = description.String
		} else {
			s.Description = ""
		}

		stmt = `SELECT tags.id, tags.name FROM tags
				JOIN activities_tags ON tags.id = activities_tags.fk_tag_id
				WHERE activities_tags.fk_activity_id = $1;`

		innerRows, err := a.Db.Query(stmt, s.ID)

		if err != nil {
			return search, models.ErrDbOperation
		}

		for innerRows.Next() {
			var t models.Tag

			if err = innerRows.Scan(
				&t.ID,
				&t.Name,
			); err != nil {
				log.Println(err)
				return search, models.ErrDbOperation
			}

			activityTags = append(activityTags, t)
		}

		s.Tags = activityTags
		search = append(search, s)
	}

	return search, nil
}

func (a *ActivityRepository) GetIntervalLog(start time.Time, end time.Time) ([]models.ActivityDayOverview, error) {

	var search []models.ActivityDayOverview

	stmt := `SELECT start_time::date, FLOOR(EXTRACT(EPOCH FROM SUM(AGE(activities.end_time, activities.start_time))))
			FROM activities
			 WHERE activities.start_time::date >= $1
			 AND activities.end_time::date < $2
			 AND activities.end_time IS NOT NULL
			 GROUP BY start_time::date
			 ORDER BY start_time::date DESC;`

	rows, err := a.Db.Query(stmt, start.Format(time.DateOnly), end.Format(time.DateOnly))

	if err != nil {
		return search, models.ErrDbOperation
	}

	for rows.Next() {
		var s models.ActivityDayOverview

		if err = rows.Scan(
			&s.Date,
			&s.TotalSec,
		); err != nil {
			return search, models.ErrDbOperation
		}

		search = append(search, s)
	}

	return search, nil
}
