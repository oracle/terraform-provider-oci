// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlattenDetails Details for the flatten operator.
type FlattenDetails struct {
	FlattenProjectionPreferences *FlattenProjectionPreferences `mandatory:"true" json:"flattenProjectionPreferences"`

	// The string of flatten attribute column name where the flatten process starts.
	FlattenAttributeRoot *string `mandatory:"true" json:"flattenAttributeRoot"`

	// The string of flatten attribute path in flattenAttributeRoot from upper level to leaf/targeted level concatenated with dot(.)
	FlattenAttributePath *string `mandatory:"true" json:"flattenAttributePath"`

	// The array of flatten columns which are the input to flatten.
	FlattenColumns []TypedObject `mandatory:"true" json:"flattenColumns"`
}

func (m FlattenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlattenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FlattenDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FlattenProjectionPreferences *FlattenProjectionPreferences `json:"flattenProjectionPreferences"`
		FlattenAttributeRoot         *string                       `json:"flattenAttributeRoot"`
		FlattenAttributePath         *string                       `json:"flattenAttributePath"`
		FlattenColumns               []typedobject                 `json:"flattenColumns"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FlattenProjectionPreferences = model.FlattenProjectionPreferences

	m.FlattenAttributeRoot = model.FlattenAttributeRoot

	m.FlattenAttributePath = model.FlattenAttributePath

	m.FlattenColumns = make([]TypedObject, len(model.FlattenColumns))
	for i, n := range model.FlattenColumns {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.FlattenColumns[i] = nn.(TypedObject)
		} else {
			m.FlattenColumns[i] = nil
		}
	}

	return
}
