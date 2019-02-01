// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// A description of the AnnouncementsService API
//

package announcementsservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// BaseAnnouncement Base for announcements and incidents
type BaseAnnouncement interface {

	// The OCID of the announcement
	GetId() *string

	// The reference JIRA ticket number
	GetReferenceTicketNumber() *string

	// Forms part of the email subject and/or the console representation (a banner or alike)
	GetSummary() *string

	// The detailed description of an announcement
	GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum

	// Lifecycle states of announcement
	GetLifecycleState() BaseAnnouncementLifecycleStateEnum

	// Show announcement as a banner
	GetIsBanner() *bool

	// The title of the first time value, e.g. Time Started
	GetTimeOneTitle() *string

	// The first time value, actual meaning depending on notification type
	GetTimeOneValue() *common.SDKTime

	// The title of the second time value, e.g. Time Ended
	GetTimeTwoTitle() *string

	// The second time value, actual meaning depending on notification type
	GetTimeTwoValue() *common.SDKTime

	// Impacted services
	GetServices() []string

	// Impacted regions
	GetAffectedRegions() []string

	// The date and time the announcement was created, in the format defined by RFC3339
	// Example: `2016-07-22T17:43:01.389+0000`
	GetTimeCreated() *common.SDKTime

	// The date and time the announcement was last updated, in the format defined by RFC3339
	// Example: `2016-07-22T17:43:01.389+0000`
	GetTimeUpdated() *common.SDKTime
}

type baseannouncement struct {
	JsonData              []byte
	Id                    *string                              `mandatory:"true" json:"id"`
	ReferenceTicketNumber *string                              `mandatory:"true" json:"referenceTicketNumber"`
	Summary               *string                              `mandatory:"true" json:"summary"`
	AnnouncementType      BaseAnnouncementAnnouncementTypeEnum `mandatory:"true" json:"announcementType"`
	LifecycleState        BaseAnnouncementLifecycleStateEnum   `mandatory:"true" json:"lifecycleState"`
	IsBanner              *bool                                `mandatory:"true" json:"isBanner"`
	TimeOneTitle          *string                              `mandatory:"false" json:"timeOneTitle"`
	TimeOneValue          *common.SDKTime                      `mandatory:"false" json:"timeOneValue"`
	TimeTwoTitle          *string                              `mandatory:"false" json:"timeTwoTitle"`
	TimeTwoValue          *common.SDKTime                      `mandatory:"false" json:"timeTwoValue"`
	Services              []string                             `mandatory:"false" json:"services"`
	AffectedRegions       []string                             `mandatory:"false" json:"affectedRegions"`
	TimeCreated           *common.SDKTime                      `mandatory:"false" json:"timeCreated"`
	TimeUpdated           *common.SDKTime                      `mandatory:"false" json:"timeUpdated"`
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
	m.AnnouncementType = s.Model.AnnouncementType
	m.LifecycleState = s.Model.LifecycleState
	m.IsBanner = s.Model.IsBanner
	m.TimeOneTitle = s.Model.TimeOneTitle
	m.TimeOneValue = s.Model.TimeOneValue
	m.TimeTwoTitle = s.Model.TimeTwoTitle
	m.TimeTwoValue = s.Model.TimeTwoValue
	m.Services = s.Model.Services
	m.AffectedRegions = s.Model.AffectedRegions
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
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
		return *m, nil
	}
}

//GetId returns Id
func (m baseannouncement) GetId() *string {
	return m.Id
}

//GetReferenceTicketNumber returns ReferenceTicketNumber
func (m baseannouncement) GetReferenceTicketNumber() *string {
	return m.ReferenceTicketNumber
}

//GetSummary returns Summary
func (m baseannouncement) GetSummary() *string {
	return m.Summary
}

//GetAnnouncementType returns AnnouncementType
func (m baseannouncement) GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum {
	return m.AnnouncementType
}

//GetLifecycleState returns LifecycleState
func (m baseannouncement) GetLifecycleState() BaseAnnouncementLifecycleStateEnum {
	return m.LifecycleState
}

//GetIsBanner returns IsBanner
func (m baseannouncement) GetIsBanner() *bool {
	return m.IsBanner
}

//GetTimeOneTitle returns TimeOneTitle
func (m baseannouncement) GetTimeOneTitle() *string {
	return m.TimeOneTitle
}

//GetTimeOneValue returns TimeOneValue
func (m baseannouncement) GetTimeOneValue() *common.SDKTime {
	return m.TimeOneValue
}

//GetTimeTwoTitle returns TimeTwoTitle
func (m baseannouncement) GetTimeTwoTitle() *string {
	return m.TimeTwoTitle
}

//GetTimeTwoValue returns TimeTwoValue
func (m baseannouncement) GetTimeTwoValue() *common.SDKTime {
	return m.TimeTwoValue
}

//GetServices returns Services
func (m baseannouncement) GetServices() []string {
	return m.Services
}

//GetAffectedRegions returns AffectedRegions
func (m baseannouncement) GetAffectedRegions() []string {
	return m.AffectedRegions
}

//GetTimeCreated returns TimeCreated
func (m baseannouncement) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m baseannouncement) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m baseannouncement) String() string {
	return common.PointerString(m)
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

var mappingBaseAnnouncementAnnouncementType = map[string]BaseAnnouncementAnnouncementTypeEnum{
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

// GetBaseAnnouncementAnnouncementTypeEnumValues Enumerates the set of values for BaseAnnouncementAnnouncementTypeEnum
func GetBaseAnnouncementAnnouncementTypeEnumValues() []BaseAnnouncementAnnouncementTypeEnum {
	values := make([]BaseAnnouncementAnnouncementTypeEnum, 0)
	for _, v := range mappingBaseAnnouncementAnnouncementType {
		values = append(values, v)
	}
	return values
}

// BaseAnnouncementLifecycleStateEnum Enum with underlying type: string
type BaseAnnouncementLifecycleStateEnum string

// Set of constants representing the allowable values for BaseAnnouncementLifecycleStateEnum
const (
	BaseAnnouncementLifecycleStateActive   BaseAnnouncementLifecycleStateEnum = "ACTIVE"
	BaseAnnouncementLifecycleStateInactive BaseAnnouncementLifecycleStateEnum = "INACTIVE"
)

var mappingBaseAnnouncementLifecycleState = map[string]BaseAnnouncementLifecycleStateEnum{
	"ACTIVE":   BaseAnnouncementLifecycleStateActive,
	"INACTIVE": BaseAnnouncementLifecycleStateInactive,
}

// GetBaseAnnouncementLifecycleStateEnumValues Enumerates the set of values for BaseAnnouncementLifecycleStateEnum
func GetBaseAnnouncementLifecycleStateEnumValues() []BaseAnnouncementLifecycleStateEnum {
	values := make([]BaseAnnouncementLifecycleStateEnum, 0)
	for _, v := range mappingBaseAnnouncementLifecycleState {
		values = append(values, v)
	}
	return values
}
