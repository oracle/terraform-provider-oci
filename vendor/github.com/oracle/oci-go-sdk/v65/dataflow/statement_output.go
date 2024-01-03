// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StatementOutput The execution output of a statement.
type StatementOutput struct {
	Data StatementOutputData `mandatory:"false" json:"data"`

	// Status of the statement output.
	Status StatementOutputStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The name of the error in the statement output.
	ErrorName *string `mandatory:"false" json:"errorName"`

	// The value of the error in the statement output.
	ErrorValue *string `mandatory:"false" json:"errorValue"`

	// The traceback of the statement output.
	Traceback []string `mandatory:"false" json:"traceback"`
}

func (m StatementOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StatementOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStatementOutputStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetStatementOutputStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StatementOutput) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Data       statementoutputdata       `json:"data"`
		Status     StatementOutputStatusEnum `json:"status"`
		ErrorName  *string                   `json:"errorName"`
		ErrorValue *string                   `json:"errorValue"`
		Traceback  []string                  `json:"traceback"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Data.UnmarshalPolymorphicJSON(model.Data.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Data = nn.(StatementOutputData)
	} else {
		m.Data = nil
	}

	m.Status = model.Status

	m.ErrorName = model.ErrorName

	m.ErrorValue = model.ErrorValue

	m.Traceback = make([]string, len(model.Traceback))
	copy(m.Traceback, model.Traceback)
	return
}

// StatementOutputStatusEnum Enum with underlying type: string
type StatementOutputStatusEnum string

// Set of constants representing the allowable values for StatementOutputStatusEnum
const (
	StatementOutputStatusOk    StatementOutputStatusEnum = "OK"
	StatementOutputStatusError StatementOutputStatusEnum = "ERROR"
)

var mappingStatementOutputStatusEnum = map[string]StatementOutputStatusEnum{
	"OK":    StatementOutputStatusOk,
	"ERROR": StatementOutputStatusError,
}

var mappingStatementOutputStatusEnumLowerCase = map[string]StatementOutputStatusEnum{
	"ok":    StatementOutputStatusOk,
	"error": StatementOutputStatusError,
}

// GetStatementOutputStatusEnumValues Enumerates the set of values for StatementOutputStatusEnum
func GetStatementOutputStatusEnumValues() []StatementOutputStatusEnum {
	values := make([]StatementOutputStatusEnum, 0)
	for _, v := range mappingStatementOutputStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStatementOutputStatusEnumStringValues Enumerates the set of values in String for StatementOutputStatusEnum
func GetStatementOutputStatusEnumStringValues() []string {
	return []string{
		"OK",
		"ERROR",
	}
}

// GetMappingStatementOutputStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatementOutputStatusEnum(val string) (StatementOutputStatusEnum, bool) {
	enum, ok := mappingStatementOutputStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
