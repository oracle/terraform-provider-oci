// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotificationTopic The properties that define a topic. For general information about topics, see
// Notifications Overview (https://docs.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
type NotificationTopic struct {

	// The name of the topic.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the topic.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The lifecycle state of the topic.
	LifecycleState NotificationTopicLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NotificationTopic) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationTopic) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotificationTopicLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNotificationTopicLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NotificationTopicLifecycleStateEnum Enum with underlying type: string
type NotificationTopicLifecycleStateEnum string

// Set of constants representing the allowable values for NotificationTopicLifecycleStateEnum
const (
	NotificationTopicLifecycleStateActive   NotificationTopicLifecycleStateEnum = "ACTIVE"
	NotificationTopicLifecycleStateDeleting NotificationTopicLifecycleStateEnum = "DELETING"
	NotificationTopicLifecycleStateCreating NotificationTopicLifecycleStateEnum = "CREATING"
)

var mappingNotificationTopicLifecycleStateEnum = map[string]NotificationTopicLifecycleStateEnum{
	"ACTIVE":   NotificationTopicLifecycleStateActive,
	"DELETING": NotificationTopicLifecycleStateDeleting,
	"CREATING": NotificationTopicLifecycleStateCreating,
}

var mappingNotificationTopicLifecycleStateEnumLowerCase = map[string]NotificationTopicLifecycleStateEnum{
	"active":   NotificationTopicLifecycleStateActive,
	"deleting": NotificationTopicLifecycleStateDeleting,
	"creating": NotificationTopicLifecycleStateCreating,
}

// GetNotificationTopicLifecycleStateEnumValues Enumerates the set of values for NotificationTopicLifecycleStateEnum
func GetNotificationTopicLifecycleStateEnumValues() []NotificationTopicLifecycleStateEnum {
	values := make([]NotificationTopicLifecycleStateEnum, 0)
	for _, v := range mappingNotificationTopicLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationTopicLifecycleStateEnumStringValues Enumerates the set of values in String for NotificationTopicLifecycleStateEnum
func GetNotificationTopicLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"CREATING",
	}
}

// GetMappingNotificationTopicLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationTopicLifecycleStateEnum(val string) (NotificationTopicLifecycleStateEnum, bool) {
	enum, ok := mappingNotificationTopicLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
