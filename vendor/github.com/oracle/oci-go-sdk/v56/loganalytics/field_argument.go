// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FieldArgument QueryString argument of type field.
type FieldArgument struct {
	Value AbstractField `mandatory:"false" json:"value"`
}

func (m FieldArgument) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FieldArgument) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFieldArgument FieldArgument
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFieldArgument
	}{
		"FIELD",
		(MarshalTypeFieldArgument)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *FieldArgument) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Value abstractfield `json:"value"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Value.UnmarshalPolymorphicJSON(model.Value.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Value = nn.(AbstractField)
	} else {
		m.Value = nil
	}

	return
}
