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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Group The group of the runbook
type Group struct {

	// The type of the group
	Type GroupTypeEnum `mandatory:"true" json:"type"`

	// The name of the group
	Name *string `mandatory:"true" json:"name"`

	Properties *ComponentProperties `mandatory:"false" json:"properties"`
}

func (m Group) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Group) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetGroupTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupTypeEnum Enum with underlying type: string
type GroupTypeEnum string

// Set of constants representing the allowable values for GroupTypeEnum
const (
	GroupTypeParallelTaskGroup     GroupTypeEnum = "PARALLEL_TASK_GROUP"
	GroupTypeParallelResourceGroup GroupTypeEnum = "PARALLEL_RESOURCE_GROUP"
	GroupTypeRollingResourceGroup  GroupTypeEnum = "ROLLING_RESOURCE_GROUP"
)

var mappingGroupTypeEnum = map[string]GroupTypeEnum{
	"PARALLEL_TASK_GROUP":     GroupTypeParallelTaskGroup,
	"PARALLEL_RESOURCE_GROUP": GroupTypeParallelResourceGroup,
	"ROLLING_RESOURCE_GROUP":  GroupTypeRollingResourceGroup,
}

var mappingGroupTypeEnumLowerCase = map[string]GroupTypeEnum{
	"parallel_task_group":     GroupTypeParallelTaskGroup,
	"parallel_resource_group": GroupTypeParallelResourceGroup,
	"rolling_resource_group":  GroupTypeRollingResourceGroup,
}

// GetGroupTypeEnumValues Enumerates the set of values for GroupTypeEnum
func GetGroupTypeEnumValues() []GroupTypeEnum {
	values := make([]GroupTypeEnum, 0)
	for _, v := range mappingGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupTypeEnumStringValues Enumerates the set of values in String for GroupTypeEnum
func GetGroupTypeEnumStringValues() []string {
	return []string{
		"PARALLEL_TASK_GROUP",
		"PARALLEL_RESOURCE_GROUP",
		"ROLLING_RESOURCE_GROUP",
	}
}

// GetMappingGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupTypeEnum(val string) (GroupTypeEnum, bool) {
	enum, ok := mappingGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
