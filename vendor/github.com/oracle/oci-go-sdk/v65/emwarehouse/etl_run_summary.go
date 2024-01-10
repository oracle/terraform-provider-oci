// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EM Warehouse API
//
// Use the EM Warehouse API to manage EM Warehouse data collection.
//

package emwarehouse

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EtlRunSummary Contains summary of a run.
type EtlRunSummary struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Data read by the dataflow run
	DataReadInBytes *int64 `mandatory:"false" json:"dataReadInBytes"`

	// Data written by the dataflow run
	DataWritten *int64 `mandatory:"false" json:"dataWritten"`

	// The current state of the etlRun.
	LifecycleState EtlRunSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The name of the ETLRun.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Details of the lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Dataflow run duration
	RunDurationInMilliseconds *int64 `mandatory:"false" json:"runDurationInMilliseconds"`

	// Time when the dataflow run was created
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the dataflow run was updated
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m EtlRunSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EtlRunSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEtlRunSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEtlRunSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EtlRunSummaryLifecycleStateEnum Enum with underlying type: string
type EtlRunSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for EtlRunSummaryLifecycleStateEnum
const (
	EtlRunSummaryLifecycleStateAccepted   EtlRunSummaryLifecycleStateEnum = "ACCEPTED"
	EtlRunSummaryLifecycleStateInProgress EtlRunSummaryLifecycleStateEnum = "IN_PROGRESS"
	EtlRunSummaryLifecycleStateCanceling  EtlRunSummaryLifecycleStateEnum = "CANCELING"
	EtlRunSummaryLifecycleStateCanceled   EtlRunSummaryLifecycleStateEnum = "CANCELED"
	EtlRunSummaryLifecycleStateFailed     EtlRunSummaryLifecycleStateEnum = "FAILED"
	EtlRunSummaryLifecycleStateSucceeded  EtlRunSummaryLifecycleStateEnum = "SUCCEEDED"
)

var mappingEtlRunSummaryLifecycleStateEnum = map[string]EtlRunSummaryLifecycleStateEnum{
	"ACCEPTED":    EtlRunSummaryLifecycleStateAccepted,
	"IN_PROGRESS": EtlRunSummaryLifecycleStateInProgress,
	"CANCELING":   EtlRunSummaryLifecycleStateCanceling,
	"CANCELED":    EtlRunSummaryLifecycleStateCanceled,
	"FAILED":      EtlRunSummaryLifecycleStateFailed,
	"SUCCEEDED":   EtlRunSummaryLifecycleStateSucceeded,
}

var mappingEtlRunSummaryLifecycleStateEnumLowerCase = map[string]EtlRunSummaryLifecycleStateEnum{
	"accepted":    EtlRunSummaryLifecycleStateAccepted,
	"in_progress": EtlRunSummaryLifecycleStateInProgress,
	"canceling":   EtlRunSummaryLifecycleStateCanceling,
	"canceled":    EtlRunSummaryLifecycleStateCanceled,
	"failed":      EtlRunSummaryLifecycleStateFailed,
	"succeeded":   EtlRunSummaryLifecycleStateSucceeded,
}

// GetEtlRunSummaryLifecycleStateEnumValues Enumerates the set of values for EtlRunSummaryLifecycleStateEnum
func GetEtlRunSummaryLifecycleStateEnumValues() []EtlRunSummaryLifecycleStateEnum {
	values := make([]EtlRunSummaryLifecycleStateEnum, 0)
	for _, v := range mappingEtlRunSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEtlRunSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for EtlRunSummaryLifecycleStateEnum
func GetEtlRunSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingEtlRunSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEtlRunSummaryLifecycleStateEnum(val string) (EtlRunSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingEtlRunSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
