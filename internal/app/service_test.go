package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/v1ack/linkShorter/internal/store"
	migration "github.com/v1ack/linkShorter/sql"
	"log"
	"testing"
)

type testSuite struct {
	suite.Suite
	store store.DataStore
}

func (t *testSuite) SetupSuite() {
	dbConnectionString := "dbname=shorter_test password=admin user=admin sslmode=disable"
	connection, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalln("Failed to init db:", err)
	}
	migration.RunMigrations(connection)
	t.store = store.CreateDbProvider(connection)
}

func TestName(t *testing.T) {
	suite.Run(t, new(testSuite))
}
