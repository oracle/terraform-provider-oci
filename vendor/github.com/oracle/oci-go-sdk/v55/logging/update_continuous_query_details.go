// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// UpdateContinuousQueryDetails Continuous query creation object.
type UpdateContinuousQueryDetails struct {

	// The continuous query expression that is run periodically.
	Query *string `mandatory:"false" json:"query"`

	// Interval in minutes that query is run periodically.
	IntervalInMinutes *int `mandatory:"false" json:"intervalInMinutes"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), the query result to trigger an event.
	Threshold *int `mandatory:"false" json:"threshold"`

	// The user-friendly query name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the custom log for continouous query.
	CustomLogId *string `mandatory:"false" json:"customLogId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	QueryStartTime ContinuousQueryStartPolicy `mandatory:"false" json:"queryStartTime"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Recommendations to act in case of a notification produced by the query.
	RecommendationText *string `mandatory:"false" json:"recommendationText"`

	// The state of the ContinuousQuerySeverity
	Severity ContinuousQuerySeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// operator used in continuous query
	Operator ContinuousQueryOperatorEnum `mandatory:"false" json:"operator,omitempty"`
}

func (m UpdateContinuousQueryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateContinuousQueryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingContinuousQuerySeverityEnum[string(m.Severity)]; !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetContinuousQuerySeverityEnumStringValues(), ",")))
	}
	if _, ok := mappingContinuousQueryOperatorEnum[string(m.Operator)]; !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetContinuousQueryOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateContinuousQueryDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Query              *string                           `json:"query"`
		IntervalInMinutes  *int                              `json:"intervalInMinutes"`
		Threshold          *int                              `json:"threshold"`
		DisplayName        *string                           `json:"displayName"`
		CustomLogId        *string                           `json:"customLogId"`
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		QueryStartTime     continuousquerystartpolicy        `json:"queryStartTime"`
		IsEnabled          *bool                             `json:"isEnabled"`
		Description        *string                           `json:"description"`
		RecommendationText *string                           `json:"recommendationText"`
		Severity           ContinuousQuerySeverityEnum       `json:"severity"`
		Operator           ContinuousQueryOperatorEnum       `json:"operator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Query = model.Query

	m.IntervalInMinutes = model.IntervalInMinutes

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	nn, e = model.QueryStartTime.UnmarshalPolymorphicJSON(model.QueryStartTime.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartTime = nn.(ContinuousQueryStartPolicy)
	} else {
		m.QueryStartTime = nil
	}

	m.IsEnabled = model.IsEnabled

	m.Description = model.Description

	m.RecommendationText = model.RecommendationText

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}
