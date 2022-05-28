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

// FailureDestinationDetails An object that represents the destination to which Oracle Functions will send an invocation record with the details of the error of the failed asynchronous function invocation.
// A notification is an example of a failure destination.
// Example: `{"kind": "NOTIFICATION", "topicId": "topic_OCID"}`
type FailureDestinationDetails interface {
}

type failuredestinationdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *failuredestinationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfailuredestinationdetails failuredestinationdetails
	s := struct {
		Model Unmarshalerfailuredestinationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *failuredestinationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STREAM":
		mm := StreamFailureDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NOTIFICATION":
		mm := NotificationFailureDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneFailureDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m failuredestinationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m failuredestinationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FailureDestinationDetailsKindEnum Enum with underlying type: string
type FailureDestinationDetailsKindEnum string

// Set of constants representing the allowable values for FailureDestinationDetailsKindEnum
const (
	FailureDestinationDetailsKindNotification FailureDestinationDetailsKindEnum = "NOTIFICATION"
	FailureDestinationDetailsKindStream       FailureDestinationDetailsKindEnum = "STREAM"
	FailureDestinationDetailsKindNone         FailureDestinationDetailsKindEnum = "NONE"
)

var mappingFailureDestinationDetailsKindEnum = map[string]FailureDestinationDetailsKindEnum{
	"NOTIFICATION": FailureDestinationDetailsKindNotification,
	"STREAM":       FailureDestinationDetailsKindStream,
	"NONE":         FailureDestinationDetailsKindNone,
}

var mappingFailureDestinationDetailsKindEnumLowerCase = map[string]FailureDestinationDetailsKindEnum{
	"notification": FailureDestinationDetailsKindNotification,
	"stream":       FailureDestinationDetailsKindStream,
	"none":         FailureDestinationDetailsKindNone,
}

// GetFailureDestinationDetailsKindEnumValues Enumerates the set of values for FailureDestinationDetailsKindEnum
func GetFailureDestinationDetailsKindEnumValues() []FailureDestinationDetailsKindEnum {
	values := make([]FailureDestinationDetailsKindEnum, 0)
	for _, v := range mappingFailureDestinationDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetFailureDestinationDetailsKindEnumStringValues Enumerates the set of values in String for FailureDestinationDetailsKindEnum
func GetFailureDestinationDetailsKindEnumStringValues() []string {
	return []string{
		"NOTIFICATION",
		"STREAM",
		"NONE",
	}
}

// GetMappingFailureDestinationDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFailureDestinationDetailsKindEnum(val string) (FailureDestinationDetailsKindEnum, bool) {
	enum, ok := mappingFailureDestinationDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
