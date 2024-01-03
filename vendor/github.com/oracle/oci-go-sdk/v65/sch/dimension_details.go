// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DimensionDetails A dimension name and value.
type DimensionDetails struct {

	// Dimension key. A valid dimension key includes only printable ASCII, excluding periods (.) and spaces.
	// Custom dimension keys are acceptable. Avoid entering confidential information.
	// Due to use by Connector Hub, the following dimension names are reserved: `connectorId`, `connectorName`, `connectorSourceType`.
	// For information on valid dimension keys and values, see MetricDataDetails.
	// Example: `type`
	Name *string `mandatory:"true" json:"name"`

	DimensionValue DimensionValueDetails `mandatory:"true" json:"dimensionValue"`
}

func (m DimensionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DimensionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DimensionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name           *string               `json:"name"`
		DimensionValue dimensionvaluedetails `json:"dimensionValue"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	nn, e = model.DimensionValue.UnmarshalPolymorphicJSON(model.DimensionValue.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DimensionValue = nn.(DimensionValueDetails)
	} else {
		m.DimensionValue = nil
	}

	return
}
