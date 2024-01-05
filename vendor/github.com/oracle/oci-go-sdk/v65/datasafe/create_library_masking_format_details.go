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

// CreateLibraryMaskingFormatDetails Details to create a library masking format, which can have one or more format entries. A format
// entry can be a basic masking format such as Random Number, or it can be a library masking format.
// The combined output of all the format entries is used for masking. It provides the flexibility
// to define a masking format that can generate different parts of a data value separately and then
// combine them to get the final data value for masking. Note that you cannot define masking
// condition in a library masking format.
type CreateLibraryMaskingFormatDetails struct {

	// The OCID of the compartment where the library masking format should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// An array of format entries. The combined output of all the format entries is used for masking.
	FormatEntries []FormatEntry `mandatory:"true" json:"formatEntries"`

	// The display name of the library masking format. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the library masking format.
	Description *string `mandatory:"false" json:"description"`

	// An array of OCIDs of the sensitive types compatible with the library masking format. It helps track the sensitive types for which the library masking format is being created.
	SensitiveTypeIds []string `mandatory:"false" json:"sensitiveTypeIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateLibraryMaskingFormatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLibraryMaskingFormatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateLibraryMaskingFormatDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		SensitiveTypeIds []string                          `json:"sensitiveTypeIds"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId    *string                           `json:"compartmentId"`
		FormatEntries    []formatentry                     `json:"formatEntries"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.SensitiveTypeIds = make([]string, len(model.SensitiveTypeIds))
	copy(m.SensitiveTypeIds, model.SensitiveTypeIds)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.FormatEntries = make([]FormatEntry, len(model.FormatEntries))
	for i, n := range model.FormatEntries {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.FormatEntries[i] = nn.(FormatEntry)
		} else {
			m.FormatEntries[i] = nil
		}
	}
	return
}
