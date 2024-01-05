// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SqlFirewallAllowedSqlDimensions The dimensions available for SQL Firewall allow SQL analytics.
type SqlFirewallAllowedSqlDimensions struct {

	// The OCID of the SQL Firewall policy corresponding to the SQL Firewall allowed SQL.
	SqlFirewallPolicyId *string `mandatory:"false" json:"sqlFirewallPolicyId"`

	// Specifies the level of SQL included for this SQL Firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallAllowedSqlDimensionsSqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// The database user name.
	DbUserName *string `mandatory:"false" json:"dbUserName"`

	// The current state of the SQL Firewall allowed SQL.
	LifecycleState SqlFirewallAllowedSqlLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m SqlFirewallAllowedSqlDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallAllowedSqlDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlFirewallAllowedSqlDimensionsSqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallAllowedSqlDimensionsSqlLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallAllowedSqlLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallAllowedSqlLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallAllowedSqlDimensionsSqlLevelEnum Enum with underlying type: string
type SqlFirewallAllowedSqlDimensionsSqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallAllowedSqlDimensionsSqlLevelEnum
const (
	SqlFirewallAllowedSqlDimensionsSqlLevelUserIssuedSql SqlFirewallAllowedSqlDimensionsSqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallAllowedSqlDimensionsSqlLevelAllSql        SqlFirewallAllowedSqlDimensionsSqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallAllowedSqlDimensionsSqlLevelEnum = map[string]SqlFirewallAllowedSqlDimensionsSqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallAllowedSqlDimensionsSqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallAllowedSqlDimensionsSqlLevelAllSql,
}

var mappingSqlFirewallAllowedSqlDimensionsSqlLevelEnumLowerCase = map[string]SqlFirewallAllowedSqlDimensionsSqlLevelEnum{
	"user_issued_sql": SqlFirewallAllowedSqlDimensionsSqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallAllowedSqlDimensionsSqlLevelAllSql,
}

// GetSqlFirewallAllowedSqlDimensionsSqlLevelEnumValues Enumerates the set of values for SqlFirewallAllowedSqlDimensionsSqlLevelEnum
func GetSqlFirewallAllowedSqlDimensionsSqlLevelEnumValues() []SqlFirewallAllowedSqlDimensionsSqlLevelEnum {
	values := make([]SqlFirewallAllowedSqlDimensionsSqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallAllowedSqlDimensionsSqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallAllowedSqlDimensionsSqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallAllowedSqlDimensionsSqlLevelEnum
func GetSqlFirewallAllowedSqlDimensionsSqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallAllowedSqlDimensionsSqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallAllowedSqlDimensionsSqlLevelEnum(val string) (SqlFirewallAllowedSqlDimensionsSqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallAllowedSqlDimensionsSqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
