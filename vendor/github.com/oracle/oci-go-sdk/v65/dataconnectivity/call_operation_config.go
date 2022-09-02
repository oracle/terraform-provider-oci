// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CallOperationConfig Holder for parameter names.
type CallOperationConfig struct {

	// The list of names of the IN/INOUT parameters.
	InFields []string `mandatory:"false" json:"inFields"`

	// The list of names of the OUT/INOUT parameters.
	OutFields []string `mandatory:"false" json:"outFields"`

	CallAttribute AbstractCallAttribute `mandatory:"false" json:"callAttribute"`

	// The List of push down operations.
	PushDownOperations []PushDownOperation `mandatory:"false" json:"pushDownOperations"`
}

func (m CallOperationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CallOperationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CallOperationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InFields           []string              `json:"inFields"`
		OutFields          []string              `json:"outFields"`
		CallAttribute      abstractcallattribute `json:"callAttribute"`
		PushDownOperations []pushdownoperation   `json:"pushDownOperations"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InFields = make([]string, len(model.InFields))
	for i, n := range model.InFields {
		m.InFields[i] = n
	}

	m.OutFields = make([]string, len(model.OutFields))
	for i, n := range model.OutFields {
		m.OutFields[i] = n
	}

	nn, e = model.CallAttribute.UnmarshalPolymorphicJSON(model.CallAttribute.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CallAttribute = nn.(AbstractCallAttribute)
	} else {
		m.CallAttribute = nil
	}

	m.PushDownOperations = make([]PushDownOperation, len(model.PushDownOperations))
	for i, n := range model.PushDownOperations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PushDownOperations[i] = nn.(PushDownOperation)
		} else {
			m.PushDownOperations[i] = nil
		}
	}

	return
}
