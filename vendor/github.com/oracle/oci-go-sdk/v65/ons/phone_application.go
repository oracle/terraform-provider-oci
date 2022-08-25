// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PhoneApplication Collection of phone numbers that can send messages to your customers.
type PhoneApplication struct {

	// The OCID of this resource that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// PhoneApplication display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the PhoneApplication was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the PhoneApplication.
	LifecycleState PhoneApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Messages endpoint
	MessagesEndpoint *string `mandatory:"true" json:"messagesEndpoint"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the PhoneApplication was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PhoneApplication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PhoneApplication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPhoneApplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPhoneApplicationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PhoneApplicationLifecycleStateEnum Enum with underlying type: string
type PhoneApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for PhoneApplicationLifecycleStateEnum
const (
	PhoneApplicationLifecycleStateCreating PhoneApplicationLifecycleStateEnum = "CREATING"
	PhoneApplicationLifecycleStateUpdating PhoneApplicationLifecycleStateEnum = "UPDATING"
	PhoneApplicationLifecycleStateActive   PhoneApplicationLifecycleStateEnum = "ACTIVE"
	PhoneApplicationLifecycleStateDeleting PhoneApplicationLifecycleStateEnum = "DELETING"
	PhoneApplicationLifecycleStateDeleted  PhoneApplicationLifecycleStateEnum = "DELETED"
	PhoneApplicationLifecycleStateFailed   PhoneApplicationLifecycleStateEnum = "FAILED"
)

var mappingPhoneApplicationLifecycleStateEnum = map[string]PhoneApplicationLifecycleStateEnum{
	"CREATING": PhoneApplicationLifecycleStateCreating,
	"UPDATING": PhoneApplicationLifecycleStateUpdating,
	"ACTIVE":   PhoneApplicationLifecycleStateActive,
	"DELETING": PhoneApplicationLifecycleStateDeleting,
	"DELETED":  PhoneApplicationLifecycleStateDeleted,
	"FAILED":   PhoneApplicationLifecycleStateFailed,
}

var mappingPhoneApplicationLifecycleStateEnumLowerCase = map[string]PhoneApplicationLifecycleStateEnum{
	"creating": PhoneApplicationLifecycleStateCreating,
	"updating": PhoneApplicationLifecycleStateUpdating,
	"active":   PhoneApplicationLifecycleStateActive,
	"deleting": PhoneApplicationLifecycleStateDeleting,
	"deleted":  PhoneApplicationLifecycleStateDeleted,
	"failed":   PhoneApplicationLifecycleStateFailed,
}

// GetPhoneApplicationLifecycleStateEnumValues Enumerates the set of values for PhoneApplicationLifecycleStateEnum
func GetPhoneApplicationLifecycleStateEnumValues() []PhoneApplicationLifecycleStateEnum {
	values := make([]PhoneApplicationLifecycleStateEnum, 0)
	for _, v := range mappingPhoneApplicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPhoneApplicationLifecycleStateEnumStringValues Enumerates the set of values in String for PhoneApplicationLifecycleStateEnum
func GetPhoneApplicationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPhoneApplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPhoneApplicationLifecycleStateEnum(val string) (PhoneApplicationLifecycleStateEnum, bool) {
	enum, ok := mappingPhoneApplicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
