// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateGroupDetails Request body for operationally managing a group.
type UpdateGroupDetails struct {

	// The type of the cursor.
	Type UpdateGroupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The time to consume from if type is AT_TIME.
	Time *common.SDKTime `mandatory:"false" json:"time"`
}

func (m UpdateGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateGroupDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUpdateGroupDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateGroupDetailsTypeEnum Enum with underlying type: string
type UpdateGroupDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateGroupDetailsTypeEnum
const (
	UpdateGroupDetailsTypeAtTime      UpdateGroupDetailsTypeEnum = "AT_TIME"
	UpdateGroupDetailsTypeLatest      UpdateGroupDetailsTypeEnum = "LATEST"
	UpdateGroupDetailsTypeTrimHorizon UpdateGroupDetailsTypeEnum = "TRIM_HORIZON"
)

var mappingUpdateGroupDetailsTypeEnum = map[string]UpdateGroupDetailsTypeEnum{
	"AT_TIME":      UpdateGroupDetailsTypeAtTime,
	"LATEST":       UpdateGroupDetailsTypeLatest,
	"TRIM_HORIZON": UpdateGroupDetailsTypeTrimHorizon,
}

var mappingUpdateGroupDetailsTypeEnumLowerCase = map[string]UpdateGroupDetailsTypeEnum{
	"at_time":      UpdateGroupDetailsTypeAtTime,
	"latest":       UpdateGroupDetailsTypeLatest,
	"trim_horizon": UpdateGroupDetailsTypeTrimHorizon,
}

// GetUpdateGroupDetailsTypeEnumValues Enumerates the set of values for UpdateGroupDetailsTypeEnum
func GetUpdateGroupDetailsTypeEnumValues() []UpdateGroupDetailsTypeEnum {
	values := make([]UpdateGroupDetailsTypeEnum, 0)
	for _, v := range mappingUpdateGroupDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateGroupDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateGroupDetailsTypeEnum
func GetUpdateGroupDetailsTypeEnumStringValues() []string {
	return []string{
		"AT_TIME",
		"LATEST",
		"TRIM_HORIZON",
	}
}

// GetMappingUpdateGroupDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateGroupDetailsTypeEnum(val string) (UpdateGroupDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateGroupDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
