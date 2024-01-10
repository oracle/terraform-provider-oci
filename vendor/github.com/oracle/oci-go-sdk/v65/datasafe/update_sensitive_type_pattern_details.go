// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateSensitiveTypePatternDetails Details to update a sensitive type with regular expressions.
type UpdateSensitiveTypePatternDetails struct {

	// The display name of the sensitive type. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The short name of the sensitive type.
	ShortName *string `mandatory:"false" json:"shortName"`

	// The description of the sensitive type.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the parent sensitive category.
	ParentCategoryId *string `mandatory:"false" json:"parentCategoryId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

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
}

// GetDisplayName returns DisplayName
func (m UpdateSensitiveTypePatternDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetShortName returns ShortName
func (m UpdateSensitiveTypePatternDetails) GetShortName() *string {
	return m.ShortName
}

// GetDescription returns Description
func (m UpdateSensitiveTypePatternDetails) GetDescription() *string {
	return m.Description
}

// GetParentCategoryId returns ParentCategoryId
func (m UpdateSensitiveTypePatternDetails) GetParentCategoryId() *string {
	return m.ParentCategoryId
}

// GetFreeformTags returns FreeformTags
func (m UpdateSensitiveTypePatternDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateSensitiveTypePatternDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateSensitiveTypePatternDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSensitiveTypePatternDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSensitiveTypePatternSearchTypeEnum(string(m.SearchType)); !ok && m.SearchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SearchType: %s. Supported values are: %s.", m.SearchType, strings.Join(GetSensitiveTypePatternSearchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateSensitiveTypePatternDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSensitiveTypePatternDetails UpdateSensitiveTypePatternDetails
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeUpdateSensitiveTypePatternDetails
	}{
		"SENSITIVE_TYPE",
		(MarshalTypeUpdateSensitiveTypePatternDetails)(m),
	}

	return json.Marshal(&s)
}
