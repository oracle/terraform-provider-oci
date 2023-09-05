// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConsumerGroupsCapabilityDetails Specifies the details for the consumer group capability.
type ConsumerGroupsCapabilityDetails struct {

	// Specifies if the primary consumer group should be automatically enabled after adding the capability.
	IsPrimaryConsumerGroupEnabled *bool `mandatory:"false" json:"isPrimaryConsumerGroupEnabled"`

	// Name of the primary consumer group. If omitted, it will be named "Primary Consumer Group".
	PrimaryConsumerGroupDisplayName *string `mandatory:"false" json:"primaryConsumerGroupDisplayName"`

	// The filter used by the primary consumer group. Only messages matching the filter will be available by consumers of the group.
	// An empty value means that all messages will be available in the group.
	PrimaryConsumerGroupFilter *string `mandatory:"false" json:"primaryConsumerGroupFilter"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue.
	// A value of 0 indicates that the DLQ is not used.
	// If the value isn't set, it will be using the value defined at the queue level.
	PrimaryConsumerGroupDeadLetterQueueDeliveryCount *int `mandatory:"false" json:"primaryConsumerGroupDeadLetterQueueDeliveryCount"`
}

func (m ConsumerGroupsCapabilityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsumerGroupsCapabilityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConsumerGroupsCapabilityDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConsumerGroupsCapabilityDetails ConsumerGroupsCapabilityDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeConsumerGroupsCapabilityDetails
	}{
		"CONSUMER_GROUPS",
		(MarshalTypeConsumerGroupsCapabilityDetails)(m),
	}

	return json.Marshal(&s)
}
