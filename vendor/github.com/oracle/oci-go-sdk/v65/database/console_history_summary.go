// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConsoleHistorySummary The details of the Db Node console history.
type ConsoleHistorySummary struct {

	// The OCID of the console history.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the console history.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the database node.
	DbNodeId *string `mandatory:"true" json:"dbNodeId"`

	// The current state of the console history.
	LifecycleState ConsoleHistorySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the console history was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The user-friendly name for the console history. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ConsoleHistorySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsoleHistorySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConsoleHistorySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConsoleHistorySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConsoleHistorySummaryLifecycleStateEnum Enum with underlying type: string
type ConsoleHistorySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ConsoleHistorySummaryLifecycleStateEnum
const (
	ConsoleHistorySummaryLifecycleStateRequested      ConsoleHistorySummaryLifecycleStateEnum = "REQUESTED"
	ConsoleHistorySummaryLifecycleStateGettingHistory ConsoleHistorySummaryLifecycleStateEnum = "GETTING_HISTORY"
	ConsoleHistorySummaryLifecycleStateSucceeded      ConsoleHistorySummaryLifecycleStateEnum = "SUCCEEDED"
	ConsoleHistorySummaryLifecycleStateFailed         ConsoleHistorySummaryLifecycleStateEnum = "FAILED"
	ConsoleHistorySummaryLifecycleStateDeleted        ConsoleHistorySummaryLifecycleStateEnum = "DELETED"
	ConsoleHistorySummaryLifecycleStateDeleting       ConsoleHistorySummaryLifecycleStateEnum = "DELETING"
)

var mappingConsoleHistorySummaryLifecycleStateEnum = map[string]ConsoleHistorySummaryLifecycleStateEnum{
	"REQUESTED":       ConsoleHistorySummaryLifecycleStateRequested,
	"GETTING_HISTORY": ConsoleHistorySummaryLifecycleStateGettingHistory,
	"SUCCEEDED":       ConsoleHistorySummaryLifecycleStateSucceeded,
	"FAILED":          ConsoleHistorySummaryLifecycleStateFailed,
	"DELETED":         ConsoleHistorySummaryLifecycleStateDeleted,
	"DELETING":        ConsoleHistorySummaryLifecycleStateDeleting,
}

var mappingConsoleHistorySummaryLifecycleStateEnumLowerCase = map[string]ConsoleHistorySummaryLifecycleStateEnum{
	"requested":       ConsoleHistorySummaryLifecycleStateRequested,
	"getting_history": ConsoleHistorySummaryLifecycleStateGettingHistory,
	"succeeded":       ConsoleHistorySummaryLifecycleStateSucceeded,
	"failed":          ConsoleHistorySummaryLifecycleStateFailed,
	"deleted":         ConsoleHistorySummaryLifecycleStateDeleted,
	"deleting":        ConsoleHistorySummaryLifecycleStateDeleting,
}

// GetConsoleHistorySummaryLifecycleStateEnumValues Enumerates the set of values for ConsoleHistorySummaryLifecycleStateEnum
func GetConsoleHistorySummaryLifecycleStateEnumValues() []ConsoleHistorySummaryLifecycleStateEnum {
	values := make([]ConsoleHistorySummaryLifecycleStateEnum, 0)
	for _, v := range mappingConsoleHistorySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConsoleHistorySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ConsoleHistorySummaryLifecycleStateEnum
func GetConsoleHistorySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"REQUESTED",
		"GETTING_HISTORY",
		"SUCCEEDED",
		"FAILED",
		"DELETED",
		"DELETING",
	}
}

// GetMappingConsoleHistorySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsoleHistorySummaryLifecycleStateEnum(val string) (ConsoleHistorySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingConsoleHistorySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
