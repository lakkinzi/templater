package models

type ValueType string

const (
	String ValueType = "String"
	Number ValueType = "Number"
	Date   ValueType = "Date"
)

func (i ValueType) String() string {
	switch i {
	case String:
		return "String"
	case Number:
		return "Number"
	case Date:
		return "Date"
	}

	return "String"
}

func WriteValueType(mode string) ValueType {
	switch mode {
	case "String":
		return String
	case "Number":
		return Number
	case "Date":
		return Date
	}
	return String
}

func GetValueTypesStrings() []string {
	return []string{String.String(), Number.String(), Date.String()}
}
