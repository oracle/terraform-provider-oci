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
	TechnologyTypeAmazonAuroraMysql     TechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	TechnologyTypeAmazonRdsMariadb      TechnologyTypeEnum = "AMAZON_RDS_MARIADB"
	TechnologyTypeAmazonRdsMysql        TechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	TechnologyTypeAmazonRdsOracle       TechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	TechnologyTypeApacheKafka           TechnologyTypeEnum = "APACHE_KAFKA"
	TechnologyTypeAzureMysql            TechnologyTypeEnum = "AZURE_MYSQL"
	TechnologyTypeGoldengate            TechnologyTypeEnum = "GOLDENGATE"
	TechnologyTypeMariadb               TechnologyTypeEnum = "MARIADB"
	TechnologyTypeMysqlServer           TechnologyTypeEnum = "MYSQL_SERVER"
	TechnologyTypeOciAutonomousDatabase TechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	TechnologyTypeOciMysql              TechnologyTypeEnum = "OCI_MYSQL"
	TechnologyTypeOciObjectStorage      TechnologyTypeEnum = "OCI_OBJECT_STORAGE"
	TechnologyTypeOciStreaming          TechnologyTypeEnum = "OCI_STREAMING"
	TechnologyTypeOracleDatabase        TechnologyTypeEnum = "ORACLE_DATABASE"
	TechnologyTypeOracleExadata         TechnologyTypeEnum = "ORACLE_EXADATA"
)

var mappingTechnologyTypeEnum = map[string]TechnologyTypeEnum{
	"AMAZON_AURORA_MYSQL":     TechnologyTypeAmazonAuroraMysql,
	"AMAZON_RDS_MARIADB":      TechnologyTypeAmazonRdsMariadb,
	"AMAZON_RDS_MYSQL":        TechnologyTypeAmazonRdsMysql,
	"AMAZON_RDS_ORACLE":       TechnologyTypeAmazonRdsOracle,
	"APACHE_KAFKA":            TechnologyTypeApacheKafka,
	"AZURE_MYSQL":             TechnologyTypeAzureMysql,
	"GOLDENGATE":              TechnologyTypeGoldengate,
	"MARIADB":                 TechnologyTypeMariadb,
	"MYSQL_SERVER":            TechnologyTypeMysqlServer,
	"OCI_AUTONOMOUS_DATABASE": TechnologyTypeOciAutonomousDatabase,
	"OCI_MYSQL":               TechnologyTypeOciMysql,
	"OCI_OBJECT_STORAGE":      TechnologyTypeOciObjectStorage,
	"OCI_STREAMING":           TechnologyTypeOciStreaming,
	"ORACLE_DATABASE":         TechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":          TechnologyTypeOracleExadata,
}

var mappingTechnologyTypeEnumLowerCase = map[string]TechnologyTypeEnum{
	"amazon_aurora_mysql":     TechnologyTypeAmazonAuroraMysql,
	"amazon_rds_mariadb":      TechnologyTypeAmazonRdsMariadb,
	"amazon_rds_mysql":        TechnologyTypeAmazonRdsMysql,
	"amazon_rds_oracle":       TechnologyTypeAmazonRdsOracle,
	"apache_kafka":            TechnologyTypeApacheKafka,
	"azure_mysql":             TechnologyTypeAzureMysql,
	"goldengate":              TechnologyTypeGoldengate,
	"mariadb":                 TechnologyTypeMariadb,
	"mysql_server":            TechnologyTypeMysqlServer,
	"oci_autonomous_database": TechnologyTypeOciAutonomousDatabase,
	"oci_mysql":               TechnologyTypeOciMysql,
	"oci_object_storage":      TechnologyTypeOciObjectStorage,
	"oci_streaming":           TechnologyTypeOciStreaming,
	"oracle_database":         TechnologyTypeOracleDatabase,
	"oracle_exadata":          TechnologyTypeOracleExadata,
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
		"AMAZON_AURORA_MYSQL",
		"AMAZON_RDS_MARIADB",
		"AMAZON_RDS_MYSQL",
		"AMAZON_RDS_ORACLE",
		"APACHE_KAFKA",
		"AZURE_MYSQL",
		"GOLDENGATE",
		"MARIADB",
		"MYSQL_SERVER",
		"OCI_AUTONOMOUS_DATABASE",
		"OCI_MYSQL",
		"OCI_OBJECT_STORAGE",
		"OCI_STREAMING",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
	}
}

// GetMappingTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTechnologyTypeEnum(val string) (TechnologyTypeEnum, bool) {
	enum, ok := mappingTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
