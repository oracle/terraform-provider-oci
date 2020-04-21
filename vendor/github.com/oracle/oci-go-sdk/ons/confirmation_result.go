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

// ConfirmationResult The confirmation details for the specified subscription.
// For information about confirming subscriptions, see
// To confirm a subscription (https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm#confirmSub).
type ConfirmationResult struct {

	// The name of the subscribed topic.
	TopicName *string `mandatory:"true" json:"topicName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic associated with the specified subscription.
	TopicId *string `mandatory:"true" json:"topicId"`

	// A locator that corresponds to the subscription protocol.
	// For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// The URL for unsubscribing from the topic.
	UnsubscribeUrl *string `mandatory:"true" json:"unsubscribeUrl"`

	// A human-readable string indicating the status of the subscription confirmation.
	Message *string `mandatory:"true" json:"message"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription specified in the request.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`
}

func (m ConfirmationResult) String() string {
	return common.PointerString(m)
}
