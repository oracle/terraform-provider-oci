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

// UpdateMessagesResult The response to a UpdateMessages request. It indicates the number of server and client failures as well as an array of entries for successful and failed actions.
type UpdateMessagesResult struct {

	// The number of messages that failed to be updated in the queue because of a server failure.
	ServerFailures *int `mandatory:"true" json:"serverFailures"`

	// The number of messages that failed to be updated in the queue because of a client failure such as an invalid receipt or invalid `visibilityInSeconds`.
	ClientFailures *int `mandatory:"true" json:"clientFailures"`

	// An array of items representing the result of each action.
	// The order is guaranteed to be the same as in the `UpdateMessagesDetails` object.
	// If a message was successfully updated in the queue, the entry includes the `id` and `visibleAfter` fields.
	// If a message failed to be updated in the queue, the entry includes the `errorCode` and `errorMessage` fields.
	Entries []UpdateMessagesResultEntry `mandatory:"true" json:"entries"`
}

func (m UpdateMessagesResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMessagesResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
