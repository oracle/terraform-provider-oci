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

// CreateArchiveFunctionSourceDetails Source details for creating an archive‑based function.
type CreateArchiveFunctionSourceDetails struct {
	ArchiveSourceDetails CreateArchiveSourceDetails `mandatory:"true" json:"archiveSourceDetails"`

	RuntimeConfig CreateRuntimeConfig `mandatory:"true" json:"runtimeConfig"`

	// The function handler that is executed when the function is invoked. The value of this field depends on the runtime used
	Handler *string `mandatory:"false" json:"handler"`
}

func (m CreateArchiveFunctionSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateArchiveFunctionSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateArchiveFunctionSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateArchiveFunctionSourceDetails CreateArchiveFunctionSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateArchiveFunctionSourceDetails
	}{
		"ARCHIVE",
		(MarshalTypeCreateArchiveFunctionSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateArchiveFunctionSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Handler              *string                    `json:"handler"`
		ArchiveSourceDetails createarchivesourcedetails `json:"archiveSourceDetails"`
		RuntimeConfig        createruntimeconfig        `json:"runtimeConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Handler = model.Handler

	nn, e = model.ArchiveSourceDetails.UnmarshalPolymorphicJSON(model.ArchiveSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArchiveSourceDetails = nn.(CreateArchiveSourceDetails)
	} else {
		m.ArchiveSourceDetails = nil
	}

	nn, e = model.RuntimeConfig.UnmarshalPolymorphicJSON(model.RuntimeConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RuntimeConfig = nn.(CreateRuntimeConfig)
	} else {
		m.RuntimeConfig = nil
	}

	return
}
