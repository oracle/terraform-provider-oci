// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// ConnectionTypeEnum Enum with underlying type: string
type ConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionTypeEnum
const (
	ConnectionTypeGoldengate            ConnectionTypeEnum = "GOLDENGATE"
	ConnectionTypeKafka                 ConnectionTypeEnum = "KAFKA"
	ConnectionTypeKafkaSchemaRegistry   ConnectionTypeEnum = "KAFKA_SCHEMA_REGISTRY"
	ConnectionTypeMysql                 ConnectionTypeEnum = "MYSQL"
	ConnectionTypeOciObjectStorage      ConnectionTypeEnum = "OCI_OBJECT_STORAGE"
	ConnectionTypeOracle                ConnectionTypeEnum = "ORACLE"
	ConnectionTypeAzureDataLakeStorage  ConnectionTypeEnum = "AZURE_DATA_LAKE_STORAGE"
	ConnectionTypePostgresql            ConnectionTypeEnum = "POSTGRESQL"
	ConnectionTypeAzureSynapseAnalytics ConnectionTypeEnum = "AZURE_SYNAPSE_ANALYTICS"
)

var mappingConnectionTypeEnum = map[string]ConnectionTypeEnum{
	"GOLDENGATE":              ConnectionTypeGoldengate,
	"KAFKA":                   ConnectionTypeKafka,
	"KAFKA_SCHEMA_REGISTRY":   ConnectionTypeKafkaSchemaRegistry,
	"MYSQL":                   ConnectionTypeMysql,
	"OCI_OBJECT_STORAGE":      ConnectionTypeOciObjectStorage,
	"ORACLE":                  ConnectionTypeOracle,
	"AZURE_DATA_LAKE_STORAGE": ConnectionTypeAzureDataLakeStorage,
	"POSTGRESQL":              ConnectionTypePostgresql,
	"AZURE_SYNAPSE_ANALYTICS": ConnectionTypeAzureSynapseAnalytics,
}

var mappingConnectionTypeEnumLowerCase = map[string]ConnectionTypeEnum{
	"goldengate":              ConnectionTypeGoldengate,
	"kafka":                   ConnectionTypeKafka,
	"kafka_schema_registry":   ConnectionTypeKafkaSchemaRegistry,
	"mysql":                   ConnectionTypeMysql,
	"oci_object_storage":      ConnectionTypeOciObjectStorage,
	"oracle":                  ConnectionTypeOracle,
	"azure_data_lake_storage": ConnectionTypeAzureDataLakeStorage,
	"postgresql":              ConnectionTypePostgresql,
	"azure_synapse_analytics": ConnectionTypeAzureSynapseAnalytics,
}

// GetConnectionTypeEnumValues Enumerates the set of values for ConnectionTypeEnum
func GetConnectionTypeEnumValues() []ConnectionTypeEnum {
	values := make([]ConnectionTypeEnum, 0)
	for _, v := range mappingConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionTypeEnum
func GetConnectionTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
		"KAFKA",
		"KAFKA_SCHEMA_REGISTRY",
		"MYSQL",
		"OCI_OBJECT_STORAGE",
		"ORACLE",
		"AZURE_DATA_LAKE_STORAGE",
		"POSTGRESQL",
		"AZURE_SYNAPSE_ANALYTICS",
	}
}

// GetMappingConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionTypeEnum(val string) (ConnectionTypeEnum, bool) {
	enum, ok := mappingConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
