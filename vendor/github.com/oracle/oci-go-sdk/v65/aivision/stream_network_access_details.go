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

// StreamNetworkAccessDetails Details about a stream Connection type
type StreamNetworkAccessDetails interface {
}

type streamnetworkaccessdetails struct {
	JsonData         []byte
	StreamAccessType string `json:"streamAccessType"`
}

// UnmarshalJSON unmarshals json
func (m *streamnetworkaccessdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamnetworkaccessdetails streamnetworkaccessdetails
	s := struct {
		Model Unmarshalerstreamnetworkaccessdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StreamAccessType = s.Model.StreamAccessType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamnetworkaccessdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StreamAccessType {
	case "PRIVATE":
		mm := PrivateStreamNetworkAccessDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for StreamNetworkAccessDetails: %s.", m.StreamAccessType)
		return *m, nil
	}
}

func (m streamnetworkaccessdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m streamnetworkaccessdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StreamNetworkAccessDetailsStreamAccessTypeEnum Enum with underlying type: string
type StreamNetworkAccessDetailsStreamAccessTypeEnum string

// Set of constants representing the allowable values for StreamNetworkAccessDetailsStreamAccessTypeEnum
const (
	StreamNetworkAccessDetailsStreamAccessTypePrivate StreamNetworkAccessDetailsStreamAccessTypeEnum = "PRIVATE"
)

var mappingStreamNetworkAccessDetailsStreamAccessTypeEnum = map[string]StreamNetworkAccessDetailsStreamAccessTypeEnum{
	"PRIVATE": StreamNetworkAccessDetailsStreamAccessTypePrivate,
}

var mappingStreamNetworkAccessDetailsStreamAccessTypeEnumLowerCase = map[string]StreamNetworkAccessDetailsStreamAccessTypeEnum{
	"private": StreamNetworkAccessDetailsStreamAccessTypePrivate,
}

// GetStreamNetworkAccessDetailsStreamAccessTypeEnumValues Enumerates the set of values for StreamNetworkAccessDetailsStreamAccessTypeEnum
func GetStreamNetworkAccessDetailsStreamAccessTypeEnumValues() []StreamNetworkAccessDetailsStreamAccessTypeEnum {
	values := make([]StreamNetworkAccessDetailsStreamAccessTypeEnum, 0)
	for _, v := range mappingStreamNetworkAccessDetailsStreamAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStreamNetworkAccessDetailsStreamAccessTypeEnumStringValues Enumerates the set of values in String for StreamNetworkAccessDetailsStreamAccessTypeEnum
func GetStreamNetworkAccessDetailsStreamAccessTypeEnumStringValues() []string {
	return []string{
		"PRIVATE",
	}
}

// GetMappingStreamNetworkAccessDetailsStreamAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStreamNetworkAccessDetailsStreamAccessTypeEnum(val string) (StreamNetworkAccessDetailsStreamAccessTypeEnum, bool) {
	enum, ok := mappingStreamNetworkAccessDetailsStreamAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
