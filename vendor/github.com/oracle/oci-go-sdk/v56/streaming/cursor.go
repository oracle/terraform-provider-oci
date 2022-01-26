// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Cursor A cursor that indicates the position in the stream from which you want to begin consuming messages and which is required by the GetMessages operation.
type Cursor struct {

	// The cursor to pass to the `GetMessages` operation.
	Value *string `mandatory:"true" json:"value"`
}

func (m Cursor) String() string {
	return common.PointerString(m)
}
