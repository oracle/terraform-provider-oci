// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataIormConfig The IORM settings of the Exadata DB system.
type ExadataIormConfig struct {

	// The current state of IORM configuration for the Exadata DB system.
	LifecycleState ExadataIormConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional information about the current `lifecycleState`.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current value for the IORM objective.
	// The default is `AUTO`.
	Objective ExadataIormConfigObjectiveEnum `mandatory:"false" json:"objective,omitempty"`

	// An array of IORM settings for all the database in
	// the Exadata DB system.
	DbPlans []DbIormConfig `mandatory:"false" json:"dbPlans"`
}

func (m ExadataIormConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataIormConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataIormConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataIormConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataIormConfigObjectiveEnum(string(m.Objective)); !ok && m.Objective != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Objective: %s. Supported values are: %s.", m.Objective, strings.Join(GetExadataIormConfigObjectiveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataIormConfigLifecycleStateEnum Enum with underlying type: string
type ExadataIormConfigLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataIormConfigLifecycleStateEnum
const (
	ExadataIormConfigLifecycleStateBootstrapping ExadataIormConfigLifecycleStateEnum = "BOOTSTRAPPING"
	ExadataIormConfigLifecycleStateEnabled       ExadataIormConfigLifecycleStateEnum = "ENABLED"
	ExadataIormConfigLifecycleStateDisabled      ExadataIormConfigLifecycleStateEnum = "DISABLED"
	ExadataIormConfigLifecycleStateUpdating      ExadataIormConfigLifecycleStateEnum = "UPDATING"
	ExadataIormConfigLifecycleStateFailed        ExadataIormConfigLifecycleStateEnum = "FAILED"
)

var mappingExadataIormConfigLifecycleStateEnum = map[string]ExadataIormConfigLifecycleStateEnum{
	"BOOTSTRAPPING": ExadataIormConfigLifecycleStateBootstrapping,
	"ENABLED":       ExadataIormConfigLifecycleStateEnabled,
	"DISABLED":      ExadataIormConfigLifecycleStateDisabled,
	"UPDATING":      ExadataIormConfigLifecycleStateUpdating,
	"FAILED":        ExadataIormConfigLifecycleStateFailed,
}

var mappingExadataIormConfigLifecycleStateEnumLowerCase = map[string]ExadataIormConfigLifecycleStateEnum{
	"bootstrapping": ExadataIormConfigLifecycleStateBootstrapping,
	"enabled":       ExadataIormConfigLifecycleStateEnabled,
	"disabled":      ExadataIormConfigLifecycleStateDisabled,
	"updating":      ExadataIormConfigLifecycleStateUpdating,
	"failed":        ExadataIormConfigLifecycleStateFailed,
}

// GetExadataIormConfigLifecycleStateEnumValues Enumerates the set of values for ExadataIormConfigLifecycleStateEnum
func GetExadataIormConfigLifecycleStateEnumValues() []ExadataIormConfigLifecycleStateEnum {
	values := make([]ExadataIormConfigLifecycleStateEnum, 0)
	for _, v := range mappingExadataIormConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataIormConfigLifecycleStateEnumStringValues Enumerates the set of values in String for ExadataIormConfigLifecycleStateEnum
func GetExadataIormConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"BOOTSTRAPPING",
		"ENABLED",
		"DISABLED",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingExadataIormConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataIormConfigLifecycleStateEnum(val string) (ExadataIormConfigLifecycleStateEnum, bool) {
	enum, ok := mappingExadataIormConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadataIormConfigObjectiveEnum Enum with underlying type: string
type ExadataIormConfigObjectiveEnum string

// Set of constants representing the allowable values for ExadataIormConfigObjectiveEnum
const (
	ExadataIormConfigObjectiveLowLatency     ExadataIormConfigObjectiveEnum = "LOW_LATENCY"
	ExadataIormConfigObjectiveHighThroughput ExadataIormConfigObjectiveEnum = "HIGH_THROUGHPUT"
	ExadataIormConfigObjectiveBalanced       ExadataIormConfigObjectiveEnum = "BALANCED"
	ExadataIormConfigObjectiveAuto           ExadataIormConfigObjectiveEnum = "AUTO"
	ExadataIormConfigObjectiveBasic          ExadataIormConfigObjectiveEnum = "BASIC"
)

var mappingExadataIormConfigObjectiveEnum = map[string]ExadataIormConfigObjectiveEnum{
	"LOW_LATENCY":     ExadataIormConfigObjectiveLowLatency,
	"HIGH_THROUGHPUT": ExadataIormConfigObjectiveHighThroughput,
	"BALANCED":        ExadataIormConfigObjectiveBalanced,
	"AUTO":            ExadataIormConfigObjectiveAuto,
	"BASIC":           ExadataIormConfigObjectiveBasic,
}

var mappingExadataIormConfigObjectiveEnumLowerCase = map[string]ExadataIormConfigObjectiveEnum{
	"low_latency":     ExadataIormConfigObjectiveLowLatency,
	"high_throughput": ExadataIormConfigObjectiveHighThroughput,
	"balanced":        ExadataIormConfigObjectiveBalanced,
	"auto":            ExadataIormConfigObjectiveAuto,
	"basic":           ExadataIormConfigObjectiveBasic,
}

// GetExadataIormConfigObjectiveEnumValues Enumerates the set of values for ExadataIormConfigObjectiveEnum
func GetExadataIormConfigObjectiveEnumValues() []ExadataIormConfigObjectiveEnum {
	values := make([]ExadataIormConfigObjectiveEnum, 0)
	for _, v := range mappingExadataIormConfigObjectiveEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataIormConfigObjectiveEnumStringValues Enumerates the set of values in String for ExadataIormConfigObjectiveEnum
func GetExadataIormConfigObjectiveEnumStringValues() []string {
	return []string{
		"LOW_LATENCY",
		"HIGH_THROUGHPUT",
		"BALANCED",
		"AUTO",
		"BASIC",
	}
}

// GetMappingExadataIormConfigObjectiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataIormConfigObjectiveEnum(val string) (ExadataIormConfigObjectiveEnum, bool) {
	enum, ok := mappingExadataIormConfigObjectiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
