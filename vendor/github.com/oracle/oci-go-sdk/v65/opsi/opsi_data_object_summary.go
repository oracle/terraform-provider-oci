// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpsiDataObjectSummary Summary of an OPSI data object.
type OpsiDataObjectSummary interface {

	// Unique identifier of OPSI data object.
	GetIdentifier() *string

	// User-friendly name of OPSI data object.
	GetDisplayName() *string

	// Description of OPSI data object.
	GetDescription() *string

	// Name of the data object, which can be used in data object queries just like how view names are used in a query.
	GetName() *string

	// Names of all the groups to which the data object belongs to.
	GetGroupNames() []string
}

type opsidataobjectsummary struct {
	JsonData       []byte
	Description    *string  `mandatory:"false" json:"description"`
	Name           *string  `mandatory:"false" json:"name"`
	GroupNames     []string `mandatory:"false" json:"groupNames"`
	Identifier     *string  `mandatory:"true" json:"identifier"`
	DisplayName    *string  `mandatory:"true" json:"displayName"`
	DataObjectType string   `json:"dataObjectType"`
}

// UnmarshalJSON unmarshals json
func (m *opsidataobjectsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleropsidataobjectsummary opsidataobjectsummary
	s := struct {
		Model Unmarshaleropsidataobjectsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Identifier = s.Model.Identifier
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.Name = s.Model.Name
	m.GroupNames = s.Model.GroupNames
	m.DataObjectType = s.Model.DataObjectType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *opsidataobjectsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataObjectType {
	case "HOST_INSIGHTS_DATA_OBJECT":
		mm := HostInsightsDataObjectSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_INSIGHTS_DATA_OBJECT":
		mm := DatabaseInsightsDataObjectSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXADATA_INSIGHTS_DATA_OBJECT":
		mm := ExadataInsightsDataObjectSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OpsiDataObjectSummary: %s.", m.DataObjectType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m opsidataobjectsummary) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m opsidataobjectsummary) GetName() *string {
	return m.Name
}

// GetGroupNames returns GroupNames
func (m opsidataobjectsummary) GetGroupNames() []string {
	return m.GroupNames
}

// GetIdentifier returns Identifier
func (m opsidataobjectsummary) GetIdentifier() *string {
	return m.Identifier
}

// GetDisplayName returns DisplayName
func (m opsidataobjectsummary) GetDisplayName() *string {
	return m.DisplayName
}

func (m opsidataobjectsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m opsidataobjectsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
