// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SendMessageDetails The message to be sent.
type SendMessageDetails interface {
}

type sendmessagedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *sendmessagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersendmessagedetails sendmessagedetails
	s := struct {
		Model Unmarshalersendmessagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sendmessagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SMS":
		mm := SendSmsDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SendMessageDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m sendmessagedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sendmessagedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SendMessageDetailsTypeEnum Enum with underlying type: string
type SendMessageDetailsTypeEnum string

// Set of constants representing the allowable values for SendMessageDetailsTypeEnum
const (
	SendMessageDetailsTypeSms SendMessageDetailsTypeEnum = "SMS"
)

var mappingSendMessageDetailsTypeEnum = map[string]SendMessageDetailsTypeEnum{
	"SMS": SendMessageDetailsTypeSms,
}

var mappingSendMessageDetailsTypeEnumLowerCase = map[string]SendMessageDetailsTypeEnum{
	"sms": SendMessageDetailsTypeSms,
}

// GetSendMessageDetailsTypeEnumValues Enumerates the set of values for SendMessageDetailsTypeEnum
func GetSendMessageDetailsTypeEnumValues() []SendMessageDetailsTypeEnum {
	values := make([]SendMessageDetailsTypeEnum, 0)
	for _, v := range mappingSendMessageDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSendMessageDetailsTypeEnumStringValues Enumerates the set of values in String for SendMessageDetailsTypeEnum
func GetSendMessageDetailsTypeEnumStringValues() []string {
	return []string{
		"SMS",
	}
}

// GetMappingSendMessageDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSendMessageDetailsTypeEnum(val string) (SendMessageDetailsTypeEnum, bool) {
	enum, ok := mappingSendMessageDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
