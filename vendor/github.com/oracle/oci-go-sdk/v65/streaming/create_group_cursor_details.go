// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGroupCursorDetails Object used to create a group cursor.
type CreateGroupCursorDetails struct {

	// The type of the cursor. This value is only used when the group is created.
	Type CreateGroupCursorDetailsTypeEnum `mandatory:"true" json:"type"`

	// Name of the consumer group.
	GroupName *string `mandatory:"true" json:"groupName"`

	// The time to consume from if type is AT_TIME.
	Time *common.SDKTime `mandatory:"false" json:"time"`

	// A unique identifier for the instance joining the consumer group. If an instanceName is not provided, a UUID will be generated and used.
	InstanceName *string `mandatory:"false" json:"instanceName"`

	// The amount of a consumer instance inactivity time, before partition reservations are released.
	TimeoutInMs *int `mandatory:"false" json:"timeoutInMs"`

	// When using consumer-groups, the default commit-on-get behaviour can be overriden by setting this value to false.
	// If disabled, a consumer must manually commit their cursors.
	CommitOnGet *bool `mandatory:"false" json:"commitOnGet"`
}

func (m CreateGroupCursorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGroupCursorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateGroupCursorDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateGroupCursorDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateGroupCursorDetailsTypeEnum Enum with underlying type: string
type CreateGroupCursorDetailsTypeEnum string

// Set of constants representing the allowable values for CreateGroupCursorDetailsTypeEnum
const (
	CreateGroupCursorDetailsTypeAtTime      CreateGroupCursorDetailsTypeEnum = "AT_TIME"
	CreateGroupCursorDetailsTypeLatest      CreateGroupCursorDetailsTypeEnum = "LATEST"
	CreateGroupCursorDetailsTypeTrimHorizon CreateGroupCursorDetailsTypeEnum = "TRIM_HORIZON"
)

var mappingCreateGroupCursorDetailsTypeEnum = map[string]CreateGroupCursorDetailsTypeEnum{
	"AT_TIME":      CreateGroupCursorDetailsTypeAtTime,
	"LATEST":       CreateGroupCursorDetailsTypeLatest,
	"TRIM_HORIZON": CreateGroupCursorDetailsTypeTrimHorizon,
}

var mappingCreateGroupCursorDetailsTypeEnumLowerCase = map[string]CreateGroupCursorDetailsTypeEnum{
	"at_time":      CreateGroupCursorDetailsTypeAtTime,
	"latest":       CreateGroupCursorDetailsTypeLatest,
	"trim_horizon": CreateGroupCursorDetailsTypeTrimHorizon,
}

// GetCreateGroupCursorDetailsTypeEnumValues Enumerates the set of values for CreateGroupCursorDetailsTypeEnum
func GetCreateGroupCursorDetailsTypeEnumValues() []CreateGroupCursorDetailsTypeEnum {
	values := make([]CreateGroupCursorDetailsTypeEnum, 0)
	for _, v := range mappingCreateGroupCursorDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateGroupCursorDetailsTypeEnumStringValues Enumerates the set of values in String for CreateGroupCursorDetailsTypeEnum
func GetCreateGroupCursorDetailsTypeEnumStringValues() []string {
	return []string{
		"AT_TIME",
		"LATEST",
		"TRIM_HORIZON",
	}
}

// GetMappingCreateGroupCursorDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateGroupCursorDetailsTypeEnum(val string) (CreateGroupCursorDetailsTypeEnum, bool) {
	enum, ok := mappingCreateGroupCursorDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
