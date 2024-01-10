// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectHarness Detailed representation of a connect harness.
type ConnectHarness struct {

	// The name of the connect harness. Avoid entering confidential information.
	// Example: `JDBCConnector`
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the connect harness.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the connect harness.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the connect harness.
	LifecycleState ConnectHarnessLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the connect harness was created, expressed in in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Any additional details about the current state of the connect harness.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}'
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ConnectHarness) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectHarness) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectHarnessLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectHarnessLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectHarnessLifecycleStateEnum Enum with underlying type: string
type ConnectHarnessLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectHarnessLifecycleStateEnum
const (
	ConnectHarnessLifecycleStateCreating ConnectHarnessLifecycleStateEnum = "CREATING"
	ConnectHarnessLifecycleStateActive   ConnectHarnessLifecycleStateEnum = "ACTIVE"
	ConnectHarnessLifecycleStateDeleting ConnectHarnessLifecycleStateEnum = "DELETING"
	ConnectHarnessLifecycleStateDeleted  ConnectHarnessLifecycleStateEnum = "DELETED"
	ConnectHarnessLifecycleStateFailed   ConnectHarnessLifecycleStateEnum = "FAILED"
	ConnectHarnessLifecycleStateUpdating ConnectHarnessLifecycleStateEnum = "UPDATING"
)

var mappingConnectHarnessLifecycleStateEnum = map[string]ConnectHarnessLifecycleStateEnum{
	"CREATING": ConnectHarnessLifecycleStateCreating,
	"ACTIVE":   ConnectHarnessLifecycleStateActive,
	"DELETING": ConnectHarnessLifecycleStateDeleting,
	"DELETED":  ConnectHarnessLifecycleStateDeleted,
	"FAILED":   ConnectHarnessLifecycleStateFailed,
	"UPDATING": ConnectHarnessLifecycleStateUpdating,
}

var mappingConnectHarnessLifecycleStateEnumLowerCase = map[string]ConnectHarnessLifecycleStateEnum{
	"creating": ConnectHarnessLifecycleStateCreating,
	"active":   ConnectHarnessLifecycleStateActive,
	"deleting": ConnectHarnessLifecycleStateDeleting,
	"deleted":  ConnectHarnessLifecycleStateDeleted,
	"failed":   ConnectHarnessLifecycleStateFailed,
	"updating": ConnectHarnessLifecycleStateUpdating,
}

// GetConnectHarnessLifecycleStateEnumValues Enumerates the set of values for ConnectHarnessLifecycleStateEnum
func GetConnectHarnessLifecycleStateEnumValues() []ConnectHarnessLifecycleStateEnum {
	values := make([]ConnectHarnessLifecycleStateEnum, 0)
	for _, v := range mappingConnectHarnessLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectHarnessLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectHarnessLifecycleStateEnum
func GetConnectHarnessLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingConnectHarnessLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectHarnessLifecycleStateEnum(val string) (ConnectHarnessLifecycleStateEnum, bool) {
	enum, ok := mappingConnectHarnessLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
