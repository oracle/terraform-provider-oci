// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PullRequestNotificationPreference The notification preference of the pull request.
type PullRequestNotificationPreference struct {

	// The ocid of pull request resource
	PullRequestId *string `mandatory:"true" json:"pullRequestId"`

	// The ocid of user.
	UserId *string `mandatory:"true" json:"userId"`

	// The override value of pull request notification preference.
	NotificationPreference PullRequestNotificationPreferenceNotificationPreferenceEnum `mandatory:"true" json:"notificationPreference"`
}

func (m PullRequestNotificationPreference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PullRequestNotificationPreference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPullRequestNotificationPreferenceNotificationPreferenceEnum(string(m.NotificationPreference)); !ok && m.NotificationPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationPreference: %s. Supported values are: %s.", m.NotificationPreference, strings.Join(GetPullRequestNotificationPreferenceNotificationPreferenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PullRequestNotificationPreferenceNotificationPreferenceEnum Enum with underlying type: string
type PullRequestNotificationPreferenceNotificationPreferenceEnum string

// Set of constants representing the allowable values for PullRequestNotificationPreferenceNotificationPreferenceEnum
const (
	PullRequestNotificationPreferenceNotificationPreferenceWatch   PullRequestNotificationPreferenceNotificationPreferenceEnum = "WATCH"
	PullRequestNotificationPreferenceNotificationPreferenceIgnore  PullRequestNotificationPreferenceNotificationPreferenceEnum = "IGNORE"
	PullRequestNotificationPreferenceNotificationPreferenceMention PullRequestNotificationPreferenceNotificationPreferenceEnum = "MENTION"
)

var mappingPullRequestNotificationPreferenceNotificationPreferenceEnum = map[string]PullRequestNotificationPreferenceNotificationPreferenceEnum{
	"WATCH":   PullRequestNotificationPreferenceNotificationPreferenceWatch,
	"IGNORE":  PullRequestNotificationPreferenceNotificationPreferenceIgnore,
	"MENTION": PullRequestNotificationPreferenceNotificationPreferenceMention,
}

var mappingPullRequestNotificationPreferenceNotificationPreferenceEnumLowerCase = map[string]PullRequestNotificationPreferenceNotificationPreferenceEnum{
	"watch":   PullRequestNotificationPreferenceNotificationPreferenceWatch,
	"ignore":  PullRequestNotificationPreferenceNotificationPreferenceIgnore,
	"mention": PullRequestNotificationPreferenceNotificationPreferenceMention,
}

// GetPullRequestNotificationPreferenceNotificationPreferenceEnumValues Enumerates the set of values for PullRequestNotificationPreferenceNotificationPreferenceEnum
func GetPullRequestNotificationPreferenceNotificationPreferenceEnumValues() []PullRequestNotificationPreferenceNotificationPreferenceEnum {
	values := make([]PullRequestNotificationPreferenceNotificationPreferenceEnum, 0)
	for _, v := range mappingPullRequestNotificationPreferenceNotificationPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetPullRequestNotificationPreferenceNotificationPreferenceEnumStringValues Enumerates the set of values in String for PullRequestNotificationPreferenceNotificationPreferenceEnum
func GetPullRequestNotificationPreferenceNotificationPreferenceEnumStringValues() []string {
	return []string{
		"WATCH",
		"IGNORE",
		"MENTION",
	}
}

// GetMappingPullRequestNotificationPreferenceNotificationPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPullRequestNotificationPreferenceNotificationPreferenceEnum(val string) (PullRequestNotificationPreferenceNotificationPreferenceEnum, bool) {
	enum, ok := mappingPullRequestNotificationPreferenceNotificationPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
