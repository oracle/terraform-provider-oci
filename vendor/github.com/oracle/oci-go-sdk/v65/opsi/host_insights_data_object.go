// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostInsightsDataObject Host insights data object.
type HostInsightsDataObject struct {

	// Unique identifier of OPSI data object.
	Identifier *string `mandatory:"true" json:"identifier"`

	// User-friendly name of OPSI data object.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Metadata of columns in a data object.
	ColumnsMetadata []DataObjectColumnMetadata `mandatory:"true" json:"columnsMetadata"`

	// Description of OPSI data object.
	Description *string `mandatory:"false" json:"description"`

	// Name of the data object, which can be used in data object queries just like how view names are used in a query.
	Name *string `mandatory:"false" json:"name"`

	// Names of all the groups to which the data object belongs to.
	GroupNames []string `mandatory:"false" json:"groupNames"`

	// Time period supported by the data object for quering data.
	// Time period is in ISO 8601 format with respect to current time. Default is last 30 days represented by P30D.
	// Examples: P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months).
	SupportedQueryTimePeriod *string `mandatory:"false" json:"supportedQueryTimePeriod"`

	// Supported query parameters by this OPSI data object that can be configured while a data object query involving this data object is executed.
	SupportedQueryParams []OpsiDataObjectSupportedQueryParam `mandatory:"false" json:"supportedQueryParams"`
}

// GetIdentifier returns Identifier
func (m HostInsightsDataObject) GetIdentifier() *string {
	return m.Identifier
}

// GetDisplayName returns DisplayName
func (m HostInsightsDataObject) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m HostInsightsDataObject) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m HostInsightsDataObject) GetName() *string {
	return m.Name
}

// GetGroupNames returns GroupNames
func (m HostInsightsDataObject) GetGroupNames() []string {
	return m.GroupNames
}

// GetSupportedQueryTimePeriod returns SupportedQueryTimePeriod
func (m HostInsightsDataObject) GetSupportedQueryTimePeriod() *string {
	return m.SupportedQueryTimePeriod
}

// GetColumnsMetadata returns ColumnsMetadata
func (m HostInsightsDataObject) GetColumnsMetadata() []DataObjectColumnMetadata {
	return m.ColumnsMetadata
}

// GetSupportedQueryParams returns SupportedQueryParams
func (m HostInsightsDataObject) GetSupportedQueryParams() []OpsiDataObjectSupportedQueryParam {
	return m.SupportedQueryParams
}

func (m HostInsightsDataObject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostInsightsDataObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostInsightsDataObject) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostInsightsDataObject HostInsightsDataObject
	s := struct {
		DiscriminatorParam string `json:"dataObjectType"`
		MarshalTypeHostInsightsDataObject
	}{
		"HOST_INSIGHTS_DATA_OBJECT",
		(MarshalTypeHostInsightsDataObject)(m),
	}

	return json.Marshal(&s)
}
