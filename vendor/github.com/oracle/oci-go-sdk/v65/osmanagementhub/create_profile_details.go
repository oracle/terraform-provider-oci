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

// CreateProfileDetails The information about new registration profile.
type CreateProfileDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the tenancy containing the registration profile.
	GetCompartmentId() *string

	// The description of the registration profile.
	GetDescription() *string

	// The OCID of the management station.
	GetManagementStationId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createprofiledetails struct {
	JsonData            []byte
	Description         *string                           `mandatory:"false" json:"description"`
	ManagementStationId *string                           `mandatory:"false" json:"managementStationId"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	ProfileType         string                            `json:"profileType"`
}

// UnmarshalJSON unmarshals json
func (m *createprofiledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateprofiledetails createprofiledetails
	s := struct {
		Model Unmarshalercreateprofiledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.ManagementStationId = s.Model.ManagementStationId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ProfileType = s.Model.ProfileType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createprofiledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ProfileType {
	case "GROUP":
		mm := CreateGroupProfileDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STATION":
		mm := CreateStationProfileDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOFTWARESOURCE":
		mm := CreateSoftwareSourceProfileDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIFECYCLE":
		mm := CreateLifecycleProfileDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateProfileDetails: %s.", m.ProfileType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createprofiledetails) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m createprofiledetails) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetFreeformTags returns FreeformTags
func (m createprofiledetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createprofiledetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetDisplayName returns DisplayName
func (m createprofiledetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createprofiledetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createprofiledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createprofiledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
