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

// SqlFirewallViolationSummary The resource represents the SQL violations collected from the target database by Oracle Data Safe.
type SqlFirewallViolationSummary struct {

	// The OCID of the SQL violation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the SQL violation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The name of the target database.
	TargetName *string `mandatory:"true" json:"targetName"`

	// The time of the SQL violation occurrence in the target database.
	OperationTime *common.SDKTime `mandatory:"true" json:"operationTime"`

	// The timestamp when this SQL violation was collected from the target database by Data Safe.
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// The action taken for this SQL violation.
	ViolationAction SqlFirewallViolationSummaryViolationActionEnum `mandatory:"true" json:"violationAction"`

	// The name of the database user.
	DbUserName *string `mandatory:"false" json:"dbUserName"`

	// The name of the operating system user for the database session.
	ClientOsUserName *string `mandatory:"false" json:"clientOsUserName"`

	// The name of the action executed by the user on the target database. For example, ALTER, CREATE, DROP.
	Operation *string `mandatory:"false" json:"operation"`

	// The SQL text caught by the firewall.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// The objects accessed by the SQL.
	SqlAccessedObjects *string `mandatory:"false" json:"sqlAccessedObjects"`

	// The name of the user that SQL was executed as.
	CurrentDbUserName *string `mandatory:"false" json:"currentDbUserName"`

	// Specifies the level of SQL for this violation.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallViolationSummarySqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// The IP address of the host machine from which the session was generated.
	ClientIp *string `mandatory:"false" json:"clientIp"`

	// The application from which the SQL violation was generated. Examples include SQL Plus or SQL Developer.
	ClientProgram *string `mandatory:"false" json:"clientProgram"`

	// Indicates whether SQL or context violation.
	ViolationCause *string `mandatory:"false" json:"violationCause"`
}

func (m SqlFirewallViolationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallViolationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallViolationSummaryViolationActionEnum(string(m.ViolationAction)); !ok && m.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", m.ViolationAction, strings.Join(GetSqlFirewallViolationSummaryViolationActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSqlFirewallViolationSummarySqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallViolationSummarySqlLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallViolationSummarySqlLevelEnum Enum with underlying type: string
type SqlFirewallViolationSummarySqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallViolationSummarySqlLevelEnum
const (
	SqlFirewallViolationSummarySqlLevelUserIssuedSql SqlFirewallViolationSummarySqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallViolationSummarySqlLevelAllSql        SqlFirewallViolationSummarySqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallViolationSummarySqlLevelEnum = map[string]SqlFirewallViolationSummarySqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallViolationSummarySqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallViolationSummarySqlLevelAllSql,
}

var mappingSqlFirewallViolationSummarySqlLevelEnumLowerCase = map[string]SqlFirewallViolationSummarySqlLevelEnum{
	"user_issued_sql": SqlFirewallViolationSummarySqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallViolationSummarySqlLevelAllSql,
}

// GetSqlFirewallViolationSummarySqlLevelEnumValues Enumerates the set of values for SqlFirewallViolationSummarySqlLevelEnum
func GetSqlFirewallViolationSummarySqlLevelEnumValues() []SqlFirewallViolationSummarySqlLevelEnum {
	values := make([]SqlFirewallViolationSummarySqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallViolationSummarySqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallViolationSummarySqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallViolationSummarySqlLevelEnum
func GetSqlFirewallViolationSummarySqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallViolationSummarySqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallViolationSummarySqlLevelEnum(val string) (SqlFirewallViolationSummarySqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallViolationSummarySqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlFirewallViolationSummaryViolationActionEnum Enum with underlying type: string
type SqlFirewallViolationSummaryViolationActionEnum string

// Set of constants representing the allowable values for SqlFirewallViolationSummaryViolationActionEnum
const (
	SqlFirewallViolationSummaryViolationActionBlocked SqlFirewallViolationSummaryViolationActionEnum = "BLOCKED"
	SqlFirewallViolationSummaryViolationActionAllowed SqlFirewallViolationSummaryViolationActionEnum = "ALLOWED"
)

var mappingSqlFirewallViolationSummaryViolationActionEnum = map[string]SqlFirewallViolationSummaryViolationActionEnum{
	"BLOCKED": SqlFirewallViolationSummaryViolationActionBlocked,
	"ALLOWED": SqlFirewallViolationSummaryViolationActionAllowed,
}

var mappingSqlFirewallViolationSummaryViolationActionEnumLowerCase = map[string]SqlFirewallViolationSummaryViolationActionEnum{
	"blocked": SqlFirewallViolationSummaryViolationActionBlocked,
	"allowed": SqlFirewallViolationSummaryViolationActionAllowed,
}

// GetSqlFirewallViolationSummaryViolationActionEnumValues Enumerates the set of values for SqlFirewallViolationSummaryViolationActionEnum
func GetSqlFirewallViolationSummaryViolationActionEnumValues() []SqlFirewallViolationSummaryViolationActionEnum {
	values := make([]SqlFirewallViolationSummaryViolationActionEnum, 0)
	for _, v := range mappingSqlFirewallViolationSummaryViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallViolationSummaryViolationActionEnumStringValues Enumerates the set of values in String for SqlFirewallViolationSummaryViolationActionEnum
func GetSqlFirewallViolationSummaryViolationActionEnumStringValues() []string {
	return []string{
		"BLOCKED",
		"ALLOWED",
	}
}

// GetMappingSqlFirewallViolationSummaryViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallViolationSummaryViolationActionEnum(val string) (SqlFirewallViolationSummaryViolationActionEnum, bool) {
	enum, ok := mappingSqlFirewallViolationSummaryViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
