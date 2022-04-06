// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

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
}

type opsidataobject struct {
	JsonData        []byte
	Identifier      *string                    `mandatory:"true" json:"identifier"`
	DisplayName     *string                    `mandatory:"true" json:"displayName"`
	ColumnsMetadata []DataObjectColumnMetadata `mandatory:"true" json:"columnsMetadata"`
	Description     *string                    `mandatory:"false" json:"description"`
	DataObjectType  string                     `json:"dataObjectType"`
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
		return *m, nil
	}
}

//GetIdentifier returns Identifier
func (m opsidataobject) GetIdentifier() *string {
	return m.Identifier
}

//GetDisplayName returns DisplayName
func (m opsidataobject) GetDisplayName() *string {
	return m.DisplayName
}

//GetColumnsMetadata returns ColumnsMetadata
func (m opsidataobject) GetColumnsMetadata() []DataObjectColumnMetadata {
	return m.ColumnsMetadata
}

//GetDescription returns Description
func (m opsidataobject) GetDescription() *string {
	return m.Description
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
