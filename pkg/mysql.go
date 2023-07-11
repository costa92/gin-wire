package pkg

import (
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

type MySQLConfig struct {
	Dsn          string `toml:"dsn" yaml:"dsn" json:"dsn" validate:"required"` // data source name
	MaxOpenCount int    `toml:"max_open_count" json:"max_open_count" validate:"required" yaml:"max_open_count"`
	MaxIdleCount int    `toml:"max_idle_count" json:"max_idle_count" validate:"required" yaml:"max_idle_count"`
	Tracing      bool   `toml:"tracing" json:"tracing" yaml:"tracing"`
}

// InitGormV2 open gorm v2 DB conn
// it is safe to use `nil` for gormLogger param, in this case,
// gorm v2 will use default gormlogger.Default impl
// nolint: gocritic
func InitGormV2(cfg *MySQLConfig) (*gorm.DB, error) {
	opt, err := mysql.ParseDSN(cfg.Dsn)
	if err != nil {
		return nil, err
	}

	if !opt.ParseTime {
		log.Println("InitGormV2: parseTime is disabled")
	}

	if opt.Loc.String() != "UTC" {
		log.Println("using non UTC timezone for parseTime", "timezone_used", opt.Loc.String())
	} else {
		log.Println("using UTC timezone for parseTime")
	}

	// if opt.Collation == "" {
	// 	opt.Collation = "utf8mb4"
	// }
	// dsn := opt.FormatDSN()

	db, err := gorm.Open(gormmysql.Open(cfg.Dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if cfg.Tracing {
		log.Println("init gorm tracing")

		plugin := otelgorm.NewPlugin(
			otelgorm.WithDBName(opt.DBName),
			otelgorm.WithAttributes(semconv.DBSystemMySQL, attribute.String("db.addr", opt.Addr)),
		)
		db.Use(plugin)
	}

	// otelConfig conn pool https://gorm.io/docs/connecting_to_the_database.html#Connection-Pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	if cfg.MaxIdleCount > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleCount)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	if cfg.MaxOpenCount > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenCount)
	}

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	// the maximum amount of time a connection may be idle
	// SetConnMaxIdleTime; added in Go 1.15
	// sqlDB.SetConnMaxIdleTime(time.Second * 3600)

	return db, nil
}
