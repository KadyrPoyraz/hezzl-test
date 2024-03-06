package postgres

import (
	"context"
	"path/filepath"
	"time"

	"github.com/KadyrPoyraz/httplayout/config"
	"github.com/KadyrPoyraz/httplayout/internal/repository/db"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/suite"
)

const pgTestDatabase = "test_coffee"
const pgMigrations = "../../../migrations"

type PostgresBaseRepositorySuite struct {
	suite.Suite

	cfg config.DBConfig
	db  *sqlx.DB

	TestPostgresClient *db.PostgresqlDB
}

func (s *PostgresBaseRepositorySuite) SetupSuite() {
	cnf, err := config.New()
	s.Require().NoError(err)

	s.cfg = cnf.DB

	s.TestPostgresClient, err = db.NewPostgresqlDB(s.cfg.TestDSN)
	s.Require().NoError(err)

	s.db = s.TestPostgresClient.DB
}

func (s *PostgresBaseRepositorySuite) SetupTest() {
	s.Require().NoError(goose.Up(s.db.DB, filepath.Join(pgMigrations, "tests")))
}

func (s *PostgresBaseRepositorySuite) TearDownTest() {
	s.Require().NoError(goose.Down(s.db.DB, filepath.Join(pgMigrations, "tests")))
}

func (s *PostgresBaseRepositorySuite) TearDownSuite() {
	s.Require().NoError(s.TestPostgresClient.Close())
}

type PostgresBaseRepositoryParallelSuite struct {
	suite.Suite

	transactionsRepo transactions.Repository

	ctx context.Context
	cfg config.Postgres
	db  *sqlx.DB

	TestPostgresClient *db.PostgresClient

	testTx transactions.Transaction
	time   time.Time
}

func (s *PostgresBaseRepositoryParallelSuite) SetupSuite() {
	cfg, err := config.NewConfig("../../../configs/local_tests")
	s.Require().NoError(err)
	s.cfg = cfg.Postgres

	testClient, err := db.NewPostgresClient(s.cfg.DSNTest, "")
	s.Require().NoError(err)
	s.TestPostgresClient = testClient
	s.db = s.TestPostgresClient.DB

	s.transactionsRepo = transactions.NewPgxRepository(s.TestPostgresClient)

	goose.SetLogger(&logger.EmptyLogger{})
}

func (s *PostgresBaseRepositoryParallelSuite) SetupTest() {
	s.ctx = context.Background()

	tx, err := s.transactionsRepo.StartTransaction(s.ctx)
	s.Require().NoError(err)
	s.testTx = tx
	s.time = time.Now()

	_, err = s.testTx.Txm().Exec(tablesToDelete)
	s.Require().NoError(err)
}

func (s *PostgresBaseRepositoryParallelSuite) TearDownTest() {
	s.testTx.Rollback()
}

func (s *PostgresBaseRepositoryParallelSuite) TearDownSuite() {
	s.Require().NoError(s.TestPostgresClient.Close())
}

func (s *PostgresBaseRepositoryParallelSuite) Tx() transactions.Transaction {
	return s.testTx
}
