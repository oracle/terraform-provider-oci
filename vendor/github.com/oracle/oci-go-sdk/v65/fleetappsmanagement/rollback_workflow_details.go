// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RollbackWorkflowDetails Rollback Workflow details.
type RollbackWorkflowDetails struct {

	// rollback Scope
	Scope RollbackWorkflowDetailsScopeEnum `mandatory:"true" json:"scope"`

	// Rollback Workflow for the runbook.
	Workflow []WorkflowGroup `mandatory:"true" json:"workflow"`
}

func (m RollbackWorkflowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RollbackWorkflowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRollbackWorkflowDetailsScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetRollbackWorkflowDetailsScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RollbackWorkflowDetailsScopeEnum Enum with underlying type: string
type RollbackWorkflowDetailsScopeEnum string

// Set of constants representing the allowable values for RollbackWorkflowDetailsScopeEnum
const (
	RollbackWorkflowDetailsScopeActionGroup RollbackWorkflowDetailsScopeEnum = "ACTION_GROUP"
	RollbackWorkflowDetailsScopeTarget      RollbackWorkflowDetailsScopeEnum = "TARGET"
)

var mappingRollbackWorkflowDetailsScopeEnum = map[string]RollbackWorkflowDetailsScopeEnum{
	"ACTION_GROUP": RollbackWorkflowDetailsScopeActionGroup,
	"TARGET":       RollbackWorkflowDetailsScopeTarget,
}

var mappingRollbackWorkflowDetailsScopeEnumLowerCase = map[string]RollbackWorkflowDetailsScopeEnum{
	"action_group": RollbackWorkflowDetailsScopeActionGroup,
	"target":       RollbackWorkflowDetailsScopeTarget,
}

// GetRollbackWorkflowDetailsScopeEnumValues Enumerates the set of values for RollbackWorkflowDetailsScopeEnum
func GetRollbackWorkflowDetailsScopeEnumValues() []RollbackWorkflowDetailsScopeEnum {
	values := make([]RollbackWorkflowDetailsScopeEnum, 0)
	for _, v := range mappingRollbackWorkflowDetailsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetRollbackWorkflowDetailsScopeEnumStringValues Enumerates the set of values in String for RollbackWorkflowDetailsScopeEnum
func GetRollbackWorkflowDetailsScopeEnumStringValues() []string {
	return []string{
		"ACTION_GROUP",
		"TARGET",
	}
}

// GetMappingRollbackWorkflowDetailsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRollbackWorkflowDetailsScopeEnum(val string) (RollbackWorkflowDetailsScopeEnum, bool) {
	enum, ok := mappingRollbackWorkflowDetailsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
