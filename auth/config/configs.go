package config

import database "github.com/muling3/auth/db/sqlc"

type AppConfig struct {
	Db *database.Queries
}
