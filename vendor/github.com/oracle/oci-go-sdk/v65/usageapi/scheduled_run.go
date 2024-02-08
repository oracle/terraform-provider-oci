// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledRun The saved schedule run.
type ScheduledRun struct {

	// The ocid representing unique shedule run
	Id *string `mandatory:"true" json:"id"`

	// The ocid representing unique shedule
	ScheduleId *string `mandatory:"true" json:"scheduleId"`

	// The time when schedule started executing
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when schedule finished executing
	TimeFinished *common.SDKTime `mandatory:"true" json:"timeFinished"`

	// Specifies if the schedule job was run successfully or not.
	LifecycleState ScheduledRunLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional details about scheduled run failure
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`
}

func (m ScheduledRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledRunLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledRunLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledRunLifecycleStateEnum Enum with underlying type: string
type ScheduledRunLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledRunLifecycleStateEnum
const (
	ScheduledRunLifecycleStateFailed    ScheduledRunLifecycleStateEnum = "FAILED"
	ScheduledRunLifecycleStateSucceeded ScheduledRunLifecycleStateEnum = "SUCCEEDED"
)

var mappingScheduledRunLifecycleStateEnum = map[string]ScheduledRunLifecycleStateEnum{
	"FAILED":    ScheduledRunLifecycleStateFailed,
	"SUCCEEDED": ScheduledRunLifecycleStateSucceeded,
}

var mappingScheduledRunLifecycleStateEnumLowerCase = map[string]ScheduledRunLifecycleStateEnum{
	"failed":    ScheduledRunLifecycleStateFailed,
	"succeeded": ScheduledRunLifecycleStateSucceeded,
}

// GetScheduledRunLifecycleStateEnumValues Enumerates the set of values for ScheduledRunLifecycleStateEnum
func GetScheduledRunLifecycleStateEnumValues() []ScheduledRunLifecycleStateEnum {
	values := make([]ScheduledRunLifecycleStateEnum, 0)
	for _, v := range mappingScheduledRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledRunLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledRunLifecycleStateEnum
func GetScheduledRunLifecycleStateEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingScheduledRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledRunLifecycleStateEnum(val string) (ScheduledRunLifecycleStateEnum, bool) {
	enum, ok := mappingScheduledRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
