// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeClusterConfigurationImportStatus The status of a Compute Cluster configuration import operation.
type ComputeClusterConfigurationImportStatus struct {

	// The unique ID for the work request.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// The OCID of the Compute Cluster.
	ComputeClusterId *string `mandatory:"false" json:"computeClusterId"`

	// The provision identifier for library.
	LibraryEntityId *string `mandatory:"false" json:"libraryEntityId"`

	// The lifecycle state of the configuration import.
	LifecycleState ComputeClusterConfigurationImportStatusLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the reason why the resource is in it's current state. Helps bubble up errors in state changes. For example, it can be used to provide actionable information for a resource in the Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// The time the configuration import started.
	StartedAt *int64 `mandatory:"false" json:"startedAt"`

	// The time the configuration import finished.
	FinishedAt *int64 `mandatory:"false" json:"finishedAt"`
}

func (m ComputeClusterConfigurationImportStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterConfigurationImportStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputeClusterConfigurationImportStatusLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComputeClusterConfigurationImportStatusLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeClusterConfigurationImportStatusLifecycleStateEnum Enum with underlying type: string
type ComputeClusterConfigurationImportStatusLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeClusterConfigurationImportStatusLifecycleStateEnum
const (
	ComputeClusterConfigurationImportStatusLifecycleStateAccepted   ComputeClusterConfigurationImportStatusLifecycleStateEnum = "ACCEPTED"
	ComputeClusterConfigurationImportStatusLifecycleStateInProgress ComputeClusterConfigurationImportStatusLifecycleStateEnum = "IN_PROGRESS"
	ComputeClusterConfigurationImportStatusLifecycleStateActive     ComputeClusterConfigurationImportStatusLifecycleStateEnum = "ACTIVE"
	ComputeClusterConfigurationImportStatusLifecycleStateFailed     ComputeClusterConfigurationImportStatusLifecycleStateEnum = "FAILED"
)

var mappingComputeClusterConfigurationImportStatusLifecycleStateEnum = map[string]ComputeClusterConfigurationImportStatusLifecycleStateEnum{
	"ACCEPTED":    ComputeClusterConfigurationImportStatusLifecycleStateAccepted,
	"IN_PROGRESS": ComputeClusterConfigurationImportStatusLifecycleStateInProgress,
	"ACTIVE":      ComputeClusterConfigurationImportStatusLifecycleStateActive,
	"FAILED":      ComputeClusterConfigurationImportStatusLifecycleStateFailed,
}

var mappingComputeClusterConfigurationImportStatusLifecycleStateEnumLowerCase = map[string]ComputeClusterConfigurationImportStatusLifecycleStateEnum{
	"accepted":    ComputeClusterConfigurationImportStatusLifecycleStateAccepted,
	"in_progress": ComputeClusterConfigurationImportStatusLifecycleStateInProgress,
	"active":      ComputeClusterConfigurationImportStatusLifecycleStateActive,
	"failed":      ComputeClusterConfigurationImportStatusLifecycleStateFailed,
}

// GetComputeClusterConfigurationImportStatusLifecycleStateEnumValues Enumerates the set of values for ComputeClusterConfigurationImportStatusLifecycleStateEnum
func GetComputeClusterConfigurationImportStatusLifecycleStateEnumValues() []ComputeClusterConfigurationImportStatusLifecycleStateEnum {
	values := make([]ComputeClusterConfigurationImportStatusLifecycleStateEnum, 0)
	for _, v := range mappingComputeClusterConfigurationImportStatusLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterConfigurationImportStatusLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeClusterConfigurationImportStatusLifecycleStateEnum
func GetComputeClusterConfigurationImportStatusLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"ACTIVE",
		"FAILED",
	}
}

// GetMappingComputeClusterConfigurationImportStatusLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterConfigurationImportStatusLifecycleStateEnum(val string) (ComputeClusterConfigurationImportStatusLifecycleStateEnum, bool) {
	enum, ok := mappingComputeClusterConfigurationImportStatusLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
