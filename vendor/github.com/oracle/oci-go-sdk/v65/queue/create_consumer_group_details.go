// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.cloud.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateConsumerGroupDetails The information about a new consumer group.
type CreateConsumerGroupDetails struct {

	// The user-friendly name of the consumer group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the associated queue.
	QueueId *string `mandatory:"true" json:"queueId"`

	// The filter used by the consumer group. Only messages matching the filter will be available by consumers of the group.
	Filter *string `mandatory:"false" json:"filter"`

	// Used to enable or disable the consumer group.
	// An enabled consumer group will have a lifecycle state of ACTIVE, while a disabled will have its state as INACTIVE.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue.
	// A value of 0 indicates that the DLQ is not used.
	// If the value isn't specified, it will be using the value defined at the queue level.
	DeadLetterQueueDeliveryCount *int `mandatory:"false" json:"deadLetterQueueDeliveryCount"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateConsumerGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateConsumerGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
