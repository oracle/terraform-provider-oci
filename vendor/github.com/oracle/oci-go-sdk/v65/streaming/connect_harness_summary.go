// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ConnectHarnessSummary Summary representation of a ConnectHarness.
type ConnectHarnessSummary struct {

	// The name of the connect harness.
	// Example: `TelemetryEvents`
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the connect harness.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the connect harness.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the connect harness.
	LifecycleState ConnectHarnessSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the connect harness was created, expressed in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2018-04-20T00:00:07.405Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ConnectHarnessSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectHarnessSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectHarnessSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectHarnessSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectHarnessSummaryLifecycleStateEnum Enum with underlying type: string
type ConnectHarnessSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ConnectHarnessSummaryLifecycleStateEnum
const (
	ConnectHarnessSummaryLifecycleStateCreating ConnectHarnessSummaryLifecycleStateEnum = "CREATING"
	ConnectHarnessSummaryLifecycleStateActive   ConnectHarnessSummaryLifecycleStateEnum = "ACTIVE"
	ConnectHarnessSummaryLifecycleStateDeleting ConnectHarnessSummaryLifecycleStateEnum = "DELETING"
	ConnectHarnessSummaryLifecycleStateDeleted  ConnectHarnessSummaryLifecycleStateEnum = "DELETED"
	ConnectHarnessSummaryLifecycleStateFailed   ConnectHarnessSummaryLifecycleStateEnum = "FAILED"
	ConnectHarnessSummaryLifecycleStateUpdating ConnectHarnessSummaryLifecycleStateEnum = "UPDATING"
)

var mappingConnectHarnessSummaryLifecycleStateEnum = map[string]ConnectHarnessSummaryLifecycleStateEnum{
	"CREATING": ConnectHarnessSummaryLifecycleStateCreating,
	"ACTIVE":   ConnectHarnessSummaryLifecycleStateActive,
	"DELETING": ConnectHarnessSummaryLifecycleStateDeleting,
	"DELETED":  ConnectHarnessSummaryLifecycleStateDeleted,
	"FAILED":   ConnectHarnessSummaryLifecycleStateFailed,
	"UPDATING": ConnectHarnessSummaryLifecycleStateUpdating,
}

var mappingConnectHarnessSummaryLifecycleStateEnumLowerCase = map[string]ConnectHarnessSummaryLifecycleStateEnum{
	"creating": ConnectHarnessSummaryLifecycleStateCreating,
	"active":   ConnectHarnessSummaryLifecycleStateActive,
	"deleting": ConnectHarnessSummaryLifecycleStateDeleting,
	"deleted":  ConnectHarnessSummaryLifecycleStateDeleted,
	"failed":   ConnectHarnessSummaryLifecycleStateFailed,
	"updating": ConnectHarnessSummaryLifecycleStateUpdating,
}

// GetConnectHarnessSummaryLifecycleStateEnumValues Enumerates the set of values for ConnectHarnessSummaryLifecycleStateEnum
func GetConnectHarnessSummaryLifecycleStateEnumValues() []ConnectHarnessSummaryLifecycleStateEnum {
	values := make([]ConnectHarnessSummaryLifecycleStateEnum, 0)
	for _, v := range mappingConnectHarnessSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectHarnessSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ConnectHarnessSummaryLifecycleStateEnum
func GetConnectHarnessSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingConnectHarnessSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectHarnessSummaryLifecycleStateEnum(val string) (ConnectHarnessSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingConnectHarnessSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
