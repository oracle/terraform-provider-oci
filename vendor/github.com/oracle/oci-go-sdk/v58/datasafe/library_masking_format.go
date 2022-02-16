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

// LibraryMaskingFormat A library masking format is a masking format stored in an Oracle Cloud Infrastructure compartment and can be used in
// multiple masking policies. If you want to use the same masking logic for multiple masking columns or even in multiple
// masking policies, you can create a library masking format and assign it to masking columns as needed. It helps you
// avoid defining the same masking logic again and again.
// Oracle Data Safe provides a set of predefined library masking formats to mask common sensitive and personal data,
// such as names, national identifiers, credit card numbers, and phone numbers. To meet your specific requirements, you
// can easily create new library masking formats and use them in your masking policies.
type LibraryMaskingFormat struct {

	// The OCID of the library masking format.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the library masking format.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the library masking format.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the library masking format was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the library masking format was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the library masking format.
	LifecycleState MaskingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies whether the library masking format is user-defined or predefined.
	Source LibraryMaskingFormatSourceEnum `mandatory:"true" json:"source"`

	// The description of the library masking format.
	Description *string `mandatory:"false" json:"description"`

	// An array of OCIDs of the sensitive types compatible with the library masking format.
	SensitiveTypeIds []string `mandatory:"false" json:"sensitiveTypeIds"`

	// An array of format entries. The combined output of all the format entries is used for masking.
	FormatEntries []FormatEntry `mandatory:"false" json:"formatEntries"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LibraryMaskingFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryMaskingFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaskingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLibraryMaskingFormatSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetLibraryMaskingFormatSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LibraryMaskingFormat) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		SensitiveTypeIds []string                          `json:"sensitiveTypeIds"`
		FormatEntries    []formatentry                     `json:"formatEntries"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		DisplayName      *string                           `json:"displayName"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState   MaskingLifecycleStateEnum         `json:"lifecycleState"`
		Source           LibraryMaskingFormatSourceEnum    `json:"source"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.SensitiveTypeIds = make([]string, len(model.SensitiveTypeIds))
	for i, n := range model.SensitiveTypeIds {
		m.SensitiveTypeIds[i] = n
	}

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.Source = model.Source

	return
}
