package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(url string) (*Postgres, error) {
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		logrus.Fatal("Unable to parse DATABASE_URL:", err)
		os.Exit(1)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		logrus.Fatal("Unable to create connection pool", err)
		os.Exit(1)
	}

	pg := &Postgres{Pool: db}
	return pg, nil
}

func (p *Postgres) Close() {
	if p != nil {
		p.Close()
	}
}
