package models

type Mode string

const (
	Builder  Mode = "Builder"
	Migrator Mode = "Migrator"
)

func (mode Mode) String() string {
	switch mode {
	case Builder:
		return "Builder"
	case Migrator:
		return "Migrator"
	}
	return "Builder"
}

func WriteMode(mode string) Mode {
	switch mode {
	case "Builder":
		return Builder
	case "Migrator":
		return Migrator
	}
	return Builder
}

func GetModels() []string {
	return []string{Builder.String(), Migrator.String()}
}
