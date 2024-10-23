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

// WorkflowGroup Workflow Group Layout.
type WorkflowGroup struct {

	// Name of the group.
	GroupName *string `mandatory:"true" json:"groupName"`

	// Workflow Group  Details.
	Type WorkflowGroupTypeEnum `mandatory:"true" json:"type"`

	// Steps within the Group.
	Steps []WorkflowComponent `mandatory:"true" json:"steps"`
}

func (m WorkflowGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkflowGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkflowGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetWorkflowGroupTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *WorkflowGroup) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		GroupName *string               `json:"groupName"`
		Type      WorkflowGroupTypeEnum `json:"type"`
		Steps     []workflowcomponent   `json:"steps"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.GroupName = model.GroupName

	m.Type = model.Type

	m.Steps = make([]WorkflowComponent, len(model.Steps))
	for i, n := range model.Steps {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Steps[i] = nn.(WorkflowComponent)
		} else {
			m.Steps[i] = nil
		}
	}
	return
}

// WorkflowGroupTypeEnum Enum with underlying type: string
type WorkflowGroupTypeEnum string

// Set of constants representing the allowable values for WorkflowGroupTypeEnum
const (
	WorkflowGroupTypeParallelResourceGroup WorkflowGroupTypeEnum = "PARALLEL_RESOURCE_GROUP"
	WorkflowGroupTypeRollingResourceGroup  WorkflowGroupTypeEnum = "ROLLING_RESOURCE_GROUP"
)

var mappingWorkflowGroupTypeEnum = map[string]WorkflowGroupTypeEnum{
	"PARALLEL_RESOURCE_GROUP": WorkflowGroupTypeParallelResourceGroup,
	"ROLLING_RESOURCE_GROUP":  WorkflowGroupTypeRollingResourceGroup,
}

var mappingWorkflowGroupTypeEnumLowerCase = map[string]WorkflowGroupTypeEnum{
	"parallel_resource_group": WorkflowGroupTypeParallelResourceGroup,
	"rolling_resource_group":  WorkflowGroupTypeRollingResourceGroup,
}

// GetWorkflowGroupTypeEnumValues Enumerates the set of values for WorkflowGroupTypeEnum
func GetWorkflowGroupTypeEnumValues() []WorkflowGroupTypeEnum {
	values := make([]WorkflowGroupTypeEnum, 0)
	for _, v := range mappingWorkflowGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkflowGroupTypeEnumStringValues Enumerates the set of values in String for WorkflowGroupTypeEnum
func GetWorkflowGroupTypeEnumStringValues() []string {
	return []string{
		"PARALLEL_RESOURCE_GROUP",
		"ROLLING_RESOURCE_GROUP",
	}
}

// GetMappingWorkflowGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkflowGroupTypeEnum(val string) (WorkflowGroupTypeEnum, bool) {
	enum, ok := mappingWorkflowGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
