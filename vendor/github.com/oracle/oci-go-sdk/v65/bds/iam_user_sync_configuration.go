// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IamUserSyncConfiguration Information about the IAM user sync configuration.
type IamUserSyncConfiguration struct {

	// whether to append POSIX attributes to IAM users
	IsPosixAttributesAdditionRequired *bool `mandatory:"true" json:"isPosixAttributesAdditionRequired"`

	// Lifecycle state of the IAM user sync config
	LifecycleState IamUserSyncConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time when this IAM user sync config was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when this IAM user sync config was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m IamUserSyncConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IamUserSyncConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIamUserSyncConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIamUserSyncConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IamUserSyncConfigurationLifecycleStateEnum Enum with underlying type: string
type IamUserSyncConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for IamUserSyncConfigurationLifecycleStateEnum
const (
	IamUserSyncConfigurationLifecycleStateCreating IamUserSyncConfigurationLifecycleStateEnum = "CREATING"
	IamUserSyncConfigurationLifecycleStateActive   IamUserSyncConfigurationLifecycleStateEnum = "ACTIVE"
	IamUserSyncConfigurationLifecycleStateInactive IamUserSyncConfigurationLifecycleStateEnum = "INACTIVE"
	IamUserSyncConfigurationLifecycleStateDeleting IamUserSyncConfigurationLifecycleStateEnum = "DELETING"
	IamUserSyncConfigurationLifecycleStateUpdating IamUserSyncConfigurationLifecycleStateEnum = "UPDATING"
	IamUserSyncConfigurationLifecycleStateFailed   IamUserSyncConfigurationLifecycleStateEnum = "FAILED"
)

var mappingIamUserSyncConfigurationLifecycleStateEnum = map[string]IamUserSyncConfigurationLifecycleStateEnum{
	"CREATING": IamUserSyncConfigurationLifecycleStateCreating,
	"ACTIVE":   IamUserSyncConfigurationLifecycleStateActive,
	"INACTIVE": IamUserSyncConfigurationLifecycleStateInactive,
	"DELETING": IamUserSyncConfigurationLifecycleStateDeleting,
	"UPDATING": IamUserSyncConfigurationLifecycleStateUpdating,
	"FAILED":   IamUserSyncConfigurationLifecycleStateFailed,
}

var mappingIamUserSyncConfigurationLifecycleStateEnumLowerCase = map[string]IamUserSyncConfigurationLifecycleStateEnum{
	"creating": IamUserSyncConfigurationLifecycleStateCreating,
	"active":   IamUserSyncConfigurationLifecycleStateActive,
	"inactive": IamUserSyncConfigurationLifecycleStateInactive,
	"deleting": IamUserSyncConfigurationLifecycleStateDeleting,
	"updating": IamUserSyncConfigurationLifecycleStateUpdating,
	"failed":   IamUserSyncConfigurationLifecycleStateFailed,
}

// GetIamUserSyncConfigurationLifecycleStateEnumValues Enumerates the set of values for IamUserSyncConfigurationLifecycleStateEnum
func GetIamUserSyncConfigurationLifecycleStateEnumValues() []IamUserSyncConfigurationLifecycleStateEnum {
	values := make([]IamUserSyncConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingIamUserSyncConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIamUserSyncConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for IamUserSyncConfigurationLifecycleStateEnum
func GetIamUserSyncConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingIamUserSyncConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIamUserSyncConfigurationLifecycleStateEnum(val string) (IamUserSyncConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingIamUserSyncConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
