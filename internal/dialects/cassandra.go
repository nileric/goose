package dialects

import (
	"github.com/pressly/goose/v3/database/dialect"
)

// NewCassandra returns a new [dialect.Querier] for Cassandra dialect.
func NewCassandra() dialect.Querier {
	return &cassandra{}
}

type cassandra struct{}

var _ dialect.Querier = (*cassandra)(nil)

func (c *cassandra) CreateTable(tableName string) string {
	panic("unimplemented")
}

func (c *cassandra) DeleteVersion(tableName string) string {
	panic("unimplemented")
}

func (c *cassandra) GetLatestVersion(tableName string) string {
	panic("unimplemented")
}

func (c *cassandra) GetMigrationByVersion(tableName string) string {
	panic("unimplemented")
}

func (c *cassandra) InsertVersion(tableName string) string {
	panic("unimplemented")
}

func (c *cassandra) ListMigrations(tableName string) string {
	panic("unimplemented")
}
