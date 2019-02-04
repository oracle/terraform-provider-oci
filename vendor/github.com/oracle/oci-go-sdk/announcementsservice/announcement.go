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

// Announcement An announcement object which represents a message targetted to a specific tenant
type Announcement struct {

	// The OCID of the announcement
	Id *string `mandatory:"true" json:"id"`

	// The reference JIRA ticket number
	ReferenceTicketNumber *string `mandatory:"true" json:"referenceTicketNumber"`

	// Forms part of the email subject and/or the console representation (a banner or alike)
	Summary *string `mandatory:"true" json:"summary"`

	// Show announcement as a banner
	IsBanner *bool `mandatory:"true" json:"isBanner"`

	// The title of the first time value, e.g. Time Started
	TimeOneTitle *string `mandatory:"false" json:"timeOneTitle"`

	// The first time value, actual meaning depending on notification type
	TimeOneValue *common.SDKTime `mandatory:"false" json:"timeOneValue"`

	// The title of the second time value, e.g. Time Ended
	TimeTwoTitle *string `mandatory:"false" json:"timeTwoTitle"`

	// The second time value, actual meaning depending on notification type
	TimeTwoValue *common.SDKTime `mandatory:"false" json:"timeTwoValue"`

	// Impacted services
	Services []string `mandatory:"false" json:"services"`

	// Impacted regions
	AffectedRegions []string `mandatory:"false" json:"affectedRegions"`

	// The date and time the announcement was created, in the format defined by RFC3339
	// Example: `2016-07-22T17:43:01.389+0000`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the announcement was last updated, in the format defined by RFC3339
	// Example: `2016-07-22T17:43:01.389+0000`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A more detailed explanation of the notification. A markdown format input
	Description *string `mandatory:"false" json:"description"`

	// A markdown format input that forms e.g. the FAQ section of a notification
	AdditionalInformation *string `mandatory:"false" json:"additionalInformation"`

	Followups []NotificationFollowupDetails `mandatory:"false" json:"followups"`

	// List of resources (possibly empty) affected by this announcement
	AffectedResources []AffectedResource `mandatory:"false" json:"affectedResources"`

	// The detailed description of an announcement
	AnnouncementType BaseAnnouncementAnnouncementTypeEnum `mandatory:"true" json:"announcementType"`

	// Lifecycle states of announcement
	LifecycleState BaseAnnouncementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m Announcement) GetId() *string {
	return m.Id
}

//GetReferenceTicketNumber returns ReferenceTicketNumber
func (m Announcement) GetReferenceTicketNumber() *string {
	return m.ReferenceTicketNumber
}

//GetSummary returns Summary
func (m Announcement) GetSummary() *string {
	return m.Summary
}

//GetTimeOneTitle returns TimeOneTitle
func (m Announcement) GetTimeOneTitle() *string {
	return m.TimeOneTitle
}

//GetTimeOneValue returns TimeOneValue
func (m Announcement) GetTimeOneValue() *common.SDKTime {
	return m.TimeOneValue
}

//GetTimeTwoTitle returns TimeTwoTitle
func (m Announcement) GetTimeTwoTitle() *string {
	return m.TimeTwoTitle
}

//GetTimeTwoValue returns TimeTwoValue
func (m Announcement) GetTimeTwoValue() *common.SDKTime {
	return m.TimeTwoValue
}

//GetServices returns Services
func (m Announcement) GetServices() []string {
	return m.Services
}

//GetAffectedRegions returns AffectedRegions
func (m Announcement) GetAffectedRegions() []string {
	return m.AffectedRegions
}

//GetAnnouncementType returns AnnouncementType
func (m Announcement) GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum {
	return m.AnnouncementType
}

//GetLifecycleState returns LifecycleState
func (m Announcement) GetLifecycleState() BaseAnnouncementLifecycleStateEnum {
	return m.LifecycleState
}

//GetIsBanner returns IsBanner
func (m Announcement) GetIsBanner() *bool {
	return m.IsBanner
}

//GetTimeCreated returns TimeCreated
func (m Announcement) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m Announcement) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m Announcement) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Announcement) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnnouncement Announcement
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAnnouncement
	}{
		"Announcement",
		(MarshalTypeAnnouncement)(m),
	}

	return json.Marshal(&s)
}
