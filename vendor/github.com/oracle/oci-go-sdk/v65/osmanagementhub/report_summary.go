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

// ReportSummary Provides summary information for a report.
type ReportSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Osmh Report.
	GetId() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The report version.
	GetReportVersion() *string

	// The date and time the Osmh Report was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The date and time the Osmh Report was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeUpdated() *common.SDKTime

	// The current state of the Osmh Report.
	GetLifecycleState() ReportLifecycleStateEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy that the managed instance resides in.
	GetTenancyId() *string

	// User-specified description for the Osmh Report.
	GetDescription() *string

	// A message that describes the current state of the Osmh Report in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	GetLifecycleDetails() *string

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type reportsummary struct {
	JsonData         []byte
	TenancyId        *string                           `mandatory:"false" json:"tenancyId"`
	Description      *string                           `mandatory:"false" json:"description"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	ReportVersion    *string                           `mandatory:"true" json:"reportVersion"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState   ReportLifecycleStateEnum          `mandatory:"true" json:"lifecycleState"`
	FreeformTags     map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	ReportType       string                            `json:"reportType"`
}

// UnmarshalJSON unmarshals json
func (m *reportsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreportsummary reportsummary
	s := struct {
		Model Unmarshalerreportsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.ReportVersion = s.Model.ReportVersion
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TenancyId = s.Model.TenancyId
	m.Description = s.Model.Description
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.SystemTags = s.Model.SystemTags
	m.ReportType = s.Model.ReportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *reportsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ReportType {
	case "ERRATA":
		mm := ErrataReportSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CVE":
		mm := CveReportSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ReportSummary: %s.", m.ReportType)
		return *m, nil
	}
}

// GetTenancyId returns TenancyId
func (m reportsummary) GetTenancyId() *string {
	return m.TenancyId
}

// GetDescription returns Description
func (m reportsummary) GetDescription() *string {
	return m.Description
}

// GetLifecycleDetails returns LifecycleDetails
func (m reportsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetSystemTags returns SystemTags
func (m reportsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m reportsummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m reportsummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m reportsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetReportVersion returns ReportVersion
func (m reportsummary) GetReportVersion() *string {
	return m.ReportVersion
}

// GetTimeCreated returns TimeCreated
func (m reportsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m reportsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m reportsummary) GetLifecycleState() ReportLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m reportsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m reportsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m reportsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m reportsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
