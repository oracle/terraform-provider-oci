// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnnouncementSubscription A subscription with the Announcements service to receive selected announcements in the format and delivery mechanisms supported by a corresponding topic endpoint configured in the Oracle Cloud Infrastructure Notifications service.
type AnnouncementSubscription struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the announcement subscription.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the announcement subscription. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the announcement subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the announcement subscription was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the announcement subscription.
	LifecycleState AnnouncementSubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the Notifications service topic that is the target for publishing announcements that match the configured announcement subscription.
	OnsTopicId *string `mandatory:"true" json:"onsTopicId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A description of the announcement subscription. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the announcement subscription was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current lifecycle state in more detail. For example, details might provide required or recommended actions for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A list of filter groups for the announcement subscription. A filter group is a combination of multiple filters applied to announcements for matching purposes.
	FilterGroups map[string]FilterGroup `mandatory:"false" json:"filterGroups"`

	// (For announcement subscriptions with SaaS configured as the platform type or Oracle Fusion Applications as the service, or both, only) The language in which the user prefers to receive emailed announcements. Specify the preference with a value that uses the x-obmcs-human-language format. For example fr-FR.
	PreferredLanguage *string `mandatory:"false" json:"preferredLanguage"`

	// The time zone in which the user prefers to receive announcements. Specify the preference with a value that uses the IANA Time Zone Database format (x-obmcs-time-zone). For example - America/Los_Angeles
	PreferredTimeZone *string `mandatory:"false" json:"preferredTimeZone"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AnnouncementSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnnouncementSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnnouncementSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAnnouncementSubscriptionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnnouncementSubscriptionLifecycleStateEnum Enum with underlying type: string
type AnnouncementSubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for AnnouncementSubscriptionLifecycleStateEnum
const (
	AnnouncementSubscriptionLifecycleStateActive  AnnouncementSubscriptionLifecycleStateEnum = "ACTIVE"
	AnnouncementSubscriptionLifecycleStateDeleted AnnouncementSubscriptionLifecycleStateEnum = "DELETED"
	AnnouncementSubscriptionLifecycleStateFailed  AnnouncementSubscriptionLifecycleStateEnum = "FAILED"
)

var mappingAnnouncementSubscriptionLifecycleStateEnum = map[string]AnnouncementSubscriptionLifecycleStateEnum{
	"ACTIVE":  AnnouncementSubscriptionLifecycleStateActive,
	"DELETED": AnnouncementSubscriptionLifecycleStateDeleted,
	"FAILED":  AnnouncementSubscriptionLifecycleStateFailed,
}

var mappingAnnouncementSubscriptionLifecycleStateEnumLowerCase = map[string]AnnouncementSubscriptionLifecycleStateEnum{
	"active":  AnnouncementSubscriptionLifecycleStateActive,
	"deleted": AnnouncementSubscriptionLifecycleStateDeleted,
	"failed":  AnnouncementSubscriptionLifecycleStateFailed,
}

// GetAnnouncementSubscriptionLifecycleStateEnumValues Enumerates the set of values for AnnouncementSubscriptionLifecycleStateEnum
func GetAnnouncementSubscriptionLifecycleStateEnumValues() []AnnouncementSubscriptionLifecycleStateEnum {
	values := make([]AnnouncementSubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingAnnouncementSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAnnouncementSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for AnnouncementSubscriptionLifecycleStateEnum
func GetAnnouncementSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAnnouncementSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnnouncementSubscriptionLifecycleStateEnum(val string) (AnnouncementSubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingAnnouncementSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
