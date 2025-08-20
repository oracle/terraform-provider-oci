// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// FileTaskArgument A file variable that holds a value
type FileTaskArgument struct {

	// Name of the input variable
	Name *string `mandatory:"true" json:"name"`

	Content InputFileContentDetails `mandatory:"false" json:"content"`
}

// GetName returns Name
func (m FileTaskArgument) GetName() *string {
	return m.Name
}

func (m FileTaskArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileTaskArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FileTaskArgument) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFileTaskArgument FileTaskArgument
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeFileTaskArgument
	}{
		"FILE",
		(MarshalTypeFileTaskArgument)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *FileTaskArgument) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Content inputfilecontentdetails `json:"content"`
		Name    *string                 `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(InputFileContentDetails)
	} else {
		m.Content = nil
	}

	m.Name = model.Name

	return
}
