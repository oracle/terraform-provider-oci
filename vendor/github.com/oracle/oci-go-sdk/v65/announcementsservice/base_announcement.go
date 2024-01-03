// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseAnnouncement Incident information that forms the basis of an announcement. Avoid entering confidential information.
type BaseAnnouncement interface {

	// The OCID of the announcement.
	GetId() *string

	// The reference Jira ticket number.
	GetReferenceTicketNumber() *string

	// A summary of the issue. A summary might appear in the console banner view of the announcement or in
	// an email subject line. Avoid entering confidential information.
	GetSummary() *string

	// Impacted Oracle Cloud Infrastructure services.
	GetServices() []string

	// Impacted regions.
	GetAffectedRegions() []string

	// The type of announcement. An announcement's type signals its severity.
	GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum

	// The current lifecycle state of the announcement.
	GetLifecycleState() BaseAnnouncementLifecycleStateEnum

	// Whether the announcement is displayed as a banner in the console.
	GetIsBanner() *bool

	// The label associated with an initial time value.
	// Example: `Time Started`
	GetTimeOneTitle() *string

	// The type of a time associated with an initial time value. If the `timeOneTitle` attribute is present, then the `timeOneTitle` attribute contains a label of `timeOneType` in English.
	// Example: `START_TIME`
	GetTimeOneType() BaseAnnouncementTimeOneTypeEnum

	// The actual value of the first time value for the event. Typically, this denotes the time an event started, but the meaning
	// can vary, depending on the announcement type. The `timeOneType` attribute describes the meaning.
	GetTimeOneValue() *common.SDKTime

	// The label associated with a second time value.
	// Example: `Time Ended`
	GetTimeTwoTitle() *string

	// The type of a time associated with second time value. If the `timeTwoTitle` attribute is present, then the `timeTwoTitle` attribute contains a label of `timeTwoType` in English.
	// Example: `END_TIME`
	GetTimeTwoType() BaseAnnouncementTimeTwoTypeEnum

	// The actual value of the second time value. Typically, this denotes the time an event ended, but the meaning
	// can vary, depending on the announcement type. The `timeTwoType` attribute describes the meaning.
	GetTimeTwoValue() *common.SDKTime

	// The date and time the announcement was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-01-01T17:43:01.389+0000`
	GetTimeCreated() *common.SDKTime

	// The date and time the announcement was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-01-01T17:43:01.389+0000`
	GetTimeUpdated() *common.SDKTime

	// The name of the environment that this announcement pertains to.
	GetEnvironmentName() *string

	// The platform type that this announcement pertains to.
	GetPlatformType() BaseAnnouncementPlatformTypeEnum

	// The sequence of connected announcements, or announcement chain, that this announcement belongs to. Related announcements share the same chain ID.
	GetChainId() *string
}

type baseannouncement struct {
	JsonData              []byte
	TimeOneTitle          *string                              `mandatory:"false" json:"timeOneTitle"`
	TimeOneType           BaseAnnouncementTimeOneTypeEnum      `mandatory:"false" json:"timeOneType,omitempty"`
	TimeOneValue          *common.SDKTime                      `mandatory:"false" json:"timeOneValue"`
	TimeTwoTitle          *string                              `mandatory:"false" json:"timeTwoTitle"`
	TimeTwoType           BaseAnnouncementTimeTwoTypeEnum      `mandatory:"false" json:"timeTwoType,omitempty"`
	TimeTwoValue          *common.SDKTime                      `mandatory:"false" json:"timeTwoValue"`
	TimeCreated           *common.SDKTime                      `mandatory:"false" json:"timeCreated"`
	TimeUpdated           *common.SDKTime                      `mandatory:"false" json:"timeUpdated"`
	EnvironmentName       *string                              `mandatory:"false" json:"environmentName"`
	PlatformType          BaseAnnouncementPlatformTypeEnum     `mandatory:"false" json:"platformType,omitempty"`
	ChainId               *string                              `mandatory:"false" json:"chainId"`
	Id                    *string                              `mandatory:"true" json:"id"`
	ReferenceTicketNumber *string                              `mandatory:"true" json:"referenceTicketNumber"`
	Summary               *string                              `mandatory:"true" json:"summary"`
	Services              []string                             `mandatory:"true" json:"services"`
	AffectedRegions       []string                             `mandatory:"true" json:"affectedRegions"`
	AnnouncementType      BaseAnnouncementAnnouncementTypeEnum `mandatory:"true" json:"announcementType"`
	LifecycleState        BaseAnnouncementLifecycleStateEnum   `mandatory:"true" json:"lifecycleState"`
	IsBanner              *bool                                `mandatory:"true" json:"isBanner"`
	Type                  string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseannouncement) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseannouncement baseannouncement
	s := struct {
		Model Unmarshalerbaseannouncement
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ReferenceTicketNumber = s.Model.ReferenceTicketNumber
	m.Summary = s.Model.Summary
	m.Services = s.Model.Services
	m.AffectedRegions = s.Model.AffectedRegions
	m.AnnouncementType = s.Model.AnnouncementType
	m.LifecycleState = s.Model.LifecycleState
	m.IsBanner = s.Model.IsBanner
	m.TimeOneTitle = s.Model.TimeOneTitle
	m.TimeOneType = s.Model.TimeOneType
	m.TimeOneValue = s.Model.TimeOneValue
	m.TimeTwoTitle = s.Model.TimeTwoTitle
	m.TimeTwoType = s.Model.TimeTwoType
	m.TimeTwoValue = s.Model.TimeTwoValue
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.EnvironmentName = s.Model.EnvironmentName
	m.PlatformType = s.Model.PlatformType
	m.ChainId = s.Model.ChainId
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseannouncement) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AnnouncementSummary":
		mm := AnnouncementSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "Announcement":
		mm := Announcement{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BaseAnnouncement: %s.", m.Type)
		return *m, nil
	}
}

// GetTimeOneTitle returns TimeOneTitle
func (m baseannouncement) GetTimeOneTitle() *string {
	return m.TimeOneTitle
}

// GetTimeOneType returns TimeOneType
func (m baseannouncement) GetTimeOneType() BaseAnnouncementTimeOneTypeEnum {
	return m.TimeOneType
}

// GetTimeOneValue returns TimeOneValue
func (m baseannouncement) GetTimeOneValue() *common.SDKTime {
	return m.TimeOneValue
}

// GetTimeTwoTitle returns TimeTwoTitle
func (m baseannouncement) GetTimeTwoTitle() *string {
	return m.TimeTwoTitle
}

// GetTimeTwoType returns TimeTwoType
func (m baseannouncement) GetTimeTwoType() BaseAnnouncementTimeTwoTypeEnum {
	return m.TimeTwoType
}

// GetTimeTwoValue returns TimeTwoValue
func (m baseannouncement) GetTimeTwoValue() *common.SDKTime {
	return m.TimeTwoValue
}

// GetTimeCreated returns TimeCreated
func (m baseannouncement) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m baseannouncement) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetEnvironmentName returns EnvironmentName
func (m baseannouncement) GetEnvironmentName() *string {
	return m.EnvironmentName
}

// GetPlatformType returns PlatformType
func (m baseannouncement) GetPlatformType() BaseAnnouncementPlatformTypeEnum {
	return m.PlatformType
}

// GetChainId returns ChainId
func (m baseannouncement) GetChainId() *string {
	return m.ChainId
}

// GetId returns Id
func (m baseannouncement) GetId() *string {
	return m.Id
}

// GetReferenceTicketNumber returns ReferenceTicketNumber
func (m baseannouncement) GetReferenceTicketNumber() *string {
	return m.ReferenceTicketNumber
}

// GetSummary returns Summary
func (m baseannouncement) GetSummary() *string {
	return m.Summary
}

// GetServices returns Services
func (m baseannouncement) GetServices() []string {
	return m.Services
}

// GetAffectedRegions returns AffectedRegions
func (m baseannouncement) GetAffectedRegions() []string {
	return m.AffectedRegions
}

// GetAnnouncementType returns AnnouncementType
func (m baseannouncement) GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum {
	return m.AnnouncementType
}

// GetLifecycleState returns LifecycleState
func (m baseannouncement) GetLifecycleState() BaseAnnouncementLifecycleStateEnum {
	return m.LifecycleState
}

// GetIsBanner returns IsBanner
func (m baseannouncement) GetIsBanner() *bool {
	return m.IsBanner
}

func (m baseannouncement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m baseannouncement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseAnnouncementAnnouncementTypeEnum(string(m.AnnouncementType)); !ok && m.AnnouncementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AnnouncementType: %s. Supported values are: %s.", m.AnnouncementType, strings.Join(GetBaseAnnouncementAnnouncementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseAnnouncementLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseAnnouncementTimeOneTypeEnum(string(m.TimeOneType)); !ok && m.TimeOneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TimeOneType: %s. Supported values are: %s.", m.TimeOneType, strings.Join(GetBaseAnnouncementTimeOneTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementTimeTwoTypeEnum(string(m.TimeTwoType)); !ok && m.TimeTwoType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TimeTwoType: %s. Supported values are: %s.", m.TimeTwoType, strings.Join(GetBaseAnnouncementTimeTwoTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetBaseAnnouncementPlatformTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseAnnouncementTimeOneTypeEnum Enum with underlying type: string
type BaseAnnouncementTimeOneTypeEnum string

// Set of constants representing the allowable values for BaseAnnouncementTimeOneTypeEnum
const (
	BaseAnnouncementTimeOneTypeActionRequiredBy BaseAnnouncementTimeOneTypeEnum = "ACTION_REQUIRED_BY"
	BaseAnnouncementTimeOneTypeNewStartTime     BaseAnnouncementTimeOneTypeEnum = "NEW_START_TIME"
	BaseAnnouncementTimeOneTypeOriginalEndTime  BaseAnnouncementTimeOneTypeEnum = "ORIGINAL_END_TIME"
	BaseAnnouncementTimeOneTypeReportDate       BaseAnnouncementTimeOneTypeEnum = "REPORT_DATE"
	BaseAnnouncementTimeOneTypeStartTime        BaseAnnouncementTimeOneTypeEnum = "START_TIME"
	BaseAnnouncementTimeOneTypeTimeDetected     BaseAnnouncementTimeOneTypeEnum = "TIME_DETECTED"
)

var mappingBaseAnnouncementTimeOneTypeEnum = map[string]BaseAnnouncementTimeOneTypeEnum{
	"ACTION_REQUIRED_BY": BaseAnnouncementTimeOneTypeActionRequiredBy,
	"NEW_START_TIME":     BaseAnnouncementTimeOneTypeNewStartTime,
	"ORIGINAL_END_TIME":  BaseAnnouncementTimeOneTypeOriginalEndTime,
	"REPORT_DATE":        BaseAnnouncementTimeOneTypeReportDate,
	"START_TIME":         BaseAnnouncementTimeOneTypeStartTime,
	"TIME_DETECTED":      BaseAnnouncementTimeOneTypeTimeDetected,
}

var mappingBaseAnnouncementTimeOneTypeEnumLowerCase = map[string]BaseAnnouncementTimeOneTypeEnum{
	"action_required_by": BaseAnnouncementTimeOneTypeActionRequiredBy,
	"new_start_time":     BaseAnnouncementTimeOneTypeNewStartTime,
	"original_end_time":  BaseAnnouncementTimeOneTypeOriginalEndTime,
	"report_date":        BaseAnnouncementTimeOneTypeReportDate,
	"start_time":         BaseAnnouncementTimeOneTypeStartTime,
	"time_detected":      BaseAnnouncementTimeOneTypeTimeDetected,
}

// GetBaseAnnouncementTimeOneTypeEnumValues Enumerates the set of values for BaseAnnouncementTimeOneTypeEnum
func GetBaseAnnouncementTimeOneTypeEnumValues() []BaseAnnouncementTimeOneTypeEnum {
	values := make([]BaseAnnouncementTimeOneTypeEnum, 0)
	for _, v := range mappingBaseAnnouncementTimeOneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseAnnouncementTimeOneTypeEnumStringValues Enumerates the set of values in String for BaseAnnouncementTimeOneTypeEnum
func GetBaseAnnouncementTimeOneTypeEnumStringValues() []string {
	return []string{
		"ACTION_REQUIRED_BY",
		"NEW_START_TIME",
		"ORIGINAL_END_TIME",
		"REPORT_DATE",
		"START_TIME",
		"TIME_DETECTED",
	}
}

// GetMappingBaseAnnouncementTimeOneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseAnnouncementTimeOneTypeEnum(val string) (BaseAnnouncementTimeOneTypeEnum, bool) {
	enum, ok := mappingBaseAnnouncementTimeOneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseAnnouncementTimeTwoTypeEnum Enum with underlying type: string
type BaseAnnouncementTimeTwoTypeEnum string

// Set of constants representing the allowable values for BaseAnnouncementTimeTwoTypeEnum
const (
	BaseAnnouncementTimeTwoTypeEndTime          BaseAnnouncementTimeTwoTypeEnum = "END_TIME"
	BaseAnnouncementTimeTwoTypeNewEndTime       BaseAnnouncementTimeTwoTypeEnum = "NEW_END_TIME"
	BaseAnnouncementTimeTwoTypeEstimatedEndTime BaseAnnouncementTimeTwoTypeEnum = "ESTIMATED_END_TIME"
)

var mappingBaseAnnouncementTimeTwoTypeEnum = map[string]BaseAnnouncementTimeTwoTypeEnum{
	"END_TIME":           BaseAnnouncementTimeTwoTypeEndTime,
	"NEW_END_TIME":       BaseAnnouncementTimeTwoTypeNewEndTime,
	"ESTIMATED_END_TIME": BaseAnnouncementTimeTwoTypeEstimatedEndTime,
}

var mappingBaseAnnouncementTimeTwoTypeEnumLowerCase = map[string]BaseAnnouncementTimeTwoTypeEnum{
	"end_time":           BaseAnnouncementTimeTwoTypeEndTime,
	"new_end_time":       BaseAnnouncementTimeTwoTypeNewEndTime,
	"estimated_end_time": BaseAnnouncementTimeTwoTypeEstimatedEndTime,
}

// GetBaseAnnouncementTimeTwoTypeEnumValues Enumerates the set of values for BaseAnnouncementTimeTwoTypeEnum
func GetBaseAnnouncementTimeTwoTypeEnumValues() []BaseAnnouncementTimeTwoTypeEnum {
	values := make([]BaseAnnouncementTimeTwoTypeEnum, 0)
	for _, v := range mappingBaseAnnouncementTimeTwoTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseAnnouncementTimeTwoTypeEnumStringValues Enumerates the set of values in String for BaseAnnouncementTimeTwoTypeEnum
func GetBaseAnnouncementTimeTwoTypeEnumStringValues() []string {
	return []string{
		"END_TIME",
		"NEW_END_TIME",
		"ESTIMATED_END_TIME",
	}
}

// GetMappingBaseAnnouncementTimeTwoTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseAnnouncementTimeTwoTypeEnum(val string) (BaseAnnouncementTimeTwoTypeEnum, bool) {
	enum, ok := mappingBaseAnnouncementTimeTwoTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseAnnouncementAnnouncementTypeEnum Enum with underlying type: string
type BaseAnnouncementAnnouncementTypeEnum string

// Set of constants representing the allowable values for BaseAnnouncementAnnouncementTypeEnum
const (
	BaseAnnouncementAnnouncementTypeActionRecommended               BaseAnnouncementAnnouncementTypeEnum = "ACTION_RECOMMENDED"
	BaseAnnouncementAnnouncementTypeActionRequired                  BaseAnnouncementAnnouncementTypeEnum = "ACTION_REQUIRED"
	BaseAnnouncementAnnouncementTypeEmergencyChange                 BaseAnnouncementAnnouncementTypeEnum = "EMERGENCY_CHANGE"
	BaseAnnouncementAnnouncementTypeEmergencyMaintenance            BaseAnnouncementAnnouncementTypeEnum = "EMERGENCY_MAINTENANCE"
	BaseAnnouncementAnnouncementTypeEmergencyMaintenanceComplete    BaseAnnouncementAnnouncementTypeEnum = "EMERGENCY_MAINTENANCE_COMPLETE"
	BaseAnnouncementAnnouncementTypeEmergencyMaintenanceExtended    BaseAnnouncementAnnouncementTypeEnum = "EMERGENCY_MAINTENANCE_EXTENDED"
	BaseAnnouncementAnnouncementTypeEmergencyMaintenanceRescheduled BaseAnnouncementAnnouncementTypeEnum = "EMERGENCY_MAINTENANCE_RESCHEDULED"
	BaseAnnouncementAnnouncementTypeInformation                     BaseAnnouncementAnnouncementTypeEnum = "INFORMATION"
	BaseAnnouncementAnnouncementTypePlannedChange                   BaseAnnouncementAnnouncementTypeEnum = "PLANNED_CHANGE"
	BaseAnnouncementAnnouncementTypePlannedChangeComplete           BaseAnnouncementAnnouncementTypeEnum = "PLANNED_CHANGE_COMPLETE"
	BaseAnnouncementAnnouncementTypePlannedChangeExtended           BaseAnnouncementAnnouncementTypeEnum = "PLANNED_CHANGE_EXTENDED"
	BaseAnnouncementAnnouncementTypePlannedChangeRescheduled        BaseAnnouncementAnnouncementTypeEnum = "PLANNED_CHANGE_RESCHEDULED"
	BaseAnnouncementAnnouncementTypeProductionEventNotification     BaseAnnouncementAnnouncementTypeEnum = "PRODUCTION_EVENT_NOTIFICATION"
	BaseAnnouncementAnnouncementTypeScheduledMaintenance            BaseAnnouncementAnnouncementTypeEnum = "SCHEDULED_MAINTENANCE"
)

var mappingBaseAnnouncementAnnouncementTypeEnum = map[string]BaseAnnouncementAnnouncementTypeEnum{
	"ACTION_RECOMMENDED":                BaseAnnouncementAnnouncementTypeActionRecommended,
	"ACTION_REQUIRED":                   BaseAnnouncementAnnouncementTypeActionRequired,
	"EMERGENCY_CHANGE":                  BaseAnnouncementAnnouncementTypeEmergencyChange,
	"EMERGENCY_MAINTENANCE":             BaseAnnouncementAnnouncementTypeEmergencyMaintenance,
	"EMERGENCY_MAINTENANCE_COMPLETE":    BaseAnnouncementAnnouncementTypeEmergencyMaintenanceComplete,
	"EMERGENCY_MAINTENANCE_EXTENDED":    BaseAnnouncementAnnouncementTypeEmergencyMaintenanceExtended,
	"EMERGENCY_MAINTENANCE_RESCHEDULED": BaseAnnouncementAnnouncementTypeEmergencyMaintenanceRescheduled,
	"INFORMATION":                       BaseAnnouncementAnnouncementTypeInformation,
	"PLANNED_CHANGE":                    BaseAnnouncementAnnouncementTypePlannedChange,
	"PLANNED_CHANGE_COMPLETE":           BaseAnnouncementAnnouncementTypePlannedChangeComplete,
	"PLANNED_CHANGE_EXTENDED":           BaseAnnouncementAnnouncementTypePlannedChangeExtended,
	"PLANNED_CHANGE_RESCHEDULED":        BaseAnnouncementAnnouncementTypePlannedChangeRescheduled,
	"PRODUCTION_EVENT_NOTIFICATION":     BaseAnnouncementAnnouncementTypeProductionEventNotification,
	"SCHEDULED_MAINTENANCE":             BaseAnnouncementAnnouncementTypeScheduledMaintenance,
}

var mappingBaseAnnouncementAnnouncementTypeEnumLowerCase = map[string]BaseAnnouncementAnnouncementTypeEnum{
	"action_recommended":                BaseAnnouncementAnnouncementTypeActionRecommended,
	"action_required":                   BaseAnnouncementAnnouncementTypeActionRequired,
	"emergency_change":                  BaseAnnouncementAnnouncementTypeEmergencyChange,
	"emergency_maintenance":             BaseAnnouncementAnnouncementTypeEmergencyMaintenance,
	"emergency_maintenance_complete":    BaseAnnouncementAnnouncementTypeEmergencyMaintenanceComplete,
	"emergency_maintenance_extended":    BaseAnnouncementAnnouncementTypeEmergencyMaintenanceExtended,
	"emergency_maintenance_rescheduled": BaseAnnouncementAnnouncementTypeEmergencyMaintenanceRescheduled,
	"information":                       BaseAnnouncementAnnouncementTypeInformation,
	"planned_change":                    BaseAnnouncementAnnouncementTypePlannedChange,
	"planned_change_complete":           BaseAnnouncementAnnouncementTypePlannedChangeComplete,
	"planned_change_extended":           BaseAnnouncementAnnouncementTypePlannedChangeExtended,
	"planned_change_rescheduled":        BaseAnnouncementAnnouncementTypePlannedChangeRescheduled,
	"production_event_notification":     BaseAnnouncementAnnouncementTypeProductionEventNotification,
	"scheduled_maintenance":             BaseAnnouncementAnnouncementTypeScheduledMaintenance,
}

// GetBaseAnnouncementAnnouncementTypeEnumValues Enumerates the set of values for BaseAnnouncementAnnouncementTypeEnum
func GetBaseAnnouncementAnnouncementTypeEnumValues() []BaseAnnouncementAnnouncementTypeEnum {
	values := make([]BaseAnnouncementAnnouncementTypeEnum, 0)
	for _, v := range mappingBaseAnnouncementAnnouncementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseAnnouncementAnnouncementTypeEnumStringValues Enumerates the set of values in String for BaseAnnouncementAnnouncementTypeEnum
func GetBaseAnnouncementAnnouncementTypeEnumStringValues() []string {
	return []string{
		"ACTION_RECOMMENDED",
		"ACTION_REQUIRED",
		"EMERGENCY_CHANGE",
		"EMERGENCY_MAINTENANCE",
		"EMERGENCY_MAINTENANCE_COMPLETE",
		"EMERGENCY_MAINTENANCE_EXTENDED",
		"EMERGENCY_MAINTENANCE_RESCHEDULED",
		"INFORMATION",
		"PLANNED_CHANGE",
		"PLANNED_CHANGE_COMPLETE",
		"PLANNED_CHANGE_EXTENDED",
		"PLANNED_CHANGE_RESCHEDULED",
		"PRODUCTION_EVENT_NOTIFICATION",
		"SCHEDULED_MAINTENANCE",
	}
}

// GetMappingBaseAnnouncementAnnouncementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseAnnouncementAnnouncementTypeEnum(val string) (BaseAnnouncementAnnouncementTypeEnum, bool) {
	enum, ok := mappingBaseAnnouncementAnnouncementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseAnnouncementLifecycleStateEnum Enum with underlying type: string
type BaseAnnouncementLifecycleStateEnum string

// Set of constants representing the allowable values for BaseAnnouncementLifecycleStateEnum
const (
	BaseAnnouncementLifecycleStateActive   BaseAnnouncementLifecycleStateEnum = "ACTIVE"
	BaseAnnouncementLifecycleStateInactive BaseAnnouncementLifecycleStateEnum = "INACTIVE"
)

var mappingBaseAnnouncementLifecycleStateEnum = map[string]BaseAnnouncementLifecycleStateEnum{
	"ACTIVE":   BaseAnnouncementLifecycleStateActive,
	"INACTIVE": BaseAnnouncementLifecycleStateInactive,
}

var mappingBaseAnnouncementLifecycleStateEnumLowerCase = map[string]BaseAnnouncementLifecycleStateEnum{
	"active":   BaseAnnouncementLifecycleStateActive,
	"inactive": BaseAnnouncementLifecycleStateInactive,
}

// GetBaseAnnouncementLifecycleStateEnumValues Enumerates the set of values for BaseAnnouncementLifecycleStateEnum
func GetBaseAnnouncementLifecycleStateEnumValues() []BaseAnnouncementLifecycleStateEnum {
	values := make([]BaseAnnouncementLifecycleStateEnum, 0)
	for _, v := range mappingBaseAnnouncementLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseAnnouncementLifecycleStateEnumStringValues Enumerates the set of values in String for BaseAnnouncementLifecycleStateEnum
func GetBaseAnnouncementLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingBaseAnnouncementLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseAnnouncementLifecycleStateEnum(val string) (BaseAnnouncementLifecycleStateEnum, bool) {
	enum, ok := mappingBaseAnnouncementLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseAnnouncementPlatformTypeEnum Enum with underlying type: string
type BaseAnnouncementPlatformTypeEnum string

// Set of constants representing the allowable values for BaseAnnouncementPlatformTypeEnum
const (
	BaseAnnouncementPlatformTypeIaas BaseAnnouncementPlatformTypeEnum = "IAAS"
	BaseAnnouncementPlatformTypeSaas BaseAnnouncementPlatformTypeEnum = "SAAS"
)

var mappingBaseAnnouncementPlatformTypeEnum = map[string]BaseAnnouncementPlatformTypeEnum{
	"IAAS": BaseAnnouncementPlatformTypeIaas,
	"SAAS": BaseAnnouncementPlatformTypeSaas,
}

var mappingBaseAnnouncementPlatformTypeEnumLowerCase = map[string]BaseAnnouncementPlatformTypeEnum{
	"iaas": BaseAnnouncementPlatformTypeIaas,
	"saas": BaseAnnouncementPlatformTypeSaas,
}

// GetBaseAnnouncementPlatformTypeEnumValues Enumerates the set of values for BaseAnnouncementPlatformTypeEnum
func GetBaseAnnouncementPlatformTypeEnumValues() []BaseAnnouncementPlatformTypeEnum {
	values := make([]BaseAnnouncementPlatformTypeEnum, 0)
	for _, v := range mappingBaseAnnouncementPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseAnnouncementPlatformTypeEnumStringValues Enumerates the set of values in String for BaseAnnouncementPlatformTypeEnum
func GetBaseAnnouncementPlatformTypeEnumStringValues() []string {
	return []string{
		"IAAS",
		"SAAS",
	}
}

// GetMappingBaseAnnouncementPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseAnnouncementPlatformTypeEnum(val string) (BaseAnnouncementPlatformTypeEnum, bool) {
	enum, ok := mappingBaseAnnouncementPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
