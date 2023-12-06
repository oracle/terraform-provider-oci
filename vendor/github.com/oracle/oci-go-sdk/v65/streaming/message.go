// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Message A message in a stream.
type Message struct {

	// The name of the stream that the message belongs to.
	Stream *string `mandatory:"true" json:"stream"`

	// The ID of the partition where the message is stored.
	Partition *string `mandatory:"true" json:"partition"`

	// The key associated with the message, expressed as a byte array.
	Key []byte `mandatory:"true" json:"key"`

	// The value associated with the message, expressed as a byte array.
	Value []byte `mandatory:"true" json:"value"`

	// The offset of the message, which uniquely identifies it within the partition.
	Offset *int64 `mandatory:"true" json:"offset"`

	// The timestamp indicating when the server appended the message to the stream.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m Message) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Message) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
