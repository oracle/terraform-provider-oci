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

// NamespaceSummary The is the namespace summary of a tenancy in Log Analytics application
type NamespaceSummary struct {

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
	LifecycleState NamespaceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m NamespaceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamespaceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNamespaceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNamespaceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NamespaceSummaryLifecycleStateEnum Enum with underlying type: string
type NamespaceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for NamespaceSummaryLifecycleStateEnum
const (
	NamespaceSummaryLifecycleStateActive   NamespaceSummaryLifecycleStateEnum = "ACTIVE"
	NamespaceSummaryLifecycleStateInactive NamespaceSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingNamespaceSummaryLifecycleStateEnum = map[string]NamespaceSummaryLifecycleStateEnum{
	"ACTIVE":   NamespaceSummaryLifecycleStateActive,
	"INACTIVE": NamespaceSummaryLifecycleStateInactive,
}

var mappingNamespaceSummaryLifecycleStateEnumLowerCase = map[string]NamespaceSummaryLifecycleStateEnum{
	"active":   NamespaceSummaryLifecycleStateActive,
	"inactive": NamespaceSummaryLifecycleStateInactive,
}

// GetNamespaceSummaryLifecycleStateEnumValues Enumerates the set of values for NamespaceSummaryLifecycleStateEnum
func GetNamespaceSummaryLifecycleStateEnumValues() []NamespaceSummaryLifecycleStateEnum {
	values := make([]NamespaceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingNamespaceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNamespaceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for NamespaceSummaryLifecycleStateEnum
func GetNamespaceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingNamespaceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamespaceSummaryLifecycleStateEnum(val string) (NamespaceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingNamespaceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
