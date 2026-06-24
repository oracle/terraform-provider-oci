// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddLogAnalyticsLogGroupDetails Input payload to ADD Log Analytics Log Group for given IntegrationInstance.
// Some actions may not be applicable to specific integration types.
type AddLogAnalyticsLogGroupDetails struct {

	// Log Group ocid.
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// Type of attachment. Supported at this include PROCESS_AUTOMATION
	AttachmentType AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum `mandatory:"false" json:"attachmentType,omitempty"`
}

func (m AddLogAnalyticsLogGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddLogAnalyticsLogGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum Enum with underlying type: string
type AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum
const (
	AddLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum = "PROCESS_AUTOMATION"
)

var mappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnum = map[string]AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum{
	"PROCESS_AUTOMATION": AddLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation,
}

var mappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumLowerCase = map[string]AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum{
	"process_automation": AddLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation,
}

// GetAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumValues Enumerates the set of values for AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum
func GetAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumValues() []AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum {
	values := make([]AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum
func GetAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"PROCESS_AUTOMATION",
	}
}

// GetMappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnum(val string) (AddLogAnalyticsLogGroupDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingAddLogAnalyticsLogGroupDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
