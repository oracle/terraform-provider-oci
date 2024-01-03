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

// ConsoleHistory The details of the Db Node console history.
type ConsoleHistory struct {

	// The OCID of the console history.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the console history.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the database node.
	DbNodeId *string `mandatory:"true" json:"dbNodeId"`

	// The current state of the console history.
	LifecycleState ConsoleHistoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m ConsoleHistory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsoleHistory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConsoleHistoryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConsoleHistoryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConsoleHistoryLifecycleStateEnum Enum with underlying type: string
type ConsoleHistoryLifecycleStateEnum string

// Set of constants representing the allowable values for ConsoleHistoryLifecycleStateEnum
const (
	ConsoleHistoryLifecycleStateRequested      ConsoleHistoryLifecycleStateEnum = "REQUESTED"
	ConsoleHistoryLifecycleStateGettingHistory ConsoleHistoryLifecycleStateEnum = "GETTING_HISTORY"
	ConsoleHistoryLifecycleStateSucceeded      ConsoleHistoryLifecycleStateEnum = "SUCCEEDED"
	ConsoleHistoryLifecycleStateFailed         ConsoleHistoryLifecycleStateEnum = "FAILED"
	ConsoleHistoryLifecycleStateDeleted        ConsoleHistoryLifecycleStateEnum = "DELETED"
	ConsoleHistoryLifecycleStateDeleting       ConsoleHistoryLifecycleStateEnum = "DELETING"
)

var mappingConsoleHistoryLifecycleStateEnum = map[string]ConsoleHistoryLifecycleStateEnum{
	"REQUESTED":       ConsoleHistoryLifecycleStateRequested,
	"GETTING_HISTORY": ConsoleHistoryLifecycleStateGettingHistory,
	"SUCCEEDED":       ConsoleHistoryLifecycleStateSucceeded,
	"FAILED":          ConsoleHistoryLifecycleStateFailed,
	"DELETED":         ConsoleHistoryLifecycleStateDeleted,
	"DELETING":        ConsoleHistoryLifecycleStateDeleting,
}

var mappingConsoleHistoryLifecycleStateEnumLowerCase = map[string]ConsoleHistoryLifecycleStateEnum{
	"requested":       ConsoleHistoryLifecycleStateRequested,
	"getting_history": ConsoleHistoryLifecycleStateGettingHistory,
	"succeeded":       ConsoleHistoryLifecycleStateSucceeded,
	"failed":          ConsoleHistoryLifecycleStateFailed,
	"deleted":         ConsoleHistoryLifecycleStateDeleted,
	"deleting":        ConsoleHistoryLifecycleStateDeleting,
}

// GetConsoleHistoryLifecycleStateEnumValues Enumerates the set of values for ConsoleHistoryLifecycleStateEnum
func GetConsoleHistoryLifecycleStateEnumValues() []ConsoleHistoryLifecycleStateEnum {
	values := make([]ConsoleHistoryLifecycleStateEnum, 0)
	for _, v := range mappingConsoleHistoryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConsoleHistoryLifecycleStateEnumStringValues Enumerates the set of values in String for ConsoleHistoryLifecycleStateEnum
func GetConsoleHistoryLifecycleStateEnumStringValues() []string {
	return []string{
		"REQUESTED",
		"GETTING_HISTORY",
		"SUCCEEDED",
		"FAILED",
		"DELETED",
		"DELETING",
	}
}

// GetMappingConsoleHistoryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsoleHistoryLifecycleStateEnum(val string) (ConsoleHistoryLifecycleStateEnum, bool) {
	enum, ok := mappingConsoleHistoryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
