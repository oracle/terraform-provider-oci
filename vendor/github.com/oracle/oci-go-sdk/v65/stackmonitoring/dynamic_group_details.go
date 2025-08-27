// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DynamicGroupDetails Dynamic Group object
type DynamicGroupDetails struct {

	// Name of dynamic Group
	Name *string `mandatory:"true" json:"name"`

	// Assignment of dynamic group in context of Stack Monitoring service. It describes the purpose of dynamic groups in Stack Monitoring.
	StackMonitoringAssignment DynamicGroupDetailsStackMonitoringAssignmentEnum `mandatory:"true" json:"stackMonitoringAssignment"`

	// Identity domain name
	Domain *string `mandatory:"false" json:"domain"`
}

func (m DynamicGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDynamicGroupDetailsStackMonitoringAssignmentEnum(string(m.StackMonitoringAssignment)); !ok && m.StackMonitoringAssignment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackMonitoringAssignment: %s. Supported values are: %s.", m.StackMonitoringAssignment, strings.Join(GetDynamicGroupDetailsStackMonitoringAssignmentEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicGroupDetailsStackMonitoringAssignmentEnum Enum with underlying type: string
type DynamicGroupDetailsStackMonitoringAssignmentEnum string

// Set of constants representing the allowable values for DynamicGroupDetailsStackMonitoringAssignmentEnum
const (
	DynamicGroupDetailsStackMonitoringAssignmentManagementAgents   DynamicGroupDetailsStackMonitoringAssignmentEnum = "MANAGEMENT_AGENTS"
	DynamicGroupDetailsStackMonitoringAssignmentMonitoredInstances DynamicGroupDetailsStackMonitoringAssignmentEnum = "MONITORED_INSTANCES"
)

var mappingDynamicGroupDetailsStackMonitoringAssignmentEnum = map[string]DynamicGroupDetailsStackMonitoringAssignmentEnum{
	"MANAGEMENT_AGENTS":   DynamicGroupDetailsStackMonitoringAssignmentManagementAgents,
	"MONITORED_INSTANCES": DynamicGroupDetailsStackMonitoringAssignmentMonitoredInstances,
}

var mappingDynamicGroupDetailsStackMonitoringAssignmentEnumLowerCase = map[string]DynamicGroupDetailsStackMonitoringAssignmentEnum{
	"management_agents":   DynamicGroupDetailsStackMonitoringAssignmentManagementAgents,
	"monitored_instances": DynamicGroupDetailsStackMonitoringAssignmentMonitoredInstances,
}

// GetDynamicGroupDetailsStackMonitoringAssignmentEnumValues Enumerates the set of values for DynamicGroupDetailsStackMonitoringAssignmentEnum
func GetDynamicGroupDetailsStackMonitoringAssignmentEnumValues() []DynamicGroupDetailsStackMonitoringAssignmentEnum {
	values := make([]DynamicGroupDetailsStackMonitoringAssignmentEnum, 0)
	for _, v := range mappingDynamicGroupDetailsStackMonitoringAssignmentEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicGroupDetailsStackMonitoringAssignmentEnumStringValues Enumerates the set of values in String for DynamicGroupDetailsStackMonitoringAssignmentEnum
func GetDynamicGroupDetailsStackMonitoringAssignmentEnumStringValues() []string {
	return []string{
		"MANAGEMENT_AGENTS",
		"MONITORED_INSTANCES",
	}
}

// GetMappingDynamicGroupDetailsStackMonitoringAssignmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicGroupDetailsStackMonitoringAssignmentEnum(val string) (DynamicGroupDetailsStackMonitoringAssignmentEnum, bool) {
	enum, ok := mappingDynamicGroupDetailsStackMonitoringAssignmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
