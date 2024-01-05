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

// SqlCollection The resource represents SQL collection for a specific database user in a target.
// SqlCollection encapsulates the SQL commands issued in the user’s database sessions, and its execution context.
type SqlCollection struct {

	// The OCID of the SQL collection.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the SQL collection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the SQL collection.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the target corresponding to the security policy deployment.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Specifies if the status of the SqlCollection. Enabled indicates that the collecting is in progress.
	Status SqlCollectionStatusEnum `mandatory:"true" json:"status"`

	// The database user name.
	DbUserName *string `mandatory:"true" json:"dbUserName"`

	// The time that the SQL collection was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SQL collection.
	LifecycleState SqlCollectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the SQL collection.
	Description *string `mandatory:"false" json:"description"`

	// The timestamp of the most recent SqlCollection start operation, in the format defined by RFC3339.
	TimeLastStarted *common.SDKTime `mandatory:"false" json:"timeLastStarted"`

	// The timestamp of the most recent SqlCollection stop operation, in the format defined by RFC3339.
	TimeLastStopped *common.SDKTime `mandatory:"false" json:"timeLastStopped"`

	// Specifies the level of SQL that will be collected.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel SqlCollectionSqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// The last date and time the SQL collection was updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Details about the current state of the SQL collection in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m SqlCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlCollectionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlCollectionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlCollectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSqlCollectionLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSqlCollectionSqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetSqlCollectionSqlLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlCollectionStatusEnum Enum with underlying type: string
type SqlCollectionStatusEnum string

// Set of constants representing the allowable values for SqlCollectionStatusEnum
const (
	SqlCollectionStatusEnabled  SqlCollectionStatusEnum = "ENABLED"
	SqlCollectionStatusDisabled SqlCollectionStatusEnum = "DISABLED"
)

var mappingSqlCollectionStatusEnum = map[string]SqlCollectionStatusEnum{
	"ENABLED":  SqlCollectionStatusEnabled,
	"DISABLED": SqlCollectionStatusDisabled,
}

var mappingSqlCollectionStatusEnumLowerCase = map[string]SqlCollectionStatusEnum{
	"enabled":  SqlCollectionStatusEnabled,
	"disabled": SqlCollectionStatusDisabled,
}

// GetSqlCollectionStatusEnumValues Enumerates the set of values for SqlCollectionStatusEnum
func GetSqlCollectionStatusEnumValues() []SqlCollectionStatusEnum {
	values := make([]SqlCollectionStatusEnum, 0)
	for _, v := range mappingSqlCollectionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlCollectionStatusEnumStringValues Enumerates the set of values in String for SqlCollectionStatusEnum
func GetSqlCollectionStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSqlCollectionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlCollectionStatusEnum(val string) (SqlCollectionStatusEnum, bool) {
	enum, ok := mappingSqlCollectionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlCollectionSqlLevelEnum Enum with underlying type: string
type SqlCollectionSqlLevelEnum string

// Set of constants representing the allowable values for SqlCollectionSqlLevelEnum
const (
	SqlCollectionSqlLevelUserIssuedSql SqlCollectionSqlLevelEnum = "USER_ISSUED_SQL"
	SqlCollectionSqlLevelAllSql        SqlCollectionSqlLevelEnum = "ALL_SQL"
)

var mappingSqlCollectionSqlLevelEnum = map[string]SqlCollectionSqlLevelEnum{
	"USER_ISSUED_SQL": SqlCollectionSqlLevelUserIssuedSql,
	"ALL_SQL":         SqlCollectionSqlLevelAllSql,
}

var mappingSqlCollectionSqlLevelEnumLowerCase = map[string]SqlCollectionSqlLevelEnum{
	"user_issued_sql": SqlCollectionSqlLevelUserIssuedSql,
	"all_sql":         SqlCollectionSqlLevelAllSql,
}

// GetSqlCollectionSqlLevelEnumValues Enumerates the set of values for SqlCollectionSqlLevelEnum
func GetSqlCollectionSqlLevelEnumValues() []SqlCollectionSqlLevelEnum {
	values := make([]SqlCollectionSqlLevelEnum, 0)
	for _, v := range mappingSqlCollectionSqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlCollectionSqlLevelEnumStringValues Enumerates the set of values in String for SqlCollectionSqlLevelEnum
func GetSqlCollectionSqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingSqlCollectionSqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlCollectionSqlLevelEnum(val string) (SqlCollectionSqlLevelEnum, bool) {
	enum, ok := mappingSqlCollectionSqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
