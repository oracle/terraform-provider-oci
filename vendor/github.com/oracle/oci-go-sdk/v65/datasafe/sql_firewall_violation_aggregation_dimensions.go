// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlFirewallViolationAggregationDimensions The details of the aggregation dimensions used for summarizing SQL violations.
type SqlFirewallViolationAggregationDimensions struct {

	// The time of the SQL violation occurrence in the target database.
	OperationTime []common.SDKTime `mandatory:"false" json:"operationTime"`

	// The name of the database user.
	DbUserName []string `mandatory:"false" json:"dbUserName"`

	// The OCID of the target database.
	TargetId []string `mandatory:"false" json:"targetId"`

	// The name of the target database.
	TargetName []string `mandatory:"false" json:"targetName"`

	// The application from which the SQL violation was generated. Examples SQL Plus or SQL Developer.
	ClientProgram []string `mandatory:"false" json:"clientProgram"`

	// The name of the action executed by the user on the target database, for example, ALTER, CREATE, DROP.
	Operation []string `mandatory:"false" json:"operation"`

	// The name of the operating system user for the database session.
	ClientOsUserName []string `mandatory:"false" json:"clientOsUserName"`

	// Indicates whether SQL or context violation.
	ViolationCause []string `mandatory:"false" json:"violationCause"`

	// The IP address of the host from which the session was spawned.
	ClientIp []string `mandatory:"false" json:"clientIp"`

	// The action taken for this SQL violation.
	ViolationAction []string `mandatory:"false" json:"violationAction"`

	// Specifies the level of SQL included for this SQL Firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel []SqlFirewallViolationAggregationDimensionsSqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`
}

func (m SqlFirewallViolationAggregationDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallViolationAggregationDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.SqlLevel {
		if _, ok := GetMappingSqlFirewallViolationAggregationDimensionsSqlLevelEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", val, strings.Join(GetSqlFirewallViolationAggregationDimensionsSqlLevelEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallViolationAggregationDimensionsSqlLevelEnum Enum with underlying type: string
type SqlFirewallViolationAggregationDimensionsSqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallViolationAggregationDimensionsSqlLevelEnum
const (
	SqlFirewallViolationAggregationDimensionsSqlLevelUserIssuedSql SqlFirewallViolationAggregationDimensionsSqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallViolationAggregationDimensionsSqlLevelAllSql        SqlFirewallViolationAggregationDimensionsSqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallViolationAggregationDimensionsSqlLevelEnum = map[string]SqlFirewallViolationAggregationDimensionsSqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallViolationAggregationDimensionsSqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallViolationAggregationDimensionsSqlLevelAllSql,
}

var mappingSqlFirewallViolationAggregationDimensionsSqlLevelEnumLowerCase = map[string]SqlFirewallViolationAggregationDimensionsSqlLevelEnum{
	"user_issued_sql": SqlFirewallViolationAggregationDimensionsSqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallViolationAggregationDimensionsSqlLevelAllSql,
}

// GetSqlFirewallViolationAggregationDimensionsSqlLevelEnumValues Enumerates the set of values for SqlFirewallViolationAggregationDimensionsSqlLevelEnum
func GetSqlFirewallViolationAggregationDimensionsSqlLevelEnumValues() []SqlFirewallViolationAggregationDimensionsSqlLevelEnum {
	values := make([]SqlFirewallViolationAggregationDimensionsSqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallViolationAggregationDimensionsSqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallViolationAggregationDimensionsSqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallViolationAggregationDimensionsSqlLevelEnum
func GetSqlFirewallViolationAggregationDimensionsSqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallViolationAggregationDimensionsSqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallViolationAggregationDimensionsSqlLevelEnum(val string) (SqlFirewallViolationAggregationDimensionsSqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallViolationAggregationDimensionsSqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
