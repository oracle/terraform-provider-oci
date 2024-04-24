// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVersionedCustomSoftwareSourceDetails Provides the information used to create a versioned custom software source.
type CreateVersionedCustomSoftwareSourceDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of vendor software sources.
	VendorSoftwareSources []Id `mandatory:"true" json:"vendorSoftwareSources"`

	// The version to assign to this custom software source.
	SoftwareSourceVersion *string `mandatory:"true" json:"softwareSourceVersion"`

	// User-friendly name for the software source. Does not have to be unique and you can change the name later. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified description for the software source. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	CustomSoftwareSourceFilter *CustomSoftwareSourceFilter `mandatory:"false" json:"customSoftwareSourceFilter"`

	// Indicates whether the service should automatically resolve package dependencies when including specific packages in the software source.
	IsAutoResolveDependencies *bool `mandatory:"false" json:"isAutoResolveDependencies"`

	// Indicates whether the service should create the software source from a list of packages provided by the user.
	IsCreatedFromPackageList *bool `mandatory:"false" json:"isCreatedFromPackageList"`

	// A property used for compatibility only. It doesn't provide a complete list of packages. See AddPackagesToSoftwareSourceDetails for providing the list of packages used to create the software source when isCreatedFromPackageList is set to true.
	Packages []string `mandatory:"false" json:"packages"`
}

// GetCompartmentId returns CompartmentId
func (m CreateVersionedCustomSoftwareSourceDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateVersionedCustomSoftwareSourceDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateVersionedCustomSoftwareSourceDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m CreateVersionedCustomSoftwareSourceDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateVersionedCustomSoftwareSourceDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateVersionedCustomSoftwareSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVersionedCustomSoftwareSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateVersionedCustomSoftwareSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateVersionedCustomSoftwareSourceDetails CreateVersionedCustomSoftwareSourceDetails
	s := struct {
		DiscriminatorParam string `json:"softwareSourceType"`
		MarshalTypeCreateVersionedCustomSoftwareSourceDetails
	}{
		"VERSIONED",
		(MarshalTypeCreateVersionedCustomSoftwareSourceDetails)(m),
	}

	return json.Marshal(&s)
}
