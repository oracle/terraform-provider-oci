// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceAnalyticsInstanceEnableOacDetails The configuration details for the enable OAC operation.
// Example: `{"attachmentType":"MANAGED","attachmentDetails":{...}}`
type ResourceAnalyticsInstanceEnableOacDetails struct {

	// The type of attachment the OAC instance is using.
	AttachmentType ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum `mandatory:"true" json:"attachmentType"`

	AttachmentDetails *ResourceAnalyticsInstanceOacAttachmentDetails `mandatory:"true" json:"attachmentDetails"`
}

func (m ResourceAnalyticsInstanceEnableOacDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceAnalyticsInstanceEnableOacDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum(string(m.AttachmentType)); !ok && m.AttachmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttachmentType: %s. Supported values are: %s.", m.AttachmentType, strings.Join(GetResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum Enum with underlying type: string
type ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum string

// Set of constants representing the allowable values for ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum
const (
	ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeManaged ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum = "MANAGED"
)

var mappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum = map[string]ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum{
	"MANAGED": ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeManaged,
}

var mappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumLowerCase = map[string]ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum{
	"managed": ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeManaged,
}

// GetResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumValues Enumerates the set of values for ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum
func GetResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumValues() []ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum {
	values := make([]ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum, 0)
	for _, v := range mappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumStringValues Enumerates the set of values in String for ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum
func GetResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
	}
}

// GetMappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum(val string) (ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum, bool) {
	enum, ok := mappingResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
