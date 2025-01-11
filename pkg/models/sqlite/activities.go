package sqlite

import (
	"database/sql"
	"time"

	"nugu.dev/basement/pkg/models"
)

type ActivityModel struct {
	DB *sql.DB
}

func (a *ActivityModel) Insert(activityType models.ActivityEnum, startTime time.Time) (int, error) {
	return 0, nil
}

func (a *ActivityModel) SetEnd(endTime time.Time) (int, error) {
	return 0, nil
}

func (a *ActivityModel) GetByYear(year int) (int, error) {
	return 0, nil
}
