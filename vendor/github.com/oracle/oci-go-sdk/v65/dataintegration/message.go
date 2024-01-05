// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Message The details of a message.
type Message struct {

	// The type of message (error, warning, or info).
	Type MessageTypeEnum `mandatory:"true" json:"type"`

	// The message code.
	Code *string `mandatory:"true" json:"code"`

	// The message text.
	Message *string `mandatory:"true" json:"message"`
}

func (m Message) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Message) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMessageTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMessageTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MessageTypeEnum Enum with underlying type: string
type MessageTypeEnum string

// Set of constants representing the allowable values for MessageTypeEnum
const (
	MessageTypeError   MessageTypeEnum = "ERROR"
	MessageTypeWarning MessageTypeEnum = "WARNING"
	MessageTypeInfo    MessageTypeEnum = "INFO"
)

var mappingMessageTypeEnum = map[string]MessageTypeEnum{
	"ERROR":   MessageTypeError,
	"WARNING": MessageTypeWarning,
	"INFO":    MessageTypeInfo,
}

var mappingMessageTypeEnumLowerCase = map[string]MessageTypeEnum{
	"error":   MessageTypeError,
	"warning": MessageTypeWarning,
	"info":    MessageTypeInfo,
}

// GetMessageTypeEnumValues Enumerates the set of values for MessageTypeEnum
func GetMessageTypeEnumValues() []MessageTypeEnum {
	values := make([]MessageTypeEnum, 0)
	for _, v := range mappingMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMessageTypeEnumStringValues Enumerates the set of values in String for MessageTypeEnum
func GetMessageTypeEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
		"INFO",
	}
}

// GetMappingMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMessageTypeEnum(val string) (MessageTypeEnum, bool) {
	enum, ok := mappingMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
