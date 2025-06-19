// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CodeCopy A copied instance of a code snippet with its license expressions.
type CodeCopy struct {

	// Array of Licenses.
	Licenses []License `mandatory:"true" json:"licenses"`

	SourceLocation CodeSnippetSourceLocation `mandatory:"true" json:"sourceLocation"`
}

func (m CodeCopy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CodeCopy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CodeCopy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Licenses       []License                 `json:"licenses"`
		SourceLocation codesnippetsourcelocation `json:"sourceLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Licenses = make([]License, len(model.Licenses))
	copy(m.Licenses, model.Licenses)
	nn, e = model.SourceLocation.UnmarshalPolymorphicJSON(model.SourceLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceLocation = nn.(CodeSnippetSourceLocation)
	} else {
		m.SourceLocation = nil
	}

	return
}
