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

// ConsoleConnectionSummary The `InstanceConsoleConnection` API provides you with console access to dbnode
// enabling you to troubleshoot malfunctioning dbnode.
type ConsoleConnectionSummary struct {

	// The OCID of the console connection.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment to contain the console connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the database node.
	DbNodeId *string `mandatory:"true" json:"dbNodeId"`

	// The SSH connection string for the console connection.
	ConnectionString *string `mandatory:"true" json:"connectionString"`

	// The SSH public key fingerprint for the console connection.
	Fingerprint *string `mandatory:"true" json:"fingerprint"`

	// The current state of the console connection.
	LifecycleState ConsoleConnectionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The SSH public key's fingerprint for the console connection service host.
	ServiceHostKeyFingerprint *string `mandatory:"false" json:"serviceHostKeyFingerprint"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ConsoleConnectionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConsoleConnectionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConsoleConnectionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConsoleConnectionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConsoleConnectionSummaryLifecycleStateEnum Enum with underlying type: string
type ConsoleConnectionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ConsoleConnectionSummaryLifecycleStateEnum
const (
	ConsoleConnectionSummaryLifecycleStateActive   ConsoleConnectionSummaryLifecycleStateEnum = "ACTIVE"
	ConsoleConnectionSummaryLifecycleStateCreating ConsoleConnectionSummaryLifecycleStateEnum = "CREATING"
	ConsoleConnectionSummaryLifecycleStateDeleted  ConsoleConnectionSummaryLifecycleStateEnum = "DELETED"
	ConsoleConnectionSummaryLifecycleStateDeleting ConsoleConnectionSummaryLifecycleStateEnum = "DELETING"
	ConsoleConnectionSummaryLifecycleStateFailed   ConsoleConnectionSummaryLifecycleStateEnum = "FAILED"
)

var mappingConsoleConnectionSummaryLifecycleStateEnum = map[string]ConsoleConnectionSummaryLifecycleStateEnum{
	"ACTIVE":   ConsoleConnectionSummaryLifecycleStateActive,
	"CREATING": ConsoleConnectionSummaryLifecycleStateCreating,
	"DELETED":  ConsoleConnectionSummaryLifecycleStateDeleted,
	"DELETING": ConsoleConnectionSummaryLifecycleStateDeleting,
	"FAILED":   ConsoleConnectionSummaryLifecycleStateFailed,
}

var mappingConsoleConnectionSummaryLifecycleStateEnumLowerCase = map[string]ConsoleConnectionSummaryLifecycleStateEnum{
	"active":   ConsoleConnectionSummaryLifecycleStateActive,
	"creating": ConsoleConnectionSummaryLifecycleStateCreating,
	"deleted":  ConsoleConnectionSummaryLifecycleStateDeleted,
	"deleting": ConsoleConnectionSummaryLifecycleStateDeleting,
	"failed":   ConsoleConnectionSummaryLifecycleStateFailed,
}

// GetConsoleConnectionSummaryLifecycleStateEnumValues Enumerates the set of values for ConsoleConnectionSummaryLifecycleStateEnum
func GetConsoleConnectionSummaryLifecycleStateEnumValues() []ConsoleConnectionSummaryLifecycleStateEnum {
	values := make([]ConsoleConnectionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingConsoleConnectionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConsoleConnectionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ConsoleConnectionSummaryLifecycleStateEnum
func GetConsoleConnectionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingConsoleConnectionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConsoleConnectionSummaryLifecycleStateEnum(val string) (ConsoleConnectionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingConsoleConnectionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
