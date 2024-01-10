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

// DataObjectRateColumnUnit Unit details of a data object column of RATE unit category.
type DataObjectRateColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Numerator DataObjectColumnUnit `mandatory:"false" json:"numerator"`

	Denominator DataObjectColumnUnit `mandatory:"false" json:"denominator"`
}

// GetDisplayName returns DisplayName
func (m DataObjectRateColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectRateColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectRateColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectRateColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectRateColumnUnit DataObjectRateColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectRateColumnUnit
	}{
		"RATE",
		(MarshalTypeDataObjectRateColumnUnit)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataObjectRateColumnUnit) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName *string              `json:"displayName"`
		Numerator   dataobjectcolumnunit `json:"numerator"`
		Denominator dataobjectcolumnunit `json:"denominator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	nn, e = model.Numerator.UnmarshalPolymorphicJSON(model.Numerator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Numerator = nn.(DataObjectColumnUnit)
	} else {
		m.Numerator = nil
	}

	nn, e = model.Denominator.UnmarshalPolymorphicJSON(model.Denominator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Denominator = nn.(DataObjectColumnUnit)
	} else {
		m.Denominator = nil
	}

	return
}
