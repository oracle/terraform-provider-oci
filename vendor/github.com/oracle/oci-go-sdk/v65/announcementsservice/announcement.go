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

// Announcement A message about an impactful operational event.
type Announcement struct {

	// The OCID of the announcement.
	Id *string `mandatory:"true" json:"id"`

	// The reference Jira ticket number.
	ReferenceTicketNumber *string `mandatory:"true" json:"referenceTicketNumber"`

	// A summary of the issue. A summary might appear in the console banner view of the announcement or in
	// an email subject line. Avoid entering confidential information.
	Summary *string `mandatory:"true" json:"summary"`

	// Impacted Oracle Cloud Infrastructure services.
	Services []string `mandatory:"true" json:"services"`

	// Impacted regions.
	AffectedRegions []string `mandatory:"true" json:"affectedRegions"`

	// Whether the announcement is displayed as a banner in the console.
	IsBanner *bool `mandatory:"true" json:"isBanner"`

	// The label associated with an initial time value.
	// Example: `Time Started`
	TimeOneTitle *string `mandatory:"false" json:"timeOneTitle"`

	// The actual value of the first time value for the event. Typically, this denotes the time an event started, but the meaning
	// can vary, depending on the announcement type. The `timeOneType` attribute describes the meaning.
	TimeOneValue *common.SDKTime `mandatory:"false" json:"timeOneValue"`

	// The label associated with a second time value.
	// Example: `Time Ended`
	TimeTwoTitle *string `mandatory:"false" json:"timeTwoTitle"`

	// The actual value of the second time value. Typically, this denotes the time an event ended, but the meaning
	// can vary, depending on the announcement type. The `timeTwoType` attribute describes the meaning.
	TimeTwoValue *common.SDKTime `mandatory:"false" json:"timeTwoValue"`

	// The date and time the announcement was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-01-01T17:43:01.389+0000`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the announcement was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-01-01T17:43:01.389+0000`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The name of the environment that this announcement pertains to.
	EnvironmentName *string `mandatory:"false" json:"environmentName"`

	// The sequence of connected announcements, or announcement chain, that this announcement belongs to. Related announcements share the same chain ID.
	ChainId *string `mandatory:"false" json:"chainId"`

	// A detailed explanation of the event, expressed by using Markdown language. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Additional information about the event, expressed by using Markdown language and included in the
	// details view of an announcement. Additional information might include remediation steps or
	// answers to frequently asked questions. Avoid entering confidential information.
	AdditionalInformation *string `mandatory:"false" json:"additionalInformation"`

	// The list of resources, if any, affected by the event described in the announcement.
	AffectedResources []AffectedResource `mandatory:"false" json:"affectedResources"`

	// The type of a time associated with an initial time value. If the `timeOneTitle` attribute is present, then the `timeOneTitle` attribute contains a label of `timeOneType` in English.
	// Example: `START_TIME`
	TimeOneType BaseAnnouncementTimeOneTypeEnum `mandatory:"false" json:"timeOneType,omitempty"`

	// The type of a time associated with second time value. If the `timeTwoTitle` attribute is present, then the `timeTwoTitle` attribute contains a label of `timeTwoType` in English.
	// Example: `END_TIME`
	TimeTwoType BaseAnnouncementTimeTwoTypeEnum `mandatory:"false" json:"timeTwoType,omitempty"`

	// The type of announcement. An announcement's type signals its severity.
	AnnouncementType BaseAnnouncementAnnouncementTypeEnum `mandatory:"true" json:"announcementType"`

	// The current lifecycle state of the announcement.
	LifecycleState BaseAnnouncementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The platform type that this announcement pertains to.
	PlatformType BaseAnnouncementPlatformTypeEnum `mandatory:"false" json:"platformType,omitempty"`
}

// GetId returns Id
func (m Announcement) GetId() *string {
	return m.Id
}

// GetReferenceTicketNumber returns ReferenceTicketNumber
func (m Announcement) GetReferenceTicketNumber() *string {
	return m.ReferenceTicketNumber
}

// GetSummary returns Summary
func (m Announcement) GetSummary() *string {
	return m.Summary
}

// GetTimeOneTitle returns TimeOneTitle
func (m Announcement) GetTimeOneTitle() *string {
	return m.TimeOneTitle
}

// GetTimeOneType returns TimeOneType
func (m Announcement) GetTimeOneType() BaseAnnouncementTimeOneTypeEnum {
	return m.TimeOneType
}

// GetTimeOneValue returns TimeOneValue
func (m Announcement) GetTimeOneValue() *common.SDKTime {
	return m.TimeOneValue
}

// GetTimeTwoTitle returns TimeTwoTitle
func (m Announcement) GetTimeTwoTitle() *string {
	return m.TimeTwoTitle
}

// GetTimeTwoType returns TimeTwoType
func (m Announcement) GetTimeTwoType() BaseAnnouncementTimeTwoTypeEnum {
	return m.TimeTwoType
}

// GetTimeTwoValue returns TimeTwoValue
func (m Announcement) GetTimeTwoValue() *common.SDKTime {
	return m.TimeTwoValue
}

// GetServices returns Services
func (m Announcement) GetServices() []string {
	return m.Services
}

// GetAffectedRegions returns AffectedRegions
func (m Announcement) GetAffectedRegions() []string {
	return m.AffectedRegions
}

// GetAnnouncementType returns AnnouncementType
func (m Announcement) GetAnnouncementType() BaseAnnouncementAnnouncementTypeEnum {
	return m.AnnouncementType
}

// GetLifecycleState returns LifecycleState
func (m Announcement) GetLifecycleState() BaseAnnouncementLifecycleStateEnum {
	return m.LifecycleState
}

// GetIsBanner returns IsBanner
func (m Announcement) GetIsBanner() *bool {
	return m.IsBanner
}

// GetTimeCreated returns TimeCreated
func (m Announcement) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m Announcement) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetEnvironmentName returns EnvironmentName
func (m Announcement) GetEnvironmentName() *string {
	return m.EnvironmentName
}

// GetPlatformType returns PlatformType
func (m Announcement) GetPlatformType() BaseAnnouncementPlatformTypeEnum {
	return m.PlatformType
}

// GetChainId returns ChainId
func (m Announcement) GetChainId() *string {
	return m.ChainId
}

func (m Announcement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Announcement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBaseAnnouncementTimeOneTypeEnum(string(m.TimeOneType)); !ok && m.TimeOneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TimeOneType: %s. Supported values are: %s.", m.TimeOneType, strings.Join(GetBaseAnnouncementTimeOneTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementTimeTwoTypeEnum(string(m.TimeTwoType)); !ok && m.TimeTwoType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TimeTwoType: %s. Supported values are: %s.", m.TimeTwoType, strings.Join(GetBaseAnnouncementTimeTwoTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementAnnouncementTypeEnum(string(m.AnnouncementType)); !ok && m.AnnouncementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AnnouncementType: %s. Supported values are: %s.", m.AnnouncementType, strings.Join(GetBaseAnnouncementAnnouncementTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseAnnouncementLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseAnnouncementPlatformTypeEnum(string(m.PlatformType)); !ok && m.PlatformType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", m.PlatformType, strings.Join(GetBaseAnnouncementPlatformTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
