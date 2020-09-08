// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateLogDetails The details to create a log object.
type CreateLogDetails struct {

	// The display name of a user-friendly name. It has to be unique within enclosing resource,
	// and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The logType that the log object is for, custom or service.
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

	// Log retention duration in days.
	RetentionDuration *int `mandatory:"false" json:"retentionDuration"`
}

func (m CreateLogDetails) String() string {
	return common.PointerString(m)
}

// CreateLogDetailsLogTypeEnum Enum with underlying type: string
type CreateLogDetailsLogTypeEnum string

// Set of constants representing the allowable values for CreateLogDetailsLogTypeEnum
const (
	CreateLogDetailsLogTypeCustom  CreateLogDetailsLogTypeEnum = "CUSTOM"
	CreateLogDetailsLogTypeService CreateLogDetailsLogTypeEnum = "SERVICE"
)

var mappingCreateLogDetailsLogType = map[string]CreateLogDetailsLogTypeEnum{
	"CUSTOM":  CreateLogDetailsLogTypeCustom,
	"SERVICE": CreateLogDetailsLogTypeService,
}

// GetCreateLogDetailsLogTypeEnumValues Enumerates the set of values for CreateLogDetailsLogTypeEnum
func GetCreateLogDetailsLogTypeEnumValues() []CreateLogDetailsLogTypeEnum {
	values := make([]CreateLogDetailsLogTypeEnum, 0)
	for _, v := range mappingCreateLogDetailsLogType {
		values = append(values, v)
	}
	return values
}
