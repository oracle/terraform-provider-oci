// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotificationPreferenceConfigCategoryDetails Notification preference
// Defines notification preferences set for compartment or tenancy.
type NotificationPreferenceConfigCategoryDetails struct {

	// The OCID of the resource.
	TopicId *string `mandatory:"true" json:"topicId"`

	// OCID of the compartment to which the resource belongs to.
	TopicCompartmentId *string `mandatory:"true" json:"topicCompartmentId"`

	Preferences *Preferences `mandatory:"false" json:"preferences"`

	// Specifies the scope of the notification preference created.
	Scope NotificationPreferenceConfigCategoryDetailsScopeEnum `mandatory:"true" json:"scope"`
}

func (m NotificationPreferenceConfigCategoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotificationPreferenceConfigCategoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNotificationPreferenceConfigCategoryDetailsScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetNotificationPreferenceConfigCategoryDetailsScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NotificationPreferenceConfigCategoryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNotificationPreferenceConfigCategoryDetails NotificationPreferenceConfigCategoryDetails
	s := struct {
		DiscriminatorParam string `json:"configCategory"`
		MarshalTypeNotificationPreferenceConfigCategoryDetails
	}{
		"NOTIFICATION_PREFERENCE",
		(MarshalTypeNotificationPreferenceConfigCategoryDetails)(m),
	}

	return json.Marshal(&s)
}

// NotificationPreferenceConfigCategoryDetailsScopeEnum Enum with underlying type: string
type NotificationPreferenceConfigCategoryDetailsScopeEnum string

// Set of constants representing the allowable values for NotificationPreferenceConfigCategoryDetailsScopeEnum
const (
	NotificationPreferenceConfigCategoryDetailsScopeCompartment NotificationPreferenceConfigCategoryDetailsScopeEnum = "COMPARTMENT"
	NotificationPreferenceConfigCategoryDetailsScopeTenancy     NotificationPreferenceConfigCategoryDetailsScopeEnum = "TENANCY"
)

var mappingNotificationPreferenceConfigCategoryDetailsScopeEnum = map[string]NotificationPreferenceConfigCategoryDetailsScopeEnum{
	"COMPARTMENT": NotificationPreferenceConfigCategoryDetailsScopeCompartment,
	"TENANCY":     NotificationPreferenceConfigCategoryDetailsScopeTenancy,
}

var mappingNotificationPreferenceConfigCategoryDetailsScopeEnumLowerCase = map[string]NotificationPreferenceConfigCategoryDetailsScopeEnum{
	"compartment": NotificationPreferenceConfigCategoryDetailsScopeCompartment,
	"tenancy":     NotificationPreferenceConfigCategoryDetailsScopeTenancy,
}

// GetNotificationPreferenceConfigCategoryDetailsScopeEnumValues Enumerates the set of values for NotificationPreferenceConfigCategoryDetailsScopeEnum
func GetNotificationPreferenceConfigCategoryDetailsScopeEnumValues() []NotificationPreferenceConfigCategoryDetailsScopeEnum {
	values := make([]NotificationPreferenceConfigCategoryDetailsScopeEnum, 0)
	for _, v := range mappingNotificationPreferenceConfigCategoryDetailsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetNotificationPreferenceConfigCategoryDetailsScopeEnumStringValues Enumerates the set of values in String for NotificationPreferenceConfigCategoryDetailsScopeEnum
func GetNotificationPreferenceConfigCategoryDetailsScopeEnumStringValues() []string {
	return []string{
		"COMPARTMENT",
		"TENANCY",
	}
}

// GetMappingNotificationPreferenceConfigCategoryDetailsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNotificationPreferenceConfigCategoryDetailsScopeEnum(val string) (NotificationPreferenceConfigCategoryDetailsScopeEnum, bool) {
	enum, ok := mappingNotificationPreferenceConfigCategoryDetailsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
