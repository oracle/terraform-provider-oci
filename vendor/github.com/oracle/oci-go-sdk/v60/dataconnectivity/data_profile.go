// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// DataProfile The data profile response.
type DataProfile struct {

	// Entity name for which prodilig is requested.
	EntityName *string `mandatory:"true" json:"entityName"`

	EntityProfileResult *EntityProfileResult `mandatory:"false" json:"entityProfileResult"`

	// Array of profiling results
	AttributeProfileResults []AttributeProfileResult `mandatory:"false" json:"attributeProfileResults"`
}

func (m DataProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DataProfile) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EntityProfileResult     *EntityProfileResult     `json:"entityProfileResult"`
		AttributeProfileResults []attributeprofileresult `json:"attributeProfileResults"`
		EntityName              *string                  `json:"entityName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EntityProfileResult = model.EntityProfileResult

	m.AttributeProfileResults = make([]AttributeProfileResult, len(model.AttributeProfileResults))
	for i, n := range model.AttributeProfileResults {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.AttributeProfileResults[i] = nn.(AttributeProfileResult)
		} else {
			m.AttributeProfileResults[i] = nil
		}
	}

	m.EntityName = model.EntityName

	return
}
