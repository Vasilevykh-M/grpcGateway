package db

import (
	"awesomeProject/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context, path *config.ConnStructDB) (*Database, error) {
	dsn := generateDsn(path)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}

func generateDsn(path *config.ConnStructDB) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", path.Host, path.Port, path.User, path.Password, path.Dbname)
}
