package main

import (
	"fmt"
	"strings"

	"ariga.io/atlas-go-sdk/recordriver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IEnum interface {
	Name() string
	All() []string
}

// New returns a new Loader.
func New(dialect string) *Loader {
	return &Loader{dialect: dialect, di: postgres.New(postgres.Config{
		DriverName: "recordriver",
		// destination db
		DSN: "gorm",
		// DSN: "postgres://d2:d2@localhost/d2?sslmode=disable",
	})}
}

// Loader is a Loader for gorm schema.
type Loader struct {
	dialect string
	di      gorm.Dialector
}

func toString(list []string) string {
	var s string

	for i, v := range list {
		if i == len(list)-1 {
			s += fmt.Sprintf("'%s'", v)
			continue
		}

		s += fmt.Sprintf("'%s',", v)
	}

	return s
}

func (l *Loader) Load(modelsOrStatements ...any) (string, error) {
	db, err := gorm.Open(l.di, &gorm.Config{})
	if err != nil {
		return "", err
	}

	// loop over modelsOrStatements
	for i := 0; i < len(modelsOrStatements); i++ {
		switch v := modelsOrStatements[i].(type) {
		case IEnum:
			stmt := fmt.Sprintf(`
			DO $$ BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN
						CREATE TYPE %s AS ENUM (%s);
				END IF;
			EXCEPTION
				WHEN duplicate_object THEN null;
			END $$
			`, v.Name(), v.Name(), toString(v.All()))
			db.Exec(stmt)

		case string:
			db.Exec(v)

		default:
			if err := db.AutoMigrate(v); err != nil {
				return "", err
			}
		}
	}

	s, ok := recordriver.Session("gorm")

	if !ok {
		return "", err
	}

	stmtExistMap := make(map[string]bool)

	var sb strings.Builder

	for _, stmt := range s.Statements {
		if stmtExistMap[stmt] {
			continue
		}

		stmtExistMap[stmt] = true

		sb.WriteString(stmt)
		sb.WriteString(";\n")
	}

	return sb.String(), nil
}
