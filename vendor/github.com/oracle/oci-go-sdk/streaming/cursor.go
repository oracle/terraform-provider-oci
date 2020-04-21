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

// Cursor A cursor that indicates the position in the stream from which you want to begin consuming messages and which is required by the GetMessages operation.
type Cursor struct {

	// The cursor to pass to the `GetMessages` operation.
	Value *string `mandatory:"true" json:"value"`
}

func (m Cursor) String() string {
	return common.PointerString(m)
}
