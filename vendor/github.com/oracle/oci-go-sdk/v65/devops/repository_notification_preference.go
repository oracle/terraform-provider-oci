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

// RepositoryNotificationPreference The notification preference of the repository.
type RepositoryNotificationPreference struct {

	// The ocid of repository resource
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The ocid of user.
	UserId *string `mandatory:"true" json:"userId"`

	// The override value of repository notification preference.
	NotificationPreference RepositoryNotificationPreferenceNotificationPreferenceEnum `mandatory:"true" json:"notificationPreference"`
}

func (m RepositoryNotificationPreference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryNotificationPreference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRepositoryNotificationPreferenceNotificationPreferenceEnum(string(m.NotificationPreference)); !ok && m.NotificationPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationPreference: %s. Supported values are: %s.", m.NotificationPreference, strings.Join(GetRepositoryNotificationPreferenceNotificationPreferenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RepositoryNotificationPreferenceNotificationPreferenceEnum Enum with underlying type: string
type RepositoryNotificationPreferenceNotificationPreferenceEnum string

// Set of constants representing the allowable values for RepositoryNotificationPreferenceNotificationPreferenceEnum
const (
	RepositoryNotificationPreferenceNotificationPreferenceWatch     RepositoryNotificationPreferenceNotificationPreferenceEnum = "WATCH"
	RepositoryNotificationPreferenceNotificationPreferenceIgnore    RepositoryNotificationPreferenceNotificationPreferenceEnum = "IGNORE"
	RepositoryNotificationPreferenceNotificationPreferenceMention   RepositoryNotificationPreferenceNotificationPreferenceEnum = "MENTION"
	RepositoryNotificationPreferenceNotificationPreferenceInherited RepositoryNotificationPreferenceNotificationPreferenceEnum = "INHERITED"
)

var mappingRepositoryNotificationPreferenceNotificationPreferenceEnum = map[string]RepositoryNotificationPreferenceNotificationPreferenceEnum{
	"WATCH":     RepositoryNotificationPreferenceNotificationPreferenceWatch,
	"IGNORE":    RepositoryNotificationPreferenceNotificationPreferenceIgnore,
	"MENTION":   RepositoryNotificationPreferenceNotificationPreferenceMention,
	"INHERITED": RepositoryNotificationPreferenceNotificationPreferenceInherited,
}

var mappingRepositoryNotificationPreferenceNotificationPreferenceEnumLowerCase = map[string]RepositoryNotificationPreferenceNotificationPreferenceEnum{
	"watch":     RepositoryNotificationPreferenceNotificationPreferenceWatch,
	"ignore":    RepositoryNotificationPreferenceNotificationPreferenceIgnore,
	"mention":   RepositoryNotificationPreferenceNotificationPreferenceMention,
	"inherited": RepositoryNotificationPreferenceNotificationPreferenceInherited,
}

// GetRepositoryNotificationPreferenceNotificationPreferenceEnumValues Enumerates the set of values for RepositoryNotificationPreferenceNotificationPreferenceEnum
func GetRepositoryNotificationPreferenceNotificationPreferenceEnumValues() []RepositoryNotificationPreferenceNotificationPreferenceEnum {
	values := make([]RepositoryNotificationPreferenceNotificationPreferenceEnum, 0)
	for _, v := range mappingRepositoryNotificationPreferenceNotificationPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryNotificationPreferenceNotificationPreferenceEnumStringValues Enumerates the set of values in String for RepositoryNotificationPreferenceNotificationPreferenceEnum
func GetRepositoryNotificationPreferenceNotificationPreferenceEnumStringValues() []string {
	return []string{
		"WATCH",
		"IGNORE",
		"MENTION",
		"INHERITED",
	}
}

// GetMappingRepositoryNotificationPreferenceNotificationPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryNotificationPreferenceNotificationPreferenceEnum(val string) (RepositoryNotificationPreferenceNotificationPreferenceEnum, bool) {
	enum, ok := mappingRepositoryNotificationPreferenceNotificationPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
