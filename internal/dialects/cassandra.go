package dialects

import (
	"fmt"

	"github.com/pressly/goose/v3/database/dialect"
)

// NewCassandra returns a new [dialect.Querier] for Cassandra dialect.
func NewCassandra() dialect.Querier {
	return &cassandra{}
}

type cassandra struct{}

var _ dialect.Querier = (*cassandra)(nil)

func (c *cassandra) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		pk int,
		version_id bigint,
		is_applied boolean,
		tstamp timestamp,
		PRIMARY KEY (pk, version_id)
	) WITH CLUSTERING ORDER BY (version_id DESC)`
	return fmt.Sprintf(q, tableName)
}

func (c *cassandra) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (pk, version_id, is_applied, tstamp) VALUES (1, ?, ?, toTimestamp(now()))`
	return fmt.Sprintf(q, tableName)
}

func (c *cassandra) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE pk=1 AND version_id=?`
	return fmt.Sprintf(q, tableName)
}

func (c *cassandra) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE pk=1 AND version_id=? LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

func (c *cassandra) ListMigrations(tableName string) string {
	// Due to CLUSTERING ORDER BY (version_id DESC), this returns descending order.
	q := `SELECT version_id, is_applied FROM %s WHERE pk=1`
	return fmt.Sprintf(q, tableName)
}

func (c *cassandra) GetLatestVersion(tableName string) string {
	q := `SELECT max(version_id) FROM %s WHERE pk=1`
	return fmt.Sprintf(q, tableName)
}
