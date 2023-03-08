package database

import (
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"slack-application/ent"
	"slack-application/log"
)

var logger = log.GetNewLogger()
var Client *ent.Client

func SetUpDatabase() {
	var err error
	Client, err = ent.Open("postgres", "host=localhost port=5432 user=shishpal dbname=postgres password=shishpal sslmode=disable")
	if err != nil {
		logger.Error("Failed to open database: {}", zap.Error(err))
		return
	}
	defer Client.Close()
}
