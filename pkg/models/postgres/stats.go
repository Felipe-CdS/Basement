package postgres

import (
	"database/sql"
	"time"

	"nugu.dev/basement/pkg/models"
)

type DayStatsModel struct {
	Db *sql.DB
}

func (a *DayStatsModel) CreateEmpty(date time.Time) (models.DayStats, error) {

	id := 0
	s := models.DayStats{}

	stmt := `SELECT id FROM day_stats WHERE date = $1`

	err := a.Db.QueryRow(stmt, date.Format(time.DateOnly)).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return s, err
	}

	if id != 0 {
		return s, models.ErrAlreadyExists
	}

	stmt = `INSERT INTO day_stats (date) VALUES ($1) RETURNING *`

	err = a.Db.QueryRow(stmt, date.Format(time.DateOnly)).Scan(
		&s.ID,
		&s.Date,
		&s.Study,
		&s.ProgrammingWork,
		&s.ProgrammingHobby,
		&s.ReadStudy,
		&s.ReadFun,
		&s.Korean,
		&s.Garbage,
	)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (a *DayStatsModel) Insert(date time.Time, duration int, aType models.ActivityEnum) (models.DayStats, error) {

	id := 0
	s := models.DayStats{}

	stmt := `SELECT id FROM day_stats WHERE date = $1`

	err := a.Db.QueryRow(stmt, date.Format(time.DateOnly)).Scan(&id)

	if err == sql.ErrNoRows {
		s, err = a.CreateEmpty(date)
		if err != nil {
			return s, err
		}
		id = s.ID
	}

	switch aType {
	case models.Study:
		stmt = `UPDATE day_stats SET study = study + $1 WHERE ID = $2 RETURNING *`
	case models.ProgrammingWork:
		stmt = `UPDATE day_stats SET programming_work = programming_work + $1 WHERE ID = $2 RETURNING *`
	case models.ProgrammingHobby:
		stmt = `UPDATE day_stats SET programming_hobby = programming_hobby + $1 WHERE ID = $2 RETURNING *`
	case models.ReadStudy:
		stmt = `UPDATE day_stats SET read_study = read_study + $1 WHERE ID = $2 RETURNING *`
	case models.ReadFun:
		stmt = `UPDATE day_stats SET read_fun = read_fun + $1 WHERE ID = $2 RETURNING *`
	case models.Korean:
		stmt = `UPDATE day_stats SET korean = korean + $1 WHERE ID = $2 RETURNING *`
	case models.Garbage:
		stmt = `UPDATE day_stats SET garbage = garbage + $1 WHERE ID = $2 RETURNING *`
	default:
		return s, models.ErrActityTypeNotFound
	}

	err = a.Db.QueryRow(stmt, duration, id).Scan(
		&s.ID,
		&s.Date,
		&s.Study,
		&s.ProgrammingWork,
		&s.ProgrammingHobby,
		&s.ReadStudy,
		&s.ReadFun,
		&s.Korean,
		&s.Garbage,
	)

	if err != nil {
		return s, err
	}

	return s, nil
}

func (a *DayStatsModel) GetByYear(year int) ([]models.DayStats, error) {

	list := []models.DayStats{}

	stmt := `SELECT *
		FROM day_stats 
		WHERE date_part('year', date) = $1
		ORDER BY date ASC`

	rows, err := a.Db.Query(stmt, year)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stat models.DayStats

		if err := rows.Scan(
			&stat.ID,
			&stat.Date,
			&stat.Study,
			&stat.ProgrammingWork,
			&stat.ProgrammingHobby,
			&stat.ReadStudy,
			&stat.ReadFun,
			&stat.Korean,
			&stat.Garbage,
		); err != nil {
			return nil, err
		}
		list = append(list, stat)
	}

	return list, nil
}
