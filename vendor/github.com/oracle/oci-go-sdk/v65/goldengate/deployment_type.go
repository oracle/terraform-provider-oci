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

// DeploymentTypeEnum Enum with underlying type: string
type DeploymentTypeEnum string

// Set of constants representing the allowable values for DeploymentTypeEnum
const (
	DeploymentTypeOgg                        DeploymentTypeEnum = "OGG"
	DeploymentTypeDatabaseOracle             DeploymentTypeEnum = "DATABASE_ORACLE"
	DeploymentTypeBigdata                    DeploymentTypeEnum = "BIGDATA"
	DeploymentTypeDatabaseMicrosoftSqlserver DeploymentTypeEnum = "DATABASE_MICROSOFT_SQLSERVER"
	DeploymentTypeDatabaseMysql              DeploymentTypeEnum = "DATABASE_MYSQL"
	DeploymentTypeDatabasePostgresql         DeploymentTypeEnum = "DATABASE_POSTGRESQL"
	DeploymentTypeDatabaseDb2Zos             DeploymentTypeEnum = "DATABASE_DB2ZOS"
	DeploymentTypeDataTransforms             DeploymentTypeEnum = "DATA_TRANSFORMS"
)

var mappingDeploymentTypeEnum = map[string]DeploymentTypeEnum{
	"OGG":                          DeploymentTypeOgg,
	"DATABASE_ORACLE":              DeploymentTypeDatabaseOracle,
	"BIGDATA":                      DeploymentTypeBigdata,
	"DATABASE_MICROSOFT_SQLSERVER": DeploymentTypeDatabaseMicrosoftSqlserver,
	"DATABASE_MYSQL":               DeploymentTypeDatabaseMysql,
	"DATABASE_POSTGRESQL":          DeploymentTypeDatabasePostgresql,
	"DATABASE_DB2ZOS":              DeploymentTypeDatabaseDb2Zos,
	"DATA_TRANSFORMS":              DeploymentTypeDataTransforms,
}

var mappingDeploymentTypeEnumLowerCase = map[string]DeploymentTypeEnum{
	"ogg":                          DeploymentTypeOgg,
	"database_oracle":              DeploymentTypeDatabaseOracle,
	"bigdata":                      DeploymentTypeBigdata,
	"database_microsoft_sqlserver": DeploymentTypeDatabaseMicrosoftSqlserver,
	"database_mysql":               DeploymentTypeDatabaseMysql,
	"database_postgresql":          DeploymentTypeDatabasePostgresql,
	"database_db2zos":              DeploymentTypeDatabaseDb2Zos,
	"data_transforms":              DeploymentTypeDataTransforms,
}

// GetDeploymentTypeEnumValues Enumerates the set of values for DeploymentTypeEnum
func GetDeploymentTypeEnumValues() []DeploymentTypeEnum {
	values := make([]DeploymentTypeEnum, 0)
	for _, v := range mappingDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentTypeEnumStringValues Enumerates the set of values in String for DeploymentTypeEnum
func GetDeploymentTypeEnumStringValues() []string {
	return []string{
		"OGG",
		"DATABASE_ORACLE",
		"BIGDATA",
		"DATABASE_MICROSOFT_SQLSERVER",
		"DATABASE_MYSQL",
		"DATABASE_POSTGRESQL",
		"DATABASE_DB2ZOS",
		"DATA_TRANSFORMS",
	}
}

// GetMappingDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentTypeEnum(val string) (DeploymentTypeEnum, bool) {
	enum, ok := mappingDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
