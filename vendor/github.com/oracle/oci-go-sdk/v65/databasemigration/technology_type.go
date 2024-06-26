// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// TechnologyTypeEnum Enum with underlying type: string
type TechnologyTypeEnum string

// Set of constants representing the allowable values for TechnologyTypeEnum
const (
	TechnologyTypeOciAutonomousDatabase TechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	TechnologyTypeOciMysql              TechnologyTypeEnum = "OCI_MYSQL"
	TechnologyTypeOracleDatabase        TechnologyTypeEnum = "ORACLE_DATABASE"
	TechnologyTypeOracleExadata         TechnologyTypeEnum = "ORACLE_EXADATA"
	TechnologyTypeAmazonRdsOracle       TechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	TechnologyTypeAmazonAuroraMysql     TechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	TechnologyTypeAmazonRdsMysql        TechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	TechnologyTypeAzureMysql            TechnologyTypeEnum = "AZURE_MYSQL"
	TechnologyTypeGoogleCloudSqlMysql   TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	TechnologyTypeMysqlServer           TechnologyTypeEnum = "MYSQL_SERVER"
)

var mappingTechnologyTypeEnum = map[string]TechnologyTypeEnum{
	"OCI_AUTONOMOUS_DATABASE": TechnologyTypeOciAutonomousDatabase,
	"OCI_MYSQL":               TechnologyTypeOciMysql,
	"ORACLE_DATABASE":         TechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":          TechnologyTypeOracleExadata,
	"AMAZON_RDS_ORACLE":       TechnologyTypeAmazonRdsOracle,
	"AMAZON_AURORA_MYSQL":     TechnologyTypeAmazonAuroraMysql,
	"AMAZON_RDS_MYSQL":        TechnologyTypeAmazonRdsMysql,
	"AZURE_MYSQL":             TechnologyTypeAzureMysql,
	"GOOGLE_CLOUD_SQL_MYSQL":  TechnologyTypeGoogleCloudSqlMysql,
	"MYSQL_SERVER":            TechnologyTypeMysqlServer,
}

var mappingTechnologyTypeEnumLowerCase = map[string]TechnologyTypeEnum{
	"oci_autonomous_database": TechnologyTypeOciAutonomousDatabase,
	"oci_mysql":               TechnologyTypeOciMysql,
	"oracle_database":         TechnologyTypeOracleDatabase,
	"oracle_exadata":          TechnologyTypeOracleExadata,
	"amazon_rds_oracle":       TechnologyTypeAmazonRdsOracle,
	"amazon_aurora_mysql":     TechnologyTypeAmazonAuroraMysql,
	"amazon_rds_mysql":        TechnologyTypeAmazonRdsMysql,
	"azure_mysql":             TechnologyTypeAzureMysql,
	"google_cloud_sql_mysql":  TechnologyTypeGoogleCloudSqlMysql,
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
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
		"AMAZON_RDS_ORACLE",
		"AMAZON_AURORA_MYSQL",
		"AMAZON_RDS_MYSQL",
		"AZURE_MYSQL",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"MYSQL_SERVER",
	}
}

// GetMappingTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTechnologyTypeEnum(val string) (TechnologyTypeEnum, bool) {
	enum, ok := mappingTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
