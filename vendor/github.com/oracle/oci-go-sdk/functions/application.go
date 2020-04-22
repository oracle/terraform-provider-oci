// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Application An application contains functions and defined attributes shared between those functions, such as network configuration and configuration. Avoid entering confidential information.
type Application struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment that contains the application.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The display name of the application. The display name is unique within the compartment containing the application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the application.
	LifecycleState ApplicationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Application configuration for functions in this application (passed as environment variables). Can be overridden by function configuration.
	// Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.
	// Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	// The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
	Config map[string]string `mandatory:"false" json:"config"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application.
	SubnetIds []string `mandatory:"false" json:"subnetIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The time the application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the application was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m Application) String() string {
	return common.PointerString(m)
}

// ApplicationLifecycleStateEnum Enum with underlying type: string
type ApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationLifecycleStateEnum
const (
	ApplicationLifecycleStateCreating ApplicationLifecycleStateEnum = "CREATING"
	ApplicationLifecycleStateActive   ApplicationLifecycleStateEnum = "ACTIVE"
	ApplicationLifecycleStateInactive ApplicationLifecycleStateEnum = "INACTIVE"
	ApplicationLifecycleStateUpdating ApplicationLifecycleStateEnum = "UPDATING"
	ApplicationLifecycleStateDeleting ApplicationLifecycleStateEnum = "DELETING"
	ApplicationLifecycleStateDeleted  ApplicationLifecycleStateEnum = "DELETED"
	ApplicationLifecycleStateFailed   ApplicationLifecycleStateEnum = "FAILED"
)

var mappingApplicationLifecycleState = map[string]ApplicationLifecycleStateEnum{
	"CREATING": ApplicationLifecycleStateCreating,
	"ACTIVE":   ApplicationLifecycleStateActive,
	"INACTIVE": ApplicationLifecycleStateInactive,
	"UPDATING": ApplicationLifecycleStateUpdating,
	"DELETING": ApplicationLifecycleStateDeleting,
	"DELETED":  ApplicationLifecycleStateDeleted,
	"FAILED":   ApplicationLifecycleStateFailed,
}

// GetApplicationLifecycleStateEnumValues Enumerates the set of values for ApplicationLifecycleStateEnum
func GetApplicationLifecycleStateEnumValues() []ApplicationLifecycleStateEnum {
	values := make([]ApplicationLifecycleStateEnum, 0)
	for _, v := range mappingApplicationLifecycleState {
		values = append(values, v)
	}
	return values
}
