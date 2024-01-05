// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NoneOutboundConnection Details required for removing Private Endpoint Outbound Connection (ReverseConnection).
type NoneOutboundConnection struct {
}

func (m NoneOutboundConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NoneOutboundConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NoneOutboundConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNoneOutboundConnection NoneOutboundConnection
	s := struct {
		DiscriminatorParam string `json:"outboundConnectionType"`
		MarshalTypeNoneOutboundConnection
	}{
		"NONE",
		(MarshalTypeNoneOutboundConnection)(m),
	}

	return json.Marshal(&s)
}
