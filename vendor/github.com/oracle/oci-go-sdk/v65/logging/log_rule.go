// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogRule A LogRule that can be used to save and share a given search result.
type LogRule struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The log rule query (https://docs.oracle.com/en-us/iaas/Content/Logging/Reference/query_language_specification.htm) that is run periodically.
	// For example,
	// search "loggroup-id" subject IN (INPUTLOOKUP 'objectstorage://bmc-logging-test/lookups/subjects.json' subject)
	// summarize count() as eventsCount by type as LogType, source, subject sort by eventsCount
	Query *string `mandatory:"true" json:"query"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	QueryStartPolicy LogRuleStartPolicy `mandatory:"true" json:"queryStartPolicy"`

	// Log rule query recurrence rule in iCalendar format RFC 5545 (https://www.rfc-editor.org/rfc/rfc5545#section-3.3.10).
	// Freq should be one of "MINUTELY", "HOURLY", "DAILY" and "WEEKLY".
	// The specified time interval should be between 5 and 10080 minutes.
	// It cannot contain "UNTIL", "COUNT", "WKST" or "BY*".
	QueryRecurrences *string `mandatory:"true" json:"queryRecurrences"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), the threshold to trigger an event.
	Threshold *int `mandatory:"true" json:"threshold"`

	// The user-friendly log rule name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the custom log for log rule.
	CustomLogId *string `mandatory:"true" json:"customLogId"`

	// The state of the LogRuleSeverity
	Severity LogRuleSeverityEnum `mandatory:"true" json:"severity"`

	// operator used in log rule
	//   1. EQUAL
	//   2. GREATER
	//   3. GREATERTHANEQUALTO
	//   4. LESS
	//   5. LESSTHANEQUALTO
	Operator LogRuleOperatorEnum `mandatory:"true" json:"operator"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Recommended actions to take in case of a notification produced by the log rule query.
	// For example,
	// when this event happens,
	//   . check the logs under <dir> and search for event.
	//   . If you find any occurrences of X open a security event in the queue https://queue
	RecommendationText *string `mandatory:"false" json:"recommendationText"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// 1. *ENABLED*    Log Rule is enabled
	// 2. *DISABLED*   Log Rule is disabled
	LogRuleStatus LogRuleLogRuleStatusEnum `mandatory:"false" json:"logRuleStatus,omitempty"`

	// The state of the LogRuleLifecycleState
	//   1. CREATING
	//   2. ACTIVE   LogRule is active and can be used by other users
	//   3. UPDATING
	//   4. INACTIVE
	//   5. DELETING
	//   6. FAILED
	LifecycleState LogRuleLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m LogRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogRuleSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetLogRuleSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetLogRuleOperatorEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLogRuleLogRuleStatusEnum(string(m.LogRuleStatus)); !ok && m.LogRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogRuleStatus: %s. Supported values are: %s.", m.LogRuleStatus, strings.Join(GetLogRuleLogRuleStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogRuleLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LogRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeLastModified   *common.SDKTime                   `json:"timeLastModified"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		RecommendationText *string                           `json:"recommendationText"`
		Description        *string                           `json:"description"`
		LogRuleStatus      LogRuleLogRuleStatusEnum          `json:"logRuleStatus"`
		LifecycleState     LogRuleLifecycleStateEnum         `json:"lifecycleState"`
		Id                 *string                           `json:"id"`
		Query              *string                           `json:"query"`
		CompartmentId      *string                           `json:"compartmentId"`
		QueryStartPolicy   logrulestartpolicy                `json:"queryStartPolicy"`
		QueryRecurrences   *string                           `json:"queryRecurrences"`
		Threshold          *int                              `json:"threshold"`
		DisplayName        *string                           `json:"displayName"`
		CustomLogId        *string                           `json:"customLogId"`
		Severity           LogRuleSeverityEnum               `json:"severity"`
		Operator           LogRuleOperatorEnum               `json:"operator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.TimeLastModified = model.TimeLastModified

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.RecommendationText = model.RecommendationText

	m.Description = model.Description

	m.LogRuleStatus = model.LogRuleStatus

	m.LifecycleState = model.LifecycleState

	m.Id = model.Id

	m.Query = model.Query

	m.CompartmentId = model.CompartmentId

	nn, e = model.QueryStartPolicy.UnmarshalPolymorphicJSON(model.QueryStartPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartPolicy = nn.(LogRuleStartPolicy)
	} else {
		m.QueryStartPolicy = nil
	}

	m.QueryRecurrences = model.QueryRecurrences

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}

// LogRuleLogRuleStatusEnum Enum with underlying type: string
type LogRuleLogRuleStatusEnum string

// Set of constants representing the allowable values for LogRuleLogRuleStatusEnum
const (
	LogRuleLogRuleStatusEnabled  LogRuleLogRuleStatusEnum = "ENABLED"
	LogRuleLogRuleStatusDisabled LogRuleLogRuleStatusEnum = "DISABLED"
)

var mappingLogRuleLogRuleStatusEnum = map[string]LogRuleLogRuleStatusEnum{
	"ENABLED":  LogRuleLogRuleStatusEnabled,
	"DISABLED": LogRuleLogRuleStatusDisabled,
}

var mappingLogRuleLogRuleStatusEnumLowerCase = map[string]LogRuleLogRuleStatusEnum{
	"enabled":  LogRuleLogRuleStatusEnabled,
	"disabled": LogRuleLogRuleStatusDisabled,
}

// GetLogRuleLogRuleStatusEnumValues Enumerates the set of values for LogRuleLogRuleStatusEnum
func GetLogRuleLogRuleStatusEnumValues() []LogRuleLogRuleStatusEnum {
	values := make([]LogRuleLogRuleStatusEnum, 0)
	for _, v := range mappingLogRuleLogRuleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleLogRuleStatusEnumStringValues Enumerates the set of values in String for LogRuleLogRuleStatusEnum
func GetLogRuleLogRuleStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingLogRuleLogRuleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleLogRuleStatusEnum(val string) (LogRuleLogRuleStatusEnum, bool) {
	enum, ok := mappingLogRuleLogRuleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LogRuleLifecycleStateEnum Enum with underlying type: string
type LogRuleLifecycleStateEnum string

// Set of constants representing the allowable values for LogRuleLifecycleStateEnum
const (
	LogRuleLifecycleStateCreating LogRuleLifecycleStateEnum = "CREATING"
	LogRuleLifecycleStateActive   LogRuleLifecycleStateEnum = "ACTIVE"
	LogRuleLifecycleStateUpdating LogRuleLifecycleStateEnum = "UPDATING"
	LogRuleLifecycleStateInactive LogRuleLifecycleStateEnum = "INACTIVE"
	LogRuleLifecycleStateDeleting LogRuleLifecycleStateEnum = "DELETING"
	LogRuleLifecycleStateFailed   LogRuleLifecycleStateEnum = "FAILED"
)

var mappingLogRuleLifecycleStateEnum = map[string]LogRuleLifecycleStateEnum{
	"CREATING": LogRuleLifecycleStateCreating,
	"ACTIVE":   LogRuleLifecycleStateActive,
	"UPDATING": LogRuleLifecycleStateUpdating,
	"INACTIVE": LogRuleLifecycleStateInactive,
	"DELETING": LogRuleLifecycleStateDeleting,
	"FAILED":   LogRuleLifecycleStateFailed,
}

var mappingLogRuleLifecycleStateEnumLowerCase = map[string]LogRuleLifecycleStateEnum{
	"creating": LogRuleLifecycleStateCreating,
	"active":   LogRuleLifecycleStateActive,
	"updating": LogRuleLifecycleStateUpdating,
	"inactive": LogRuleLifecycleStateInactive,
	"deleting": LogRuleLifecycleStateDeleting,
	"failed":   LogRuleLifecycleStateFailed,
}

// GetLogRuleLifecycleStateEnumValues Enumerates the set of values for LogRuleLifecycleStateEnum
func GetLogRuleLifecycleStateEnumValues() []LogRuleLifecycleStateEnum {
	values := make([]LogRuleLifecycleStateEnum, 0)
	for _, v := range mappingLogRuleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleLifecycleStateEnumStringValues Enumerates the set of values in String for LogRuleLifecycleStateEnum
func GetLogRuleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingLogRuleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleLifecycleStateEnum(val string) (LogRuleLifecycleStateEnum, bool) {
	enum, ok := mappingLogRuleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LogRuleSeverityEnum Enum with underlying type: string
type LogRuleSeverityEnum string

// Set of constants representing the allowable values for LogRuleSeverityEnum
const (
	LogRuleSeverityCritical      LogRuleSeverityEnum = "CRITICAL"
	LogRuleSeverityHigh          LogRuleSeverityEnum = "HIGH"
	LogRuleSeverityMedium        LogRuleSeverityEnum = "MEDIUM"
	LogRuleSeverityLow           LogRuleSeverityEnum = "LOW"
	LogRuleSeverityMinor         LogRuleSeverityEnum = "MINOR"
	LogRuleSeverityInformational LogRuleSeverityEnum = "INFORMATIONAL"
	LogRuleSeverityNone          LogRuleSeverityEnum = "NONE"
)

var mappingLogRuleSeverityEnum = map[string]LogRuleSeverityEnum{
	"CRITICAL":      LogRuleSeverityCritical,
	"HIGH":          LogRuleSeverityHigh,
	"MEDIUM":        LogRuleSeverityMedium,
	"LOW":           LogRuleSeverityLow,
	"MINOR":         LogRuleSeverityMinor,
	"INFORMATIONAL": LogRuleSeverityInformational,
	"NONE":          LogRuleSeverityNone,
}

var mappingLogRuleSeverityEnumLowerCase = map[string]LogRuleSeverityEnum{
	"critical":      LogRuleSeverityCritical,
	"high":          LogRuleSeverityHigh,
	"medium":        LogRuleSeverityMedium,
	"low":           LogRuleSeverityLow,
	"minor":         LogRuleSeverityMinor,
	"informational": LogRuleSeverityInformational,
	"none":          LogRuleSeverityNone,
}

// GetLogRuleSeverityEnumValues Enumerates the set of values for LogRuleSeverityEnum
func GetLogRuleSeverityEnumValues() []LogRuleSeverityEnum {
	values := make([]LogRuleSeverityEnum, 0)
	for _, v := range mappingLogRuleSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleSeverityEnumStringValues Enumerates the set of values in String for LogRuleSeverityEnum
func GetLogRuleSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
		"INFORMATIONAL",
		"NONE",
	}
}

// GetMappingLogRuleSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleSeverityEnum(val string) (LogRuleSeverityEnum, bool) {
	enum, ok := mappingLogRuleSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LogRuleOperatorEnum Enum with underlying type: string
type LogRuleOperatorEnum string

// Set of constants representing the allowable values for LogRuleOperatorEnum
const (
	LogRuleOperatorEqual              LogRuleOperatorEnum = "EQUAL"
	LogRuleOperatorGreater            LogRuleOperatorEnum = "GREATER"
	LogRuleOperatorGreaterthanequalto LogRuleOperatorEnum = "GREATERTHANEQUALTO"
	LogRuleOperatorLess               LogRuleOperatorEnum = "LESS"
	LogRuleOperatorLessthanequalto    LogRuleOperatorEnum = "LESSTHANEQUALTO"
)

var mappingLogRuleOperatorEnum = map[string]LogRuleOperatorEnum{
	"EQUAL":              LogRuleOperatorEqual,
	"GREATER":            LogRuleOperatorGreater,
	"GREATERTHANEQUALTO": LogRuleOperatorGreaterthanequalto,
	"LESS":               LogRuleOperatorLess,
	"LESSTHANEQUALTO":    LogRuleOperatorLessthanequalto,
}

var mappingLogRuleOperatorEnumLowerCase = map[string]LogRuleOperatorEnum{
	"equal":              LogRuleOperatorEqual,
	"greater":            LogRuleOperatorGreater,
	"greaterthanequalto": LogRuleOperatorGreaterthanequalto,
	"less":               LogRuleOperatorLess,
	"lessthanequalto":    LogRuleOperatorLessthanequalto,
}

// GetLogRuleOperatorEnumValues Enumerates the set of values for LogRuleOperatorEnum
func GetLogRuleOperatorEnumValues() []LogRuleOperatorEnum {
	values := make([]LogRuleOperatorEnum, 0)
	for _, v := range mappingLogRuleOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleOperatorEnumStringValues Enumerates the set of values in String for LogRuleOperatorEnum
func GetLogRuleOperatorEnumStringValues() []string {
	return []string{
		"EQUAL",
		"GREATER",
		"GREATERTHANEQUALTO",
		"LESS",
		"LESSTHANEQUALTO",
	}
}

// GetMappingLogRuleOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleOperatorEnum(val string) (LogRuleOperatorEnum, bool) {
	enum, ok := mappingLogRuleOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
