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

// Log Represents a log object.
type Log struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Log group OCID.
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The logType that the log object is for, whether custom or service.
	LogType LogLogTypeEnum `mandatory:"true" json:"logType"`

	// The pipeline state.
	LifecycleState LogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

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

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Log retention duration in 30-day increments (30, 60, 90 and so on).
	RetentionDuration *int `mandatory:"false" json:"retentionDuration"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m Log) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Log) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogLogTypeEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetLogLogTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogLogTypeEnum Enum with underlying type: string
type LogLogTypeEnum string

// Set of constants representing the allowable values for LogLogTypeEnum
const (
	LogLogTypeCustom  LogLogTypeEnum = "CUSTOM"
	LogLogTypeService LogLogTypeEnum = "SERVICE"
)

var mappingLogLogTypeEnum = map[string]LogLogTypeEnum{
	"CUSTOM":  LogLogTypeCustom,
	"SERVICE": LogLogTypeService,
}

// GetLogLogTypeEnumValues Enumerates the set of values for LogLogTypeEnum
func GetLogLogTypeEnumValues() []LogLogTypeEnum {
	values := make([]LogLogTypeEnum, 0)
	for _, v := range mappingLogLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogLogTypeEnumStringValues Enumerates the set of values in String for LogLogTypeEnum
func GetLogLogTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"SERVICE",
	}
}

// GetMappingLogLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogLogTypeEnum(val string) (LogLogTypeEnum, bool) {
	mappingLogLogTypeEnumIgnoreCase := make(map[string]LogLogTypeEnum)
	for k, v := range mappingLogLogTypeEnum {
		mappingLogLogTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLogLogTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
