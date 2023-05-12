package main

import (
	"context"
	"example.com/m/v2/internal/auditlog"
	"example.com/m/v2/internal/goodrussians"
	"os/signal"
	"syscall"
)

func main() {
	var (
		ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

		config = MustReadConfig()
	)
	defer stop()

	var (
		mongoClient      = MustCreateMongoClient(config)
		mongoDatabase    = mongoClient.Database(config.Mongo.Database)
		auditLogsStorage = auditlog.NewStorageMongo(mongoDatabase.Collection(config.Mongo.AuditLogsCollection))
		auditLog         = auditlog.NewAuditLog(auditLogsStorage)
	)

	err := mongoClient.Connect(ctx)
	if err != nil {
		panic("connect to database")
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic("ping to database")
	}

	goodRussiansCounter := goodrussians.NewGoodRussiansCounter(auditLog)

	goodRussiansCounter.Count(ctx)
}
