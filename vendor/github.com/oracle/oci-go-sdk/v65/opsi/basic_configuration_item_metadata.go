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

// BasicConfigurationItemMetadata Basic configuration item metadata.
type BasicConfigurationItemMetadata struct {

	// User-friendly display name for the configuration item.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of configuration item .
	Description *string `mandatory:"false" json:"description"`

	// Data type of configuration item.
	// Examples: STRING, BOOLEAN, NUMBER
	DataType *string `mandatory:"false" json:"dataType"`

	UnitDetails *ConfigurationItemUnitDetails `mandatory:"false" json:"unitDetails"`

	ValueInputDetails ConfigurationItemAllowedValueDetails `mandatory:"false" json:"valueInputDetails"`
}

func (m BasicConfigurationItemMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicConfigurationItemMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicConfigurationItemMetadata) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicConfigurationItemMetadata BasicConfigurationItemMetadata
	s := struct {
		DiscriminatorParam string `json:"configItemType"`
		MarshalTypeBasicConfigurationItemMetadata
	}{
		"BASIC",
		(MarshalTypeBasicConfigurationItemMetadata)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *BasicConfigurationItemMetadata) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                              `json:"displayName"`
		Description       *string                              `json:"description"`
		DataType          *string                              `json:"dataType"`
		UnitDetails       *ConfigurationItemUnitDetails        `json:"unitDetails"`
		ValueInputDetails configurationitemallowedvaluedetails `json:"valueInputDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.DataType = model.DataType

	m.UnitDetails = model.UnitDetails

	nn, e = model.ValueInputDetails.UnmarshalPolymorphicJSON(model.ValueInputDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValueInputDetails = nn.(ConfigurationItemAllowedValueDetails)
	} else {
		m.ValueInputDetails = nil
	}

	return
}
