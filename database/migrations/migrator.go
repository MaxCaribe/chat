package migrations

import (
	"chat/database"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const filePath = "./database/migrations"

var createDBVersionsTable = `
	CREATE TABLE IF NOT EXISTS db_versions (
		version varchar(14)
	)
`

var selectMigratedVersions = `
	SELECT * FROM db_versions
`

var insertDBVersion = `
	INSERT INTO db_versions(version) VALUES ($1)
`

func Migrate() {
	pool := database.Connection()
	ctx := context.Background()

	pool.Exec(ctx, createDBVersionsTable)
	versions, _ := pool.Query(ctx, selectMigratedVersions)
	defer versions.Close()

	files, _ := os.ReadDir(filePath)
	versionsDB := rowsToVersions(versions)
	slices.Sort(versionsDB)
	for _, file := range uniqMigrations(versionsDB, files) {
		sql, _ := os.ReadFile(filePath + "/" + file.Name())
		pool.Exec(ctx, string(sql))
		pool.Exec(ctx, insertDBVersion, versionFromFile(file))
	}
}

func uniqMigrations(versionsDB []string, filesMigration []os.DirEntry) []os.DirEntry {
	allVersions := make(map[string]os.DirEntry)
	for _, file := range filesMigration {
		version := versionFromFile(file)
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}
		allVersions[version] = file
	}
	for _, version := range versionsDB {
		delete(allVersions, version)
	}
	var migrations []os.DirEntry
	for _, file := range allVersions {
		migrations = append(migrations, file)
	}

	return migrations
}

func rowsToVersions(rows pgx.Rows) []string {
	var versionsDB []string
	var memo string
	for rows.Next() {
		err := rows.Scan(memo)
		if err != nil {
			log.Fatal(err)
		}
		versionsDB = append(versionsDB, memo)
	}
	return versionsDB
}

func versionFromFile(file os.DirEntry) string {
	version, _, _ := strings.Cut(file.Name(), "_")
	return version
}
