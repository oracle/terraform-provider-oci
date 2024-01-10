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

// CreateSqlCollectionDetails Details for SQL collection creation.
type CreateSqlCollectionDetails struct {

	// The OCID of the compartment containing the SQL collection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target corresponding to the security policy deployment.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The database user name.
	DbUserName *string `mandatory:"true" json:"dbUserName"`

	// The display name of the SQL collection. The name does not have to be unique, and it is changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the SQL collection.
	Description *string `mandatory:"false" json:"description"`

	// Specifies if the SqlCollection has to be started after creation. Enabled indicates that the SqlCollection will be started after creation.
	Status CreateSqlCollectionDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Specifies the level of SQL that will be collected.
	// USER_ISSUED_SQL - User issued SQL statements only.
	// ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units.
	SqlLevel CreateSqlCollectionDetailsSqlLevelEnum `mandatory:"false" json:"sqlLevel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSqlCollectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSqlCollectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateSqlCollectionDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCreateSqlCollectionDetailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateSqlCollectionDetailsSqlLevelEnum(string(m.SqlLevel)); !ok && m.SqlLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlLevel: %s. Supported values are: %s.", m.SqlLevel, strings.Join(GetCreateSqlCollectionDetailsSqlLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSqlCollectionDetailsStatusEnum Enum with underlying type: string
type CreateSqlCollectionDetailsStatusEnum string

// Set of constants representing the allowable values for CreateSqlCollectionDetailsStatusEnum
const (
	CreateSqlCollectionDetailsStatusEnabled  CreateSqlCollectionDetailsStatusEnum = "ENABLED"
	CreateSqlCollectionDetailsStatusDisabled CreateSqlCollectionDetailsStatusEnum = "DISABLED"
)

var mappingCreateSqlCollectionDetailsStatusEnum = map[string]CreateSqlCollectionDetailsStatusEnum{
	"ENABLED":  CreateSqlCollectionDetailsStatusEnabled,
	"DISABLED": CreateSqlCollectionDetailsStatusDisabled,
}

var mappingCreateSqlCollectionDetailsStatusEnumLowerCase = map[string]CreateSqlCollectionDetailsStatusEnum{
	"enabled":  CreateSqlCollectionDetailsStatusEnabled,
	"disabled": CreateSqlCollectionDetailsStatusDisabled,
}

// GetCreateSqlCollectionDetailsStatusEnumValues Enumerates the set of values for CreateSqlCollectionDetailsStatusEnum
func GetCreateSqlCollectionDetailsStatusEnumValues() []CreateSqlCollectionDetailsStatusEnum {
	values := make([]CreateSqlCollectionDetailsStatusEnum, 0)
	for _, v := range mappingCreateSqlCollectionDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSqlCollectionDetailsStatusEnumStringValues Enumerates the set of values in String for CreateSqlCollectionDetailsStatusEnum
func GetCreateSqlCollectionDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCreateSqlCollectionDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSqlCollectionDetailsStatusEnum(val string) (CreateSqlCollectionDetailsStatusEnum, bool) {
	enum, ok := mappingCreateSqlCollectionDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateSqlCollectionDetailsSqlLevelEnum Enum with underlying type: string
type CreateSqlCollectionDetailsSqlLevelEnum string

// Set of constants representing the allowable values for CreateSqlCollectionDetailsSqlLevelEnum
const (
	CreateSqlCollectionDetailsSqlLevelUserIssuedSql CreateSqlCollectionDetailsSqlLevelEnum = "USER_ISSUED_SQL"
	CreateSqlCollectionDetailsSqlLevelAllSql        CreateSqlCollectionDetailsSqlLevelEnum = "ALL_SQL"
)

var mappingCreateSqlCollectionDetailsSqlLevelEnum = map[string]CreateSqlCollectionDetailsSqlLevelEnum{
	"USER_ISSUED_SQL": CreateSqlCollectionDetailsSqlLevelUserIssuedSql,
	"ALL_SQL":         CreateSqlCollectionDetailsSqlLevelAllSql,
}

var mappingCreateSqlCollectionDetailsSqlLevelEnumLowerCase = map[string]CreateSqlCollectionDetailsSqlLevelEnum{
	"user_issued_sql": CreateSqlCollectionDetailsSqlLevelUserIssuedSql,
	"all_sql":         CreateSqlCollectionDetailsSqlLevelAllSql,
}

// GetCreateSqlCollectionDetailsSqlLevelEnumValues Enumerates the set of values for CreateSqlCollectionDetailsSqlLevelEnum
func GetCreateSqlCollectionDetailsSqlLevelEnumValues() []CreateSqlCollectionDetailsSqlLevelEnum {
	values := make([]CreateSqlCollectionDetailsSqlLevelEnum, 0)
	for _, v := range mappingCreateSqlCollectionDetailsSqlLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSqlCollectionDetailsSqlLevelEnumStringValues Enumerates the set of values in String for CreateSqlCollectionDetailsSqlLevelEnum
func GetCreateSqlCollectionDetailsSqlLevelEnumStringValues() []string {
	return []string{
		"USER_ISSUED_SQL",
		"ALL_SQL",
	}
}

// GetMappingCreateSqlCollectionDetailsSqlLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSqlCollectionDetailsSqlLevelEnum(val string) (CreateSqlCollectionDetailsSqlLevelEnum, bool) {
	enum, ok := mappingCreateSqlCollectionDetailsSqlLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
