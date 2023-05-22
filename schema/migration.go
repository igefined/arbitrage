package schema

import (
	"embed"
	"fmt"
	"regexp"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/igdotog/core/logger"
)

func Migrate(log *logger.Logger, fs *embed.FS, dbUrl string) {
	source, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Errorf("Failed to read migrations source: %s", err)

		return
	}

	instance, err := migrate.NewWithSourceInstance("iofs", source, makeMigrateUrl(dbUrl))
	if err != nil {
		log.Errorf("Failed to initialization the migrate instance: %s", err)

		return
	}

	err = instance.Up()

	switch err {
	case nil:
		log.Infof("The migration schema: The schema successfully upgraded!")
	case migrate.ErrNoChange:
		log.Infof("The migration schema: The schema not changed")
	default:
		log.Errorf("Could not apply the migration schema: %s", err)
	}
}

func makeMigrateUrl(dbUrl string) string {
	urlRe := regexp.MustCompile(`^[^\\?]+`)
	url := urlRe.FindString(dbUrl)

	sslModeRe := regexp.MustCompile("(sslmode=)[a-zA-Z0-9]+")
	sslMode := sslModeRe.FindString(dbUrl)

	return fmt.Sprintf("%s?%s", url, sslMode)
}
