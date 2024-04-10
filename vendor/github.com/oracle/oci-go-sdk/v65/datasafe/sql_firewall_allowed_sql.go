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

// SqlFirewallAllowedSql The resource represents a SQL Firewall allowed SQL in Data Safe.
type SqlFirewallAllowedSql struct {

	// The OCID of the SQL Firewall allowed SQL.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the SQL Firewall allowed SQL.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the SQL Firewall allowed SQL.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the SQL Firewall policy corresponding to the SQL Firewall allowed SQL.
	SqlFirewallPolicyId *string `mandatory:"true" json:"sqlFirewallPolicyId"`

	// The database user name.
	DbUserName *string `mandatory:"true" json:"dbUserName"`

	// The SQL text of the SQL Firewall allowed SQL.
	SqlText *string `mandatory:"true" json:"sqlText"`

	// Specifies the level of SQL included for this SQL Firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallAllowedSqlSqlLevelEnum `mandatory:"true" json:"sqlLevel"`

	// Version of the associated SQL Firewall policy. This identifies whether the allowed SQLs were added in the same batch or not.
	Version *float32 `mandatory:"true" json:"version"`

	// The time the the SQL Firewall allowed SQL was collected from the target database, in the format defined by RFC3339.
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// The current state of the SQL Firewall allowed SQL.
	LifecycleState SqlFirewallAllowedSqlLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the SQL Firewall allowed SQL.
	Description *string `mandatory:"false" json:"description"`

	// The name of the user that SQL was executed as.
	CurrentUser *string `mandatory:"false" json:"currentUser"`

	// The objects accessed by the SQL.
	SqlAccessedObjects []string `mandatory:"false" json:"sqlAccessedObjects"`

	// The last date and time the SQL Firewall allowed SQL was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SqlFirewallAllowedSql) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallAllowedSql) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallAllowedSqlSqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallAllowedSqlSqlLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallAllowedSqlLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallAllowedSqlLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallAllowedSqlSqlLevelEnum Enum with underlying type: string
type SqlFirewallAllowedSqlSqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallAllowedSqlSqlLevelEnum
const (
	SqlFirewallAllowedSqlSqlLevelUserIssuedSql SqlFirewallAllowedSqlSqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallAllowedSqlSqlLevelAllSql        SqlFirewallAllowedSqlSqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallAllowedSqlSqlLevelEnum = map[string]SqlFirewallAllowedSqlSqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallAllowedSqlSqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallAllowedSqlSqlLevelAllSql,
}

var mappingSqlFirewallAllowedSqlSqlLevelEnumLowerCase = map[string]SqlFirewallAllowedSqlSqlLevelEnum{
	"user_issued_sql": SqlFirewallAllowedSqlSqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallAllowedSqlSqlLevelAllSql,
}

// GetSqlFirewallAllowedSqlSqlLevelEnumValues Enumerates the set of values for SqlFirewallAllowedSqlSqlLevelEnum
func GetSqlFirewallAllowedSqlSqlLevelEnumValues() []SqlFirewallAllowedSqlSqlLevelEnum {
	values := make([]SqlFirewallAllowedSqlSqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallAllowedSqlSqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallAllowedSqlSqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallAllowedSqlSqlLevelEnum
func GetSqlFirewallAllowedSqlSqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallAllowedSqlSqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallAllowedSqlSqlLevelEnum(val string) (SqlFirewallAllowedSqlSqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallAllowedSqlSqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
