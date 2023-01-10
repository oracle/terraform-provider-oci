// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// PutMessagesResultEntry Represents the result of a PutMessages request, whether it was successful or not.
// If a message was successfully appended to the stream, the entry includes the `offset`, `partition`, and `timestamp`.
// If the message failed to be appended to the stream, the entry includes the `error` and `errorMessage`.
type PutMessagesResultEntry struct {

	// The ID of the partition where the message was stored.
	Partition *string `mandatory:"false" json:"partition"`

	// The offset of the message in the partition.
	Offset *int64 `mandatory:"false" json:"offset"`

	// The timestamp indicating when the server appended the message to the stream.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The error code, in case the message was not successfully appended to the stream.
	Error *string `mandatory:"false" json:"error"`

	// A human-readable error message associated with the error code.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m PutMessagesResultEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PutMessagesResultEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
