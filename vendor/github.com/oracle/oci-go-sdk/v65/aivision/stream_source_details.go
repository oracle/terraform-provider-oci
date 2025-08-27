// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamSourceDetails Details about a stream source
type StreamSourceDetails interface {
	GetStreamNetworkAccessDetails() StreamNetworkAccessDetails
}

type streamsourcedetails struct {
	JsonData                   []byte
	StreamNetworkAccessDetails streamnetworkaccessdetails `mandatory:"true" json:"streamNetworkAccessDetails"`
	SourceType                 string                     `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *streamsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamsourcedetails streamsourcedetails
	s := struct {
		Model Unmarshalerstreamsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StreamNetworkAccessDetails = s.Model.StreamNetworkAccessDetails
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "RTSP":
		mm := RtspSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for StreamSourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

// GetStreamNetworkAccessDetails returns StreamNetworkAccessDetails
func (m streamsourcedetails) GetStreamNetworkAccessDetails() streamnetworkaccessdetails {
	return m.StreamNetworkAccessDetails
}

func (m streamsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streamsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamSourceDetailsSourceTypeEnum Enum with underlying type: string
type StreamSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for StreamSourceDetailsSourceTypeEnum
const (
	StreamSourceDetailsSourceTypeRtsp StreamSourceDetailsSourceTypeEnum = "RTSP"
)

var mappingStreamSourceDetailsSourceTypeEnum = map[string]StreamSourceDetailsSourceTypeEnum{
	"RTSP": StreamSourceDetailsSourceTypeRtsp,
}

var mappingStreamSourceDetailsSourceTypeEnumLowerCase = map[string]StreamSourceDetailsSourceTypeEnum{
	"rtsp": StreamSourceDetailsSourceTypeRtsp,
}

// GetStreamSourceDetailsSourceTypeEnumValues Enumerates the set of values for StreamSourceDetailsSourceTypeEnum
func GetStreamSourceDetailsSourceTypeEnumValues() []StreamSourceDetailsSourceTypeEnum {
	values := make([]StreamSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingStreamSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for StreamSourceDetailsSourceTypeEnum
func GetStreamSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"RTSP",
	}
}

// GetMappingStreamSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamSourceDetailsSourceTypeEnum(val string) (StreamSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingStreamSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
