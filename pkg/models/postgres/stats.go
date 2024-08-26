package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"nugu.dev/basement/pkg/models"
)

type DayStatsModel struct {
	Db *sql.DB
}

func (a *DayStatsModel) Insert(date time.Time, duration int, aType models.ActivityEnum) (string, error) {

	var id string

	stmt := `SELECT id FROM day_stats WHERE date = $1`

	err := a.Db.QueryRow(stmt, date.Format(time.DateOnly)).Scan(&id)

	if err == sql.ErrNoRows {

		stmt = `INSERT INTO day_stats (id, date) VALUES ($1, $2) RETURNING id`

		err = a.Db.QueryRow(
			stmt,
			uuid.New(),
			time.Now(),
		).Scan(&id)

		if err != nil {
			return "", err
		}
	}

	switch aType {
	case models.Study:
		stmt = `UPDATE day_stats SET study = study + $1 WHERE ID = $2`
	case models.ProgrammingWork:
		stmt = `UPDATE day_stats SET programming_work = programming_work + $1 WHERE ID = $2`
	case models.ProgrammingHobby:
		stmt = `UPDATE day_stats SET programming_hobby = programming_hobby + $1 WHERE ID = $2`
	case models.ReadStudy:
		stmt = `UPDATE day_stats SET read_study = read_study + $1 WHERE ID = $2`
	case models.ReadFun:
		stmt = `UPDATE day_stats SET read_fun = read_fun + $1 WHERE ID = $2`
	case models.Garbage:
		stmt = `UPDATE day_stats SET garbage = garbage + $1 WHERE ID = $2`
	}

	_, err = a.Db.Exec(stmt, duration, id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (a *DayStatsModel) GetByYear(year int) ([]models.DayStats, error) {

	list := []models.DayStats{}

	stmt := `SELECT date, study, programming_work, programming_hobby, read_study, read_fun, garbage
		FROM day_stats 
		WHERE date_part('year', date) = $1`

	rows, err := a.Db.Query(stmt, year)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stat models.DayStats

		if err := rows.Scan(
			&stat.Date,
			&stat.Study,
			&stat.ProgrammingWork,
			&stat.ProgrammingHobby,
			&stat.ReadStudy,
			&stat.ReadFun,
			&stat.Garbage,
		); err != nil {
			return nil, err
		}
		list = append(list, stat)
	}

	return list, nil
}
