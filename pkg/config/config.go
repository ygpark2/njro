package config

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	cfg        Configuration
	configLock sync.RWMutex
	once       sync.Once

	// Version and build variables populated by govvv at compile-time.
	Version    = "untouched"
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
)

const versionMsg = `
version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

// Configuration holds application-wide configuration.
type Configuration struct {
	Database *DatabaseConfiguration   `yaml:"database" json:"database"`
	Email    *EmailConfiguration      `yaml:"email" json:"email"`
	Services map[string]ServiceConfig `yaml:"services" json:"services"`
	Features *FeaturesConfiguration   `yaml:"features" json:"features"`
}

// DatabaseConfiguration holds RDBMS connection parameters.
type DatabaseConfiguration struct {
	Dialect      string `yaml:"dialect" json:"dialect"` // postgres, sqlite3, mysql
	Host         string `yaml:"host" json:"host"`
	Port         uint32 `yaml:"port" json:"port"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	Database     string `yaml:"database" json:"database"`
	Logging      bool   `yaml:"logging" json:"logging"`
	MaxOpenConns uint32 `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns uint32 `yaml:"max_idle_conns" json:"max_idle_conns"`
}

// DSN generates a database connection string based on dialect and credentials.
func (d *DatabaseConfiguration) DSN() (string, error) {
	if d == nil {
		return "", fmt.Errorf("database configuration is nil")
	}

	dialect := strings.ToLower(d.Dialect)
	if dialect == "" || dialect == "postgres" || dialect == "postgre" || dialect == "1" {
		port := d.Port
		if port == 0 {
			port = 5432
		}
		host := d.Host
		if host == "" {
			host = "localhost"
		}
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			host, port, d.Username, d.Database, d.Password,
		), nil
	}

	if dialect == "sqlite" || dialect == "sqlite3" || dialect == "3" {
		if d.Database == "" {
			return "file::memory:?cache=shared&_fk=1", nil
		}
		return d.Database, nil
	}

	if dialect == "mysql" || dialect == "2" {
		port := d.Port
		if port == 0 {
			port = 3306
		}
		return fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			d.Username, d.Password, d.Host, port, d.Database,
		), nil
	}

	return "", fmt.Errorf("unsupported database dialect: %s", d.Dialect)
}

// EmailConfiguration holds mail server credentials.
type EmailConfiguration struct {
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	EmailServer string `yaml:"email_server" json:"email_server"`
	Port        uint32 `yaml:"port" json:"port"`
	From        string `yaml:"from" json:"from"`
}

// ServiceConfig holds service discovery information.
type ServiceConfig struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Version  string `yaml:"version" json:"version"`
	Deadline uint32 `yaml:"deadline" json:"deadline"`
}

// FeaturesConfiguration holds toggleable features like TLS, metrics, tracing.
type FeaturesConfiguration struct {
	TLS struct {
		Enabled    bool   `yaml:"enabled" json:"enabled"`
		CertFile   string `yaml:"cert_file" json:"cert_file"`
		KeyFile    string `yaml:"key_file" json:"key_file"`
		CaFile     string `yaml:"ca_file" json:"ca_file"`
		Servername string `yaml:"servername" json:"servername"`
	} `yaml:"tls" json:"tls"`
	Metrics struct {
		Enabled bool   `yaml:"enabled" json:"enabled"`
		Address string `yaml:"address" json:"address"`
	} `yaml:"metrics" json:"metrics"`
}

func init() {
	Load()
}

// Load reads configuration from YAML file or environment overrides.
func Load() {
	configLock.Lock()
	defer configLock.Unlock()

	// Default fallback config
	cfg = Configuration{
		Database: &DatabaseConfiguration{
			Dialect:  "sqlite3",
			Database: "file::memory:?cache=shared&_fk=1",
			Host:     "",
			Port:     5432,
		},
		Email: &EmailConfiguration{
			EmailServer: "smtp.gmail.com",
			Port:        587,
		},
		Services: make(map[string]ServiceConfig),
		Features: &FeaturesConfiguration{},
	}

	// Determine config file path
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = os.Getenv("CONFIGOR_FILE_PATH")
	}

	candidates := []string{
		configPath,
		"config/config.yaml",
		"/config/config.yaml",
		"../config/config.yaml",
		"../../config/config.yaml",
	}

	for _, p := range candidates {
		if p == "" {
			continue
		}
		data, err := os.ReadFile(p)
		if err == nil {
			_ = yaml.Unmarshal(data, &cfg)
			break
		}
	}

	// Environment variable overrides
	if host := os.Getenv("DB_HOST"); host != "" {
		cfg.Database.Host = host
	}
	if user := os.Getenv("DB_USER"); user != "" {
		cfg.Database.Username = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		cfg.Database.Password = pass
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		cfg.Database.Database = name
	}
	if dialect := os.Getenv("DB_DIALECT"); dialect != "" {
		cfg.Database.Dialect = dialect
	}
}

// GetConfig returns the active configuration.
func GetConfig() Configuration {
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg
}

// GetBuildInfo returns compile-time metadata.
func GetBuildInfo() string {
	return fmt.Sprintf(
		versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler,
		runtime.GOOS, runtime.GOARCH, GitCommit, GitBranch, GitState, GitSummary,
	)
}

// ParsedTarget represents a structured network target.
type ParsedTarget struct {
	Scheme    string
	Authority string
	Endpoint  string
}

func split2(s, sep string) (string, string, bool) {
	spl := strings.SplitN(s, sep, 2)
	if len(spl) < 2 {
		return "", "", false
	}
	return spl[0], spl[1], true
}

// ParseTarget splits target into Scheme, Authority, and Endpoint.
func ParseTarget(target string) (ret ParsedTarget) {
	var ok bool
	ret.Scheme, ret.Endpoint, ok = split2(target, "://")
	if !ok {
		return ParsedTarget{Endpoint: target}
	}
	ret.Authority, ret.Endpoint, ok = split2(ret.Endpoint, "/")
	if !ok {
		return ParsedTarget{Endpoint: target}
	}
	return ret
}
