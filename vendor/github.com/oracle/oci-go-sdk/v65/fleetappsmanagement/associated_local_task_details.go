// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

	// Task type.
	Type AssociatedLocalTaskDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The OS for the task.
	OsType OsTypeEnum `mandatory:"false" json:"osType,omitempty"`
}

func (m AssociatedLocalTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedLocalTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssociatedLocalTaskDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAssociatedLocalTaskDetailsTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOsTypeEnum(string(m.OsType)); !ok && m.OsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsType: %s. Supported values are: %s.", m.OsType, strings.Join(GetOsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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
		Description            *string                            `json:"description"`
		Type                   AssociatedLocalTaskDetailsTypeEnum `json:"type"`
		Platform               *string                            `json:"platform"`
		IsCopyToLibraryEnabled *bool                              `json:"isCopyToLibraryEnabled"`
		OsType                 OsTypeEnum                         `json:"osType"`
		Properties             *Properties                        `json:"properties"`
		IsDiscoveryOutputTask  *bool                              `json:"isDiscoveryOutputTask"`
		IsApplySubjectTask     *bool                              `json:"isApplySubjectTask"`
		Name                   *string                            `json:"name"`
		ExecutionDetails       executiondetails                   `json:"executionDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Type = model.Type

	m.Platform = model.Platform

	m.IsCopyToLibraryEnabled = model.IsCopyToLibraryEnabled

	m.OsType = model.OsType

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

	return
}

// AssociatedLocalTaskDetailsTypeEnum Enum with underlying type: string
type AssociatedLocalTaskDetailsTypeEnum string

// Set of constants representing the allowable values for AssociatedLocalTaskDetailsTypeEnum
const (
	AssociatedLocalTaskDetailsTypeUserDefined   AssociatedLocalTaskDetailsTypeEnum = "USER_DEFINED"
	AssociatedLocalTaskDetailsTypeOracleDefined AssociatedLocalTaskDetailsTypeEnum = "ORACLE_DEFINED"
)

var mappingAssociatedLocalTaskDetailsTypeEnum = map[string]AssociatedLocalTaskDetailsTypeEnum{
	"USER_DEFINED":   AssociatedLocalTaskDetailsTypeUserDefined,
	"ORACLE_DEFINED": AssociatedLocalTaskDetailsTypeOracleDefined,
}

var mappingAssociatedLocalTaskDetailsTypeEnumLowerCase = map[string]AssociatedLocalTaskDetailsTypeEnum{
	"user_defined":   AssociatedLocalTaskDetailsTypeUserDefined,
	"oracle_defined": AssociatedLocalTaskDetailsTypeOracleDefined,
}

// GetAssociatedLocalTaskDetailsTypeEnumValues Enumerates the set of values for AssociatedLocalTaskDetailsTypeEnum
func GetAssociatedLocalTaskDetailsTypeEnumValues() []AssociatedLocalTaskDetailsTypeEnum {
	values := make([]AssociatedLocalTaskDetailsTypeEnum, 0)
	for _, v := range mappingAssociatedLocalTaskDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedLocalTaskDetailsTypeEnumStringValues Enumerates the set of values in String for AssociatedLocalTaskDetailsTypeEnum
func GetAssociatedLocalTaskDetailsTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"ORACLE_DEFINED",
	}
}

// GetMappingAssociatedLocalTaskDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedLocalTaskDetailsTypeEnum(val string) (AssociatedLocalTaskDetailsTypeEnum, bool) {
	enum, ok := mappingAssociatedLocalTaskDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
