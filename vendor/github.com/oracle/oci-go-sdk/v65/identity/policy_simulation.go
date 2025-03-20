// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PolicySimulation Resource object containing details for a given Policy Simulation task.
type PolicySimulation struct {

	// Id of the Policy Simulation task.
	Id *string `mandatory:"true" json:"id"`

	// compartmentId for which the Policy Simulation task is being run.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Status of the Policy Simulation task.
	LifecycleState PolicySimulationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Start of the time window for which the Policy Simulation task should process the logs. Must be less than endTime.
	StartTime *common.SDKTime `mandatory:"true" json:"startTime"`

	// End of the time window for which the Policy Simulation task should process the logs.Must be less than now.
	EndTime *common.SDKTime `mandatory:"true" json:"endTime"`

	// Time when the Policy Simulation task was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when the Policy Simulation task was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	UpdatedPoliciesPath *ObjectStorageFileDownloadLocation `mandatory:"true" json:"updatedPoliciesPath"`

	OutputFilesPath *ObjectStorageFileUploadLocation `mandatory:"true" json:"outputFilesPath"`

	// Message detailing failure during Policy Simulation task.
	FailureMessage *string `mandatory:"false" json:"failureMessage"`
}

func (m PolicySimulation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PolicySimulation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPolicySimulationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPolicySimulationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PolicySimulationLifecycleStateEnum Enum with underlying type: string
type PolicySimulationLifecycleStateEnum string

// Set of constants representing the allowable values for PolicySimulationLifecycleStateEnum
const (
	PolicySimulationLifecycleStateAccepted   PolicySimulationLifecycleStateEnum = "ACCEPTED"
	PolicySimulationLifecycleStateWaiting    PolicySimulationLifecycleStateEnum = "WAITING"
	PolicySimulationLifecycleStateInProgress PolicySimulationLifecycleStateEnum = "IN_PROGRESS"
	PolicySimulationLifecycleStateFailed     PolicySimulationLifecycleStateEnum = "FAILED"
	PolicySimulationLifecycleStateSucceeded  PolicySimulationLifecycleStateEnum = "SUCCEEDED"
	PolicySimulationLifecycleStateCanceling  PolicySimulationLifecycleStateEnum = "CANCELING"
)

var mappingPolicySimulationLifecycleStateEnum = map[string]PolicySimulationLifecycleStateEnum{
	"ACCEPTED":    PolicySimulationLifecycleStateAccepted,
	"WAITING":     PolicySimulationLifecycleStateWaiting,
	"IN_PROGRESS": PolicySimulationLifecycleStateInProgress,
	"FAILED":      PolicySimulationLifecycleStateFailed,
	"SUCCEEDED":   PolicySimulationLifecycleStateSucceeded,
	"CANCELING":   PolicySimulationLifecycleStateCanceling,
}

var mappingPolicySimulationLifecycleStateEnumLowerCase = map[string]PolicySimulationLifecycleStateEnum{
	"accepted":    PolicySimulationLifecycleStateAccepted,
	"waiting":     PolicySimulationLifecycleStateWaiting,
	"in_progress": PolicySimulationLifecycleStateInProgress,
	"failed":      PolicySimulationLifecycleStateFailed,
	"succeeded":   PolicySimulationLifecycleStateSucceeded,
	"canceling":   PolicySimulationLifecycleStateCanceling,
}

// GetPolicySimulationLifecycleStateEnumValues Enumerates the set of values for PolicySimulationLifecycleStateEnum
func GetPolicySimulationLifecycleStateEnumValues() []PolicySimulationLifecycleStateEnum {
	values := make([]PolicySimulationLifecycleStateEnum, 0)
	for _, v := range mappingPolicySimulationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicySimulationLifecycleStateEnumStringValues Enumerates the set of values in String for PolicySimulationLifecycleStateEnum
func GetPolicySimulationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
	}
}

// GetMappingPolicySimulationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicySimulationLifecycleStateEnum(val string) (PolicySimulationLifecycleStateEnum, bool) {
	enum, ok := mappingPolicySimulationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
