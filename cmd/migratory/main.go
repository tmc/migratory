package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmc/migratory"
)

var (
	flagMigrationsPath = flag.String("dir", "migrations", "path to migrations directory")
	flagDSN            = flag.String("dsn", "", "DSN for database connection")
	flagTarget         = flag.String("target", "latest", "target database state to move to")
	flagDryRun         = flag.Bool("dryrun", false, "dry run mode")
)

func main() {
	flag.Parse()

	if *flagDSN == "" {
		*flagDSN = os.Getenv("MIGRATORY_DSN")
	}
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if *flagDSN == "" {
		return fmt.Errorf("migratory: missing DSN")
	}
	m, err := migratory.New(*flagMigrationsPath, *flagDSN, *flagTarget)
	if err != nil {
		return err
	}
	if err := m.Run(*flagDryRun); err != nil {
		return err
	}
	return nil
}
