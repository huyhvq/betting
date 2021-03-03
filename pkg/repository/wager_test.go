package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/huyhvq/betting/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type WagerSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository WagerRepository
	person     *model.Wager
}

func (s *WagerSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)
	s.repository = NewWager(s.DB)
}

func (s *WagerSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(WagerSuite))
}

func (s *WagerSuite) Test_repository_Get() {
	var (
		id  = 2
		twv = 1000
	)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `wagers`")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "total_wager_value"}).
			AddRow(id, twv))

	res, err := s.repository.GetByID(uint(id))

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.Wager{
		Model: gorm.Model{
			ID: uint(id),
		},
		TotalWagerValue: 1000,
	}, res))
}
