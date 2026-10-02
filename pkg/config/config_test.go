package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ygpark2/njro/pkg/config"
)

func TestGetConfig_Defaults(t *testing.T) {
	cfg := config.GetConfig()
	require.NotNil(t, cfg.Database)
	assert.NotEmpty(t, cfg.Database.Dialect)

	dsn, err := cfg.Database.DSN()
	require.NoError(t, err)
	assert.NotEmpty(t, dsn)
}

func TestDatabase_DSN(t *testing.T) {
	// 1. Postgres
	pg := &config.DatabaseConfiguration{
		Dialect:  "postgres",
		Host:     "10.0.0.1",
		Port:     5432,
		Username: "postgres",
		Password: "secretpassword",
		Database: "mboard",
	}
	dsn, err := pg.DSN()
	require.NoError(t, err)
	assert.Equal(t, "host=10.0.0.1 port=5432 user=postgres dbname=mboard sslmode=disable password=secretpassword", dsn)

	// 2. SQLite
	sqlite := &config.DatabaseConfiguration{
		Dialect:  "sqlite3",
		Database: "file:test.db?cache=shared",
	}
	dsn, err = sqlite.DSN()
	require.NoError(t, err)
	assert.Equal(t, "file:test.db?cache=shared", dsn)

	// 3. MySQL
	mysql := &config.DatabaseConfiguration{
		Dialect:  "mysql",
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "password",
		Database: "mydb",
	}
	dsn, err = mysql.DSN()
	require.NoError(t, err)
	assert.Equal(t, "root:password@(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local", dsn)
}

func TestEnvOverride(t *testing.T) {
	os.Setenv("DB_HOST", "custom-host.internal")
	os.Setenv("DB_USER", "custom-user")
	t.Cleanup(func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_USER")
		config.Load()
	})

	config.Load()
	cfg := config.GetConfig()
	assert.Equal(t, "custom-host.internal", cfg.Database.Host)
	assert.Equal(t, "custom-user", cfg.Database.Username)
}

func TestParseTargetString(t *testing.T) {
	for _, test := range []struct {
		targetStr string
		want      config.ParsedTarget
	}{
		{targetStr: "", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: ""}},
		{targetStr: ":///", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: ""}},
		{targetStr: "a:///", want: config.ParsedTarget{Scheme: "a", Authority: "", Endpoint: ""}},
		{targetStr: "://a/", want: config.ParsedTarget{Scheme: "", Authority: "a", Endpoint: ""}},
		{targetStr: ":///a", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a"}},
		{targetStr: "a://b/", want: config.ParsedTarget{Scheme: "a", Authority: "b", Endpoint: ""}},
		{targetStr: "a:///b", want: config.ParsedTarget{Scheme: "a", Authority: "", Endpoint: "b"}},
		{targetStr: "://a/b", want: config.ParsedTarget{Scheme: "", Authority: "a", Endpoint: "b"}},
		{targetStr: "a://b/c", want: config.ParsedTarget{Scheme: "a", Authority: "b", Endpoint: "c"}},
		{targetStr: "dns:///google.com", want: config.ParsedTarget{Scheme: "dns", Authority: "", Endpoint: "google.com"}},
		{targetStr: "dns:///google.com:8080", want: config.ParsedTarget{Scheme: "dns", Authority: "", Endpoint: "google.com:8080"}},
		{targetStr: "dns://a.server.com/google.com", want: config.ParsedTarget{Scheme: "dns", Authority: "a.server.com", Endpoint: "google.com"}},
		{targetStr: "dns://a.server.com/google.com/?a=b", want: config.ParsedTarget{Scheme: "dns", Authority: "a.server.com", Endpoint: "google.com/?a=b"}},

		{targetStr: "/", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "/"}},
		{targetStr: "google.com", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "google.com"}},
		{targetStr: "google.com/?a=b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "google.com/?a=b"}},
		{targetStr: "/unix/socket/address", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "/unix/socket/address"}},
		{targetStr: "unix:///tmp/mysrv.sock", want: config.ParsedTarget{Scheme: "unix", Authority: "", Endpoint: "tmp/mysrv.sock"}},

		// If we can only parse part of the target.
		{targetStr: "://", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "://"}},
		{targetStr: "unix://domain", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "unix://domain"}},
		{targetStr: "a:b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a:b"}},
		{targetStr: "a/b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a/b"}},
		{targetStr: "a:/b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a:/b"}},
		{targetStr: "a//b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a//b"}},
		{targetStr: "a://b", want: config.ParsedTarget{Scheme: "", Authority: "", Endpoint: "a://b"}},
	} {
		got := config.ParseTarget(test.targetStr)
		assert.Equal(t, test.want, got, "parsing target %s", test.targetStr)
	}
}

func TestBuildInfo(t *testing.T) {
	info := config.GetBuildInfo()
	assert.Contains(t, info, "version")
	assert.Contains(t, info, "go version")
}
