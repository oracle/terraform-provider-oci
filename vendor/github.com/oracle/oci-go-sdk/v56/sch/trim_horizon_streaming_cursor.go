// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TrimHorizonStreamingCursor `TRIM_HORIZON` cursor type. Sets the starting point for consuming the stream at the oldest available message in the stream. For more information about Streaming cursors, see Using Cursors (https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors).
type TrimHorizonStreamingCursor struct {
}

func (m TrimHorizonStreamingCursor) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TrimHorizonStreamingCursor) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTrimHorizonStreamingCursor TrimHorizonStreamingCursor
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeTrimHorizonStreamingCursor
	}{
		"TRIM_HORIZON",
		(MarshalTypeTrimHorizonStreamingCursor)(m),
	}

	return json.Marshal(&s)
}
