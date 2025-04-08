// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamingSourceDetailsResponse The Streaming source response. Private metadata is included when the target is a stream accessed through
// a private endpoint (https://docs.oracle.com/iaas/Content/Streaming/Concepts/streamsecurity.htm#private_endpoints).
type StreamingSourceDetailsResponse struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	StreamId *string `mandatory:"true" json:"streamId"`

	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`

	Cursor StreamingCursorDetails `mandatory:"false" json:"cursor"`
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m StreamingSourceDetailsResponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m StreamingSourceDetailsResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamingSourceDetailsResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamingSourceDetailsResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamingSourceDetailsResponse StreamingSourceDetailsResponse
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStreamingSourceDetailsResponse
	}{
		"streaming",
		(MarshalTypeStreamingSourceDetailsResponse)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *StreamingSourceDetailsResponse) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PrivateEndpointMetadata *PrivateEndpointMetadata `json:"privateEndpointMetadata"`
		Cursor                  streamingcursordetails   `json:"cursor"`
		StreamId                *string                  `json:"streamId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.PrivateEndpointMetadata = model.PrivateEndpointMetadata

	nn, e = model.Cursor.UnmarshalPolymorphicJSON(model.Cursor.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Cursor = nn.(StreamingCursorDetails)
	} else {
		m.Cursor = nil
	}

	m.StreamId = model.StreamId

	return
}
