package store

import (
	"context"
	"database/sql"
)

type RoleStore struct {
	db *sql.DB
}

func (rs *RoleStore) GetRoleByName(ctx context.Context, slug string) (*Role, error) {
	query := `SELECT * FROM roles WHERE name=$1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	role := &Role{}

	err := rs.db.QueryRowContext(
		ctx,
		query,
		slug,
	).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
		&role.Description,
	)

	if err != nil {
		return nil, err
	}

	return role, nil

}
