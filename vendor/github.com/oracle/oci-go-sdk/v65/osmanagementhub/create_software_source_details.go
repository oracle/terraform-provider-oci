// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSoftwareSourceDetails Description of a software source to be created.
type CreateSoftwareSourceDetails interface {

	// The OCID of the tenancy containing the software source.
	GetCompartmentId() *string

	// User friendly name for the software source.
	GetDisplayName() *string

	// Information specified by the user about the software source.
	GetDescription() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createsoftwaresourcedetails struct {
	JsonData           []byte
	Description        *string                           `mandatory:"false" json:"description"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	SoftwareSourceType string                            `json:"softwareSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createsoftwaresourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatesoftwaresourcedetails createsoftwaresourcedetails
	s := struct {
		Model Unmarshalercreatesoftwaresourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SoftwareSourceType = s.Model.SoftwareSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createsoftwaresourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SoftwareSourceType {
	case "CUSTOM":
		mm := CreateCustomSoftwareSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSIONED":
		mm := CreateVersionedCustomSoftwareSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateSoftwareSourceDetails: %s.", m.SoftwareSourceType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createsoftwaresourcedetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m createsoftwaresourcedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createsoftwaresourcedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createsoftwaresourcedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m createsoftwaresourcedetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m createsoftwaresourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createsoftwaresourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
