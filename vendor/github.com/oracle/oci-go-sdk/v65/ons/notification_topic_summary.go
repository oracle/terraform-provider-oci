// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotificationTopicSummary A summary of the properties that define a topic.
type NotificationTopicSummary struct {

	// The name of the topic.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the topic.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of the topic.
	LifecycleState NotificationTopicSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the topic was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The endpoint for managing subscriptions or publishing messages to the topic.
	ApiEndpoint *string `mandatory:"true" json:"apiEndpoint"`

	// A unique short topic Id. This is used only for SMS subscriptions.
	ShortTopicId *string `mandatory:"false" json:"shortTopicId"`

	// The description of the topic.
	Description *string `mandatory:"false" json:"description"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NotificationTopicSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationTopicSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotificationTopicSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNotificationTopicSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NotificationTopicSummaryLifecycleStateEnum Enum with underlying type: string
type NotificationTopicSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for NotificationTopicSummaryLifecycleStateEnum
const (
	NotificationTopicSummaryLifecycleStateActive   NotificationTopicSummaryLifecycleStateEnum = "ACTIVE"
	NotificationTopicSummaryLifecycleStateDeleting NotificationTopicSummaryLifecycleStateEnum = "DELETING"
	NotificationTopicSummaryLifecycleStateCreating NotificationTopicSummaryLifecycleStateEnum = "CREATING"
)

var mappingNotificationTopicSummaryLifecycleStateEnum = map[string]NotificationTopicSummaryLifecycleStateEnum{
	"ACTIVE":   NotificationTopicSummaryLifecycleStateActive,
	"DELETING": NotificationTopicSummaryLifecycleStateDeleting,
	"CREATING": NotificationTopicSummaryLifecycleStateCreating,
}

var mappingNotificationTopicSummaryLifecycleStateEnumLowerCase = map[string]NotificationTopicSummaryLifecycleStateEnum{
	"active":   NotificationTopicSummaryLifecycleStateActive,
	"deleting": NotificationTopicSummaryLifecycleStateDeleting,
	"creating": NotificationTopicSummaryLifecycleStateCreating,
}

// GetNotificationTopicSummaryLifecycleStateEnumValues Enumerates the set of values for NotificationTopicSummaryLifecycleStateEnum
func GetNotificationTopicSummaryLifecycleStateEnumValues() []NotificationTopicSummaryLifecycleStateEnum {
	values := make([]NotificationTopicSummaryLifecycleStateEnum, 0)
	for _, v := range mappingNotificationTopicSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationTopicSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for NotificationTopicSummaryLifecycleStateEnum
func GetNotificationTopicSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"CREATING",
	}
}

// GetMappingNotificationTopicSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationTopicSummaryLifecycleStateEnum(val string) (NotificationTopicSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingNotificationTopicSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
