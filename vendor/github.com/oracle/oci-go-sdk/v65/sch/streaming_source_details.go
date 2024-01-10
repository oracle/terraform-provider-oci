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

// StreamingSourceDetails The Streaming source.
type StreamingSourceDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	StreamId *string `mandatory:"true" json:"streamId"`

	Cursor StreamingCursorDetails `mandatory:"false" json:"cursor"`
}

func (m StreamingSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamingSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamingSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamingSourceDetails StreamingSourceDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStreamingSourceDetails
	}{
		"streaming",
		(MarshalTypeStreamingSourceDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *StreamingSourceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Cursor   streamingcursordetails `json:"cursor"`
		StreamId *string                `json:"streamId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
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
