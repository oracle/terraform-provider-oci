// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see https://docs.oracle.com/en-us/iaas/Content/Logging/Concepts/loggingoverview.htm.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogRuleSummary A LogRule that can be used to save and share a given search result.
type LogRuleSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The log rule query (https://docs.oracle.com/en-us/iaas/Content/Logging/Reference/query_language_specification.htm) that is run periodically.
	// For example,
	// search "loggroup-id" subject IN (INPUTLOOKUP 'objectstorage://bmc-logging-test/lookups/subjects.json' subject)
	// summarize count() as eventsCount by type as LogType, source, subject sort by eventsCount
	Query *string `mandatory:"true" json:"query"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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
	Operator LogRuleOperatorEnum `mandatory:"true" json:"operator"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	QueryStartPolicy LogRuleStartPolicy `mandatory:"false" json:"queryStartPolicy"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// 1. *ENABLED*    Log Rule is enabled
	// 2. *DISABLED*   Log Rule is disabled
	LogRuleStatus LogRuleSummaryLogRuleStatusEnum `mandatory:"false" json:"logRuleStatus,omitempty"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Recommended actions to take in case of a notification produced by the log rule query.
	// For example,
	// when this event happens,
	//   . check the logs under <dir> and search for event.
	//   . If you find any occurrences of X open a security event in the queue https://queue
	RecommendationText *string `mandatory:"false" json:"recommendationText"`

	// lifecycle state of log rule
	LifecycleState LogRuleLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m LogRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogRuleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogRuleSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetLogRuleSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetLogRuleOperatorEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLogRuleSummaryLogRuleStatusEnum(string(m.LogRuleStatus)); !ok && m.LogRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogRuleStatus: %s. Supported values are: %s.", m.LogRuleStatus, strings.Join(GetLogRuleSummaryLogRuleStatusEnumStringValues(), ",")))
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
func (m *LogRuleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeLastModified   *common.SDKTime                   `json:"timeLastModified"`
		QueryStartPolicy   logrulestartpolicy                `json:"queryStartPolicy"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		LogRuleStatus      LogRuleSummaryLogRuleStatusEnum   `json:"logRuleStatus"`
		Description        *string                           `json:"description"`
		RecommendationText *string                           `json:"recommendationText"`
		LifecycleState     LogRuleLifecycleStateEnum         `json:"lifecycleState"`
		Id                 *string                           `json:"id"`
		Query              *string                           `json:"query"`
		CompartmentId      *string                           `json:"compartmentId"`
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

	nn, e = model.QueryStartPolicy.UnmarshalPolymorphicJSON(model.QueryStartPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartPolicy = nn.(LogRuleStartPolicy)
	} else {
		m.QueryStartPolicy = nil
	}

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.LogRuleStatus = model.LogRuleStatus

	m.Description = model.Description

	m.RecommendationText = model.RecommendationText

	m.LifecycleState = model.LifecycleState

	m.Id = model.Id

	m.Query = model.Query

	m.CompartmentId = model.CompartmentId

	m.QueryRecurrences = model.QueryRecurrences

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}

// LogRuleSummaryLogRuleStatusEnum Enum with underlying type: string
type LogRuleSummaryLogRuleStatusEnum string

// Set of constants representing the allowable values for LogRuleSummaryLogRuleStatusEnum
const (
	LogRuleSummaryLogRuleStatusEnabled  LogRuleSummaryLogRuleStatusEnum = "ENABLED"
	LogRuleSummaryLogRuleStatusDisabled LogRuleSummaryLogRuleStatusEnum = "DISABLED"
)

var mappingLogRuleSummaryLogRuleStatusEnum = map[string]LogRuleSummaryLogRuleStatusEnum{
	"ENABLED":  LogRuleSummaryLogRuleStatusEnabled,
	"DISABLED": LogRuleSummaryLogRuleStatusDisabled,
}

var mappingLogRuleSummaryLogRuleStatusEnumLowerCase = map[string]LogRuleSummaryLogRuleStatusEnum{
	"enabled":  LogRuleSummaryLogRuleStatusEnabled,
	"disabled": LogRuleSummaryLogRuleStatusDisabled,
}

// GetLogRuleSummaryLogRuleStatusEnumValues Enumerates the set of values for LogRuleSummaryLogRuleStatusEnum
func GetLogRuleSummaryLogRuleStatusEnumValues() []LogRuleSummaryLogRuleStatusEnum {
	values := make([]LogRuleSummaryLogRuleStatusEnum, 0)
	for _, v := range mappingLogRuleSummaryLogRuleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLogRuleSummaryLogRuleStatusEnumStringValues Enumerates the set of values in String for LogRuleSummaryLogRuleStatusEnum
func GetLogRuleSummaryLogRuleStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingLogRuleSummaryLogRuleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogRuleSummaryLogRuleStatusEnum(val string) (LogRuleSummaryLogRuleStatusEnum, bool) {
	enum, ok := mappingLogRuleSummaryLogRuleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
