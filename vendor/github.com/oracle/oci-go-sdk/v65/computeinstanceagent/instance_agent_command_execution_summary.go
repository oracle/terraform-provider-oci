// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceAgentCommandExecutionSummary Execution details for a command.
type InstanceAgentCommandExecutionSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the command.
	InstanceAgentCommandId *string `mandatory:"true" json:"instanceAgentCommandId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The command delivery state.
	//  * `VISIBLE` - The command is visible to the instance.
	//  * `PENDING` - The command is pending acknowledgment from the instance.
	//  * `ACKED` - The command has been received and acknowledged by the instance.
	//  * `ACKED_CANCELED` - The canceled command has been received and acknowledged by the instance.
	//  * `EXPIRED` - The instance has not requested for commands and the command's delivery has expired.
	DeliveryState InstanceAgentCommandExecutionSummaryDeliveryStateEnum `mandatory:"true" json:"deliveryState"`

	// The command execution lifecycle state.
	// * `ACCEPTED` - The command has been accepted to run.
	// * `IN_PROGRESS` - The command is in progress.
	// * `SUCCEEDED` - The command was successfully executed.
	// * `FAILED` - The command failed to execute.
	// * `TIMED_OUT` - The command execution timed out.
	// * `CANCELED` - The command execution was canceled.
	LifecycleState InstanceAgentCommandExecutionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the command was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the command was last updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A large, non-consecutive number that Oracle Cloud Agent assigns to each created command.
	SequenceNumber *int64 `mandatory:"true" json:"sequenceNumber"`

	// The execution output from a command.
	Content InstanceAgentCommandExecutionOutputContent `mandatory:"true" json:"content"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m InstanceAgentCommandExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceAgentCommandExecutionSummaryDeliveryStateEnum(string(m.DeliveryState)); !ok && m.DeliveryState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeliveryState: %s. Supported values are: %s.", m.DeliveryState, strings.Join(GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstanceAgentCommandExecutionSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string                                                `json:"displayName"`
		InstanceAgentCommandId *string                                                `json:"instanceAgentCommandId"`
		InstanceId             *string                                                `json:"instanceId"`
		DeliveryState          InstanceAgentCommandExecutionSummaryDeliveryStateEnum  `json:"deliveryState"`
		LifecycleState         InstanceAgentCommandExecutionSummaryLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated            *common.SDKTime                                        `json:"timeCreated"`
		TimeUpdated            *common.SDKTime                                        `json:"timeUpdated"`
		SequenceNumber         *int64                                                 `json:"sequenceNumber"`
		Content                instanceagentcommandexecutionoutputcontent             `json:"content"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.InstanceAgentCommandId = model.InstanceAgentCommandId

	m.InstanceId = model.InstanceId

	m.DeliveryState = model.DeliveryState

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.SequenceNumber = model.SequenceNumber

	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(InstanceAgentCommandExecutionOutputContent)
	} else {
		m.Content = nil
	}

	return
}

// InstanceAgentCommandExecutionSummaryDeliveryStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionSummaryDeliveryStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionSummaryDeliveryStateEnum
const (
	InstanceAgentCommandExecutionSummaryDeliveryStateVisible       InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "VISIBLE"
	InstanceAgentCommandExecutionSummaryDeliveryStatePending       InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "PENDING"
	InstanceAgentCommandExecutionSummaryDeliveryStateAcked         InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "ACKED"
	InstanceAgentCommandExecutionSummaryDeliveryStateAckedCanceled InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "ACKED_CANCELED"
	InstanceAgentCommandExecutionSummaryDeliveryStateExpired       InstanceAgentCommandExecutionSummaryDeliveryStateEnum = "EXPIRED"
)

var mappingInstanceAgentCommandExecutionSummaryDeliveryStateEnum = map[string]InstanceAgentCommandExecutionSummaryDeliveryStateEnum{
	"VISIBLE":        InstanceAgentCommandExecutionSummaryDeliveryStateVisible,
	"PENDING":        InstanceAgentCommandExecutionSummaryDeliveryStatePending,
	"ACKED":          InstanceAgentCommandExecutionSummaryDeliveryStateAcked,
	"ACKED_CANCELED": InstanceAgentCommandExecutionSummaryDeliveryStateAckedCanceled,
	"EXPIRED":        InstanceAgentCommandExecutionSummaryDeliveryStateExpired,
}

var mappingInstanceAgentCommandExecutionSummaryDeliveryStateEnumLowerCase = map[string]InstanceAgentCommandExecutionSummaryDeliveryStateEnum{
	"visible":        InstanceAgentCommandExecutionSummaryDeliveryStateVisible,
	"pending":        InstanceAgentCommandExecutionSummaryDeliveryStatePending,
	"acked":          InstanceAgentCommandExecutionSummaryDeliveryStateAcked,
	"acked_canceled": InstanceAgentCommandExecutionSummaryDeliveryStateAckedCanceled,
	"expired":        InstanceAgentCommandExecutionSummaryDeliveryStateExpired,
}

// GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionSummaryDeliveryStateEnum
func GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumValues() []InstanceAgentCommandExecutionSummaryDeliveryStateEnum {
	values := make([]InstanceAgentCommandExecutionSummaryDeliveryStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionSummaryDeliveryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumStringValues Enumerates the set of values in String for InstanceAgentCommandExecutionSummaryDeliveryStateEnum
func GetInstanceAgentCommandExecutionSummaryDeliveryStateEnumStringValues() []string {
	return []string{
		"VISIBLE",
		"PENDING",
		"ACKED",
		"ACKED_CANCELED",
		"EXPIRED",
	}
}

// GetMappingInstanceAgentCommandExecutionSummaryDeliveryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandExecutionSummaryDeliveryStateEnum(val string) (InstanceAgentCommandExecutionSummaryDeliveryStateEnum, bool) {
	enum, ok := mappingInstanceAgentCommandExecutionSummaryDeliveryStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InstanceAgentCommandExecutionSummaryLifecycleStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionSummaryLifecycleStateEnum
const (
	InstanceAgentCommandExecutionSummaryLifecycleStateAccepted   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "ACCEPTED"
	InstanceAgentCommandExecutionSummaryLifecycleStateInProgress InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "IN_PROGRESS"
	InstanceAgentCommandExecutionSummaryLifecycleStateSucceeded  InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "SUCCEEDED"
	InstanceAgentCommandExecutionSummaryLifecycleStateFailed     InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "FAILED"
	InstanceAgentCommandExecutionSummaryLifecycleStateTimedOut   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "TIMED_OUT"
	InstanceAgentCommandExecutionSummaryLifecycleStateCanceled   InstanceAgentCommandExecutionSummaryLifecycleStateEnum = "CANCELED"
)

var mappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum = map[string]InstanceAgentCommandExecutionSummaryLifecycleStateEnum{
	"ACCEPTED":    InstanceAgentCommandExecutionSummaryLifecycleStateAccepted,
	"IN_PROGRESS": InstanceAgentCommandExecutionSummaryLifecycleStateInProgress,
	"SUCCEEDED":   InstanceAgentCommandExecutionSummaryLifecycleStateSucceeded,
	"FAILED":      InstanceAgentCommandExecutionSummaryLifecycleStateFailed,
	"TIMED_OUT":   InstanceAgentCommandExecutionSummaryLifecycleStateTimedOut,
	"CANCELED":    InstanceAgentCommandExecutionSummaryLifecycleStateCanceled,
}

var mappingInstanceAgentCommandExecutionSummaryLifecycleStateEnumLowerCase = map[string]InstanceAgentCommandExecutionSummaryLifecycleStateEnum{
	"accepted":    InstanceAgentCommandExecutionSummaryLifecycleStateAccepted,
	"in_progress": InstanceAgentCommandExecutionSummaryLifecycleStateInProgress,
	"succeeded":   InstanceAgentCommandExecutionSummaryLifecycleStateSucceeded,
	"failed":      InstanceAgentCommandExecutionSummaryLifecycleStateFailed,
	"timed_out":   InstanceAgentCommandExecutionSummaryLifecycleStateTimedOut,
	"canceled":    InstanceAgentCommandExecutionSummaryLifecycleStateCanceled,
}

// GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionSummaryLifecycleStateEnum
func GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumValues() []InstanceAgentCommandExecutionSummaryLifecycleStateEnum {
	values := make([]InstanceAgentCommandExecutionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for InstanceAgentCommandExecutionSummaryLifecycleStateEnum
func GetInstanceAgentCommandExecutionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
		"CANCELED",
	}
}

// GetMappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandExecutionSummaryLifecycleStateEnum(val string) (InstanceAgentCommandExecutionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingInstanceAgentCommandExecutionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
