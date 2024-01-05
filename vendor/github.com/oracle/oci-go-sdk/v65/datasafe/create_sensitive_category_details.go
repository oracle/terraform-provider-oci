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

// CreateSensitiveCategoryDetails Details to create a new sensitive category.
type CreateSensitiveCategoryDetails struct {

	// The OCID of the compartment where the sensitive type should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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
}

// GetDisplayName returns DisplayName
func (m CreateSensitiveCategoryDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateSensitiveCategoryDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetShortName returns ShortName
func (m CreateSensitiveCategoryDetails) GetShortName() *string {
	return m.ShortName
}

// GetDescription returns Description
func (m CreateSensitiveCategoryDetails) GetDescription() *string {
	return m.Description
}

// GetParentCategoryId returns ParentCategoryId
func (m CreateSensitiveCategoryDetails) GetParentCategoryId() *string {
	return m.ParentCategoryId
}

// GetFreeformTags returns FreeformTags
func (m CreateSensitiveCategoryDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateSensitiveCategoryDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateSensitiveCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSensitiveCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSensitiveCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSensitiveCategoryDetails CreateSensitiveCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeCreateSensitiveCategoryDetails
	}{
		"SENSITIVE_CATEGORY",
		(MarshalTypeCreateSensitiveCategoryDetails)(m),
	}

	return json.Marshal(&s)
}
