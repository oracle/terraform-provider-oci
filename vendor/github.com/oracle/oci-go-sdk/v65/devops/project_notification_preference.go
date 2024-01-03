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

// ProjectNotificationPreference The notification preference of the project.
type ProjectNotificationPreference struct {

	// The ocid of project resource
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The ocid of user.
	UserId *string `mandatory:"true" json:"userId"`

	// The override value of project notification preference.
	NotificationPreference ProjectNotificationPreferenceNotificationPreferenceEnum `mandatory:"true" json:"notificationPreference"`
}

func (m ProjectNotificationPreference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProjectNotificationPreference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProjectNotificationPreferenceNotificationPreferenceEnum(string(m.NotificationPreference)); !ok && m.NotificationPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationPreference: %s. Supported values are: %s.", m.NotificationPreference, strings.Join(GetProjectNotificationPreferenceNotificationPreferenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProjectNotificationPreferenceNotificationPreferenceEnum Enum with underlying type: string
type ProjectNotificationPreferenceNotificationPreferenceEnum string

// Set of constants representing the allowable values for ProjectNotificationPreferenceNotificationPreferenceEnum
const (
	ProjectNotificationPreferenceNotificationPreferenceWatch   ProjectNotificationPreferenceNotificationPreferenceEnum = "WATCH"
	ProjectNotificationPreferenceNotificationPreferenceIgnore  ProjectNotificationPreferenceNotificationPreferenceEnum = "IGNORE"
	ProjectNotificationPreferenceNotificationPreferenceMention ProjectNotificationPreferenceNotificationPreferenceEnum = "MENTION"
)

var mappingProjectNotificationPreferenceNotificationPreferenceEnum = map[string]ProjectNotificationPreferenceNotificationPreferenceEnum{
	"WATCH":   ProjectNotificationPreferenceNotificationPreferenceWatch,
	"IGNORE":  ProjectNotificationPreferenceNotificationPreferenceIgnore,
	"MENTION": ProjectNotificationPreferenceNotificationPreferenceMention,
}

var mappingProjectNotificationPreferenceNotificationPreferenceEnumLowerCase = map[string]ProjectNotificationPreferenceNotificationPreferenceEnum{
	"watch":   ProjectNotificationPreferenceNotificationPreferenceWatch,
	"ignore":  ProjectNotificationPreferenceNotificationPreferenceIgnore,
	"mention": ProjectNotificationPreferenceNotificationPreferenceMention,
}

// GetProjectNotificationPreferenceNotificationPreferenceEnumValues Enumerates the set of values for ProjectNotificationPreferenceNotificationPreferenceEnum
func GetProjectNotificationPreferenceNotificationPreferenceEnumValues() []ProjectNotificationPreferenceNotificationPreferenceEnum {
	values := make([]ProjectNotificationPreferenceNotificationPreferenceEnum, 0)
	for _, v := range mappingProjectNotificationPreferenceNotificationPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetProjectNotificationPreferenceNotificationPreferenceEnumStringValues Enumerates the set of values in String for ProjectNotificationPreferenceNotificationPreferenceEnum
func GetProjectNotificationPreferenceNotificationPreferenceEnumStringValues() []string {
	return []string{
		"WATCH",
		"IGNORE",
		"MENTION",
	}
}

// GetMappingProjectNotificationPreferenceNotificationPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProjectNotificationPreferenceNotificationPreferenceEnum(val string) (ProjectNotificationPreferenceNotificationPreferenceEnum, bool) {
	enum, ok := mappingProjectNotificationPreferenceNotificationPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
