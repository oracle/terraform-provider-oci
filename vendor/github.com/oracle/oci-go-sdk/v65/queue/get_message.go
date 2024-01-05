// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GetMessage A message consumed from a queue.
type GetMessage struct {

	// The ID of the message. This ID is only used for tracing and debugging purposes and isn't used as a parameter in any request.
	Id *int64 `mandatory:"true" json:"id"`

	// The content of the message.
	Content *string `mandatory:"true" json:"content"`

	// A receipt is a base64urlencode opaque token, uniquely representing a message.
	// The receipt can be used to delete a message or update its visibility.
	Receipt *string `mandatory:"true" json:"receipt"`

	// The number of times that the message has been delivered to a consumer.
	DeliveryCount *int `mandatory:"true" json:"deliveryCount"`

	// The time after which the message will be visible to other consumers, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	VisibleAfter *common.SDKTime `mandatory:"true" json:"visibleAfter"`

	// The time after which the message will be automatically deleted, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	ExpireAfter *common.SDKTime `mandatory:"true" json:"expireAfter"`

	Metadata *MessageMetadata `mandatory:"false" json:"metadata"`
}

func (m GetMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GetMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
