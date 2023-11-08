// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SensitiveTypePattern Details of the sensitive type.
type SensitiveTypePattern struct {

	// The OCID of the sensitive type.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the sensitive type.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the sensitive type.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the sensitive type was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the sensitive type was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The short name of the sensitive type.
	ShortName *string `mandatory:"false" json:"shortName"`

	// The description of the sensitive type.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the parent sensitive category.
	ParentCategoryId *string `mandatory:"false" json:"parentCategoryId"`

	// Specifies whether the sensitive type is common. Common sensitive types belong to
	// library sensitive types which are frequently used to perform sensitive data discovery.
	IsCommon *bool `mandatory:"false" json:"isCommon"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A regular expression to be used by data discovery for matching column names.
	NamePattern *string `mandatory:"false" json:"namePattern"`

	// A regular expression to be used by data discovery for matching column comments.
	CommentPattern *string `mandatory:"false" json:"commentPattern"`

	// A regular expression to be used by data discovery for matching column data values.
	DataPattern *string `mandatory:"false" json:"dataPattern"`

	// The OCID of the library masking format that should be used to mask the sensitive columns associated with the sensitive type.
	DefaultMaskingFormatId *string `mandatory:"false" json:"defaultMaskingFormatId"`

	// The search type indicating how the column name, comment and data patterns should be used by data discovery.
	// Learn more (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/sensitive-types.html#GUID-1D1AD98E-B93F-4FF2-80AE-CB7D8A14F6CC).
	SearchType SensitiveTypePatternSearchTypeEnum `mandatory:"false" json:"searchType,omitempty"`

	// The current state of the sensitive type.
	LifecycleState DiscoveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies whether the sensitive type is user-defined or predefined.
	Source SensitiveTypeSourceEnum `mandatory:"true" json:"source"`
}

// GetId returns Id
func (m SensitiveTypePattern) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m SensitiveTypePattern) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m SensitiveTypePattern) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m SensitiveTypePattern) GetLifecycleState() DiscoveryLifecycleStateEnum {
	return m.LifecycleState
}

// GetShortName returns ShortName
func (m SensitiveTypePattern) GetShortName() *string {
	return m.ShortName
}

// GetSource returns Source
func (m SensitiveTypePattern) GetSource() SensitiveTypeSourceEnum {
	return m.Source
}

// GetTimeCreated returns TimeCreated
func (m SensitiveTypePattern) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m SensitiveTypePattern) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDescription returns Description
func (m SensitiveTypePattern) GetDescription() *string {
	return m.Description
}

// GetParentCategoryId returns ParentCategoryId
func (m SensitiveTypePattern) GetParentCategoryId() *string {
	return m.ParentCategoryId
}

// GetIsCommon returns IsCommon
func (m SensitiveTypePattern) GetIsCommon() *bool {
	return m.IsCommon
}

// GetFreeformTags returns FreeformTags
func (m SensitiveTypePattern) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SensitiveTypePattern) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SensitiveTypePattern) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SensitiveTypePattern) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveTypePattern) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSensitiveTypePatternSearchTypeEnum(string(m.SearchType)); !ok && m.SearchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SearchType: %s. Supported values are: %s.", m.SearchType, strings.Join(GetSensitiveTypePatternSearchTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveTypeSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetSensitiveTypeSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SensitiveTypePattern) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSensitiveTypePattern SensitiveTypePattern
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeSensitiveTypePattern
	}{
		"SENSITIVE_TYPE",
		(MarshalTypeSensitiveTypePattern)(m),
	}

	return json.Marshal(&s)
}

// SensitiveTypePatternSearchTypeEnum Enum with underlying type: string
type SensitiveTypePatternSearchTypeEnum string

// Set of constants representing the allowable values for SensitiveTypePatternSearchTypeEnum
const (
	SensitiveTypePatternSearchTypeOr  SensitiveTypePatternSearchTypeEnum = "OR"
	SensitiveTypePatternSearchTypeAnd SensitiveTypePatternSearchTypeEnum = "AND"
)

var mappingSensitiveTypePatternSearchTypeEnum = map[string]SensitiveTypePatternSearchTypeEnum{
	"OR":  SensitiveTypePatternSearchTypeOr,
	"AND": SensitiveTypePatternSearchTypeAnd,
}

var mappingSensitiveTypePatternSearchTypeEnumLowerCase = map[string]SensitiveTypePatternSearchTypeEnum{
	"or":  SensitiveTypePatternSearchTypeOr,
	"and": SensitiveTypePatternSearchTypeAnd,
}

// GetSensitiveTypePatternSearchTypeEnumValues Enumerates the set of values for SensitiveTypePatternSearchTypeEnum
func GetSensitiveTypePatternSearchTypeEnumValues() []SensitiveTypePatternSearchTypeEnum {
	values := make([]SensitiveTypePatternSearchTypeEnum, 0)
	for _, v := range mappingSensitiveTypePatternSearchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveTypePatternSearchTypeEnumStringValues Enumerates the set of values in String for SensitiveTypePatternSearchTypeEnum
func GetSensitiveTypePatternSearchTypeEnumStringValues() []string {
	return []string{
		"OR",
		"AND",
	}
}

// GetMappingSensitiveTypePatternSearchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveTypePatternSearchTypeEnum(val string) (SensitiveTypePatternSearchTypeEnum, bool) {
	enum, ok := mappingSensitiveTypePatternSearchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
