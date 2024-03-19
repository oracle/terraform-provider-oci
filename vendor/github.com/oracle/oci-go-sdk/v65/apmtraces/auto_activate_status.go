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

// AutoActivateStatus Status of autoactivation for the given data key in the APM Domain.
type AutoActivateStatus struct {

	// State of autoactivation in this APM Domain.  If "ON" auto-activate is set to true, if "OFF" auto-activate is set to false.
	State *string `mandatory:"true" json:"state"`

	// Data key type for which auto-activate needs needs to be turned on or off.
	DataKey AutoActivateStatusDataKeyEnum `mandatory:"true" json:"dataKey"`
}

func (m AutoActivateStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoActivateStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoActivateStatusDataKeyEnum(string(m.DataKey)); !ok && m.DataKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataKey: %s. Supported values are: %s.", m.DataKey, strings.Join(GetAutoActivateStatusDataKeyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoActivateStatusDataKeyEnum Enum with underlying type: string
type AutoActivateStatusDataKeyEnum string

// Set of constants representing the allowable values for AutoActivateStatusDataKeyEnum
const (
	AutoActivateStatusDataKeyPrivateDataKey AutoActivateStatusDataKeyEnum = "PRIVATE_DATA_KEY"
	AutoActivateStatusDataKeyPublicDataKey  AutoActivateStatusDataKeyEnum = "PUBLIC_DATA_KEY"
)

var mappingAutoActivateStatusDataKeyEnum = map[string]AutoActivateStatusDataKeyEnum{
	"PRIVATE_DATA_KEY": AutoActivateStatusDataKeyPrivateDataKey,
	"PUBLIC_DATA_KEY":  AutoActivateStatusDataKeyPublicDataKey,
}

var mappingAutoActivateStatusDataKeyEnumLowerCase = map[string]AutoActivateStatusDataKeyEnum{
	"private_data_key": AutoActivateStatusDataKeyPrivateDataKey,
	"public_data_key":  AutoActivateStatusDataKeyPublicDataKey,
}

// GetAutoActivateStatusDataKeyEnumValues Enumerates the set of values for AutoActivateStatusDataKeyEnum
func GetAutoActivateStatusDataKeyEnumValues() []AutoActivateStatusDataKeyEnum {
	values := make([]AutoActivateStatusDataKeyEnum, 0)
	for _, v := range mappingAutoActivateStatusDataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoActivateStatusDataKeyEnumStringValues Enumerates the set of values in String for AutoActivateStatusDataKeyEnum
func GetAutoActivateStatusDataKeyEnumStringValues() []string {
	return []string{
		"PRIVATE_DATA_KEY",
		"PUBLIC_DATA_KEY",
	}
}

// GetMappingAutoActivateStatusDataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoActivateStatusDataKeyEnum(val string) (AutoActivateStatusDataKeyEnum, bool) {
	enum, ok := mappingAutoActivateStatusDataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
