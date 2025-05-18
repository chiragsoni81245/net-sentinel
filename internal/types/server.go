package types

import (
	"database/sql"

	"github.com/chiragsoni81245/net-sentinel/internal/config"
)

type Server struct {
    Config *config.Config
    DB     *sql.DB
}
