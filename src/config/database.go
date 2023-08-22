package config

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kiramishima/shining_guardian/src/domain"
	"go.uber.org/zap"
)

// NewDatabase creates an instance of DB
func NewDatabase(cfg *domain.Configuration, logger *zap.SugaredLogger) (*sqlx.DB, error) {

	db, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	// seteamos el numero maximo de conexiones abiertas. 0 indica sin limite
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// seteamos el numero maximo de conexiones inactivas. 0 indica sin limite
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// usamos time.ParseDuration para convertir el string de duracion a time.Duration
	duration, err := time.ParseDuration(cfg.MaxIdleTime)

	if err != nil {
		return nil, err
	}
	// Seteamos el timeout para las inactivas
	db.SetConnMaxIdleTime(duration)

	// creamos el contexto con 5 segundos de timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// PingContext
	status := "up"
	err = db.PingContext(ctx)
	if err != nil {
		status = "down"
		return nil, err
	}
	logger.Debugf("Status DB: %s", status)
	return db, nil
}
