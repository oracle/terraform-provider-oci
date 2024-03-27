// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EventContent Information collected during the event, such as logs.
type EventContent interface {
}

type eventcontent struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *eventcontent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalereventcontent eventcontent
	s := struct {
		Model Unmarshalereventcontent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *eventcontent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "KERNEL":
		mm := KernelEventContent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXPLOIT_ATTEMPT":
		mm := ExploitAttemptEventContent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for EventContent: %s.", m.Type)
		return *m, nil
	}
}

func (m eventcontent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m eventcontent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EventContentTypeEnum Enum with underlying type: string
type EventContentTypeEnum string

// Set of constants representing the allowable values for EventContentTypeEnum
const (
	EventContentTypeKernel         EventContentTypeEnum = "KERNEL"
	EventContentTypeExploitAttempt EventContentTypeEnum = "EXPLOIT_ATTEMPT"
)

var mappingEventContentTypeEnum = map[string]EventContentTypeEnum{
	"KERNEL":          EventContentTypeKernel,
	"EXPLOIT_ATTEMPT": EventContentTypeExploitAttempt,
}

var mappingEventContentTypeEnumLowerCase = map[string]EventContentTypeEnum{
	"kernel":          EventContentTypeKernel,
	"exploit_attempt": EventContentTypeExploitAttempt,
}

// GetEventContentTypeEnumValues Enumerates the set of values for EventContentTypeEnum
func GetEventContentTypeEnumValues() []EventContentTypeEnum {
	values := make([]EventContentTypeEnum, 0)
	for _, v := range mappingEventContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEventContentTypeEnumStringValues Enumerates the set of values in String for EventContentTypeEnum
func GetEventContentTypeEnumStringValues() []string {
	return []string{
		"KERNEL",
		"EXPLOIT_ATTEMPT",
	}
}

// GetMappingEventContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEventContentTypeEnum(val string) (EventContentTypeEnum, bool) {
	enum, ok := mappingEventContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
