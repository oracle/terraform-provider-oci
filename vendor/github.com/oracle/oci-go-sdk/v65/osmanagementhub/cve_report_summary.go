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

// CveReportSummary An object that defines a CVE report Summary.
type CveReportSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Osmh Report.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The report version.
	ReportVersion *string `mandatory:"true" json:"reportVersion"`

	// The date and time the Osmh Report was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Osmh Report was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The list of cve names.
	Cves []string `mandatory:"true" json:"cves"`

	// User-specified description for the Osmh Report.
	Description *string `mandatory:"false" json:"description"`

	// A message that describes the current state of the Osmh Report in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the Osmh Report.
	LifecycleState ReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of operating system types.
	OsFamilies []OsFamilyEnum `mandatory:"false" json:"osFamilies,omitempty"`
}

// GetId returns Id
func (m CveReportSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m CveReportSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m CveReportSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetReportVersion returns ReportVersion
func (m CveReportSummary) GetReportVersion() *string {
	return m.ReportVersion
}

// GetDescription returns Description
func (m CveReportSummary) GetDescription() *string {
	return m.Description
}

// GetTimeCreated returns TimeCreated
func (m CveReportSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CveReportSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetOsFamilies returns OsFamilies
func (m CveReportSummary) GetOsFamilies() []OsFamilyEnum {
	return m.OsFamilies
}

// GetLifecycleState returns LifecycleState
func (m CveReportSummary) GetLifecycleState() ReportLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m CveReportSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m CveReportSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CveReportSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m CveReportSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m CveReportSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CveReportSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReportLifecycleStateEnumStringValues(), ",")))
	}
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
func (m CveReportSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCveReportSummary CveReportSummary
	s := struct {
		DiscriminatorParam string `json:"reportType"`
		MarshalTypeCveReportSummary
	}{
		"CVE",
		(MarshalTypeCveReportSummary)(m),
	}

	return json.Marshal(&s)
}
