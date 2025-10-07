// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChangeIotDomainDataRetentionPeriodDetails The configuration details for data retention periods.
type ChangeIotDomainDataRetentionPeriodDetails struct {

	// The type of data retention period to apply. Allowed values are RAW_DATA, REJECTED_DATA, HISTORIZED_DATA, and RAW_COMMAND_DATA.
	Type ChangeIotDomainDataRetentionPeriodDetailsTypeEnum `mandatory:"true" json:"type"`

	// The duration (in days) for which data will be retained in the IoT domain.
	DataRetentionPeriodInDays *int `mandatory:"true" json:"dataRetentionPeriodInDays"`
}

func (m ChangeIotDomainDataRetentionPeriodDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeIotDomainDataRetentionPeriodDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChangeIotDomainDataRetentionPeriodDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetChangeIotDomainDataRetentionPeriodDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeIotDomainDataRetentionPeriodDetailsTypeEnum Enum with underlying type: string
type ChangeIotDomainDataRetentionPeriodDetailsTypeEnum string

// Set of constants representing the allowable values for ChangeIotDomainDataRetentionPeriodDetailsTypeEnum
const (
	ChangeIotDomainDataRetentionPeriodDetailsTypeRawData        ChangeIotDomainDataRetentionPeriodDetailsTypeEnum = "RAW_DATA"
	ChangeIotDomainDataRetentionPeriodDetailsTypeRejectedData   ChangeIotDomainDataRetentionPeriodDetailsTypeEnum = "REJECTED_DATA"
	ChangeIotDomainDataRetentionPeriodDetailsTypeHistorizedData ChangeIotDomainDataRetentionPeriodDetailsTypeEnum = "HISTORIZED_DATA"
	ChangeIotDomainDataRetentionPeriodDetailsTypeRawCommandData ChangeIotDomainDataRetentionPeriodDetailsTypeEnum = "RAW_COMMAND_DATA"
)

var mappingChangeIotDomainDataRetentionPeriodDetailsTypeEnum = map[string]ChangeIotDomainDataRetentionPeriodDetailsTypeEnum{
	"RAW_DATA":         ChangeIotDomainDataRetentionPeriodDetailsTypeRawData,
	"REJECTED_DATA":    ChangeIotDomainDataRetentionPeriodDetailsTypeRejectedData,
	"HISTORIZED_DATA":  ChangeIotDomainDataRetentionPeriodDetailsTypeHistorizedData,
	"RAW_COMMAND_DATA": ChangeIotDomainDataRetentionPeriodDetailsTypeRawCommandData,
}

var mappingChangeIotDomainDataRetentionPeriodDetailsTypeEnumLowerCase = map[string]ChangeIotDomainDataRetentionPeriodDetailsTypeEnum{
	"raw_data":         ChangeIotDomainDataRetentionPeriodDetailsTypeRawData,
	"rejected_data":    ChangeIotDomainDataRetentionPeriodDetailsTypeRejectedData,
	"historized_data":  ChangeIotDomainDataRetentionPeriodDetailsTypeHistorizedData,
	"raw_command_data": ChangeIotDomainDataRetentionPeriodDetailsTypeRawCommandData,
}

// GetChangeIotDomainDataRetentionPeriodDetailsTypeEnumValues Enumerates the set of values for ChangeIotDomainDataRetentionPeriodDetailsTypeEnum
func GetChangeIotDomainDataRetentionPeriodDetailsTypeEnumValues() []ChangeIotDomainDataRetentionPeriodDetailsTypeEnum {
	values := make([]ChangeIotDomainDataRetentionPeriodDetailsTypeEnum, 0)
	for _, v := range mappingChangeIotDomainDataRetentionPeriodDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeIotDomainDataRetentionPeriodDetailsTypeEnumStringValues Enumerates the set of values in String for ChangeIotDomainDataRetentionPeriodDetailsTypeEnum
func GetChangeIotDomainDataRetentionPeriodDetailsTypeEnumStringValues() []string {
	return []string{
		"RAW_DATA",
		"REJECTED_DATA",
		"HISTORIZED_DATA",
		"RAW_COMMAND_DATA",
	}
}

// GetMappingChangeIotDomainDataRetentionPeriodDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeIotDomainDataRetentionPeriodDetailsTypeEnum(val string) (ChangeIotDomainDataRetentionPeriodDetailsTypeEnum, bool) {
	enum, ok := mappingChangeIotDomainDataRetentionPeriodDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
