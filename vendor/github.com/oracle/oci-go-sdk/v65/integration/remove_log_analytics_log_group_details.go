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

// RemoveLogAnalyticsLogGroupDetails Input payload to remove Log Analytics Log Group for given IntegrationInstance or its supported attachments.
// Some actions may not be applicable to specific integration types.
type RemoveLogAnalyticsLogGroupDetails struct {

	// Type of attachment. Supported at this include PROCESS_AUTOMATION
	AttachmentType RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum `mandatory:"false" json:"attachmentType,omitempty"`
}

func (m RemoveLogAnalyticsLogGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveLogAnalyticsLogGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum Enum with underlying type: string
type RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum
const (
	RemoveLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum = "PROCESS_AUTOMATION"
)

var mappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum = map[string]RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum{
	"PROCESS_AUTOMATION": RemoveLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation,
}

var mappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumLowerCase = map[string]RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum{
	"process_automation": RemoveLogAnalyticsLogGroupDetailsAttachmentTypeProcessAutomation,
}

// GetRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumValues Enumerates the set of values for RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum
func GetRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumValues() []RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum {
	values := make([]RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum
func GetRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"PROCESS_AUTOMATION",
	}
}

// GetMappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum(val string) (RemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingRemoveLogAnalyticsLogGroupDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
