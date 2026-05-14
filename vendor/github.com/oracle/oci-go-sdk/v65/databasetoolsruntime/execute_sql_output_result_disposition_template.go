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

// ExecuteSqlOutputResultDispositionTemplate Template defining how the result of a command should be stored.
type ExecuteSqlOutputResultDispositionTemplate struct {

	// Commands matching this statement type will use this result disposition.
	StatementType StatementTypeEnum `mandatory:"true" json:"statementType"`

	ObjectTemplate ExecuteSqlOutputDispositionDetails `mandatory:"true" json:"objectTemplate"`
}

func (m ExecuteSqlOutputResultDispositionTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlOutputResultDispositionTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStatementTypeEnum(string(m.StatementType)); !ok && m.StatementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StatementType: %s. Supported values are: %s.", m.StatementType, strings.Join(GetStatementTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExecuteSqlOutputResultDispositionTemplate) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StatementType  StatementTypeEnum                  `json:"statementType"`
		ObjectTemplate executesqloutputdispositiondetails `json:"objectTemplate"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StatementType = model.StatementType

	nn, e = model.ObjectTemplate.UnmarshalPolymorphicJSON(model.ObjectTemplate.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ObjectTemplate = nn.(ExecuteSqlOutputDispositionDetails)
	} else {
		m.ObjectTemplate = nil
	}

	return
}
