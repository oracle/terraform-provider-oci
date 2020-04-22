// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PutMessagesResult The response to a PutMessages request. It indicates the number
// of failed messages as well as an array of results for successful and failed messages.
type PutMessagesResult struct {

	// The number of messages that failed to be added to the stream.
	Failures *int `mandatory:"true" json:"failures"`

	// An array of items representing the result of each message.
	// The order is guaranteed to be the same as in the `PutMessagesDetails` object.
	// If a message was successfully appended to the stream, the entry includes the `offset`, `partition`, and `timestamp`.
	// If a message failed to be appended to the stream, the entry includes the `error` and `errorMessage`.
	Entries []PutMessagesResultEntry `mandatory:"true" json:"entries"`
}

func (m PutMessagesResult) String() string {
	return common.PointerString(m)
}
