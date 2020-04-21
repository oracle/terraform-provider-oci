// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Subscription The subscription's configuration. For general information about subscriptions, see
// Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
type Subscription struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// The protocol used for the subscription.
	// For information about subscription protocols, see
	// To create a subscription (https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub).
	Protocol *string `mandatory:"true" json:"protocol"`

	// A locator that corresponds to the subscription protocol.
	// For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this suscription was created.
	CreatedTime *int64 `mandatory:"false" json:"createdTime"`

	// The delivery policy of the subscription. Stored as a JSON string.
	DeliverPolicy *string `mandatory:"false" json:"deliverPolicy"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleStateEnum
const (
	SubscriptionLifecycleStatePending SubscriptionLifecycleStateEnum = "PENDING"
	SubscriptionLifecycleStateActive  SubscriptionLifecycleStateEnum = "ACTIVE"
	SubscriptionLifecycleStateDeleted SubscriptionLifecycleStateEnum = "DELETED"
)

var mappingSubscriptionLifecycleState = map[string]SubscriptionLifecycleStateEnum{
	"PENDING": SubscriptionLifecycleStatePending,
	"ACTIVE":  SubscriptionLifecycleStateActive,
	"DELETED": SubscriptionLifecycleStateDeleted,
}

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleState {
		values = append(values, v)
	}
	return values
}
