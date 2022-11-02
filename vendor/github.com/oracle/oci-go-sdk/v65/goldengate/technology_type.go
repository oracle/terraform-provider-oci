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

// TechnologyTypeEnum Enum with underlying type: string
type TechnologyTypeEnum string

// Set of constants representing the allowable values for TechnologyTypeEnum
const (
	TechnologyTypeOciAutonomousDatabase TechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	TechnologyTypeOciMysql              TechnologyTypeEnum = "OCI_MYSQL"
	TechnologyTypeOciObjectStorage      TechnologyTypeEnum = "OCI_OBJECT_STORAGE"
	TechnologyTypeOciStreaming          TechnologyTypeEnum = "OCI_STREAMING"
	TechnologyTypeOracleDatabase        TechnologyTypeEnum = "ORACLE_DATABASE"
	TechnologyTypeOracleExadata         TechnologyTypeEnum = "ORACLE_EXADATA"
	TechnologyTypeAmazonRdsOracle       TechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	TechnologyTypeAmazonAuroraMysql     TechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	TechnologyTypeAmazonRdsMariadb      TechnologyTypeEnum = "AMAZON_RDS_MARIADB"
	TechnologyTypeAmazonRdsMysql        TechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	TechnologyTypeApacheKafka           TechnologyTypeEnum = "APACHE_KAFKA"
	TechnologyTypeAzureMysql            TechnologyTypeEnum = "AZURE_MYSQL"
	TechnologyTypeGoldengate            TechnologyTypeEnum = "GOLDENGATE"
	TechnologyTypeGoogleCloudSqlMysql   TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	TechnologyTypeMariadb               TechnologyTypeEnum = "MARIADB"
	TechnologyTypeMysqlServer           TechnologyTypeEnum = "MYSQL_SERVER"
)

var mappingTechnologyTypeEnum = map[string]TechnologyTypeEnum{
	"OCI_AUTONOMOUS_DATABASE": TechnologyTypeOciAutonomousDatabase,
	"OCI_MYSQL":               TechnologyTypeOciMysql,
	"OCI_OBJECT_STORAGE":      TechnologyTypeOciObjectStorage,
	"OCI_STREAMING":           TechnologyTypeOciStreaming,
	"ORACLE_DATABASE":         TechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":          TechnologyTypeOracleExadata,
	"AMAZON_RDS_ORACLE":       TechnologyTypeAmazonRdsOracle,
	"AMAZON_AURORA_MYSQL":     TechnologyTypeAmazonAuroraMysql,
	"AMAZON_RDS_MARIADB":      TechnologyTypeAmazonRdsMariadb,
	"AMAZON_RDS_MYSQL":        TechnologyTypeAmazonRdsMysql,
	"APACHE_KAFKA":            TechnologyTypeApacheKafka,
	"AZURE_MYSQL":             TechnologyTypeAzureMysql,
	"GOLDENGATE":              TechnologyTypeGoldengate,
	"GOOGLE_CLOUD_SQL_MYSQL":  TechnologyTypeGoogleCloudSqlMysql,
	"MARIADB":                 TechnologyTypeMariadb,
	"MYSQL_SERVER":            TechnologyTypeMysqlServer,
}

var mappingTechnologyTypeEnumLowerCase = map[string]TechnologyTypeEnum{
	"oci_autonomous_database": TechnologyTypeOciAutonomousDatabase,
	"oci_mysql":               TechnologyTypeOciMysql,
	"oci_object_storage":      TechnologyTypeOciObjectStorage,
	"oci_streaming":           TechnologyTypeOciStreaming,
	"oracle_database":         TechnologyTypeOracleDatabase,
	"oracle_exadata":          TechnologyTypeOracleExadata,
	"amazon_rds_oracle":       TechnologyTypeAmazonRdsOracle,
	"amazon_aurora_mysql":     TechnologyTypeAmazonAuroraMysql,
	"amazon_rds_mariadb":      TechnologyTypeAmazonRdsMariadb,
	"amazon_rds_mysql":        TechnologyTypeAmazonRdsMysql,
	"apache_kafka":            TechnologyTypeApacheKafka,
	"azure_mysql":             TechnologyTypeAzureMysql,
	"goldengate":              TechnologyTypeGoldengate,
	"google_cloud_sql_mysql":  TechnologyTypeGoogleCloudSqlMysql,
	"mariadb":                 TechnologyTypeMariadb,
	"mysql_server":            TechnologyTypeMysqlServer,
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
		"OCI_AUTONOMOUS_DATABASE",
		"OCI_MYSQL",
		"OCI_OBJECT_STORAGE",
		"OCI_STREAMING",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
		"AMAZON_RDS_ORACLE",
		"AMAZON_AURORA_MYSQL",
		"AMAZON_RDS_MARIADB",
		"AMAZON_RDS_MYSQL",
		"APACHE_KAFKA",
		"AZURE_MYSQL",
		"GOLDENGATE",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"MARIADB",
		"MYSQL_SERVER",
	}
}

// GetMappingTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTechnologyTypeEnum(val string) (TechnologyTypeEnum, bool) {
	enum, ok := mappingTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
