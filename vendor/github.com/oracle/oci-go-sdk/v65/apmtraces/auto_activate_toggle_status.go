// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoActivateToggleStatus Response of a auto-activate toggle operation.
type AutoActivateToggleStatus struct {

	// Status of this operation.
	Status AutoActivateToggleStatusStatusEnum `mandatory:"true" json:"status"`

	// State of autoactivation in this APM Domain.  If "ON" auto-activate is set to true, if "OFF" auto-activate is set to false.
	State *string `mandatory:"true" json:"state"`

	// Data key type for which auto-activate needs needs to be turned on or off.
	DataKey AutoActivateToggleStatusDataKeyEnum `mandatory:"true" json:"dataKey"`
}

func (m AutoActivateToggleStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoActivateToggleStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoActivateToggleStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAutoActivateToggleStatusStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoActivateToggleStatusDataKeyEnum(string(m.DataKey)); !ok && m.DataKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataKey: %s. Supported values are: %s.", m.DataKey, strings.Join(GetAutoActivateToggleStatusDataKeyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoActivateToggleStatusStatusEnum Enum with underlying type: string
type AutoActivateToggleStatusStatusEnum string

// Set of constants representing the allowable values for AutoActivateToggleStatusStatusEnum
const (
	AutoActivateToggleStatusStatusSuccess AutoActivateToggleStatusStatusEnum = "SUCCESS"
)

var mappingAutoActivateToggleStatusStatusEnum = map[string]AutoActivateToggleStatusStatusEnum{
	"SUCCESS": AutoActivateToggleStatusStatusSuccess,
}

var mappingAutoActivateToggleStatusStatusEnumLowerCase = map[string]AutoActivateToggleStatusStatusEnum{
	"success": AutoActivateToggleStatusStatusSuccess,
}

// GetAutoActivateToggleStatusStatusEnumValues Enumerates the set of values for AutoActivateToggleStatusStatusEnum
func GetAutoActivateToggleStatusStatusEnumValues() []AutoActivateToggleStatusStatusEnum {
	values := make([]AutoActivateToggleStatusStatusEnum, 0)
	for _, v := range mappingAutoActivateToggleStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoActivateToggleStatusStatusEnumStringValues Enumerates the set of values in String for AutoActivateToggleStatusStatusEnum
func GetAutoActivateToggleStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
	}
}

// GetMappingAutoActivateToggleStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoActivateToggleStatusStatusEnum(val string) (AutoActivateToggleStatusStatusEnum, bool) {
	enum, ok := mappingAutoActivateToggleStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutoActivateToggleStatusDataKeyEnum Enum with underlying type: string
type AutoActivateToggleStatusDataKeyEnum string

// Set of constants representing the allowable values for AutoActivateToggleStatusDataKeyEnum
const (
	AutoActivateToggleStatusDataKeyPrivateDataKey AutoActivateToggleStatusDataKeyEnum = "PRIVATE_DATA_KEY"
	AutoActivateToggleStatusDataKeyPublicDataKey  AutoActivateToggleStatusDataKeyEnum = "PUBLIC_DATA_KEY"
)

var mappingAutoActivateToggleStatusDataKeyEnum = map[string]AutoActivateToggleStatusDataKeyEnum{
	"PRIVATE_DATA_KEY": AutoActivateToggleStatusDataKeyPrivateDataKey,
	"PUBLIC_DATA_KEY":  AutoActivateToggleStatusDataKeyPublicDataKey,
}

var mappingAutoActivateToggleStatusDataKeyEnumLowerCase = map[string]AutoActivateToggleStatusDataKeyEnum{
	"private_data_key": AutoActivateToggleStatusDataKeyPrivateDataKey,
	"public_data_key":  AutoActivateToggleStatusDataKeyPublicDataKey,
}

// GetAutoActivateToggleStatusDataKeyEnumValues Enumerates the set of values for AutoActivateToggleStatusDataKeyEnum
func GetAutoActivateToggleStatusDataKeyEnumValues() []AutoActivateToggleStatusDataKeyEnum {
	values := make([]AutoActivateToggleStatusDataKeyEnum, 0)
	for _, v := range mappingAutoActivateToggleStatusDataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoActivateToggleStatusDataKeyEnumStringValues Enumerates the set of values in String for AutoActivateToggleStatusDataKeyEnum
func GetAutoActivateToggleStatusDataKeyEnumStringValues() []string {
	return []string{
		"PRIVATE_DATA_KEY",
		"PUBLIC_DATA_KEY",
	}
}

// GetMappingAutoActivateToggleStatusDataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoActivateToggleStatusDataKeyEnum(val string) (AutoActivateToggleStatusDataKeyEnum, bool) {
	enum, ok := mappingAutoActivateToggleStatusDataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
