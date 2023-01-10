// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// A description of the Queue API
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PurgeQueueDetails Purge parameters.
type PurgeQueueDetails struct {

	// Type of the purge to perform:
	// - NORMAL - purge only normal queue
	// - DLQ - purge only DLQ
	// - BOTH - purge both normal queue and DLQ
	PurgeType PurgeQueueDetailsPurgeTypeEnum `mandatory:"true" json:"purgeType"`
}

func (m PurgeQueueDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PurgeQueueDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPurgeQueueDetailsPurgeTypeEnum(string(m.PurgeType)); !ok && m.PurgeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PurgeType: %s. Supported values are: %s.", m.PurgeType, strings.Join(GetPurgeQueueDetailsPurgeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PurgeQueueDetailsPurgeTypeEnum Enum with underlying type: string
type PurgeQueueDetailsPurgeTypeEnum string

// Set of constants representing the allowable values for PurgeQueueDetailsPurgeTypeEnum
const (
	PurgeQueueDetailsPurgeTypeNormal PurgeQueueDetailsPurgeTypeEnum = "NORMAL"
	PurgeQueueDetailsPurgeTypeDlq    PurgeQueueDetailsPurgeTypeEnum = "DLQ"
	PurgeQueueDetailsPurgeTypeBoth   PurgeQueueDetailsPurgeTypeEnum = "BOTH"
)

var mappingPurgeQueueDetailsPurgeTypeEnum = map[string]PurgeQueueDetailsPurgeTypeEnum{
	"NORMAL": PurgeQueueDetailsPurgeTypeNormal,
	"DLQ":    PurgeQueueDetailsPurgeTypeDlq,
	"BOTH":   PurgeQueueDetailsPurgeTypeBoth,
}

var mappingPurgeQueueDetailsPurgeTypeEnumLowerCase = map[string]PurgeQueueDetailsPurgeTypeEnum{
	"normal": PurgeQueueDetailsPurgeTypeNormal,
	"dlq":    PurgeQueueDetailsPurgeTypeDlq,
	"both":   PurgeQueueDetailsPurgeTypeBoth,
}

// GetPurgeQueueDetailsPurgeTypeEnumValues Enumerates the set of values for PurgeQueueDetailsPurgeTypeEnum
func GetPurgeQueueDetailsPurgeTypeEnumValues() []PurgeQueueDetailsPurgeTypeEnum {
	values := make([]PurgeQueueDetailsPurgeTypeEnum, 0)
	for _, v := range mappingPurgeQueueDetailsPurgeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPurgeQueueDetailsPurgeTypeEnumStringValues Enumerates the set of values in String for PurgeQueueDetailsPurgeTypeEnum
func GetPurgeQueueDetailsPurgeTypeEnumStringValues() []string {
	return []string{
		"NORMAL",
		"DLQ",
		"BOTH",
	}
}

// GetMappingPurgeQueueDetailsPurgeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPurgeQueueDetailsPurgeTypeEnum(val string) (PurgeQueueDetailsPurgeTypeEnum, bool) {
	enum, ok := mappingPurgeQueueDetailsPurgeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
