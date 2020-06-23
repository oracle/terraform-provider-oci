// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Message The details of a message.
type Message struct {

	// The type of message (error, warning, or info).
	Type MessageTypeEnum `mandatory:"true" json:"type"`

	// The message code
	Code *string `mandatory:"true" json:"code"`

	// The message text
	Message *string `mandatory:"true" json:"message"`
}

func (m Message) String() string {
	return common.PointerString(m)
}

// MessageTypeEnum Enum with underlying type: string
type MessageTypeEnum string

// Set of constants representing the allowable values for MessageTypeEnum
const (
	MessageTypeError   MessageTypeEnum = "ERROR"
	MessageTypeWarning MessageTypeEnum = "WARNING"
	MessageTypeInfo    MessageTypeEnum = "INFO"
)

var mappingMessageType = map[string]MessageTypeEnum{
	"ERROR":   MessageTypeError,
	"WARNING": MessageTypeWarning,
	"INFO":    MessageTypeInfo,
}

// GetMessageTypeEnumValues Enumerates the set of values for MessageTypeEnum
func GetMessageTypeEnumValues() []MessageTypeEnum {
	values := make([]MessageTypeEnum, 0)
	for _, v := range mappingMessageType {
		values = append(values, v)
	}
	return values
}
