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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FieldArgument QueryString argument of type field.
type FieldArgument struct {
	Value AbstractField `mandatory:"false" json:"value"`
}

func (m FieldArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FieldArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
