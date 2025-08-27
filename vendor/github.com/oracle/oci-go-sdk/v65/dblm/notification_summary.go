// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotificationSummary Notifications
type NotificationSummary struct {

	// Notification type
	NotificationType NotificationSummaryNotificationTypeEnum `mandatory:"true" json:"notificationType"`

	// Notification text
	NotificationText *string `mandatory:"true" json:"notificationText"`

	// Notification identifier.
	Id *string `mandatory:"true" json:"id"`

	// Published date
	TimePublished *common.SDKTime `mandatory:"true" json:"timePublished"`
}

func (m NotificationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotificationSummaryNotificationTypeEnum(string(m.NotificationType)); !ok && m.NotificationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationType: %s. Supported values are: %s.", m.NotificationType, strings.Join(GetNotificationSummaryNotificationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NotificationSummaryNotificationTypeEnum Enum with underlying type: string
type NotificationSummaryNotificationTypeEnum string

// Set of constants representing the allowable values for NotificationSummaryNotificationTypeEnum
const (
	NotificationSummaryNotificationTypeCve      NotificationSummaryNotificationTypeEnum = "CVE"
	NotificationSummaryNotificationTypeAdvisory NotificationSummaryNotificationTypeEnum = "ADVISORY"
	NotificationSummaryNotificationTypePatch    NotificationSummaryNotificationTypeEnum = "PATCH"
)

var mappingNotificationSummaryNotificationTypeEnum = map[string]NotificationSummaryNotificationTypeEnum{
	"CVE":      NotificationSummaryNotificationTypeCve,
	"ADVISORY": NotificationSummaryNotificationTypeAdvisory,
	"PATCH":    NotificationSummaryNotificationTypePatch,
}

var mappingNotificationSummaryNotificationTypeEnumLowerCase = map[string]NotificationSummaryNotificationTypeEnum{
	"cve":      NotificationSummaryNotificationTypeCve,
	"advisory": NotificationSummaryNotificationTypeAdvisory,
	"patch":    NotificationSummaryNotificationTypePatch,
}

// GetNotificationSummaryNotificationTypeEnumValues Enumerates the set of values for NotificationSummaryNotificationTypeEnum
func GetNotificationSummaryNotificationTypeEnumValues() []NotificationSummaryNotificationTypeEnum {
	values := make([]NotificationSummaryNotificationTypeEnum, 0)
	for _, v := range mappingNotificationSummaryNotificationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationSummaryNotificationTypeEnumStringValues Enumerates the set of values in String for NotificationSummaryNotificationTypeEnum
func GetNotificationSummaryNotificationTypeEnumStringValues() []string {
	return []string{
		"CVE",
		"ADVISORY",
		"PATCH",
	}
}

// GetMappingNotificationSummaryNotificationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationSummaryNotificationTypeEnum(val string) (NotificationSummaryNotificationTypeEnum, bool) {
	enum, ok := mappingNotificationSummaryNotificationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
