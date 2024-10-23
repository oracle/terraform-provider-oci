// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Details The details of the task.
type Details struct {
	ExecutionDetails ExecutionDetails `mandatory:"true" json:"executionDetails"`

	// The OS for the task
	OsType OsTypeEnum `mandatory:"true" json:"osType"`

	// The scope of the task
	Scope TaskScopeEnum `mandatory:"true" json:"scope"`

	// The platform of the runbook.
	Platform *string `mandatory:"false" json:"platform"`

	Properties *Properties `mandatory:"false" json:"properties"`

	// Is this a discovery output task?
	IsDiscoveryOutputTask *bool `mandatory:"false" json:"isDiscoveryOutputTask"`

	// Is this an Apply Subject Task?
	// Set this to true for a Patch Execution Task which applies patches(subjects) on a target.
	IsApplySubjectTask *bool `mandatory:"false" json:"isApplySubjectTask"`
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
		Platform              *string          `json:"platform"`
		Properties            *Properties      `json:"properties"`
		IsDiscoveryOutputTask *bool            `json:"isDiscoveryOutputTask"`
		IsApplySubjectTask    *bool            `json:"isApplySubjectTask"`
		ExecutionDetails      executiondetails `json:"executionDetails"`
		OsType                OsTypeEnum       `json:"osType"`
		Scope                 TaskScopeEnum    `json:"scope"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Platform = model.Platform

	m.Properties = model.Properties

	m.IsDiscoveryOutputTask = model.IsDiscoveryOutputTask

	m.IsApplySubjectTask = model.IsApplySubjectTask

	nn, e = model.ExecutionDetails.UnmarshalPolymorphicJSON(model.ExecutionDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ExecutionDetails = nn.(ExecutionDetails)
	} else {
		m.ExecutionDetails = nil
	}

	m.OsType = model.OsType

	m.Scope = model.Scope

	return
}
