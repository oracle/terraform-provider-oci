// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Namespace This is the namespace details of a tenancy in Log Analytics application
type Namespace struct {

	// This is the namespace name of a tenancy
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The is the tenancy ID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This indicates if the tenancy is onboarded to Log Analytics
	IsOnboarded *bool `mandatory:"true" json:"isOnboarded"`

	// This indicates if the log set feature is enabled for the tenancy
	IsLogSetEnabled *bool `mandatory:"false" json:"isLogSetEnabled"`

	// This indicates if data has ever been ingested for the tenancy in Log Analytics
	IsDataEverIngested *bool `mandatory:"false" json:"isDataEverIngested"`

	// This indicates if old data can be archived for a tenancy
	IsArchivingEnabled *bool `mandatory:"false" json:"isArchivingEnabled"`

	// The current state of the compartment.
	LifecycleState NamespaceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Namespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Namespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNamespaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNamespaceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NamespaceLifecycleStateEnum Enum with underlying type: string
type NamespaceLifecycleStateEnum string

// Set of constants representing the allowable values for NamespaceLifecycleStateEnum
const (
	NamespaceLifecycleStateActive   NamespaceLifecycleStateEnum = "ACTIVE"
	NamespaceLifecycleStateInactive NamespaceLifecycleStateEnum = "INACTIVE"
)

var mappingNamespaceLifecycleStateEnum = map[string]NamespaceLifecycleStateEnum{
	"ACTIVE":   NamespaceLifecycleStateActive,
	"INACTIVE": NamespaceLifecycleStateInactive,
}

var mappingNamespaceLifecycleStateEnumLowerCase = map[string]NamespaceLifecycleStateEnum{
	"active":   NamespaceLifecycleStateActive,
	"inactive": NamespaceLifecycleStateInactive,
}

// GetNamespaceLifecycleStateEnumValues Enumerates the set of values for NamespaceLifecycleStateEnum
func GetNamespaceLifecycleStateEnumValues() []NamespaceLifecycleStateEnum {
	values := make([]NamespaceLifecycleStateEnum, 0)
	for _, v := range mappingNamespaceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNamespaceLifecycleStateEnumStringValues Enumerates the set of values in String for NamespaceLifecycleStateEnum
func GetNamespaceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingNamespaceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamespaceLifecycleStateEnum(val string) (NamespaceLifecycleStateEnum, bool) {
	enum, ok := mappingNamespaceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
