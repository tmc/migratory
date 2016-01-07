package migratory

type Plan struct {
	Forwards  []Migration
	Backwards []Migration
}
