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

// CreateContinuousQueryDetails Continuous query creation object.
type CreateContinuousQueryDetails struct {

	// The continuous query expression that is run periodically.
	// For example,
	// search "loggroup-id" subject IN (INPUTLOOKUP 'objectstorage://bmc-logging-test/lookups/subjects.json' subject)
	// summarize count() as eventsCount by type as LogType, source, subject sort by eventsCount
	Query *string `mandatory:"true" json:"query"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Interval in minutes that query is run periodically.
	IntervalInMinutes *int `mandatory:"true" json:"intervalInMinutes"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), the query result to trigger an event.
	Threshold *int `mandatory:"true" json:"threshold"`

	// The user-friendly query name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the custom log for continuous query.
	CustomLogId *string `mandatory:"true" json:"customLogId"`

	QueryStartTime ContinuousQueryStartPolicy `mandatory:"true" json:"queryStartTime"`

	// The state of the ContinuousQuerySeverity
	Severity ContinuousQuerySeverityEnum `mandatory:"true" json:"severity"`

	// operator used in continuous query
	Operator ContinuousQueryOperatorEnum `mandatory:"true" json:"operator"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Recommendations to act in case of a notification produced by the query.
	// For example,
	// when this event happens,
	// . check the logs under <dir> and search for event.
	// . If you find any occurrences of X open a security event in the queue https://queue
	RecommendationText *string `mandatory:"false" json:"recommendationText"`
}

func (m CreateContinuousQueryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContinuousQueryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContinuousQuerySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetContinuousQuerySeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContinuousQueryOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetContinuousQueryOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateContinuousQueryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		IsEnabled          *bool                             `json:"isEnabled"`
		Description        *string                           `json:"description"`
		RecommendationText *string                           `json:"recommendationText"`
		Query              *string                           `json:"query"`
		CompartmentId      *string                           `json:"compartmentId"`
		IntervalInMinutes  *int                              `json:"intervalInMinutes"`
		Threshold          *int                              `json:"threshold"`
		DisplayName        *string                           `json:"displayName"`
		CustomLogId        *string                           `json:"customLogId"`
		QueryStartTime     continuousquerystartpolicy        `json:"queryStartTime"`
		Severity           ContinuousQuerySeverityEnum       `json:"severity"`
		Operator           ContinuousQueryOperatorEnum       `json:"operator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.IsEnabled = model.IsEnabled

	m.Description = model.Description

	m.RecommendationText = model.RecommendationText

	m.Query = model.Query

	m.CompartmentId = model.CompartmentId

	m.IntervalInMinutes = model.IntervalInMinutes

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	nn, e = model.QueryStartTime.UnmarshalPolymorphicJSON(model.QueryStartTime.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartTime = nn.(ContinuousQueryStartPolicy)
	} else {
		m.QueryStartTime = nil
	}

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}
