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

// OpsiDataObject OPSI data object.
type OpsiDataObject interface {

	// Unique identifier of OPSI data object.
	GetIdentifier() *string

	// User-friendly name of OPSI data object.
	GetDisplayName() *string

	// Metadata of columns in a data object.
	GetColumnsMetadata() []DataObjectColumnMetadata

	// Description of OPSI data object.
	GetDescription() *string

	// Name of the data object, which can be used in data object queries just like how view names are used in a query.
	GetName() *string

	// Names of all the groups to which the data object belongs to.
	GetGroupNames() []string

	// Time period supported by the data object for quering data.
	// Time period is in ISO 8601 format with respect to current time. Default is last 30 days represented by P30D.
	// Examples: P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months).
	GetSupportedQueryTimePeriod() *string

	// Supported query parameters by this OPSI data object that can be configured while a data object query involving this data object is executed.
	GetSupportedQueryParams() []OpsiDataObjectSupportedQueryParam
}

type opsidataobject struct {
	JsonData                 []byte
	Description              *string                             `mandatory:"false" json:"description"`
	Name                     *string                             `mandatory:"false" json:"name"`
	GroupNames               []string                            `mandatory:"false" json:"groupNames"`
	SupportedQueryTimePeriod *string                             `mandatory:"false" json:"supportedQueryTimePeriod"`
	SupportedQueryParams     []OpsiDataObjectSupportedQueryParam `mandatory:"false" json:"supportedQueryParams"`
	Identifier               *string                             `mandatory:"true" json:"identifier"`
	DisplayName              *string                             `mandatory:"true" json:"displayName"`
	ColumnsMetadata          []DataObjectColumnMetadata          `mandatory:"true" json:"columnsMetadata"`
	DataObjectType           string                              `json:"dataObjectType"`
}

// UnmarshalJSON unmarshals json
func (m *opsidataobject) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleropsidataobject opsidataobject
	s := struct {
		Model Unmarshaleropsidataobject
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Identifier = s.Model.Identifier
	m.DisplayName = s.Model.DisplayName
	m.ColumnsMetadata = s.Model.ColumnsMetadata
	m.Description = s.Model.Description
	m.Name = s.Model.Name
	m.GroupNames = s.Model.GroupNames
	m.SupportedQueryTimePeriod = s.Model.SupportedQueryTimePeriod
	m.SupportedQueryParams = s.Model.SupportedQueryParams
	m.DataObjectType = s.Model.DataObjectType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *opsidataobject) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataObjectType {
	case "HOST_INSIGHTS_DATA_OBJECT":
		mm := HostInsightsDataObject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADATA_INSIGHTS_DATA_OBJECT":
		mm := ExadataInsightsDataObject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_INSIGHTS_DATA_OBJECT":
		mm := DatabaseInsightsDataObject{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OpsiDataObject: %s.", m.DataObjectType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m opsidataobject) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m opsidataobject) GetName() *string {
	return m.Name
}

// GetGroupNames returns GroupNames
func (m opsidataobject) GetGroupNames() []string {
	return m.GroupNames
}

// GetSupportedQueryTimePeriod returns SupportedQueryTimePeriod
func (m opsidataobject) GetSupportedQueryTimePeriod() *string {
	return m.SupportedQueryTimePeriod
}

// GetSupportedQueryParams returns SupportedQueryParams
func (m opsidataobject) GetSupportedQueryParams() []OpsiDataObjectSupportedQueryParam {
	return m.SupportedQueryParams
}

// GetIdentifier returns Identifier
func (m opsidataobject) GetIdentifier() *string {
	return m.Identifier
}

// GetDisplayName returns DisplayName
func (m opsidataobject) GetDisplayName() *string {
	return m.DisplayName
}

// GetColumnsMetadata returns ColumnsMetadata
func (m opsidataobject) GetColumnsMetadata() []DataObjectColumnMetadata {
	return m.ColumnsMetadata
}

func (m opsidataobject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m opsidataobject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
