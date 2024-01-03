// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogSummary Log object configuration summary.
type LogSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Log group OCID.
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// The user-friendly display name. This must be unique within the enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The pipeline state.
	LifecycleState LogLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The logType that the log object is for, whether custom or service.
	LogType LogSummaryLogTypeEnum `mandatory:"true" json:"logType"`

	// Whether or not this resource is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Configuration *Configuration `mandatory:"false" json:"configuration"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Log retention duration in 30-day increments (30, 60, 90 and so on until 180).
	RetentionDuration *int `mandatory:"false" json:"retentionDuration"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m LogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogSummaryLogTypeEnum(string(m.LogType)); !ok && m.LogType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", m.LogType, strings.Join(GetLogSummaryLogTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LogSummaryLogTypeEnum Enum with underlying type: string
type LogSummaryLogTypeEnum string

// Set of constants representing the allowable values for LogSummaryLogTypeEnum
const (
	LogSummaryLogTypeCustom  LogSummaryLogTypeEnum = "CUSTOM"
	LogSummaryLogTypeService LogSummaryLogTypeEnum = "SERVICE"
)

var mappingLogSummaryLogTypeEnum = map[string]LogSummaryLogTypeEnum{
	"CUSTOM":  LogSummaryLogTypeCustom,
	"SERVICE": LogSummaryLogTypeService,
}

var mappingLogSummaryLogTypeEnumLowerCase = map[string]LogSummaryLogTypeEnum{
	"custom":  LogSummaryLogTypeCustom,
	"service": LogSummaryLogTypeService,
}

// GetLogSummaryLogTypeEnumValues Enumerates the set of values for LogSummaryLogTypeEnum
func GetLogSummaryLogTypeEnumValues() []LogSummaryLogTypeEnum {
	values := make([]LogSummaryLogTypeEnum, 0)
	for _, v := range mappingLogSummaryLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogSummaryLogTypeEnumStringValues Enumerates the set of values in String for LogSummaryLogTypeEnum
func GetLogSummaryLogTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"SERVICE",
	}
}

// GetMappingLogSummaryLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogSummaryLogTypeEnum(val string) (LogSummaryLogTypeEnum, bool) {
	enum, ok := mappingLogSummaryLogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
