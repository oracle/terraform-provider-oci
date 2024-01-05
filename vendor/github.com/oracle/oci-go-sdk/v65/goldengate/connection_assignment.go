// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectionAssignment Represents the metadata description of a connection assignment.
// Before you can use a connection as a GoldenGate source or target, you must assign it to a deployment.
type ConnectionAssignment struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection assignment being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	ConnectionId *string `mandatory:"true" json:"connectionId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Possible lifecycle states for connection assignments.
	LifecycleState ConnectionAssignmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Credential store alias.
	AliasName *string `mandatory:"false" json:"aliasName"`
}

func (m ConnectionAssignment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectionAssignment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectionAssignmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionAssignmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionAssignmentLifecycleStateEnum Enum with underlying type: string
type ConnectionAssignmentLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectionAssignmentLifecycleStateEnum
const (
	ConnectionAssignmentLifecycleStateCreating ConnectionAssignmentLifecycleStateEnum = "CREATING"
	ConnectionAssignmentLifecycleStateActive   ConnectionAssignmentLifecycleStateEnum = "ACTIVE"
	ConnectionAssignmentLifecycleStateFailed   ConnectionAssignmentLifecycleStateEnum = "FAILED"
	ConnectionAssignmentLifecycleStateUpdating ConnectionAssignmentLifecycleStateEnum = "UPDATING"
	ConnectionAssignmentLifecycleStateDeleting ConnectionAssignmentLifecycleStateEnum = "DELETING"
)

var mappingConnectionAssignmentLifecycleStateEnum = map[string]ConnectionAssignmentLifecycleStateEnum{
	"CREATING": ConnectionAssignmentLifecycleStateCreating,
	"ACTIVE":   ConnectionAssignmentLifecycleStateActive,
	"FAILED":   ConnectionAssignmentLifecycleStateFailed,
	"UPDATING": ConnectionAssignmentLifecycleStateUpdating,
	"DELETING": ConnectionAssignmentLifecycleStateDeleting,
}

var mappingConnectionAssignmentLifecycleStateEnumLowerCase = map[string]ConnectionAssignmentLifecycleStateEnum{
	"creating": ConnectionAssignmentLifecycleStateCreating,
	"active":   ConnectionAssignmentLifecycleStateActive,
	"failed":   ConnectionAssignmentLifecycleStateFailed,
	"updating": ConnectionAssignmentLifecycleStateUpdating,
	"deleting": ConnectionAssignmentLifecycleStateDeleting,
}

// GetConnectionAssignmentLifecycleStateEnumValues Enumerates the set of values for ConnectionAssignmentLifecycleStateEnum
func GetConnectionAssignmentLifecycleStateEnumValues() []ConnectionAssignmentLifecycleStateEnum {
	values := make([]ConnectionAssignmentLifecycleStateEnum, 0)
	for _, v := range mappingConnectionAssignmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionAssignmentLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectionAssignmentLifecycleStateEnum
func GetConnectionAssignmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"FAILED",
		"UPDATING",
		"DELETING",
	}
}

// GetMappingConnectionAssignmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionAssignmentLifecycleStateEnum(val string) (ConnectionAssignmentLifecycleStateEnum, bool) {
	enum, ok := mappingConnectionAssignmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
