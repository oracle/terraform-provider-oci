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

// StreamingTargetDetails The stream used for the Streaming target.
// For configuration instructions, see
// To create a service connector (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/managingconnectors.htm#create).
type StreamingTargetDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	StreamId *string `mandatory:"true" json:"streamId"`
}

func (m StreamingTargetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m StreamingTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamingTargetDetails StreamingTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStreamingTargetDetails
	}{
		"streaming",
		(MarshalTypeStreamingTargetDetails)(m),
	}

	return json.Marshal(&s)
}
