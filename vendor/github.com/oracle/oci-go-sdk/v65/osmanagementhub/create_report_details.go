// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateReportDetails The data to create a Osmh Report.
type CreateReportDetails interface {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the OsmhReporting in.
	GetCompartmentId() *string

	// User-specified description for the Osmh Report.
	GetDescription() *string

	// List of operating system types.
	GetOsFamilies() []OsFamilyEnum

	// The compartment ids.
	GetCompartmentIds() []string

	// Indicates if sub-compartments are included in the report.
	GetIsSubCompartmentIncluded() *bool

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createreportdetails struct {
	JsonData                 []byte
	Description              *string                           `mandatory:"false" json:"description"`
	OsFamilies               []OsFamilyEnum                    `mandatory:"false" json:"osFamilies,omitempty"`
	CompartmentIds           []string                          `mandatory:"false" json:"compartmentIds"`
	IsSubCompartmentIncluded *bool                             `mandatory:"false" json:"isSubCompartmentIncluded"`
	FreeformTags             map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DisplayName              *string                           `mandatory:"true" json:"displayName"`
	CompartmentId            *string                           `mandatory:"true" json:"compartmentId"`
	ReportType               string                            `json:"reportType"`
}

// UnmarshalJSON unmarshals json
func (m *createreportdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatereportdetails createreportdetails
	s := struct {
		Model Unmarshalercreatereportdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.OsFamilies = s.Model.OsFamilies
	m.CompartmentIds = s.Model.CompartmentIds
	m.IsSubCompartmentIncluded = s.Model.IsSubCompartmentIncluded
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ReportType = s.Model.ReportType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createreportdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ReportType {
	case "CVE":
		mm := CreateCveReportDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ERRATA":
		mm := CreateErrataReportDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateReportDetails: %s.", m.ReportType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createreportdetails) GetDescription() *string {
	return m.Description
}

// GetOsFamilies returns OsFamilies
func (m createreportdetails) GetOsFamilies() []OsFamilyEnum {
	return m.OsFamilies
}

// GetCompartmentIds returns CompartmentIds
func (m createreportdetails) GetCompartmentIds() []string {
	return m.CompartmentIds
}

// GetIsSubCompartmentIncluded returns IsSubCompartmentIncluded
func (m createreportdetails) GetIsSubCompartmentIncluded() *bool {
	return m.IsSubCompartmentIncluded
}

// GetFreeformTags returns FreeformTags
func (m createreportdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createreportdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetDisplayName returns DisplayName
func (m createreportdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m createreportdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createreportdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createreportdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.OsFamilies {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && string(val) != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamilies: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
