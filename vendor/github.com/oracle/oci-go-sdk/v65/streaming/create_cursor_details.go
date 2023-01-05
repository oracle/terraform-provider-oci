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

// CreateCursorDetails Object used to create a cursor to consume messages in a stream.
type CreateCursorDetails struct {

	// The partition to get messages from.
	Partition *string `mandatory:"true" json:"partition"`

	// The type of cursor, which determines the starting point from which the stream will be consumed:
	// - `AFTER_OFFSET:` The partition position immediately following the offset you specify. (Offsets are assigned when you successfully append a message to a partition in a stream.)
	// - `AT_OFFSET:` The exact partition position indicated by the offset you specify.
	// - `AT_TIME:` A specific point in time.
	// - `LATEST:` The most recent message in the partition that was added after the cursor was created.
	// - `TRIM_HORIZON:` The oldest message in the partition that is within the retention period window.
	Type CreateCursorDetailsTypeEnum `mandatory:"true" json:"type"`

	// The offset to consume from if the cursor type is `AT_OFFSET` or `AFTER_OFFSET`.
	Offset *int64 `mandatory:"false" json:"offset"`

	// The time to consume from if the cursor type is `AT_TIME`, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	Time *common.SDKTime `mandatory:"false" json:"time"`
}

func (m CreateCursorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCursorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateCursorDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateCursorDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCursorDetailsTypeEnum Enum with underlying type: string
type CreateCursorDetailsTypeEnum string

// Set of constants representing the allowable values for CreateCursorDetailsTypeEnum
const (
	CreateCursorDetailsTypeAfterOffset CreateCursorDetailsTypeEnum = "AFTER_OFFSET"
	CreateCursorDetailsTypeAtOffset    CreateCursorDetailsTypeEnum = "AT_OFFSET"
	CreateCursorDetailsTypeAtTime      CreateCursorDetailsTypeEnum = "AT_TIME"
	CreateCursorDetailsTypeLatest      CreateCursorDetailsTypeEnum = "LATEST"
	CreateCursorDetailsTypeTrimHorizon CreateCursorDetailsTypeEnum = "TRIM_HORIZON"
)

var mappingCreateCursorDetailsTypeEnum = map[string]CreateCursorDetailsTypeEnum{
	"AFTER_OFFSET": CreateCursorDetailsTypeAfterOffset,
	"AT_OFFSET":    CreateCursorDetailsTypeAtOffset,
	"AT_TIME":      CreateCursorDetailsTypeAtTime,
	"LATEST":       CreateCursorDetailsTypeLatest,
	"TRIM_HORIZON": CreateCursorDetailsTypeTrimHorizon,
}

var mappingCreateCursorDetailsTypeEnumLowerCase = map[string]CreateCursorDetailsTypeEnum{
	"after_offset": CreateCursorDetailsTypeAfterOffset,
	"at_offset":    CreateCursorDetailsTypeAtOffset,
	"at_time":      CreateCursorDetailsTypeAtTime,
	"latest":       CreateCursorDetailsTypeLatest,
	"trim_horizon": CreateCursorDetailsTypeTrimHorizon,
}

// GetCreateCursorDetailsTypeEnumValues Enumerates the set of values for CreateCursorDetailsTypeEnum
func GetCreateCursorDetailsTypeEnumValues() []CreateCursorDetailsTypeEnum {
	values := make([]CreateCursorDetailsTypeEnum, 0)
	for _, v := range mappingCreateCursorDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCursorDetailsTypeEnumStringValues Enumerates the set of values in String for CreateCursorDetailsTypeEnum
func GetCreateCursorDetailsTypeEnumStringValues() []string {
	return []string{
		"AFTER_OFFSET",
		"AT_OFFSET",
		"AT_TIME",
		"LATEST",
		"TRIM_HORIZON",
	}
}

// GetMappingCreateCursorDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCursorDetailsTypeEnum(val string) (CreateCursorDetailsTypeEnum, bool) {
	enum, ok := mappingCreateCursorDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
