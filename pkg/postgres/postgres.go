package postgres

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

func New(url string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	pg := &Postgres{db}
	return pg, nil
}

func (p *Postgres) Close() {
	if p != nil {
		p.Close()
	}
}
