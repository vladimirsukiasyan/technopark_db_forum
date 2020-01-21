package service

import (
	"log"

	"github.com/vladimirsukiasyan/technopark_db_forum/internal/models"
	"gopkg.in/jackc/pgx.v2"
)

const postgres = "postgres"

var Error models.Error

type ForumPgsql struct {
	db *pgx.ConnPool
}

func NewForumPgsql(config *pgx.ConnConfig) *ForumPgsql {
	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     *config,
		MaxConnections: 50,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}
	p, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &ForumPgsql{
		db: p,
	}
}
