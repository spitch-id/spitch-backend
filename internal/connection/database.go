package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spitch-id/spitch-backend/internal/config"
)

func NewDatabase(env *config.Env) *pgxpool.Pool {
	// dbCfg := cfg.Database

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		env.DATABASE_USER,
		env.DATABASE_PASS,
		env.DATABASE_HOST,
		env.DATABASE_PORT,
		env.DATABASE_NAME,
		env.DATABASE_SSLMODE,
		env.DATABASE_TIMEZONE,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse dsn: %v", err)
	}

	config.MaxConns = int32(env.DATABASE_MAX_CONNECTIONS)
	config.MinConns = int32(env.DATABASE_POOL_IDLE)
	config.MaxConnLifetime = time.Second * time.Duration(env.DATABASE_MAXLIFETIME_CONNECTIONS)

	// config.ConnConfig.Tracer = NewPgxTracer(log)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Infof("Connected to PostgreSQL at %s:%s", env.DATABASE_HOST, env.DATABASE_PORT)
	return pool
}

// func NewPgxTracer(log *logger) *tracelog.TraceLog {
// 	return &tracelog.TraceLog{
// 		Logger: tracelog.LoggerFunc(func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
// 			switch level {
// 			case tracelog.LogLevelDebug:
// 				log.WithFields(logrus.Fields(data)).Debug(msg)
// 			case tracelog.LogLevelInfo:
// 				log.WithFields(logrus.Fields(data)).Info(msg)
// 			case tracelog.LogLevelWarn:
// 				log.WithFields(logrus.Fields(data)).Warn(msg)
// 			case tracelog.LogLevelError:
// 				log.WithFields(logrus.Fields(data)).Error(msg)
// 			}
// 		}),
// 		LogLevel: tracelog.LogLevelInfo,
// 	}
// }
