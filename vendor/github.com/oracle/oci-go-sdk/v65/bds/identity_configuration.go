// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// IdentityConfiguration Details about the identity configuration
type IdentityConfiguration struct {

	// The id of the UPST config
	Id *string `mandatory:"true" json:"id"`

	// the display name of the identity configuration
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Identity domain to use for identity config
	IdentityDomainId *string `mandatory:"true" json:"identityDomainId"`

	// identity domain confidential application ID for the identity config
	ConfidentialApplicationId *string `mandatory:"true" json:"confidentialApplicationId"`

	// Lifecycle state of the identity configuration
	LifecycleState IdentityConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time when this identity configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when this identity configuration config was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	IamUserSyncConfiguration *IamUserSyncConfiguration `mandatory:"false" json:"iamUserSyncConfiguration"`

	UpstConfiguration *UpstConfiguration `mandatory:"false" json:"upstConfiguration"`
}

func (m IdentityConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdentityConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIdentityConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityConfigurationLifecycleStateEnum Enum with underlying type: string
type IdentityConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for IdentityConfigurationLifecycleStateEnum
const (
	IdentityConfigurationLifecycleStateActive  IdentityConfigurationLifecycleStateEnum = "ACTIVE"
	IdentityConfigurationLifecycleStateDeleted IdentityConfigurationLifecycleStateEnum = "DELETED"
)

var mappingIdentityConfigurationLifecycleStateEnum = map[string]IdentityConfigurationLifecycleStateEnum{
	"ACTIVE":  IdentityConfigurationLifecycleStateActive,
	"DELETED": IdentityConfigurationLifecycleStateDeleted,
}

var mappingIdentityConfigurationLifecycleStateEnumLowerCase = map[string]IdentityConfigurationLifecycleStateEnum{
	"active":  IdentityConfigurationLifecycleStateActive,
	"deleted": IdentityConfigurationLifecycleStateDeleted,
}

// GetIdentityConfigurationLifecycleStateEnumValues Enumerates the set of values for IdentityConfigurationLifecycleStateEnum
func GetIdentityConfigurationLifecycleStateEnumValues() []IdentityConfigurationLifecycleStateEnum {
	values := make([]IdentityConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingIdentityConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for IdentityConfigurationLifecycleStateEnum
func GetIdentityConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingIdentityConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityConfigurationLifecycleStateEnum(val string) (IdentityConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingIdentityConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
