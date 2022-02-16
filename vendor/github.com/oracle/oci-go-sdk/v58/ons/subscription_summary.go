// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SubscriptionSummary The subscription's configuration summary.
type SubscriptionSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
	TopicId *string `mandatory:"true" json:"topicId"`

	// The protocol used for the subscription.
	// Allowed values:
	//   * `CUSTOM_HTTPS`
	//   * `EMAIL`
	//   * `HTTPS` (deprecated; for PagerDuty endpoints, use `PAGERDUTY`)
	//   * `ORACLE_FUNCTIONS`
	//   * `PAGERDUTY`
	//   * `SLACK`
	//   * `SMS`
	// For information about subscription protocols, see
	// To create a subscription (https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#createSub).
	Protocol *string `mandatory:"true" json:"protocol"`

	// A locator that corresponds to the subscription protocol.
	// For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE.
	LifecycleState SubscriptionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this suscription was created.
	CreatedTime *int64 `mandatory:"false" json:"createdTime"`

	DeliveryPolicy *DeliveryPolicy `mandatory:"false" json:"deliveryPolicy"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `mandatory:"false" json:"etag"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubscriptionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriptionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubscriptionSummaryLifecycleStateEnum Enum with underlying type: string
type SubscriptionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionSummaryLifecycleStateEnum
const (
	SubscriptionSummaryLifecycleStatePending SubscriptionSummaryLifecycleStateEnum = "PENDING"
	SubscriptionSummaryLifecycleStateActive  SubscriptionSummaryLifecycleStateEnum = "ACTIVE"
	SubscriptionSummaryLifecycleStateDeleted SubscriptionSummaryLifecycleStateEnum = "DELETED"
)

var mappingSubscriptionSummaryLifecycleStateEnum = map[string]SubscriptionSummaryLifecycleStateEnum{
	"PENDING": SubscriptionSummaryLifecycleStatePending,
	"ACTIVE":  SubscriptionSummaryLifecycleStateActive,
	"DELETED": SubscriptionSummaryLifecycleStateDeleted,
}

// GetSubscriptionSummaryLifecycleStateEnumValues Enumerates the set of values for SubscriptionSummaryLifecycleStateEnum
func GetSubscriptionSummaryLifecycleStateEnumValues() []SubscriptionSummaryLifecycleStateEnum {
	values := make([]SubscriptionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SubscriptionSummaryLifecycleStateEnum
func GetSubscriptionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PENDING",
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingSubscriptionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionSummaryLifecycleStateEnum(val string) (SubscriptionSummaryLifecycleStateEnum, bool) {
	mappingSubscriptionSummaryLifecycleStateEnumIgnoreCase := make(map[string]SubscriptionSummaryLifecycleStateEnum)
	for k, v := range mappingSubscriptionSummaryLifecycleStateEnum {
		mappingSubscriptionSummaryLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSubscriptionSummaryLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
