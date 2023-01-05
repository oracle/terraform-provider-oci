// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// PublicLoggingDetails Log resource
type PublicLoggingDetails struct {

	// OCID of log object
	LogId *string `mandatory:"true" json:"logId"`

	// Tenancy of log object
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// Phone application ocid
	Resource *string `mandatory:"true" json:"resource"`

	// Category
	Category *string `mandatory:"true" json:"category"`

	// Log parameters
	Parameters map[string]string `mandatory:"false" json:"parameters"`

	// The time when log object was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time when log object was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Log object.
	LifecycleState PublicLoggingDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m PublicLoggingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicLoggingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPublicLoggingDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPublicLoggingDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicLoggingDetailsLifecycleStateEnum Enum with underlying type: string
type PublicLoggingDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for PublicLoggingDetailsLifecycleStateEnum
const (
	PublicLoggingDetailsLifecycleStateCreating PublicLoggingDetailsLifecycleStateEnum = "CREATING"
	PublicLoggingDetailsLifecycleStateUpdating PublicLoggingDetailsLifecycleStateEnum = "UPDATING"
	PublicLoggingDetailsLifecycleStateActive   PublicLoggingDetailsLifecycleStateEnum = "ACTIVE"
	PublicLoggingDetailsLifecycleStateDeleted  PublicLoggingDetailsLifecycleStateEnum = "DELETED"
	PublicLoggingDetailsLifecycleStateDeleting PublicLoggingDetailsLifecycleStateEnum = "DELETING"
	PublicLoggingDetailsLifecycleStateFailed   PublicLoggingDetailsLifecycleStateEnum = "FAILED"
)

var mappingPublicLoggingDetailsLifecycleStateEnum = map[string]PublicLoggingDetailsLifecycleStateEnum{
	"CREATING": PublicLoggingDetailsLifecycleStateCreating,
	"UPDATING": PublicLoggingDetailsLifecycleStateUpdating,
	"ACTIVE":   PublicLoggingDetailsLifecycleStateActive,
	"DELETED":  PublicLoggingDetailsLifecycleStateDeleted,
	"DELETING": PublicLoggingDetailsLifecycleStateDeleting,
	"FAILED":   PublicLoggingDetailsLifecycleStateFailed,
}

var mappingPublicLoggingDetailsLifecycleStateEnumLowerCase = map[string]PublicLoggingDetailsLifecycleStateEnum{
	"creating": PublicLoggingDetailsLifecycleStateCreating,
	"updating": PublicLoggingDetailsLifecycleStateUpdating,
	"active":   PublicLoggingDetailsLifecycleStateActive,
	"deleted":  PublicLoggingDetailsLifecycleStateDeleted,
	"deleting": PublicLoggingDetailsLifecycleStateDeleting,
	"failed":   PublicLoggingDetailsLifecycleStateFailed,
}

// GetPublicLoggingDetailsLifecycleStateEnumValues Enumerates the set of values for PublicLoggingDetailsLifecycleStateEnum
func GetPublicLoggingDetailsLifecycleStateEnumValues() []PublicLoggingDetailsLifecycleStateEnum {
	values := make([]PublicLoggingDetailsLifecycleStateEnum, 0)
	for _, v := range mappingPublicLoggingDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicLoggingDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for PublicLoggingDetailsLifecycleStateEnum
func GetPublicLoggingDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingPublicLoggingDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicLoggingDetailsLifecycleStateEnum(val string) (PublicLoggingDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingPublicLoggingDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
