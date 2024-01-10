// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PutMessagesResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
