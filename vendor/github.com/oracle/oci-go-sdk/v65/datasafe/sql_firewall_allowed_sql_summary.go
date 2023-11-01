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

// SqlFirewallAllowedSqlSummary The resource represents a SQL firewall allowed SQL in Data Safe.
type SqlFirewallAllowedSqlSummary struct {

	// The OCID of the SQL firewall allowed SQL.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the SQL firewall allowed SQL.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the SQL firewall allowed SQL.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the SQL firewall policy corresponding to the SQL firewall allowed SQL.
	SqlFirewallPolicyId *string `mandatory:"true" json:"sqlFirewallPolicyId"`

	// The database user name.
	DbUserName *string `mandatory:"true" json:"dbUserName"`

	// The SQL text of the SQL firewall allowed SQL.
	SqlText *string `mandatory:"true" json:"sqlText"`

	// Specifies the level of SQL included for this SQL firewall policy.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlFirewallAllowedSqlSummarySqlLevelEnum `mandatory:"true" json:"sqlLevel"`

	// Version of the associated SQL firewall policy. This identifies whether the allowed SQLs were added in the same batch or not.
	Version *float32 `mandatory:"true" json:"version"`

	// The time the the SQL firewall allowed SQL was collected from the target database, in the format defined by RFC3339.
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// The current state of the SQL firewall allowed SQL.
	LifecycleState SqlFirewallAllowedSqlLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the SQL firewall allowed SQL.
	Description *string `mandatory:"false" json:"description"`

	// The name of the user that SQL was executed as.
	CurrentUser *string `mandatory:"false" json:"currentUser"`

	// The objects accessed by the SQL.
	SqlAccessedObjects []string `mandatory:"false" json:"sqlAccessedObjects"`

	// The last date and time the SQL firewall allowed SQL was updated, in the format defined by RFC3339.
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

func (m SqlFirewallAllowedSqlSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlFirewallAllowedSqlSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlFirewallAllowedSqlSummarySqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlFirewallAllowedSqlSummarySqlLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlFirewallAllowedSqlLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlFirewallAllowedSqlLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlFirewallAllowedSqlSummarySqlLevelEnum Enum with underlying type: string
type SqlFirewallAllowedSqlSummarySqlLevelEnum string

// Set of constants representing the allowable values for SqlFirewallAllowedSqlSummarySqlLevelEnum
const (
	SqlFirewallAllowedSqlSummarySqlLevelUserIssuedSql SqlFirewallAllowedSqlSummarySqlLevelEnum = "USER_ISSUED_SQL"
	SqlFirewallAllowedSqlSummarySqlLevelAllSql        SqlFirewallAllowedSqlSummarySqlLevelEnum = "ALL_SQL"
)

var mappingSqlFirewallAllowedSqlSummarySqlLevelEnum = map[string]SqlFirewallAllowedSqlSummarySqlLevelEnum{
	"USER_ISSUED_SQL": SqlFirewallAllowedSqlSummarySqlLevelUserIssuedSql,
	"ALL_SQL":         SqlFirewallAllowedSqlSummarySqlLevelAllSql,
}

var mappingSqlFirewallAllowedSqlSummarySqlLevelEnumLowerCase = map[string]SqlFirewallAllowedSqlSummarySqlLevelEnum{
	"user_issued_sql": SqlFirewallAllowedSqlSummarySqlLevelUserIssuedSql,
	"all_sql":         SqlFirewallAllowedSqlSummarySqlLevelAllSql,
}

// GetSqlFirewallAllowedSqlSummarySqlLevelEnumValues Enumerates the set of values for SqlFirewallAllowedSqlSummarySqlLevelEnum
func GetSqlFirewallAllowedSqlSummarySqlLevelEnumValues() []SqlFirewallAllowedSqlSummarySqlLevelEnum {
	values := make([]SqlFirewallAllowedSqlSummarySqlLevelEnum, 0)
	for _, v := range mappingSqlFirewallAllowedSqlSummarySqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallAllowedSqlSummarySqlLevelEnumStringValues Enumerates the set of values in String for SqlFirewallAllowedSqlSummarySqlLevelEnum
func GetSqlFirewallAllowedSqlSummarySqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlFirewallAllowedSqlSummarySqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallAllowedSqlSummarySqlLevelEnum(val string) (SqlFirewallAllowedSqlSummarySqlLevelEnum, bool) {
	enum, ok := mappingSqlFirewallAllowedSqlSummarySqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
