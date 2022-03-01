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
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// ContinuousQuery A ContinuousQuery that can be used to save and share a given search result.
type ContinuousQuery struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The continuous query expression that is run periodically.
	Query *string `mandatory:"true" json:"query"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the resource was last modified.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	QueryStartTime ContinuousQueryStartPolicy `mandatory:"true" json:"queryStartTime"`

	// Interval in minutes that query is run periodically.
	IntervalInMinutes *int `mandatory:"true" json:"intervalInMinutes"`

	// The integer value that must be exceeded, fall below or equal to (depending on the operator), the query result to trigger an event.
	Threshold *int `mandatory:"true" json:"threshold"`

	// The user-friendly query name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the custom log for continouous query.
	CustomLogId *string `mandatory:"true" json:"customLogId"`

	// The state of the ContinuousQuerySeverity
	Severity ContinuousQuerySeverityEnum `mandatory:"true" json:"severity"`

	// operator used in continuous query
	//   1. EQUAL
	//   2. GREATER
	//   3. GREATERTHANEQUALTO
	//   4. LESS
	//   5. LESSTHANEQUALTO
	Operator ContinuousQueryOperatorEnum `mandatory:"true" json:"operator"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Recommendations to act in case of a notification produced by the query.
	RecommendationText *string `mandatory:"false" json:"recommendationText"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The state of the ContinuousQueryLifecycleState
	//   1. CREATING
	//   2. ACTIVE   ContinuousQuery is active and can be used by other users
	//   3. UPDATING
	//   4. INACTIVE
	//   5. DELETING
	//   6. FAILED
	LifecycleState ContinuousQueryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ContinuousQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContinuousQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContinuousQuerySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetContinuousQuerySeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContinuousQueryOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetContinuousQueryOperatorEnumStringValues(), ",")))
	}

	if _, ok := GetMappingContinuousQueryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetContinuousQueryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ContinuousQuery) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags        map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags       map[string]string                 `json:"freeformTags"`
		RecommendationText *string                           `json:"recommendationText"`
		Description        *string                           `json:"description"`
		IsEnabled          *bool                             `json:"isEnabled"`
		LifecycleState     ContinuousQueryLifecycleStateEnum `json:"lifecycleState"`
		Id                 *string                           `json:"id"`
		Query              *string                           `json:"query"`
		CompartmentId      *string                           `json:"compartmentId"`
		TimeCreated        *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                   `json:"timeUpdated"`
		QueryStartTime     continuousquerystartpolicy        `json:"queryStartTime"`
		IntervalInMinutes  *int                              `json:"intervalInMinutes"`
		Threshold          *int                              `json:"threshold"`
		DisplayName        *string                           `json:"displayName"`
		CustomLogId        *string                           `json:"customLogId"`
		Severity           ContinuousQuerySeverityEnum       `json:"severity"`
		Operator           ContinuousQueryOperatorEnum       `json:"operator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.RecommendationText = model.RecommendationText

	m.Description = model.Description

	m.IsEnabled = model.IsEnabled

	m.LifecycleState = model.LifecycleState

	m.Id = model.Id

	m.Query = model.Query

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	nn, e = model.QueryStartTime.UnmarshalPolymorphicJSON(model.QueryStartTime.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.QueryStartTime = nn.(ContinuousQueryStartPolicy)
	} else {
		m.QueryStartTime = nil
	}

	m.IntervalInMinutes = model.IntervalInMinutes

	m.Threshold = model.Threshold

	m.DisplayName = model.DisplayName

	m.CustomLogId = model.CustomLogId

	m.Severity = model.Severity

	m.Operator = model.Operator

	return
}

// ContinuousQueryLifecycleStateEnum Enum with underlying type: string
type ContinuousQueryLifecycleStateEnum string

// Set of constants representing the allowable values for ContinuousQueryLifecycleStateEnum
const (
	ContinuousQueryLifecycleStateCreating ContinuousQueryLifecycleStateEnum = "CREATING"
	ContinuousQueryLifecycleStateActive   ContinuousQueryLifecycleStateEnum = "ACTIVE"
	ContinuousQueryLifecycleStateUpdating ContinuousQueryLifecycleStateEnum = "UPDATING"
	ContinuousQueryLifecycleStateInactive ContinuousQueryLifecycleStateEnum = "INACTIVE"
	ContinuousQueryLifecycleStateDeleting ContinuousQueryLifecycleStateEnum = "DELETING"
	ContinuousQueryLifecycleStateFailed   ContinuousQueryLifecycleStateEnum = "FAILED"
)

var mappingContinuousQueryLifecycleStateEnum = map[string]ContinuousQueryLifecycleStateEnum{
	"CREATING": ContinuousQueryLifecycleStateCreating,
	"ACTIVE":   ContinuousQueryLifecycleStateActive,
	"UPDATING": ContinuousQueryLifecycleStateUpdating,
	"INACTIVE": ContinuousQueryLifecycleStateInactive,
	"DELETING": ContinuousQueryLifecycleStateDeleting,
	"FAILED":   ContinuousQueryLifecycleStateFailed,
}

var mappingContinuousQueryLifecycleStateEnumLowerCase = map[string]ContinuousQueryLifecycleStateEnum{
	"creating": ContinuousQueryLifecycleStateCreating,
	"active":   ContinuousQueryLifecycleStateActive,
	"updating": ContinuousQueryLifecycleStateUpdating,
	"inactive": ContinuousQueryLifecycleStateInactive,
	"deleting": ContinuousQueryLifecycleStateDeleting,
	"failed":   ContinuousQueryLifecycleStateFailed,
}

// GetContinuousQueryLifecycleStateEnumValues Enumerates the set of values for ContinuousQueryLifecycleStateEnum
func GetContinuousQueryLifecycleStateEnumValues() []ContinuousQueryLifecycleStateEnum {
	values := make([]ContinuousQueryLifecycleStateEnum, 0)
	for _, v := range mappingContinuousQueryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetContinuousQueryLifecycleStateEnumStringValues Enumerates the set of values in String for ContinuousQueryLifecycleStateEnum
func GetContinuousQueryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingContinuousQueryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContinuousQueryLifecycleStateEnum(val string) (ContinuousQueryLifecycleStateEnum, bool) {
	enum, ok := mappingContinuousQueryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ContinuousQuerySeverityEnum Enum with underlying type: string
type ContinuousQuerySeverityEnum string

// Set of constants representing the allowable values for ContinuousQuerySeverityEnum
const (
	ContinuousQuerySeverityCritical ContinuousQuerySeverityEnum = "CRITICAL"
	ContinuousQuerySeverityHigh     ContinuousQuerySeverityEnum = "HIGH"
	ContinuousQuerySeverityMedium   ContinuousQuerySeverityEnum = "MEDIUM"
	ContinuousQuerySeverityLow      ContinuousQuerySeverityEnum = "LOW"
	ContinuousQuerySeverityMinor    ContinuousQuerySeverityEnum = "MINOR"
	ContinuousQuerySeverityNone     ContinuousQuerySeverityEnum = "NONE"
)

var mappingContinuousQuerySeverityEnum = map[string]ContinuousQuerySeverityEnum{
	"CRITICAL": ContinuousQuerySeverityCritical,
	"HIGH":     ContinuousQuerySeverityHigh,
	"MEDIUM":   ContinuousQuerySeverityMedium,
	"LOW":      ContinuousQuerySeverityLow,
	"MINOR":    ContinuousQuerySeverityMinor,
	"NONE":     ContinuousQuerySeverityNone,
}

var mappingContinuousQuerySeverityEnumLowerCase = map[string]ContinuousQuerySeverityEnum{
	"critical": ContinuousQuerySeverityCritical,
	"high":     ContinuousQuerySeverityHigh,
	"medium":   ContinuousQuerySeverityMedium,
	"low":      ContinuousQuerySeverityLow,
	"minor":    ContinuousQuerySeverityMinor,
	"none":     ContinuousQuerySeverityNone,
}

// GetContinuousQuerySeverityEnumValues Enumerates the set of values for ContinuousQuerySeverityEnum
func GetContinuousQuerySeverityEnumValues() []ContinuousQuerySeverityEnum {
	values := make([]ContinuousQuerySeverityEnum, 0)
	for _, v := range mappingContinuousQuerySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetContinuousQuerySeverityEnumStringValues Enumerates the set of values in String for ContinuousQuerySeverityEnum
func GetContinuousQuerySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
		"NONE",
	}
}

// GetMappingContinuousQuerySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContinuousQuerySeverityEnum(val string) (ContinuousQuerySeverityEnum, bool) {
	enum, ok := mappingContinuousQuerySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ContinuousQueryOperatorEnum Enum with underlying type: string
type ContinuousQueryOperatorEnum string

// Set of constants representing the allowable values for ContinuousQueryOperatorEnum
const (
	ContinuousQueryOperatorEqual              ContinuousQueryOperatorEnum = "EQUAL"
	ContinuousQueryOperatorGreater            ContinuousQueryOperatorEnum = "GREATER"
	ContinuousQueryOperatorGreaterthanequalto ContinuousQueryOperatorEnum = "GREATERTHANEQUALTO"
	ContinuousQueryOperatorLess               ContinuousQueryOperatorEnum = "LESS"
	ContinuousQueryOperatorLessthanequalto    ContinuousQueryOperatorEnum = "LESSTHANEQUALTO"
)

var mappingContinuousQueryOperatorEnum = map[string]ContinuousQueryOperatorEnum{
	"EQUAL":              ContinuousQueryOperatorEqual,
	"GREATER":            ContinuousQueryOperatorGreater,
	"GREATERTHANEQUALTO": ContinuousQueryOperatorGreaterthanequalto,
	"LESS":               ContinuousQueryOperatorLess,
	"LESSTHANEQUALTO":    ContinuousQueryOperatorLessthanequalto,
}

var mappingContinuousQueryOperatorEnumLowerCase = map[string]ContinuousQueryOperatorEnum{
	"equal":              ContinuousQueryOperatorEqual,
	"greater":            ContinuousQueryOperatorGreater,
	"greaterthanequalto": ContinuousQueryOperatorGreaterthanequalto,
	"less":               ContinuousQueryOperatorLess,
	"lessthanequalto":    ContinuousQueryOperatorLessthanequalto,
}

// GetContinuousQueryOperatorEnumValues Enumerates the set of values for ContinuousQueryOperatorEnum
func GetContinuousQueryOperatorEnumValues() []ContinuousQueryOperatorEnum {
	values := make([]ContinuousQueryOperatorEnum, 0)
	for _, v := range mappingContinuousQueryOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetContinuousQueryOperatorEnumStringValues Enumerates the set of values in String for ContinuousQueryOperatorEnum
func GetContinuousQueryOperatorEnumStringValues() []string {
	return []string{
		"EQUAL",
		"GREATER",
		"GREATERTHANEQUALTO",
		"LESS",
		"LESSTHANEQUALTO",
	}
}

// GetMappingContinuousQueryOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContinuousQueryOperatorEnum(val string) (ContinuousQueryOperatorEnum, bool) {
	enum, ok := mappingContinuousQueryOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
