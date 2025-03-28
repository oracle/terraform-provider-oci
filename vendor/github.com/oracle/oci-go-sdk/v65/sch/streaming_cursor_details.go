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

// StreamingCursorDetails The read setting (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector-streaming-source.htm), which determines where in the stream to start moving data.
// For configuration instructions, see
// Creating a Connector with a Streaming Source (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector-streaming-source.htm).
type StreamingCursorDetails interface {
}

type streamingcursordetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *streamingcursordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamingcursordetails streamingcursordetails
	s := struct {
		Model Unmarshalerstreamingcursordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamingcursordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "TRIM_HORIZON":
		mm := TrimHorizonStreamingCursor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LATEST":
		mm := LatestStreamingCursor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for StreamingCursorDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m streamingcursordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streamingcursordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamingCursorDetailsKindEnum Enum with underlying type: string
type StreamingCursorDetailsKindEnum string

// Set of constants representing the allowable values for StreamingCursorDetailsKindEnum
const (
	StreamingCursorDetailsKindLatest      StreamingCursorDetailsKindEnum = "LATEST"
	StreamingCursorDetailsKindTrimHorizon StreamingCursorDetailsKindEnum = "TRIM_HORIZON"
)

var mappingStreamingCursorDetailsKindEnum = map[string]StreamingCursorDetailsKindEnum{
	"LATEST":       StreamingCursorDetailsKindLatest,
	"TRIM_HORIZON": StreamingCursorDetailsKindTrimHorizon,
}

var mappingStreamingCursorDetailsKindEnumLowerCase = map[string]StreamingCursorDetailsKindEnum{
	"latest":       StreamingCursorDetailsKindLatest,
	"trim_horizon": StreamingCursorDetailsKindTrimHorizon,
}

// GetStreamingCursorDetailsKindEnumValues Enumerates the set of values for StreamingCursorDetailsKindEnum
func GetStreamingCursorDetailsKindEnumValues() []StreamingCursorDetailsKindEnum {
	values := make([]StreamingCursorDetailsKindEnum, 0)
	for _, v := range mappingStreamingCursorDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamingCursorDetailsKindEnumStringValues Enumerates the set of values in String for StreamingCursorDetailsKindEnum
func GetStreamingCursorDetailsKindEnumStringValues() []string {
	return []string{
		"LATEST",
		"TRIM_HORIZON",
	}
}

// GetMappingStreamingCursorDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamingCursorDetailsKindEnum(val string) (StreamingCursorDetailsKindEnum, bool) {
	enum, ok := mappingStreamingCursorDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
