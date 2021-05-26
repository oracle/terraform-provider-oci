// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
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

var mappingExadataIormConfigLifecycleState = map[string]ExadataIormConfigLifecycleStateEnum{
	"BOOTSTRAPPING": ExadataIormConfigLifecycleStateBootstrapping,
	"ENABLED":       ExadataIormConfigLifecycleStateEnabled,
	"DISABLED":      ExadataIormConfigLifecycleStateDisabled,
	"UPDATING":      ExadataIormConfigLifecycleStateUpdating,
	"FAILED":        ExadataIormConfigLifecycleStateFailed,
}

// GetExadataIormConfigLifecycleStateEnumValues Enumerates the set of values for ExadataIormConfigLifecycleStateEnum
func GetExadataIormConfigLifecycleStateEnumValues() []ExadataIormConfigLifecycleStateEnum {
	values := make([]ExadataIormConfigLifecycleStateEnum, 0)
	for _, v := range mappingExadataIormConfigLifecycleState {
		values = append(values, v)
	}
	return values
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

var mappingExadataIormConfigObjective = map[string]ExadataIormConfigObjectiveEnum{
	"LOW_LATENCY":     ExadataIormConfigObjectiveLowLatency,
	"HIGH_THROUGHPUT": ExadataIormConfigObjectiveHighThroughput,
	"BALANCED":        ExadataIormConfigObjectiveBalanced,
	"AUTO":            ExadataIormConfigObjectiveAuto,
	"BASIC":           ExadataIormConfigObjectiveBasic,
}

// GetExadataIormConfigObjectiveEnumValues Enumerates the set of values for ExadataIormConfigObjectiveEnum
func GetExadataIormConfigObjectiveEnumValues() []ExadataIormConfigObjectiveEnum {
	values := make([]ExadataIormConfigObjectiveEnum, 0)
	for _, v := range mappingExadataIormConfigObjective {
		values = append(values, v)
	}
	return values
}
