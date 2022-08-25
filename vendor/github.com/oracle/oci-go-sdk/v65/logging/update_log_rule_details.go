// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateLogRuleDetails Log rule update object request details.
type UpdateLogRuleDetails struct {

	// The log rule query (https://docs.oracle.com/en-us/iaas/Content/Logging/Reference/query_language_specification.htm) that is run periodically.
	// For example,
	// search "loggroup-id" subject IN (INPUTLOOKUP 'objectstorage://bmc-logging-test/lookups/subjects.json' subject)
	// summarize count() as eventsCount by type as LogType, source, subject sort by eventsCount
	Query *string `mandatory:"false" json:"query"`

	// Log rule query recurrence rule in iCalendar format RFC 5545 (https://www.rfc-editor.org/rfc/rfc5545#section-3.3.10).
	// Freq should be one of "MINUTELY", "HOURLY", "DAILY" and "WEEKLY".
	// The specified time interval should be between 5 and 10080 minutes.
	// It cannot contain "UNTIL", "COUNT", "WKST" or "BY*".
	QueryRecurrences *string `mandatory:"false" json:"queryRecurrences"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), the threshold to trigger an event.
	Threshold *int `mandatory:"false" json:"threshold"`

	// The user-friendly log rule name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the custom log for log rule.
	CustomLogId *string `mandatory:"false" json:"customLogId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	QueryStartPolicy LogRuleStartPolicy `mandatory:"false" json:"queryStartPolicy"`

	// 1. *ENABLED*    Log Rule is enabled
	// 2. *DISABLED*   Log Rule is disabled
	LogRuleStatus UpdateLogRuleDetailsLogRuleStatusEnum `mandatory:"false" json:"logRuleStatus,omitempty"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Recommended actions to take in case of a notification produced by the log rule query.
	// For example,
	// when this event happens,
	//   . check the logs under <dir> and search for event.
	//   . If you find any occurrences of X open a security event in the queue https://queue
	RecommendationText *string `mandatory:"false" json:"recommendationText"`

	// The state of the LogRuleSeverity
	Severity LogRuleSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// operator used in log rule
	Operator LogRuleOperatorEnum `mandatory:"false" json:"operator,omitempty"`
}

func (m UpdateLogRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLogRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateLogRuleDetailsLogRuleStatusEnum(string(m.LogRuleStatus)); !ok && m.LogRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogRuleStatus: %s. Supported values are: %s.", m.LogRuleStatus, strings.Join(GetUpdateLogRuleDetailsLogRuleStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetLogRuleSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetLogRuleOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateLogRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Query              *string                               `json:"query"`
		QueryRecurrences   *string                               `json:"queryRecurrences"`
		Threshold          *int                                  `json:"threshold"`
		DisplayName        *string                               `json:"displayName"`
		CustomLogId        *string                               `json:"customLogId"`
		DefinedTags        map[string]map[string]interface{}     `json:"definedTags"`
		FreeformTags       map[string]string                     `json:"freeformTags"`
		QueryStartPolicy   logrulestartpolicy                    `json:"queryStartPolicy"`
		LogRuleStatus      UpdateLogRuleDetailsLogRuleStatusEnum `json:"logRuleStatus"`
		Description        *string                               `json:"description"`
		RecommendationText *string                               `json:"recommendationText"`
		Severity           LogRuleSeverityEnum                   `json:"severity"`
		Operator           LogRuleOperatorEnum                   `json:"operator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Query = model.Query

	m.QueryRecurrences = model.QueryRecurrences

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	nn, e = model.QueryStartPolicy.UnmarshalPolymorphicJSON(model.QueryStartPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartPolicy = nn.(LogRuleStartPolicy)
	} else {
		m.QueryStartPolicy = nil
	}

	m.LogRuleStatus = model.LogRuleStatus

	m.Description = model.Description

	m.RecommendationText = model.RecommendationText

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}

// UpdateLogRuleDetailsLogRuleStatusEnum Enum with underlying type: string
type UpdateLogRuleDetailsLogRuleStatusEnum string

// Set of constants representing the allowable values for UpdateLogRuleDetailsLogRuleStatusEnum
const (
	UpdateLogRuleDetailsLogRuleStatusEnabled  UpdateLogRuleDetailsLogRuleStatusEnum = "ENABLED"
	UpdateLogRuleDetailsLogRuleStatusDisabled UpdateLogRuleDetailsLogRuleStatusEnum = "DISABLED"
)

var mappingUpdateLogRuleDetailsLogRuleStatusEnum = map[string]UpdateLogRuleDetailsLogRuleStatusEnum{
	"ENABLED":  UpdateLogRuleDetailsLogRuleStatusEnabled,
	"DISABLED": UpdateLogRuleDetailsLogRuleStatusDisabled,
}

var mappingUpdateLogRuleDetailsLogRuleStatusEnumLowerCase = map[string]UpdateLogRuleDetailsLogRuleStatusEnum{
	"enabled":  UpdateLogRuleDetailsLogRuleStatusEnabled,
	"disabled": UpdateLogRuleDetailsLogRuleStatusDisabled,
}

// GetUpdateLogRuleDetailsLogRuleStatusEnumValues Enumerates the set of values for UpdateLogRuleDetailsLogRuleStatusEnum
func GetUpdateLogRuleDetailsLogRuleStatusEnumValues() []UpdateLogRuleDetailsLogRuleStatusEnum {
	values := make([]UpdateLogRuleDetailsLogRuleStatusEnum, 0)
	for _, v := range mappingUpdateLogRuleDetailsLogRuleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateLogRuleDetailsLogRuleStatusEnumStringValues Enumerates the set of values in String for UpdateLogRuleDetailsLogRuleStatusEnum
func GetUpdateLogRuleDetailsLogRuleStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateLogRuleDetailsLogRuleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateLogRuleDetailsLogRuleStatusEnum(val string) (UpdateLogRuleDetailsLogRuleStatusEnum, bool) {
	enum, ok := mappingUpdateLogRuleDetailsLogRuleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
