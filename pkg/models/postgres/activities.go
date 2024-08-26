package postgres

import (
	"database/sql"
	"time"

	"nugu.dev/basement/pkg/models"
)

type ActivityModel struct {
	Db *sql.DB
}

func (a *ActivityModel) Insert(activityType models.ActivityEnum, startTime time.Time) (int, error) {

	// TODO: IF LAST ONE IS NOT FINISHED RETURN ERROR
	var prev, next int

	stmt := `SELECT id FROM activities ORDER BY id DESC`
	query := a.Db.QueryRow(stmt).Scan(&prev)

	if query == sql.ErrNoRows {
		prev = 0
	}

	stmt = `INSERT INTO activities (id, activity_type, start_time) VALUES ($1, $2, $3) RETURNING id`

	err := a.Db.QueryRow(stmt, prev+1, activityType, startTime).Scan(&next)

	if err != nil {
		return 0, err
	}

	return next, nil
}

func (a *ActivityModel) SetEnd(endTime time.Time) (int, models.ActivityEnum, error) {

	// TODO: IF LAST ONE IS FINISHED RETURN ERROR

	var id int

	stmt := `SELECT id FROM activities ORDER BY id DESC`
	query := a.Db.QueryRow(stmt).Scan(&id)

	if query == sql.ErrNoRows {
		return 0, "", query
	}

	var startTime time.Time
	var aType models.ActivityEnum

	stmt = `UPDATE activities SET end_time = $1 WHERE ID = $2 RETURNING start_time, activity_type`

	err := a.Db.QueryRow(stmt, endTime, id).Scan(&startTime, &aType)

	if err != nil {
		return 0, "", query
	}

	return int(endTime.Sub(startTime).Seconds()), aType, nil
}

func (a *ActivityModel) GetByYear(year int) (int, error) {

	return 0, nil
}
