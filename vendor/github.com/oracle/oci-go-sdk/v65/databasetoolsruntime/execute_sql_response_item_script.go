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

// ExecuteSqlResponseItemScript A statement execution response for script type requests.
type ExecuteSqlResponseItemScript struct {

	// The sequence number of the statement. Sequence number of the first statement is 1.
	StatementId *int `mandatory:"false" json:"statementId"`

	// The Statement type.
	StatementType StatementTypeEnum `mandatory:"false" json:"statementType,omitempty"`

	StatementPos *ExecuteSqlResponseItemStatementPos `mandatory:"false" json:"statementPos"`

	// Text of statements executed.
	StatementText *string `mandatory:"false" json:"statementText"`

	Error *ExecuteSqlResponseItemError `mandatory:"false" json:"error"`

	// Output from DBMS_OUTPUT package. Server output must be enabled (e.g., SET SERVEROUTPUT ON).
	DbmsOutput *string `mandatory:"false" json:"dbmsOutput"`

	Properties *ExecuteSqlResponseItemProperties `mandatory:"false" json:"properties"`

	// Responses generated when executing the statements.
	Responses []string `mandatory:"false" json:"responses"`

	// Results generated when executing the statements.
	Results []int `mandatory:"false" json:"results"`

	ResultSetObject ExecuteSqlOutputDispositionDetails `mandatory:"false" json:"resultSetObject"`

	ResultSet *ExecuteSqlResponseItemResultSet `mandatory:"false" json:"resultSet"`

	// Array of objects specifying the bind information.
	Binds []ExecuteSqlBind `mandatory:"false" json:"binds"`
}

func (m ExecuteSqlResponseItemScript) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseItemScript) ValidateEnumValue() (bool, error) {
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
func (m *ExecuteSqlResponseItemScript) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StatementId     *int                                `json:"statementId"`
		StatementType   StatementTypeEnum                   `json:"statementType"`
		StatementPos    *ExecuteSqlResponseItemStatementPos `json:"statementPos"`
		StatementText   *string                             `json:"statementText"`
		Error           *ExecuteSqlResponseItemError        `json:"error"`
		DbmsOutput      *string                             `json:"dbmsOutput"`
		Properties      *ExecuteSqlResponseItemProperties   `json:"properties"`
		Responses       []string                            `json:"responses"`
		Results         []int                               `json:"results"`
		ResultSetObject executesqloutputdispositiondetails  `json:"resultSetObject"`
		ResultSet       *ExecuteSqlResponseItemResultSet    `json:"resultSet"`
		Binds           []ExecuteSqlBind                    `json:"binds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StatementId = model.StatementId

	m.StatementType = model.StatementType

	m.StatementPos = model.StatementPos

	m.StatementText = model.StatementText

	m.Error = model.Error

	m.DbmsOutput = model.DbmsOutput

	m.Properties = model.Properties

	m.Responses = make([]string, len(model.Responses))
	copy(m.Responses, model.Responses)
	m.Results = make([]int, len(model.Results))
	copy(m.Results, model.Results)
	nn, e = model.ResultSetObject.UnmarshalPolymorphicJSON(model.ResultSetObject.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultSetObject = nn.(ExecuteSqlOutputDispositionDetails)
	} else {
		m.ResultSetObject = nil
	}

	m.ResultSet = model.ResultSet

	m.Binds = make([]ExecuteSqlBind, len(model.Binds))
	copy(m.Binds, model.Binds)
	return
}
