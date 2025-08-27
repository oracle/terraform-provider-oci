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

// UpdateThirdPartySoftwareSourceDetails Provides the information used to update a third-party software source.
type UpdateThirdPartySoftwareSourceDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-friendly name for the software source.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified description of the software source.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// URL for the third-party software source.
	Url *string `mandatory:"false" json:"url"`

	// URI of the GPG key for this software source.
	GpgKeyUrl *string `mandatory:"false" json:"gpgKeyUrl"`

	// Whether signature verification should be done for the software source.
	IsGpgCheckEnabled *bool `mandatory:"false" json:"isGpgCheckEnabled"`

	// Whether SSL validation needs to be turned on
	IsSslVerifyEnabled *bool `mandatory:"false" json:"isSslVerifyEnabled"`

	// Advanced repository options for the software source
	AdvancedRepoOptions *string `mandatory:"false" json:"advancedRepoOptions"`

	// Whether this software source can be synced to a management station
	IsMirrorSyncAllowed *bool `mandatory:"false" json:"isMirrorSyncAllowed"`
}

// GetCompartmentId returns CompartmentId
func (m UpdateThirdPartySoftwareSourceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m UpdateThirdPartySoftwareSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateThirdPartySoftwareSourceDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateThirdPartySoftwareSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateThirdPartySoftwareSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateThirdPartySoftwareSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateThirdPartySoftwareSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateThirdPartySoftwareSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateThirdPartySoftwareSourceDetails UpdateThirdPartySoftwareSourceDetails
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeUpdateThirdPartySoftwareSourceDetails
	}{
		"THIRD_PARTY",
		(MarshalTypeUpdateThirdPartySoftwareSourceDetails)(m),
	}

	return json.Marshal(&s)
}
