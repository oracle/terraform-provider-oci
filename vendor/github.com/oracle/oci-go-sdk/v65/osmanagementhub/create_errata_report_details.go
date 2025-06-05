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

// CreateErrataReportDetails An object that provides data to create an Errata report.
type CreateErrataReportDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the OsmhReporting in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified description for the Osmh Report.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if sub-compartments are included in the report.
	IsSubCompartmentIncluded *bool `mandatory:"false" json:"isSubCompartmentIncluded"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The start issue date to filter by
	TimeStartIssueDate *common.SDKTime `mandatory:"false" json:"timeStartIssueDate"`

	// The end issue date to filter by
	TimeEndIssueDate *common.SDKTime `mandatory:"false" json:"timeEndIssueDate"`

	// The managed instance ids.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The managed instance group ids.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The software source ids.
	SoftwareSourceIds []string `mandatory:"false" json:"softwareSourceIds"`

	// The dynamic set ids.
	DynamicSetIds []string `mandatory:"false" json:"dynamicSetIds"`

	// The compartment ids.
	CompartmentIds []string `mandatory:"false" json:"compartmentIds"`

	// List of operating system vendors.
	Vendors []VendorNameEnum `mandatory:"false" json:"vendors,omitempty"`

	// List of errata types.
	ClassificationTypes []ClassificationTypesEnum `mandatory:"false" json:"classificationTypes,omitempty"`

	// List of errata severities.
	AdvisorySeverities []AdvisorySeverityEnum `mandatory:"false" json:"advisorySeverities,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateErrataReportDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateErrataReportDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateErrataReportDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetIsSubCompartmentIncluded returns IsSubCompartmentIncluded
func (m CreateErrataReportDetails) GetIsSubCompartmentIncluded() *bool {
	return m.IsSubCompartmentIncluded
}

// GetFreeformTags returns FreeformTags
func (m CreateErrataReportDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateErrataReportDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateErrataReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateErrataReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Vendors {
		if _, ok := GetMappingVendorNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Vendors: %s. Supported values are: %s.", val, strings.Join(GetVendorNameEnumStringValues(), ",")))
		}
	}

	for _, val := range m.ClassificationTypes {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationTypes: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	for _, val := range m.AdvisorySeverities {
		if _, ok := GetMappingAdvisorySeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisorySeverities: %s. Supported values are: %s.", val, strings.Join(GetAdvisorySeverityEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateErrataReportDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateErrataReportDetails CreateErrataReportDetails
	s := struct {
		DiscriminatorParam string `json:"reportType"`
		MarshalTypeCreateErrataReportDetails
	}{
		"ERRATA",
		(MarshalTypeCreateErrataReportDetails)(m),
	}

	return json.Marshal(&s)
}
