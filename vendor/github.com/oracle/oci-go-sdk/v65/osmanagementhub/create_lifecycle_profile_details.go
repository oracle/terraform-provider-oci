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

// CreateLifecycleProfileDetails Provides the information used to create a lifecycle environment registration profile.
type CreateLifecycleProfileDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the registration profile.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage that the instance will be associated with.
	LifecycleStageId *string `mandatory:"true" json:"lifecycleStageId"`

	// User-specified description of the registration profile.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station to associate with an instance once registered. Associating with a management station applies only to non-OCI instances.
	ManagementStationId *string `mandatory:"false" json:"managementStationId"`

	// Indicates if the profile is set as the default. There is exactly one default profile for a specified architecture, OS family, registration type, and vendor. When registering an instance with the corresonding characteristics, the default profile is used, unless another profile is specified.
	IsDefaultProfile *bool `mandatory:"false" json:"isDefaultProfile"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The type of instance to register.
	RegistrationType ProfileRegistrationTypeEnum `mandatory:"false" json:"registrationType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateLifecycleProfileDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CreateLifecycleProfileDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDescription returns Description
func (m CreateLifecycleProfileDetails) GetDescription() *string {
	return m.Description
}

// GetManagementStationId returns ManagementStationId
func (m CreateLifecycleProfileDetails) GetManagementStationId() *string {
	return m.ManagementStationId
}

// GetRegistrationType returns RegistrationType
func (m CreateLifecycleProfileDetails) GetRegistrationType() ProfileRegistrationTypeEnum {
	return m.RegistrationType
}

// GetIsDefaultProfile returns IsDefaultProfile
func (m CreateLifecycleProfileDetails) GetIsDefaultProfile() *bool {
	return m.IsDefaultProfile
}

// GetFreeformTags returns FreeformTags
func (m CreateLifecycleProfileDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateLifecycleProfileDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLifecycleProfileDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLifecycleProfileDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProfileRegistrationTypeEnum(string(m.RegistrationType)); !ok && m.RegistrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationType: %s. Supported values are: %s.", m.RegistrationType, strings.Join(GetProfileRegistrationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateLifecycleProfileDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLifecycleProfileDetails CreateLifecycleProfileDetails
	s := struct {
		DiscriminatorParam string `json:"profileType"`
		MarshalTypeCreateLifecycleProfileDetails
	}{
		"LIFECYCLE",
		(MarshalTypeCreateLifecycleProfileDetails)(m),
	}

	return json.Marshal(&s)
}
