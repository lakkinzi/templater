package models

type ServerBuildOperation string
type ServerBuildOperations []ServerBuildOperation

const (
	CreateModel       ServerBuildOperation = "CreateModel"
	CreateApi         ServerBuildOperation = "CreateApi"
	CreateRouting     ServerBuildOperation = "CreateRouting"
	CreateServiceMany ServerBuildOperation = "CreateServiceMany"
)

func (ServerBuildOperation ServerBuildOperation) String() string {
	switch ServerBuildOperation {
	case CreateModel:
		return "CreateModel"
	case CreateApi:
		return "CreateApi"
	case CreateRouting:
		return "CreateRouting"
	case CreateServiceMany:
		return "CreateServiceMany"
	}
	return "CreateModel"
}

func GetServerBuildOperations() []string {
	return []string{CreateModel.String(), CreateApi.String(), CreateRouting.String(), CreateServiceMany.String()}
}

func StringsToServerBuildOperations(operationsStrings []string) ServerBuildOperations {
	var operations ServerBuildOperations
	for _, operationString := range operationsStrings {
		if CreateModel.String() == operationString {
			operations = append(operations, CreateModel)
		}
		if CreateApi.String() == operationString {
			operations = append(operations, CreateApi)
		}
		if CreateRouting.String() == operationString {
			operations = append(operations, CreateRouting)
		}
		if CreateServiceMany.String() == operationString {
			operations = append(operations, CreateServiceMany)
		}
	}
	return operations
}
