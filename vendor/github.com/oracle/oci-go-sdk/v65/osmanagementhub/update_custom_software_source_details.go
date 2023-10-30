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

// UpdateCustomSoftwareSourceDetails Information for updating a custom or software source.
type UpdateCustomSoftwareSourceDetails struct {

	// The OCID of the tenancy containing the software source.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User friendly name for the software source.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Information specified by the user about the software source.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// List of vendor software sources.
	VendorSoftwareSources []Id `mandatory:"false" json:"vendorSoftwareSources"`

	CustomSoftwareSourceFilter *CustomSoftwareSourceFilter `mandatory:"false" json:"customSoftwareSourceFilter"`

	// Indicates whether service should automatically update the custom software source for the user.
	IsAutomaticallyUpdated *bool `mandatory:"false" json:"isAutomaticallyUpdated"`
}

// GetCompartmentId returns CompartmentId
func (m UpdateCustomSoftwareSourceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m UpdateCustomSoftwareSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateCustomSoftwareSourceDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateCustomSoftwareSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateCustomSoftwareSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateCustomSoftwareSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCustomSoftwareSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateCustomSoftwareSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateCustomSoftwareSourceDetails UpdateCustomSoftwareSourceDetails
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeUpdateCustomSoftwareSourceDetails
	}{
		"CUSTOM",
		(MarshalTypeUpdateCustomSoftwareSourceDetails)(m),
	}

	return json.Marshal(&s)
}
