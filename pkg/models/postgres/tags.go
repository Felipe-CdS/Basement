package postgres

import (
	"database/sql"
	"log"

	"nugu.dev/basement/pkg/models"
)

type TagRepository struct {
	Db *sql.DB
}

func (r *TagRepository) GetActivityTags() ([]models.Tag, error) {

	var tags []models.Tag

	stmt := `SELECT id, name FROM tags WHERE type = 'activity';`

	rows, err := r.Db.Query(stmt)

	if err != nil {
		if err == sql.ErrNoRows {
			return tags, models.ErrNotFound
		}
		log.Println(err)
		return tags, models.ErrDbOperation
	}

	for rows.Next() {
		var t models.Tag

		if err := rows.Scan(
			&t.ID,
			&t.Name,
		); err != nil {
			return tags, models.ErrDbOperation
		}

		tags = append(tags, t)

	}

	return tags, nil
}
