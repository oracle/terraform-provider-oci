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

// AssociatedLocalTaskDetails The details of the local task.
// The local tasks are specific to a single runbook.
type AssociatedLocalTaskDetails struct {
	ExecutionDetails ExecutionDetails `mandatory:"true" json:"executionDetails"`

	// The description of the task.
	Description *string `mandatory:"false" json:"description"`

	// The platform of the runbook.
	Platform *string `mandatory:"false" json:"platform"`

	// Make a copy of this task in Library
	IsCopyToLibraryEnabled *bool `mandatory:"false" json:"isCopyToLibraryEnabled"`

	Properties *Properties `mandatory:"false" json:"properties"`

	// Is this a discovery output task?
	IsDiscoveryOutputTask *bool `mandatory:"false" json:"isDiscoveryOutputTask"`

	// Is this an Apply Subject Task? Ex. Patch Execution Task
	IsApplySubjectTask *bool `mandatory:"false" json:"isApplySubjectTask"`

	// The name of the task
	Name *string `mandatory:"false" json:"name"`

	// The OS for the task.
	OsType OsTypeEnum `mandatory:"true" json:"osType"`
}

func (m AssociatedLocalTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedLocalTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetOsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AssociatedLocalTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAssociatedLocalTaskDetails AssociatedLocalTaskDetails
	s := struct {
		DiscriminatorParam string `json:"scope"`
		MarshalTypeAssociatedLocalTaskDetails
	}{
		"LOCAL",
		(MarshalTypeAssociatedLocalTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AssociatedLocalTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string          `json:"description"`
		Platform               *string          `json:"platform"`
		IsCopyToLibraryEnabled *bool            `json:"isCopyToLibraryEnabled"`
		Properties             *Properties      `json:"properties"`
		IsDiscoveryOutputTask  *bool            `json:"isDiscoveryOutputTask"`
		IsApplySubjectTask     *bool            `json:"isApplySubjectTask"`
		Name                   *string          `json:"name"`
		ExecutionDetails       executiondetails `json:"executionDetails"`
		OsType                 OsTypeEnum       `json:"osType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Platform = model.Platform

	m.IsCopyToLibraryEnabled = model.IsCopyToLibraryEnabled

	m.Properties = model.Properties

	m.IsDiscoveryOutputTask = model.IsDiscoveryOutputTask

	m.IsApplySubjectTask = model.IsApplySubjectTask

	m.Name = model.Name

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

	return
}
