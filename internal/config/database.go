package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}

func NewDatabase(env *Env, log *logrus.Logger) *gorm.DB {
	var db *gorm.DB
	var err error

	switch env.DatabaseDriver {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			env.DatabaseHost, env.DatabaseUsername, env.DatabasePassword, env.DatabaseName, env.DatabasePort, env.DatabaseSSLMode, env.DatabaseTimeZone,
		)

		config := gorm.Config{
			Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
				SlowThreshold:             time.Second * 5,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				LogLevel:                  logger.Info,
			}),
			NowFunc: func() time.Time {
				loc, _ := time.LoadLocation(env.DatabaseTimeZone)
				return time.Now().In(loc)
			},
		}

		db, err = gorm.Open(postgres.Open(dsn), &config)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		connection, err := db.DB()
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		connection.SetMaxIdleConns(env.MaxIdleConns)
		connection.SetConnMaxIdleTime(time.Duration(env.IdleTimeout) * time.Second)
		connection.SetMaxOpenConns(env.MaxOpenConns)
		connection.SetConnMaxLifetime(time.Duration(env.MaxLifetime) * time.Second)
	}

	return db
}
