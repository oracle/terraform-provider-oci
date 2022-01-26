// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConditionGroup Condition configured on a target
type ConditionGroup struct {

	// compartment associated with condition
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Condition Condition `mandatory:"true" json:"condition"`
}

func (m ConditionGroup) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConditionGroup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId *string   `json:"compartmentId"`
		Condition     condition `json:"condition"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	nn, e = model.Condition.UnmarshalPolymorphicJSON(model.Condition.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Condition = nn.(Condition)
	} else {
		m.Condition = nil
	}

	return
}
