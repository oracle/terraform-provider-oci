// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TrimHorizonStreamingCursor `TRIM_HORIZON` cursor type. Sets the starting point for consuming the stream at the oldest available message in the stream. For more information about Streaming cursors, see Using Cursors (https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors).
type TrimHorizonStreamingCursor struct {
}

func (m TrimHorizonStreamingCursor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TrimHorizonStreamingCursor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
