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

// CreateCveReportDetails An object that provides data to create a CVE report.
type CreateCveReportDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the OsmhReporting in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of cve names.
	Cves []string `mandatory:"true" json:"cves"`

	// User-specified description for the Osmh Report.
	Description *string `mandatory:"false" json:"description"`

	// The compartment ids.
	CompartmentIds []string `mandatory:"false" json:"compartmentIds"`

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

	// List of operating system types.
	OsFamilies []OsFamilyEnum `mandatory:"false" json:"osFamilies,omitempty"`
}

// GetDisplayName returns DisplayName
func (m CreateCveReportDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateCveReportDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateCveReportDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetOsFamilies returns OsFamilies
func (m CreateCveReportDetails) GetOsFamilies() []OsFamilyEnum {
	return m.OsFamilies
}

// GetCompartmentIds returns CompartmentIds
func (m CreateCveReportDetails) GetCompartmentIds() []string {
	return m.CompartmentIds
}

// GetIsSubCompartmentIncluded returns IsSubCompartmentIncluded
func (m CreateCveReportDetails) GetIsSubCompartmentIncluded() *bool {
	return m.IsSubCompartmentIncluded
}

// GetFreeformTags returns FreeformTags
func (m CreateCveReportDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateCveReportDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateCveReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCveReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.OsFamilies {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamilies: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCveReportDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCveReportDetails CreateCveReportDetails
	s := struct {
		DiscriminatorParam string `json:"reportType"`
		MarshalTypeCreateCveReportDetails
	}{
		"CVE",
		(MarshalTypeCreateCveReportDetails)(m),
	}

	return json.Marshal(&s)
}
