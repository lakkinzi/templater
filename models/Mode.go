package models

type Mode string

const (
	BuilderMode  Mode = "Builder"
	MigratorMode Mode = "Migrator"
)

func (mode Mode) String() string {
	switch mode {
	case BuilderMode:
		return "Builder"
	case MigratorMode:
		return "Migrator"
	}
	return "Builder"
}

func WriteMode(mode string) Mode {
	switch mode {
	case "Builder":
		return BuilderMode
	case "Migrator":
		return MigratorMode
	}
	return BuilderMode
}

func GetModels() []string {
	return []string{BuilderMode.String(), MigratorMode.String()}
}
