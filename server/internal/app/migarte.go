package app

import (
	"errors"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 10
	defaultTimeout  = time.Second * 5
)

func init() {
	dbURL, ok := os.LookupEnv("PG_URL")
	logrus.Info(dbURL)
	if !ok || len(dbURL) == 0 {
		logrus.Fatal("migrate: env variable not declared: PG_URL")
		return
	}

	dbURL += "?sslmode=disable"

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", dbURL)
		if err == nil {
			break
		}

		logrus.Printf("Migrate: pgdb is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		logrus.Fatalf("Migrate: pgdb connect error: %s", err)
		return
	}

	err = m.Up()
	defer func() { _, _ = m.Close() }()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Fatalf("Migrate: up error: %s", err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		logrus.Printf("Migrate: no change")
		return
	}

	logrus.Printf("Migrate: up success")
}
