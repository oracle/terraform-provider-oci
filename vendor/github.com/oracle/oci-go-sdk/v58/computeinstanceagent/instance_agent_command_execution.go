// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InstanceAgentCommandExecution A command's execution summary.
type InstanceAgentCommandExecution struct {

	// The OCID of the command
	InstanceAgentCommandId *string `mandatory:"true" json:"instanceAgentCommandId"`

	// The OCID of the instance
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Specifies the command delivery state.
	//  * `VISIBLE` - The command is visible to instance.
	//  * `PENDING` - The command is pending ack from the instance.
	//  * `ACKED` - The command has been received and acked by the instance.
	//  * `ACKED_CANCELED` - The canceled command has been received and acked by the instance.
	//  * `EXPIRED` - The instance has not requested for commands and its delivery has expired.
	DeliveryState InstanceAgentCommandExecutionDeliveryStateEnum `mandatory:"true" json:"deliveryState"`

	// command execution life cycle state.
	// * `ACCEPTED` - The command execution has been accepted to run.
	// * `IN_PROGRESS` - The command execution is in progress.
	// * `SUCCEEDED` - The command execution is successful.
	// * `FAILED` - The command execution has failed.
	// * `TIMED_OUT` - The command execution has timedout.
	// * `CANCELED` - The command execution has canceled.
	LifecycleState InstanceAgentCommandExecutionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The command creation date
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The command last updated at date.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The large non-consecutive number that Run Command Service assigns to each created command.
	SequenceNumber *int64 `mandatory:"true" json:"sequenceNumber"`

	Content InstanceAgentCommandExecutionOutputContent `mandatory:"true" json:"content"`

	// The user friendly display name of the command.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m InstanceAgentCommandExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceAgentCommandExecutionDeliveryStateEnum(string(m.DeliveryState)); !ok && m.DeliveryState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeliveryState: %s. Supported values are: %s.", m.DeliveryState, strings.Join(GetInstanceAgentCommandExecutionDeliveryStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstanceAgentCommandExecutionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInstanceAgentCommandExecutionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstanceAgentCommandExecution) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string                                         `json:"displayName"`
		InstanceAgentCommandId *string                                         `json:"instanceAgentCommandId"`
		InstanceId             *string                                         `json:"instanceId"`
		DeliveryState          InstanceAgentCommandExecutionDeliveryStateEnum  `json:"deliveryState"`
		LifecycleState         InstanceAgentCommandExecutionLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated            *common.SDKTime                                 `json:"timeCreated"`
		TimeUpdated            *common.SDKTime                                 `json:"timeUpdated"`
		SequenceNumber         *int64                                          `json:"sequenceNumber"`
		Content                instanceagentcommandexecutionoutputcontent      `json:"content"`
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

// InstanceAgentCommandExecutionDeliveryStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionDeliveryStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionDeliveryStateEnum
const (
	InstanceAgentCommandExecutionDeliveryStateVisible       InstanceAgentCommandExecutionDeliveryStateEnum = "VISIBLE"
	InstanceAgentCommandExecutionDeliveryStatePending       InstanceAgentCommandExecutionDeliveryStateEnum = "PENDING"
	InstanceAgentCommandExecutionDeliveryStateAcked         InstanceAgentCommandExecutionDeliveryStateEnum = "ACKED"
	InstanceAgentCommandExecutionDeliveryStateAckedCanceled InstanceAgentCommandExecutionDeliveryStateEnum = "ACKED_CANCELED"
	InstanceAgentCommandExecutionDeliveryStateExpired       InstanceAgentCommandExecutionDeliveryStateEnum = "EXPIRED"
)

var mappingInstanceAgentCommandExecutionDeliveryStateEnum = map[string]InstanceAgentCommandExecutionDeliveryStateEnum{
	"VISIBLE":        InstanceAgentCommandExecutionDeliveryStateVisible,
	"PENDING":        InstanceAgentCommandExecutionDeliveryStatePending,
	"ACKED":          InstanceAgentCommandExecutionDeliveryStateAcked,
	"ACKED_CANCELED": InstanceAgentCommandExecutionDeliveryStateAckedCanceled,
	"EXPIRED":        InstanceAgentCommandExecutionDeliveryStateExpired,
}

// GetInstanceAgentCommandExecutionDeliveryStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionDeliveryStateEnum
func GetInstanceAgentCommandExecutionDeliveryStateEnumValues() []InstanceAgentCommandExecutionDeliveryStateEnum {
	values := make([]InstanceAgentCommandExecutionDeliveryStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionDeliveryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandExecutionDeliveryStateEnumStringValues Enumerates the set of values in String for InstanceAgentCommandExecutionDeliveryStateEnum
func GetInstanceAgentCommandExecutionDeliveryStateEnumStringValues() []string {
	return []string{
		"VISIBLE",
		"PENDING",
		"ACKED",
		"ACKED_CANCELED",
		"EXPIRED",
	}
}

// GetMappingInstanceAgentCommandExecutionDeliveryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandExecutionDeliveryStateEnum(val string) (InstanceAgentCommandExecutionDeliveryStateEnum, bool) {
	mappingInstanceAgentCommandExecutionDeliveryStateEnumIgnoreCase := make(map[string]InstanceAgentCommandExecutionDeliveryStateEnum)
	for k, v := range mappingInstanceAgentCommandExecutionDeliveryStateEnum {
		mappingInstanceAgentCommandExecutionDeliveryStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingInstanceAgentCommandExecutionDeliveryStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// InstanceAgentCommandExecutionLifecycleStateEnum Enum with underlying type: string
type InstanceAgentCommandExecutionLifecycleStateEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionLifecycleStateEnum
const (
	InstanceAgentCommandExecutionLifecycleStateAccepted   InstanceAgentCommandExecutionLifecycleStateEnum = "ACCEPTED"
	InstanceAgentCommandExecutionLifecycleStateInProgress InstanceAgentCommandExecutionLifecycleStateEnum = "IN_PROGRESS"
	InstanceAgentCommandExecutionLifecycleStateSucceeded  InstanceAgentCommandExecutionLifecycleStateEnum = "SUCCEEDED"
	InstanceAgentCommandExecutionLifecycleStateFailed     InstanceAgentCommandExecutionLifecycleStateEnum = "FAILED"
	InstanceAgentCommandExecutionLifecycleStateTimedOut   InstanceAgentCommandExecutionLifecycleStateEnum = "TIMED_OUT"
	InstanceAgentCommandExecutionLifecycleStateCanceled   InstanceAgentCommandExecutionLifecycleStateEnum = "CANCELED"
)

var mappingInstanceAgentCommandExecutionLifecycleStateEnum = map[string]InstanceAgentCommandExecutionLifecycleStateEnum{
	"ACCEPTED":    InstanceAgentCommandExecutionLifecycleStateAccepted,
	"IN_PROGRESS": InstanceAgentCommandExecutionLifecycleStateInProgress,
	"SUCCEEDED":   InstanceAgentCommandExecutionLifecycleStateSucceeded,
	"FAILED":      InstanceAgentCommandExecutionLifecycleStateFailed,
	"TIMED_OUT":   InstanceAgentCommandExecutionLifecycleStateTimedOut,
	"CANCELED":    InstanceAgentCommandExecutionLifecycleStateCanceled,
}

// GetInstanceAgentCommandExecutionLifecycleStateEnumValues Enumerates the set of values for InstanceAgentCommandExecutionLifecycleStateEnum
func GetInstanceAgentCommandExecutionLifecycleStateEnumValues() []InstanceAgentCommandExecutionLifecycleStateEnum {
	values := make([]InstanceAgentCommandExecutionLifecycleStateEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandExecutionLifecycleStateEnumStringValues Enumerates the set of values in String for InstanceAgentCommandExecutionLifecycleStateEnum
func GetInstanceAgentCommandExecutionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
		"CANCELED",
	}
}

// GetMappingInstanceAgentCommandExecutionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandExecutionLifecycleStateEnum(val string) (InstanceAgentCommandExecutionLifecycleStateEnum, bool) {
	mappingInstanceAgentCommandExecutionLifecycleStateEnumIgnoreCase := make(map[string]InstanceAgentCommandExecutionLifecycleStateEnum)
	for k, v := range mappingInstanceAgentCommandExecutionLifecycleStateEnum {
		mappingInstanceAgentCommandExecutionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingInstanceAgentCommandExecutionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
