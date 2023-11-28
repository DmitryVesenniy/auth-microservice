package config

import "time"

const (
	DefaultEnv = ".env"

	EngineTypeSqlite   = "sqlite"
	EngineTypePostgres = "postgres"
	EngineTypeMySQL    = "mysql"
	EngineTypeRedis    = "redis"
	EngineTypeMongo    = "mongo"
)

type Config struct {
	MigrationsPath  string        `yaml:"migrations_path" env-default:"migrations"`
	Env             string        `yaml:"env" env-default:".env"`
	StorageFile     string        `yaml:"storage_path" env-required:"true"`
	TokenTTL        time.Duration `yaml:"token_ttl" env-default:"10m"`
	TokenExpiredTTL time.Duration `yaml:"token_expired_ttl" env-default:"24h"`
	GRPCServer      GRPCConfig    `yaml:"grpc"`
	Database        StorageConfig `yaml:"database"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type StorageConfig struct {
	EngineType string `yaml:"engine_type" env-default:"sqlite"`
	Name       string `yaml:"name"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
}
