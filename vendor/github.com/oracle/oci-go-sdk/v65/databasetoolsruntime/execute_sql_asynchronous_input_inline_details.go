// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlAsynchronousInputInlineDetails Inline scipt input.
type ExecuteSqlAsynchronousInputInlineDetails struct {
	Content ExecuteSqlInputDetails `mandatory:"true" json:"content"`
}

func (m ExecuteSqlAsynchronousInputInlineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlAsynchronousInputInlineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlAsynchronousInputInlineDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlAsynchronousInputInlineDetails ExecuteSqlAsynchronousInputInlineDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlAsynchronousInputInlineDetails
	}{
		"INLINE",
		(MarshalTypeExecuteSqlAsynchronousInputInlineDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExecuteSqlAsynchronousInputInlineDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Content executesqlinputdetails `json:"content"`
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
		m.Content = nn.(ExecuteSqlInputDetails)
	} else {
		m.Content = nil
	}

	return
}
