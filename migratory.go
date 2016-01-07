package migratory

import (
	"database/sql"
	"fmt"
	"strings"
)

type Migratory struct {
	dir    string
	dsn    string
	target string

	db *sql.DB
}

func New(path, DSN, target string) (*Migratory, error) {
	return &Migratory{
		dir:    path,
		dsn:    DSN,
		target: target,
	}, nil
}

func (m *Migratory) Run(dryRun bool) error {
	if dryRun {
		return m.dryRun()
	}
	return m.run()
}

func (m *Migratory) run() error {
	return fmt.Errorf("migratory: not implemented")
}

func (m *Migratory) dryRun() error {
	if err := m.preFlight(); err != nil {
		return err
	}
	plan, err := m.constructPlan()
	if err != nil {
		return err
	}
	fmt.Println(plan)
	return nil
}

func (m *Migratory) preFlight() error {
	if err := m.preFlightDB(); err != nil {
		return err
	}
	return nil
}

func (m *Migratory) preFlightDB() error {
	driver, dataSource, err := splitDSN(m.dsn)
	if err != nil {
		return err
	}
	m.db, err = sql.Open(driver, dataSource)
	if err != nil {
		return err
	}
	return m.db.Ping()
}

func (m *Migratory) constructPlan() (*Plan, error) {
	return nil, fmt.Errorf("migratory: plan: not implemented")
}

func splitDSN(dsn string) (driverName string, dataSource string, err error) {
	parts := strings.Split(dsn, "://")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("migratory: expected '://' in DSN, got '%s'", dsn)
	}
	return parts[0], parts[1], nil

}
