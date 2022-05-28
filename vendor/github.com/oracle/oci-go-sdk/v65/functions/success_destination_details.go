// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SuccessDestinationDetails An object that represents the destination to which Oracle Functions will send the response body of the successful asynchronous function invocation.
// A stream is an example of a success destination.
// Example: `{"kind": "STREAM", "streamId": "stream_OCID"}`
type SuccessDestinationDetails interface {
}

type successdestinationdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *successdestinationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersuccessdestinationdetails successdestinationdetails
	s := struct {
		Model Unmarshalersuccessdestinationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *successdestinationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STREAM":
		mm := StreamSuccessDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECTSTORAGE":
		mm := ObjectStorageSuccessDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneSuccessDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m successdestinationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m successdestinationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SuccessDestinationDetailsKindEnum Enum with underlying type: string
type SuccessDestinationDetailsKindEnum string

// Set of constants representing the allowable values for SuccessDestinationDetailsKindEnum
const (
	SuccessDestinationDetailsKindObjectstorage SuccessDestinationDetailsKindEnum = "OBJECTSTORAGE"
	SuccessDestinationDetailsKindStream        SuccessDestinationDetailsKindEnum = "STREAM"
	SuccessDestinationDetailsKindNone          SuccessDestinationDetailsKindEnum = "NONE"
)

var mappingSuccessDestinationDetailsKindEnum = map[string]SuccessDestinationDetailsKindEnum{
	"OBJECTSTORAGE": SuccessDestinationDetailsKindObjectstorage,
	"STREAM":        SuccessDestinationDetailsKindStream,
	"NONE":          SuccessDestinationDetailsKindNone,
}

var mappingSuccessDestinationDetailsKindEnumLowerCase = map[string]SuccessDestinationDetailsKindEnum{
	"objectstorage": SuccessDestinationDetailsKindObjectstorage,
	"stream":        SuccessDestinationDetailsKindStream,
	"none":          SuccessDestinationDetailsKindNone,
}

// GetSuccessDestinationDetailsKindEnumValues Enumerates the set of values for SuccessDestinationDetailsKindEnum
func GetSuccessDestinationDetailsKindEnumValues() []SuccessDestinationDetailsKindEnum {
	values := make([]SuccessDestinationDetailsKindEnum, 0)
	for _, v := range mappingSuccessDestinationDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetSuccessDestinationDetailsKindEnumStringValues Enumerates the set of values in String for SuccessDestinationDetailsKindEnum
func GetSuccessDestinationDetailsKindEnumStringValues() []string {
	return []string{
		"OBJECTSTORAGE",
		"STREAM",
		"NONE",
	}
}

// GetMappingSuccessDestinationDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSuccessDestinationDetailsKindEnum(val string) (SuccessDestinationDetailsKindEnum, bool) {
	enum, ok := mappingSuccessDestinationDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
