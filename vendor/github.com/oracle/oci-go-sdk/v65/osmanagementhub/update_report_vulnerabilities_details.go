// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateReportVulnerabilitiesDetails Vulnerability details for update.
type UpdateReportVulnerabilitiesDetails struct {
	VulnerabilityDetails VulnerabilityDetails `mandatory:"true" json:"vulnerabilityDetails"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates whether the managed instances should use the required Software Source to execute the vulnerability
	// update (even if it is not attached to it).
	ShouldUseMissingSoftwareSources *bool `mandatory:"false" json:"shouldUseMissingSoftwareSources"`

	// The managed instance ids.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The managed instance group ids.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The dynamic set ids.
	DynamicSetIds []string `mandatory:"false" json:"dynamicSetIds"`

	// The compartment ids.
	CompartmentIds []string `mandatory:"false" json:"compartmentIds"`
}

func (m UpdateReportVulnerabilitiesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateReportVulnerabilitiesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateReportVulnerabilitiesDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                     *string              `json:"displayName"`
		ShouldUseMissingSoftwareSources *bool                `json:"shouldUseMissingSoftwareSources"`
		ManagedInstanceIds              []string             `json:"managedInstanceIds"`
		ManagedInstanceGroupIds         []string             `json:"managedInstanceGroupIds"`
		DynamicSetIds                   []string             `json:"dynamicSetIds"`
		CompartmentIds                  []string             `json:"compartmentIds"`
		VulnerabilityDetails            vulnerabilitydetails `json:"vulnerabilityDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.ShouldUseMissingSoftwareSources = model.ShouldUseMissingSoftwareSources

	m.ManagedInstanceIds = make([]string, len(model.ManagedInstanceIds))
	copy(m.ManagedInstanceIds, model.ManagedInstanceIds)
	m.ManagedInstanceGroupIds = make([]string, len(model.ManagedInstanceGroupIds))
	copy(m.ManagedInstanceGroupIds, model.ManagedInstanceGroupIds)
	m.DynamicSetIds = make([]string, len(model.DynamicSetIds))
	copy(m.DynamicSetIds, model.DynamicSetIds)
	m.CompartmentIds = make([]string, len(model.CompartmentIds))
	copy(m.CompartmentIds, model.CompartmentIds)
	nn, e = model.VulnerabilityDetails.UnmarshalPolymorphicJSON(model.VulnerabilityDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.VulnerabilityDetails = nn.(VulnerabilityDetails)
	} else {
		m.VulnerabilityDetails = nil
	}

	return
}
