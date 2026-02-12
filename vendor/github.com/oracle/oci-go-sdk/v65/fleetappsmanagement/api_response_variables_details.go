// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApiResponseVariablesDetails Details of variables from api response
type ApiResponseVariablesDetails struct {

	// Name of the variable.
	Name *string `mandatory:"true" json:"name"`

	ValueExtractionDetails ValueExtractionDetails `mandatory:"true" json:"valueExtractionDetails"`
}

func (m ApiResponseVariablesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiResponseVariablesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ApiResponseVariablesDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name                   *string                `json:"name"`
		ValueExtractionDetails valueextractiondetails `json:"valueExtractionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	nn, e = model.ValueExtractionDetails.UnmarshalPolymorphicJSON(model.ValueExtractionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ValueExtractionDetails = nn.(ValueExtractionDetails)
	} else {
		m.ValueExtractionDetails = nil
	}

	return
}
