// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateArchiveFunctionSourceDetails Source details for updating an archive‑based function.
type UpdateArchiveFunctionSourceDetails struct {
	ArchiveSourceDetails UpdateArchiveSourceDetails `mandatory:"false" json:"archiveSourceDetails"`

	// The function handler that is executed when the function is invoked. The value of this field depends on the runtime used
	Handler *string `mandatory:"false" json:"handler"`

	RuntimeConfig UpdateRuntimeConfig `mandatory:"false" json:"runtimeConfig"`
}

func (m UpdateArchiveFunctionSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateArchiveFunctionSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateArchiveFunctionSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateArchiveFunctionSourceDetails UpdateArchiveFunctionSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUpdateArchiveFunctionSourceDetails
	}{
		"ARCHIVE",
		(MarshalTypeUpdateArchiveFunctionSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateArchiveFunctionSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ArchiveSourceDetails updatearchivesourcedetails `json:"archiveSourceDetails"`
		Handler              *string                    `json:"handler"`
		RuntimeConfig        updateruntimeconfig        `json:"runtimeConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ArchiveSourceDetails.UnmarshalPolymorphicJSON(model.ArchiveSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArchiveSourceDetails = nn.(UpdateArchiveSourceDetails)
	} else {
		m.ArchiveSourceDetails = nil
	}

	m.Handler = model.Handler

	nn, e = model.RuntimeConfig.UnmarshalPolymorphicJSON(model.RuntimeConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RuntimeConfig = nn.(UpdateRuntimeConfig)
	} else {
		m.RuntimeConfig = nil
	}

	return
}
