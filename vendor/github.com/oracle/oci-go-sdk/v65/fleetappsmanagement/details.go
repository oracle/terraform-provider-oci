// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Details The details of the task.
type Details struct {

	// The OS for the task
	OsType OsTypeEnum `mandatory:"true" json:"osType"`

	// The scope of the task
	Scope TaskScopeEnum `mandatory:"true" json:"scope"`

	ExecutionDetails ExecutionDetails `mandatory:"false" json:"executionDetails"`

	// The platform of the runbook.
	Platform *string `mandatory:"false" json:"platform"`

	Properties *Properties `mandatory:"false" json:"properties"`
}

func (m Details) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Details) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetOsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetTaskScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Details) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ExecutionDetails executiondetails `json:"executionDetails"`
		Platform         *string          `json:"platform"`
		Properties       *Properties      `json:"properties"`
		OsType           OsTypeEnum       `json:"osType"`
		Scope            TaskScopeEnum    `json:"scope"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ExecutionDetails.UnmarshalPolymorphicJSON(model.ExecutionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExecutionDetails = nn.(ExecutionDetails)
	} else {
		m.ExecutionDetails = nil
	}

	m.Platform = model.Platform

	m.Properties = model.Properties

	m.OsType = model.OsType

	m.Scope = model.Scope

	return
}
