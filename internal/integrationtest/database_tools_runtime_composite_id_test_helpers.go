package integrationtest

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func checkImportedDatabaseToolsRuntimeCompositeID(resourceName string, parser func(string) (map[string]string, error)) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		if len(states) == 0 {
			return fmt.Errorf("%s: expected at least one imported state", resourceName)
		}

		importedState := states[0]
		parsedAttributes, err := parser(importedState.ID)
		if err != nil {
			return fmt.Errorf("%s: failed to parse imported composite id %q: %w", resourceName, importedState.ID, err)
		}

		for attributeName, expectedValue := range parsedAttributes {
			actualValue, exists := importedState.Attributes[attributeName]
			if !exists {
				return fmt.Errorf("%s: imported state is missing attribute %q", resourceName, attributeName)
			}
			if actualValue != expectedValue {
				return fmt.Errorf("%s: imported attribute %q expected %q, got %q", resourceName, attributeName, expectedValue, actualValue)
			}
		}

		return nil
	}
}

func parseDatabaseToolsRuntimeCompositeIDToAttributes(compositeID string, pattern string, indexes map[string]int, expectedParts int) (map[string]string, error) {
	parts := strings.Split(compositeID, "/")
	match, _ := regexp.MatchString(pattern, compositeID)
	if !match || len(parts) != expectedParts {
		return nil, fmt.Errorf("illegal compositeId %s encountered", compositeID)
	}

	attributes := map[string]string{}
	for attributeName, index := range indexes {
		attributes[attributeName], _ = url.PathUnescape(parts[index])
	}

	return attributes, nil
}

func parseDatabaseToolsRuntimeConnectionCredentialCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsConnections/.*/credentials/.*",
		map[string]int{
			"database_tools_connection_id": 1,
			"key":                          3,
		},
		4,
	)
}

func parseDatabaseToolsRuntimeConnectionCredentialPublicSynonymCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsConnections/.*/credentials/.*/publicSynonyms/.*",
		map[string]int{
			"database_tools_connection_id": 1,
			"credential_key":               3,
			"key":                          5,
		},
		6,
	)
}

func parseDatabaseToolsRuntimeConnectionCredentialExecuteGranteeCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsConnections/.*/credentials/.*/executeGrantees/.*",
		map[string]int{
			"database_tools_connection_id": 1,
			"credential_key":               3,
			"key":                          5,
		},
		6,
	)
}

func parseDatabaseToolsRuntimeConnectionPropertySetCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsConnections/.*/propertySets/.*",
		map[string]int{
			"database_tools_connection_id": 1,
			"property_set_key":             3,
		},
		4,
	)
}

func parseDatabaseToolsRuntimeDatabaseApiGatewayConfigGlobalCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsDatabaseApiGatewayConfigs/.*/globals/.*",
		map[string]int{
			"database_tools_database_api_gateway_config_id": 1,
			"global_key": 3,
		},
		4,
	)
}

func parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*",
		map[string]int{
			"database_tools_database_api_gateway_config_id": 1,
			"key": 3,
		},
		4,
	)
}

func parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolApiSpecCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*/apiSpecs/.*",
		map[string]int{
			"database_tools_database_api_gateway_config_id": 1,
			"pool_key": 3,
			"key":      5,
		},
		6,
	)
}

func parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolAutoApiSpecCompositeIDToAttributes(compositeID string) (map[string]string, error) {
	return parseDatabaseToolsRuntimeCompositeIDToAttributes(
		compositeID,
		"databaseToolsDatabaseApiGatewayConfigs/.*/pools/.*/autoApiSpecs/.*",
		map[string]int{
			"database_tools_database_api_gateway_config_id": 1,
			"pool_key": 3,
			"key":      5,
		},
		6,
	)
}
