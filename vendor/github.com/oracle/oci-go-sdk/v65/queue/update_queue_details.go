// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateQueueDetails The information to be updated.
type UpdateQueueDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the queue.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The default visibility timeout of the messages consumed from the queue, in seconds.
	VisibilityInSeconds *int `mandatory:"false" json:"visibilityInSeconds"`

	// The default polling timeout of the messages in the queue, in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// The percentage of allocated queue resources that can be consumed by a single channel. For example, if a queue has a storage limit of 2Gb, and a single channel consumption limit is 0.1 (10%), that means data size of a single channel  can't exceed 200Mb. Consumption limit of 100% (default) means that a single channel can consume up-to all allocated queue's resources.
	ChannelConsumptionLimit *int `mandatory:"false" json:"channelConsumptionLimit"`

	// The number of times a message can be delivered to a consumer before being moved to the dead letter queue.
	// A value of 0 indicates that the DLQ is not used.
	// Changing that value to a lower threshold does not retroactively move in-flight messages in the dead letter queue.
	DeadLetterQueueDeliveryCount *int `mandatory:"false" json:"deadLetterQueueDeliveryCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom encryption key to be used to encrypt messages content. A string with a length of 0 means the custom key should be removed from queue.
	CustomEncryptionKeyId *string `mandatory:"false" json:"customEncryptionKeyId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateQueueDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateQueueDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
