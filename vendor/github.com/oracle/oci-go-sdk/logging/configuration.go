// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// Configuration Log object configuration.
type Configuration struct {
	Source Source `mandatory:"true" json:"source"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	Archiving *Archiving `mandatory:"false" json:"archiving"`
}

func (m Configuration) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Configuration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CompartmentId *string    `json:"compartmentId"`
		Archiving     *Archiving `json:"archiving"`
		Source        source     `json:"source"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CompartmentId = model.CompartmentId

	m.Archiving = model.Archiving

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(Source)
	} else {
		m.Source = nil
	}

	return
}
