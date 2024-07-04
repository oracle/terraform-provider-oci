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

// ResourcePrincipalConfiguration Resource Principal Session Token Details.
type ResourcePrincipalConfiguration struct {

	// The id of the ResourcePrincipalConfiguration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the bdsInstance which is the parent resource id.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Life span in hours of each resource principal session token.
	SessionTokenLifeSpanDurationInHours *int `mandatory:"true" json:"sessionTokenLifeSpanDurationInHours"`

	// The state of the ResourcePrincipalConfiguration.
	LifecycleState ResourcePrincipalConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the ResourcePrincipalConfiguration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the ResourcePrincipalConfiguration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// the time the resource principal session token was refreshed, shown as an rfc 3339 formatted datetime string.
	TimeTokenRefreshed *common.SDKTime `mandatory:"false" json:"timeTokenRefreshed"`

	// the time the resource principal session token will expired, shown as an rfc 3339 formatted datetime string.
	TimeTokenExpiry *common.SDKTime `mandatory:"false" json:"timeTokenExpiry"`
}

func (m ResourcePrincipalConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourcePrincipalConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourcePrincipalConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourcePrincipalConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourcePrincipalConfigurationLifecycleStateEnum Enum with underlying type: string
type ResourcePrincipalConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ResourcePrincipalConfigurationLifecycleStateEnum
const (
	ResourcePrincipalConfigurationLifecycleStateCreating ResourcePrincipalConfigurationLifecycleStateEnum = "CREATING"
	ResourcePrincipalConfigurationLifecycleStateActive   ResourcePrincipalConfigurationLifecycleStateEnum = "ACTIVE"
	ResourcePrincipalConfigurationLifecycleStateUpdating ResourcePrincipalConfigurationLifecycleStateEnum = "UPDATING"
	ResourcePrincipalConfigurationLifecycleStateDeleting ResourcePrincipalConfigurationLifecycleStateEnum = "DELETING"
	ResourcePrincipalConfigurationLifecycleStateDeleted  ResourcePrincipalConfigurationLifecycleStateEnum = "DELETED"
	ResourcePrincipalConfigurationLifecycleStateFailed   ResourcePrincipalConfigurationLifecycleStateEnum = "FAILED"
)

var mappingResourcePrincipalConfigurationLifecycleStateEnum = map[string]ResourcePrincipalConfigurationLifecycleStateEnum{
	"CREATING": ResourcePrincipalConfigurationLifecycleStateCreating,
	"ACTIVE":   ResourcePrincipalConfigurationLifecycleStateActive,
	"UPDATING": ResourcePrincipalConfigurationLifecycleStateUpdating,
	"DELETING": ResourcePrincipalConfigurationLifecycleStateDeleting,
	"DELETED":  ResourcePrincipalConfigurationLifecycleStateDeleted,
	"FAILED":   ResourcePrincipalConfigurationLifecycleStateFailed,
}

var mappingResourcePrincipalConfigurationLifecycleStateEnumLowerCase = map[string]ResourcePrincipalConfigurationLifecycleStateEnum{
	"creating": ResourcePrincipalConfigurationLifecycleStateCreating,
	"active":   ResourcePrincipalConfigurationLifecycleStateActive,
	"updating": ResourcePrincipalConfigurationLifecycleStateUpdating,
	"deleting": ResourcePrincipalConfigurationLifecycleStateDeleting,
	"deleted":  ResourcePrincipalConfigurationLifecycleStateDeleted,
	"failed":   ResourcePrincipalConfigurationLifecycleStateFailed,
}

// GetResourcePrincipalConfigurationLifecycleStateEnumValues Enumerates the set of values for ResourcePrincipalConfigurationLifecycleStateEnum
func GetResourcePrincipalConfigurationLifecycleStateEnumValues() []ResourcePrincipalConfigurationLifecycleStateEnum {
	values := make([]ResourcePrincipalConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingResourcePrincipalConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResourcePrincipalConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for ResourcePrincipalConfigurationLifecycleStateEnum
func GetResourcePrincipalConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingResourcePrincipalConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourcePrincipalConfigurationLifecycleStateEnum(val string) (ResourcePrincipalConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingResourcePrincipalConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
