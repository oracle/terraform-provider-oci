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

// Report An object that defines an Errata or CVE report.
type Report interface {

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

	// The current state of the OsmhReporting.
	GetLifecycleState() ReportLifecycleStateEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy that the managed instance resides in.
	GetTenancyId() *string

	// User-specified description for the Osmh Report.
	GetDescription() *string

	// The compartment ids.
	GetCompartmentIds() []string

	// Indicates if sub-compartments are included in the report.
	GetIsSubCompartmentIncluded() *bool

	// List of operating system types.
	GetOsFamilies() []OsFamilyEnum

	// A message that describes the current state of the OsmhReporting in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	GetLifecycleDetails() *string
}

type report struct {
	JsonData                 []byte
	TenancyId                *string                           `mandatory:"false" json:"tenancyId"`
	Description              *string                           `mandatory:"false" json:"description"`
	CompartmentIds           []string                          `mandatory:"false" json:"compartmentIds"`
	IsSubCompartmentIncluded *bool                             `mandatory:"false" json:"isSubCompartmentIncluded"`
	OsFamilies               []OsFamilyEnum                    `mandatory:"false" json:"osFamilies,omitempty"`
	LifecycleDetails         *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                       *string                           `mandatory:"true" json:"id"`
	DisplayName              *string                           `mandatory:"true" json:"displayName"`
	CompartmentId            *string                           `mandatory:"true" json:"compartmentId"`
	ReportVersion            *string                           `mandatory:"true" json:"reportVersion"`
	TimeCreated              *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated              *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState           ReportLifecycleStateEnum          `mandatory:"true" json:"lifecycleState"`
	FreeformTags             map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	SystemTags               map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
	ReportType               string                            `json:"reportType"`
}

// UnmarshalJSON unmarshals json
func (m *report) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerreport report
	s := struct {
		Model Unmarshalerreport
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
	m.SystemTags = s.Model.SystemTags
	m.TenancyId = s.Model.TenancyId
	m.Description = s.Model.Description
	m.CompartmentIds = s.Model.CompartmentIds
	m.IsSubCompartmentIncluded = s.Model.IsSubCompartmentIncluded
	m.OsFamilies = s.Model.OsFamilies
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.ReportType = s.Model.ReportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *report) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ReportType {
	case "ERRATA":
		mm := ErrataReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CVE":
		mm := CveReport{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Report: %s.", m.ReportType)
		return *m, nil
	}
}

// GetTenancyId returns TenancyId
func (m report) GetTenancyId() *string {
	return m.TenancyId
}

// GetDescription returns Description
func (m report) GetDescription() *string {
	return m.Description
}

// GetCompartmentIds returns CompartmentIds
func (m report) GetCompartmentIds() []string {
	return m.CompartmentIds
}

// GetIsSubCompartmentIncluded returns IsSubCompartmentIncluded
func (m report) GetIsSubCompartmentIncluded() *bool {
	return m.IsSubCompartmentIncluded
}

// GetOsFamilies returns OsFamilies
func (m report) GetOsFamilies() []OsFamilyEnum {
	return m.OsFamilies
}

// GetLifecycleDetails returns LifecycleDetails
func (m report) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m report) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m report) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m report) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetReportVersion returns ReportVersion
func (m report) GetReportVersion() *string {
	return m.ReportVersion
}

// GetTimeCreated returns TimeCreated
func (m report) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m report) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m report) GetLifecycleState() ReportLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m report) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m report) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m report) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m report) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m report) ValidateEnumValue() (bool, error) {
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportLifecycleStateEnum Enum with underlying type: string
type ReportLifecycleStateEnum string

// Set of constants representing the allowable values for ReportLifecycleStateEnum
const (
	ReportLifecycleStateCreating ReportLifecycleStateEnum = "CREATING"
	ReportLifecycleStateUpdating ReportLifecycleStateEnum = "UPDATING"
	ReportLifecycleStateActive   ReportLifecycleStateEnum = "ACTIVE"
	ReportLifecycleStateDeleting ReportLifecycleStateEnum = "DELETING"
	ReportLifecycleStateDeleted  ReportLifecycleStateEnum = "DELETED"
	ReportLifecycleStateFailed   ReportLifecycleStateEnum = "FAILED"
)

var mappingReportLifecycleStateEnum = map[string]ReportLifecycleStateEnum{
	"CREATING": ReportLifecycleStateCreating,
	"UPDATING": ReportLifecycleStateUpdating,
	"ACTIVE":   ReportLifecycleStateActive,
	"DELETING": ReportLifecycleStateDeleting,
	"DELETED":  ReportLifecycleStateDeleted,
	"FAILED":   ReportLifecycleStateFailed,
}

var mappingReportLifecycleStateEnumLowerCase = map[string]ReportLifecycleStateEnum{
	"creating": ReportLifecycleStateCreating,
	"updating": ReportLifecycleStateUpdating,
	"active":   ReportLifecycleStateActive,
	"deleting": ReportLifecycleStateDeleting,
	"deleted":  ReportLifecycleStateDeleted,
	"failed":   ReportLifecycleStateFailed,
}

// GetReportLifecycleStateEnumValues Enumerates the set of values for ReportLifecycleStateEnum
func GetReportLifecycleStateEnumValues() []ReportLifecycleStateEnum {
	values := make([]ReportLifecycleStateEnum, 0)
	for _, v := range mappingReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReportLifecycleStateEnumStringValues Enumerates the set of values in String for ReportLifecycleStateEnum
func GetReportLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportLifecycleStateEnum(val string) (ReportLifecycleStateEnum, bool) {
	enum, ok := mappingReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
