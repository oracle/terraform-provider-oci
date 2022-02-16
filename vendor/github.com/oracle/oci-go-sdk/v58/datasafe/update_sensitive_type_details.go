// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateSensitiveTypeDetails Details to update a sensitive type.
type UpdateSensitiveTypeDetails interface {

	// The display name of the sensitive type. The name does not have to be unique, and it's changeable.
	GetDisplayName() *string

	// The short name of the sensitive type.
	GetShortName() *string

	// The description of the sensitive type.
	GetDescription() *string

	// The OCID of the parent sensitive category.
	GetParentCategoryId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatesensitivetypedetails struct {
	JsonData         []byte
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	ShortName        *string                           `mandatory:"false" json:"shortName"`
	Description      *string                           `mandatory:"false" json:"description"`
	ParentCategoryId *string                           `mandatory:"false" json:"parentCategoryId"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	EntityType       string                            `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *updatesensitivetypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatesensitivetypedetails updatesensitivetypedetails
	s := struct {
		Model Unmarshalerupdatesensitivetypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.ShortName = s.Model.ShortName
	m.Description = s.Model.Description
	m.ParentCategoryId = s.Model.ParentCategoryId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatesensitivetypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "SENSITIVE_CATEGORY":
		mm := UpdateSensitiveCategoryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_TYPE":
		mm := UpdateSensitiveTypePatternDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatesensitivetypedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetShortName returns ShortName
func (m updatesensitivetypedetails) GetShortName() *string {
	return m.ShortName
}

//GetDescription returns Description
func (m updatesensitivetypedetails) GetDescription() *string {
	return m.Description
}

//GetParentCategoryId returns ParentCategoryId
func (m updatesensitivetypedetails) GetParentCategoryId() *string {
	return m.ParentCategoryId
}

//GetFreeformTags returns FreeformTags
func (m updatesensitivetypedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatesensitivetypedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatesensitivetypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatesensitivetypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
