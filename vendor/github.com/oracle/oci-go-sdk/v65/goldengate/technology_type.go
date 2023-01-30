// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TechnologyTypeEnum Enum with underlying type: string
type TechnologyTypeEnum string

// Set of constants representing the allowable values for TechnologyTypeEnum
const (
	TechnologyTypeGoldengate               TechnologyTypeEnum = "GOLDENGATE"
	TechnologyTypeOciAutonomousDatabase    TechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	TechnologyTypeOciMysql                 TechnologyTypeEnum = "OCI_MYSQL"
	TechnologyTypeOciObjectStorage         TechnologyTypeEnum = "OCI_OBJECT_STORAGE"
	TechnologyTypeOciStreaming             TechnologyTypeEnum = "OCI_STREAMING"
	TechnologyTypeOracleDatabase           TechnologyTypeEnum = "ORACLE_DATABASE"
	TechnologyTypeOracleExadata            TechnologyTypeEnum = "ORACLE_EXADATA"
	TechnologyTypeAmazonRdsOracle          TechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	TechnologyTypeAmazonAuroraMysql        TechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	TechnologyTypeAmazonAuroraPostgresql   TechnologyTypeEnum = "AMAZON_AURORA_POSTGRESQL"
	TechnologyTypeAmazonRdsMariadb         TechnologyTypeEnum = "AMAZON_RDS_MARIADB"
	TechnologyTypeAmazonRdsMysql           TechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	TechnologyTypeAmazonRdsPostgresql      TechnologyTypeEnum = "AMAZON_RDS_POSTGRESQL"
	TechnologyTypeApacheKafka              TechnologyTypeEnum = "APACHE_KAFKA"
	TechnologyTypeAzureDataLakeStorage     TechnologyTypeEnum = "AZURE_DATA_LAKE_STORAGE"
	TechnologyTypeAzureEventHubs           TechnologyTypeEnum = "AZURE_EVENT_HUBS"
	TechnologyTypeAzureMysql               TechnologyTypeEnum = "AZURE_MYSQL"
	TechnologyTypeAzurePostgresql          TechnologyTypeEnum = "AZURE_POSTGRESQL"
	TechnologyTypeAzureSynapseAnalytics    TechnologyTypeEnum = "AZURE_SYNAPSE_ANALYTICS"
	TechnologyTypeConfluentKafka           TechnologyTypeEnum = "CONFLUENT_KAFKA"
	TechnologyTypeConfluentSchemaRegistry  TechnologyTypeEnum = "CONFLUENT_SCHEMA_REGISTRY"
	TechnologyTypeGoogleCloudSqlMysql      TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	TechnologyTypeGoogleCloudSqlPostgresql TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_POSTGRESQL"
	TechnologyTypeMariadb                  TechnologyTypeEnum = "MARIADB"
	TechnologyTypeMysqlServer              TechnologyTypeEnum = "MYSQL_SERVER"
	TechnologyTypePostgresqlServer         TechnologyTypeEnum = "POSTGRESQL_SERVER"
)

var mappingTechnologyTypeEnum = map[string]TechnologyTypeEnum{
	"GOLDENGATE":                  TechnologyTypeGoldengate,
	"OCI_AUTONOMOUS_DATABASE":     TechnologyTypeOciAutonomousDatabase,
	"OCI_MYSQL":                   TechnologyTypeOciMysql,
	"OCI_OBJECT_STORAGE":          TechnologyTypeOciObjectStorage,
	"OCI_STREAMING":               TechnologyTypeOciStreaming,
	"ORACLE_DATABASE":             TechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":              TechnologyTypeOracleExadata,
	"AMAZON_RDS_ORACLE":           TechnologyTypeAmazonRdsOracle,
	"AMAZON_AURORA_MYSQL":         TechnologyTypeAmazonAuroraMysql,
	"AMAZON_AURORA_POSTGRESQL":    TechnologyTypeAmazonAuroraPostgresql,
	"AMAZON_RDS_MARIADB":          TechnologyTypeAmazonRdsMariadb,
	"AMAZON_RDS_MYSQL":            TechnologyTypeAmazonRdsMysql,
	"AMAZON_RDS_POSTGRESQL":       TechnologyTypeAmazonRdsPostgresql,
	"APACHE_KAFKA":                TechnologyTypeApacheKafka,
	"AZURE_DATA_LAKE_STORAGE":     TechnologyTypeAzureDataLakeStorage,
	"AZURE_EVENT_HUBS":            TechnologyTypeAzureEventHubs,
	"AZURE_MYSQL":                 TechnologyTypeAzureMysql,
	"AZURE_POSTGRESQL":            TechnologyTypeAzurePostgresql,
	"AZURE_SYNAPSE_ANALYTICS":     TechnologyTypeAzureSynapseAnalytics,
	"CONFLUENT_KAFKA":             TechnologyTypeConfluentKafka,
	"CONFLUENT_SCHEMA_REGISTRY":   TechnologyTypeConfluentSchemaRegistry,
	"GOOGLE_CLOUD_SQL_MYSQL":      TechnologyTypeGoogleCloudSqlMysql,
	"GOOGLE_CLOUD_SQL_POSTGRESQL": TechnologyTypeGoogleCloudSqlPostgresql,
	"MARIADB":                     TechnologyTypeMariadb,
	"MYSQL_SERVER":                TechnologyTypeMysqlServer,
	"POSTGRESQL_SERVER":           TechnologyTypePostgresqlServer,
}

var mappingTechnologyTypeEnumLowerCase = map[string]TechnologyTypeEnum{
	"goldengate":                  TechnologyTypeGoldengate,
	"oci_autonomous_database":     TechnologyTypeOciAutonomousDatabase,
	"oci_mysql":                   TechnologyTypeOciMysql,
	"oci_object_storage":          TechnologyTypeOciObjectStorage,
	"oci_streaming":               TechnologyTypeOciStreaming,
	"oracle_database":             TechnologyTypeOracleDatabase,
	"oracle_exadata":              TechnologyTypeOracleExadata,
	"amazon_rds_oracle":           TechnologyTypeAmazonRdsOracle,
	"amazon_aurora_mysql":         TechnologyTypeAmazonAuroraMysql,
	"amazon_aurora_postgresql":    TechnologyTypeAmazonAuroraPostgresql,
	"amazon_rds_mariadb":          TechnologyTypeAmazonRdsMariadb,
	"amazon_rds_mysql":            TechnologyTypeAmazonRdsMysql,
	"amazon_rds_postgresql":       TechnologyTypeAmazonRdsPostgresql,
	"apache_kafka":                TechnologyTypeApacheKafka,
	"azure_data_lake_storage":     TechnologyTypeAzureDataLakeStorage,
	"azure_event_hubs":            TechnologyTypeAzureEventHubs,
	"azure_mysql":                 TechnologyTypeAzureMysql,
	"azure_postgresql":            TechnologyTypeAzurePostgresql,
	"azure_synapse_analytics":     TechnologyTypeAzureSynapseAnalytics,
	"confluent_kafka":             TechnologyTypeConfluentKafka,
	"confluent_schema_registry":   TechnologyTypeConfluentSchemaRegistry,
	"google_cloud_sql_mysql":      TechnologyTypeGoogleCloudSqlMysql,
	"google_cloud_sql_postgresql": TechnologyTypeGoogleCloudSqlPostgresql,
	"mariadb":                     TechnologyTypeMariadb,
	"mysql_server":                TechnologyTypeMysqlServer,
	"postgresql_server":           TechnologyTypePostgresqlServer,
}

// GetTechnologyTypeEnumValues Enumerates the set of values for TechnologyTypeEnum
func GetTechnologyTypeEnumValues() []TechnologyTypeEnum {
	values := make([]TechnologyTypeEnum, 0)
	for _, v := range mappingTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTechnologyTypeEnumStringValues Enumerates the set of values in String for TechnologyTypeEnum
func GetTechnologyTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
		"OCI_AUTONOMOUS_DATABASE",
		"OCI_MYSQL",
		"OCI_OBJECT_STORAGE",
		"OCI_STREAMING",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
		"AMAZON_RDS_ORACLE",
		"AMAZON_AURORA_MYSQL",
		"AMAZON_AURORA_POSTGRESQL",
		"AMAZON_RDS_MARIADB",
		"AMAZON_RDS_MYSQL",
		"AMAZON_RDS_POSTGRESQL",
		"APACHE_KAFKA",
		"AZURE_DATA_LAKE_STORAGE",
		"AZURE_EVENT_HUBS",
		"AZURE_MYSQL",
		"AZURE_POSTGRESQL",
		"AZURE_SYNAPSE_ANALYTICS",
		"CONFLUENT_KAFKA",
		"CONFLUENT_SCHEMA_REGISTRY",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"GOOGLE_CLOUD_SQL_POSTGRESQL",
		"MARIADB",
		"MYSQL_SERVER",
		"POSTGRESQL_SERVER",
	}
}

// GetMappingTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTechnologyTypeEnum(val string) (TechnologyTypeEnum, bool) {
	enum, ok := mappingTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
