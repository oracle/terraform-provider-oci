// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AttributeSet Represents an attribute set. An attribute set is a collection of data attributes defined by the user. i.e an attribute set of ip addresses, os user names or database privileged users.
type AttributeSet struct {

	// The OCID of an attribute set.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment where the attribute set is stored.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of an attribute set. The name does not have to be unique, and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of an attribute set.
	LifecycleState AttributeSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time an attribute set was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of attribute set.
	AttributeSetType AttributeSetAttributeSetTypeEnum `mandatory:"true" json:"attributeSetType"`

	// The list of values in an attribute set
	AttributeSetValues []string `mandatory:"true" json:"attributeSetValues"`

	// Description of an attribute set.
	Description *string `mandatory:"false" json:"description"`

	// The date and time an attribute set was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A boolean flag indicating to list user defined or seeded attribute sets.
	IsUserDefined *bool `mandatory:"false" json:"isUserDefined"`

	// Indicates whether the attribute set is in use by other resource.
	InUse AttributeSetInUseEnum `mandatory:"false" json:"inUse,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AttributeSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAttributeSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeSetAttributeSetTypeEnum(string(m.AttributeSetType)); !ok && m.AttributeSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeSetType: %s. Supported values are: %s.", m.AttributeSetType, strings.Join(GetAttributeSetAttributeSetTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAttributeSetInUseEnum(string(m.InUse)); !ok && m.InUse != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InUse: %s. Supported values are: %s.", m.InUse, strings.Join(GetAttributeSetInUseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeSetLifecycleStateEnum Enum with underlying type: string
type AttributeSetLifecycleStateEnum string

// Set of constants representing the allowable values for AttributeSetLifecycleStateEnum
const (
	AttributeSetLifecycleStateCreating AttributeSetLifecycleStateEnum = "CREATING"
	AttributeSetLifecycleStateActive   AttributeSetLifecycleStateEnum = "ACTIVE"
	AttributeSetLifecycleStateFailed   AttributeSetLifecycleStateEnum = "FAILED"
	AttributeSetLifecycleStateDeleting AttributeSetLifecycleStateEnum = "DELETING"
	AttributeSetLifecycleStateUpdating AttributeSetLifecycleStateEnum = "UPDATING"
)

var mappingAttributeSetLifecycleStateEnum = map[string]AttributeSetLifecycleStateEnum{
	"CREATING": AttributeSetLifecycleStateCreating,
	"ACTIVE":   AttributeSetLifecycleStateActive,
	"FAILED":   AttributeSetLifecycleStateFailed,
	"DELETING": AttributeSetLifecycleStateDeleting,
	"UPDATING": AttributeSetLifecycleStateUpdating,
}

var mappingAttributeSetLifecycleStateEnumLowerCase = map[string]AttributeSetLifecycleStateEnum{
	"creating": AttributeSetLifecycleStateCreating,
	"active":   AttributeSetLifecycleStateActive,
	"failed":   AttributeSetLifecycleStateFailed,
	"deleting": AttributeSetLifecycleStateDeleting,
	"updating": AttributeSetLifecycleStateUpdating,
}

// GetAttributeSetLifecycleStateEnumValues Enumerates the set of values for AttributeSetLifecycleStateEnum
func GetAttributeSetLifecycleStateEnumValues() []AttributeSetLifecycleStateEnum {
	values := make([]AttributeSetLifecycleStateEnum, 0)
	for _, v := range mappingAttributeSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeSetLifecycleStateEnumStringValues Enumerates the set of values in String for AttributeSetLifecycleStateEnum
func GetAttributeSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"UPDATING",
	}
}

// GetMappingAttributeSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeSetLifecycleStateEnum(val string) (AttributeSetLifecycleStateEnum, bool) {
	enum, ok := mappingAttributeSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeSetAttributeSetTypeEnum Enum with underlying type: string
type AttributeSetAttributeSetTypeEnum string

// Set of constants representing the allowable values for AttributeSetAttributeSetTypeEnum
const (
	AttributeSetAttributeSetTypeIpAddress      AttributeSetAttributeSetTypeEnum = "IP_ADDRESS"
	AttributeSetAttributeSetTypeClientProgram  AttributeSetAttributeSetTypeEnum = "CLIENT_PROGRAM"
	AttributeSetAttributeSetTypeOsUser         AttributeSetAttributeSetTypeEnum = "OS_USER"
	AttributeSetAttributeSetTypeDatabaseUser   AttributeSetAttributeSetTypeEnum = "DATABASE_USER"
	AttributeSetAttributeSetTypeDatabaseObject AttributeSetAttributeSetTypeEnum = "DATABASE_OBJECT"
)

var mappingAttributeSetAttributeSetTypeEnum = map[string]AttributeSetAttributeSetTypeEnum{
	"IP_ADDRESS":      AttributeSetAttributeSetTypeIpAddress,
	"CLIENT_PROGRAM":  AttributeSetAttributeSetTypeClientProgram,
	"OS_USER":         AttributeSetAttributeSetTypeOsUser,
	"DATABASE_USER":   AttributeSetAttributeSetTypeDatabaseUser,
	"DATABASE_OBJECT": AttributeSetAttributeSetTypeDatabaseObject,
}

var mappingAttributeSetAttributeSetTypeEnumLowerCase = map[string]AttributeSetAttributeSetTypeEnum{
	"ip_address":      AttributeSetAttributeSetTypeIpAddress,
	"client_program":  AttributeSetAttributeSetTypeClientProgram,
	"os_user":         AttributeSetAttributeSetTypeOsUser,
	"database_user":   AttributeSetAttributeSetTypeDatabaseUser,
	"database_object": AttributeSetAttributeSetTypeDatabaseObject,
}

// GetAttributeSetAttributeSetTypeEnumValues Enumerates the set of values for AttributeSetAttributeSetTypeEnum
func GetAttributeSetAttributeSetTypeEnumValues() []AttributeSetAttributeSetTypeEnum {
	values := make([]AttributeSetAttributeSetTypeEnum, 0)
	for _, v := range mappingAttributeSetAttributeSetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeSetAttributeSetTypeEnumStringValues Enumerates the set of values in String for AttributeSetAttributeSetTypeEnum
func GetAttributeSetAttributeSetTypeEnumStringValues() []string {
	return []string{
		"IP_ADDRESS",
		"CLIENT_PROGRAM",
		"OS_USER",
		"DATABASE_USER",
		"DATABASE_OBJECT",
	}
}

// GetMappingAttributeSetAttributeSetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeSetAttributeSetTypeEnum(val string) (AttributeSetAttributeSetTypeEnum, bool) {
	enum, ok := mappingAttributeSetAttributeSetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AttributeSetInUseEnum Enum with underlying type: string
type AttributeSetInUseEnum string

// Set of constants representing the allowable values for AttributeSetInUseEnum
const (
	AttributeSetInUseYes AttributeSetInUseEnum = "YES"
	AttributeSetInUseNo  AttributeSetInUseEnum = "NO"
)

var mappingAttributeSetInUseEnum = map[string]AttributeSetInUseEnum{
	"YES": AttributeSetInUseYes,
	"NO":  AttributeSetInUseNo,
}

var mappingAttributeSetInUseEnumLowerCase = map[string]AttributeSetInUseEnum{
	"yes": AttributeSetInUseYes,
	"no":  AttributeSetInUseNo,
}

// GetAttributeSetInUseEnumValues Enumerates the set of values for AttributeSetInUseEnum
func GetAttributeSetInUseEnumValues() []AttributeSetInUseEnum {
	values := make([]AttributeSetInUseEnum, 0)
	for _, v := range mappingAttributeSetInUseEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeSetInUseEnumStringValues Enumerates the set of values in String for AttributeSetInUseEnum
func GetAttributeSetInUseEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingAttributeSetInUseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeSetInUseEnum(val string) (AttributeSetInUseEnum, bool) {
	enum, ok := mappingAttributeSetInUseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
