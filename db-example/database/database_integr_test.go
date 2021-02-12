// +build integration

package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	// WARN: really important otherwise "database/sql" is not able to find the "postgres" driver and test fails!
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/bygui86/go-testing/db-example/logging"
)

func TestNew_Integr_Success(t *testing.T) {
	logging.SugaredLog.Debugf("DB_HOST: %s", os.Getenv("DB_HOST"))
	logging.SugaredLog.Debugf("DB_PORT: %s", os.Getenv("DB_PORT"))
	logging.SugaredLog.Debugf("DB_USERNAME: %s", os.Getenv("DB_USERNAME"))
	logging.SugaredLog.Debugf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
	logging.SugaredLog.Debugf("DB_NAME: %s", os.Getenv("DB_NAME"))
	logging.SugaredLog.Debugf("DB_SSL_MODE: %s", os.Getenv("DB_SSL_MODE"))

	db, dbErr := New()
	assert.NoError(t, dbErr)

	logging.SugaredLog.Debugf("[pre-ping] DB conn open: %d", db.Stats().OpenConnections)
	logging.SugaredLog.Debugf("[pre-ping] DB conn idle: %d", db.Stats().Idle)
	logging.SugaredLog.Debugf("[pre-ping] DB conn inUse: %d", db.Stats().InUse)

	pingErr := ping(db)
	assert.NoError(t, pingErr)

	logging.SugaredLog.Debugf("[post-ping] DB conn open: %d", db.Stats().OpenConnections)
	logging.SugaredLog.Debugf("[post-ping] DB conn idle: %d", db.Stats().Idle)
	logging.SugaredLog.Debugf("[post-ping] DB conn inUse: %d", db.Stats().InUse)

	assert.Equal(t, 1, db.Stats().OpenConnections)
	assert.Equal(t, 1, db.Stats().Idle)
}

func TestNewWithWrappedTracing_Integr_Success(t *testing.T) {
	logging.SugaredLog.Debugf("DB_HOST: %s", os.Getenv("DB_HOST"))
	logging.SugaredLog.Debugf("DB_PORT: %s", os.Getenv("DB_PORT"))
	logging.SugaredLog.Debugf("DB_USERNAME: %s", os.Getenv("DB_USERNAME"))
	logging.SugaredLog.Debugf("DB_PASSWORD: %s", os.Getenv("DB_PASSWORD"))
	logging.SugaredLog.Debugf("DB_NAME: %s", os.Getenv("DB_NAME"))
	logging.SugaredLog.Debugf("DB_SSL_MODE: %s", os.Getenv("DB_SSL_MODE"))

	db, dbErr := NewWithWrappedTracing()
	assert.NoError(t, dbErr)

	logging.SugaredLog.Debugf("[pre-ping] DB conn open: %d", db.Stats().OpenConnections)
	logging.SugaredLog.Debugf("[pre-ping] DB conn idle: %d", db.Stats().Idle)
	logging.SugaredLog.Debugf("[pre-ping] DB conn inUse: %d", db.Stats().InUse)

	pingErr := ping(db)
	assert.NoError(t, pingErr)

	logging.SugaredLog.Debugf("[post-ping] DB conn open: %d", db.Stats().OpenConnections)
	logging.SugaredLog.Debugf("[post-ping] DB conn idle: %d", db.Stats().Idle)
	logging.SugaredLog.Debugf("[post-ping] DB conn inUse: %d", db.Stats().InUse)

	assert.Equal(t, 1, db.Stats().OpenConnections)
	assert.Equal(t, 1, db.Stats().Idle)
}

func TestInitDb_Integr_Success(t *testing.T) {
	db, dbErr := New()
	require.NoError(t, dbErr)

	pingErr := ping(db)
	require.NoError(t, pingErr)

	initErr := InitDb(db)
	assert.NoError(t, initErr)

	_, tableErr := db.Exec("SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'products'")
	assert.NoError(t, tableErr)
}
