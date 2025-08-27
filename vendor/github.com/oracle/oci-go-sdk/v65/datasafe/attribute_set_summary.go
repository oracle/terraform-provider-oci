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

// AttributeSetSummary Summary details of an attribute set.
type AttributeSetSummary struct {

	// The OCID of an attribute set.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains attribute set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of an attribute set. The name does not have to be unique, and is changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of an attribute set.
	LifecycleState AttributeSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time an attribute set was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The type of attribute set.
	AttributeSetType AttributeSetAttributeSetTypeEnum `mandatory:"true" json:"attributeSetType"`

	// Description of an attribute set.
	Description *string `mandatory:"false" json:"description"`

	// The date and time an attribute set was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Indicates whether the attribute set is user defined or pre defined in Data Safe. Values can either be 'true' or 'false'.
	IsUserDefined *bool `mandatory:"false" json:"isUserDefined"`

	// Indicates whether the attribute set is in use by other resource.
	InUse AttributeSetSummaryInUseEnum `mandatory:"false" json:"inUse,omitempty"`

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

func (m AttributeSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttributeSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAttributeSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAttributeSetAttributeSetTypeEnum(string(m.AttributeSetType)); !ok && m.AttributeSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeSetType: %s. Supported values are: %s.", m.AttributeSetType, strings.Join(GetAttributeSetAttributeSetTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAttributeSetSummaryInUseEnum(string(m.InUse)); !ok && m.InUse != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InUse: %s. Supported values are: %s.", m.InUse, strings.Join(GetAttributeSetSummaryInUseEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttributeSetSummaryInUseEnum Enum with underlying type: string
type AttributeSetSummaryInUseEnum string

// Set of constants representing the allowable values for AttributeSetSummaryInUseEnum
const (
	AttributeSetSummaryInUseYes AttributeSetSummaryInUseEnum = "YES"
	AttributeSetSummaryInUseNo  AttributeSetSummaryInUseEnum = "NO"
)

var mappingAttributeSetSummaryInUseEnum = map[string]AttributeSetSummaryInUseEnum{
	"YES": AttributeSetSummaryInUseYes,
	"NO":  AttributeSetSummaryInUseNo,
}

var mappingAttributeSetSummaryInUseEnumLowerCase = map[string]AttributeSetSummaryInUseEnum{
	"yes": AttributeSetSummaryInUseYes,
	"no":  AttributeSetSummaryInUseNo,
}

// GetAttributeSetSummaryInUseEnumValues Enumerates the set of values for AttributeSetSummaryInUseEnum
func GetAttributeSetSummaryInUseEnumValues() []AttributeSetSummaryInUseEnum {
	values := make([]AttributeSetSummaryInUseEnum, 0)
	for _, v := range mappingAttributeSetSummaryInUseEnum {
		values = append(values, v)
	}
	return values
}

// GetAttributeSetSummaryInUseEnumStringValues Enumerates the set of values in String for AttributeSetSummaryInUseEnum
func GetAttributeSetSummaryInUseEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingAttributeSetSummaryInUseEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttributeSetSummaryInUseEnum(val string) (AttributeSetSummaryInUseEnum, bool) {
	enum, ok := mappingAttributeSetSummaryInUseEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
