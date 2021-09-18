package models

type Operations struct {
	ServerOperations ServerOperations
}

type ServerOperations struct {
	ServerBuildOperations ServerBuildOperations
	MigrateOperations     MigrateOperations
}

type MigrateOperation string
type MigrateOperations []*MigrateOperation

const (
	CreateMigration MigrateOperation = "CreateMigration"
)
