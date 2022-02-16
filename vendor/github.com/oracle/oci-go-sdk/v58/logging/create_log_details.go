// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateLogDetails The details to create a log object.
type CreateLogDetails struct {

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The logType that the log object is for, whether custom or service.
	LogType CreateLogDetailsLogTypeEnum `mandatory:"true" json:"logType"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Configuration *Configuration `mandatory:"false" json:"configuration"`

	// Log retention duration in 30-day increments (30, 60, 90 and so on).
	RetentionDuration *int `mandatory:"false" json:"retentionDuration"`
}

func (m CreateLogDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLogDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateLogDetailsLogTypeEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetCreateLogDetailsLogTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateLogDetailsLogTypeEnum Enum with underlying type: string
type CreateLogDetailsLogTypeEnum string

// Set of constants representing the allowable values for CreateLogDetailsLogTypeEnum
const (
	CreateLogDetailsLogTypeCustom  CreateLogDetailsLogTypeEnum = "CUSTOM"
	CreateLogDetailsLogTypeService CreateLogDetailsLogTypeEnum = "SERVICE"
)

var mappingCreateLogDetailsLogTypeEnum = map[string]CreateLogDetailsLogTypeEnum{
	"CUSTOM":  CreateLogDetailsLogTypeCustom,
	"SERVICE": CreateLogDetailsLogTypeService,
}

// GetCreateLogDetailsLogTypeEnumValues Enumerates the set of values for CreateLogDetailsLogTypeEnum
func GetCreateLogDetailsLogTypeEnumValues() []CreateLogDetailsLogTypeEnum {
	values := make([]CreateLogDetailsLogTypeEnum, 0)
	for _, v := range mappingCreateLogDetailsLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateLogDetailsLogTypeEnumStringValues Enumerates the set of values in String for CreateLogDetailsLogTypeEnum
func GetCreateLogDetailsLogTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"SERVICE",
	}
}

// GetMappingCreateLogDetailsLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateLogDetailsLogTypeEnum(val string) (CreateLogDetailsLogTypeEnum, bool) {
	mappingCreateLogDetailsLogTypeEnumIgnoreCase := make(map[string]CreateLogDetailsLogTypeEnum)
	for k, v := range mappingCreateLogDetailsLogTypeEnum {
		mappingCreateLogDetailsLogTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateLogDetailsLogTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
