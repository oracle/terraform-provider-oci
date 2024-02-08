// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Rule A generic rule object - represents an ingest time rule or a scheduled task.
type Rule struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The kind of rule - either an ingest time rule or a scheduled task.
	Kind RuleKindEnum `mandatory:"true" json:"kind"`

	// The ingest time rule or scheduled task display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current state of the logging analytics rule.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A flag indicating whether or not the ingest time rule or scheduled task is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The most recent task execution status.
	LastExecutionStatus RuleLastExecutionStatusEnum `mandatory:"false" json:"lastExecutionStatus,omitempty"`

	// The date and time the scheduled task last executed, in the format defined by RFC3339.
	TimeLastExecuted *common.SDKTime `mandatory:"false" json:"timeLastExecuted"`
}

func (m Rule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Rule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRuleKindEnum(string(m.Kind)); !ok && m.Kind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kind: %s. Supported values are: %s.", m.Kind, strings.Join(GetRuleKindEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuleLastExecutionStatusEnum(string(m.LastExecutionStatus)); !ok && m.LastExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastExecutionStatus: %s. Supported values are: %s.", m.LastExecutionStatus, strings.Join(GetRuleLastExecutionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleLastExecutionStatusEnum Enum with underlying type: string
type RuleLastExecutionStatusEnum string

// Set of constants representing the allowable values for RuleLastExecutionStatusEnum
const (
	RuleLastExecutionStatusFailed    RuleLastExecutionStatusEnum = "FAILED"
	RuleLastExecutionStatusSucceeded RuleLastExecutionStatusEnum = "SUCCEEDED"
)

var mappingRuleLastExecutionStatusEnum = map[string]RuleLastExecutionStatusEnum{
	"FAILED":    RuleLastExecutionStatusFailed,
	"SUCCEEDED": RuleLastExecutionStatusSucceeded,
}

var mappingRuleLastExecutionStatusEnumLowerCase = map[string]RuleLastExecutionStatusEnum{
	"failed":    RuleLastExecutionStatusFailed,
	"succeeded": RuleLastExecutionStatusSucceeded,
}

// GetRuleLastExecutionStatusEnumValues Enumerates the set of values for RuleLastExecutionStatusEnum
func GetRuleLastExecutionStatusEnumValues() []RuleLastExecutionStatusEnum {
	values := make([]RuleLastExecutionStatusEnum, 0)
	for _, v := range mappingRuleLastExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleLastExecutionStatusEnumStringValues Enumerates the set of values in String for RuleLastExecutionStatusEnum
func GetRuleLastExecutionStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingRuleLastExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleLastExecutionStatusEnum(val string) (RuleLastExecutionStatusEnum, bool) {
	enum, ok := mappingRuleLastExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
