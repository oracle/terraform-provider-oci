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

// BdsLakeConfiguration The lake configuration information.
type BdsLakeConfiguration struct {

	// The ID of the lake configuration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the lake.
	LakeId *string `mandatory:"true" json:"lakeId"`

	// The current state of the lake configuration lifecycle.
	LifecycleState BdsLakeConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the configuration was created. It is displayed in an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the lake configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the BDS API key used for the lake configuration.
	BdsApiKeyId *string `mandatory:"false" json:"bdsApiKeyId"`

	// The time when the configuration was updated. It is displayed in an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsLakeConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsLakeConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsLakeConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsLakeConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsLakeConfigurationLifecycleStateEnum Enum with underlying type: string
type BdsLakeConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for BdsLakeConfigurationLifecycleStateEnum
const (
	BdsLakeConfigurationLifecycleStateCreating   BdsLakeConfigurationLifecycleStateEnum = "CREATING"
	BdsLakeConfigurationLifecycleStateActivating BdsLakeConfigurationLifecycleStateEnum = "ACTIVATING"
	BdsLakeConfigurationLifecycleStateActive     BdsLakeConfigurationLifecycleStateEnum = "ACTIVE"
	BdsLakeConfigurationLifecycleStateInactive   BdsLakeConfigurationLifecycleStateEnum = "INACTIVE"
	BdsLakeConfigurationLifecycleStateUpdating   BdsLakeConfigurationLifecycleStateEnum = "UPDATING"
	BdsLakeConfigurationLifecycleStateFailed     BdsLakeConfigurationLifecycleStateEnum = "FAILED"
	BdsLakeConfigurationLifecycleStateDeleting   BdsLakeConfigurationLifecycleStateEnum = "DELETING"
	BdsLakeConfigurationLifecycleStateDeleted    BdsLakeConfigurationLifecycleStateEnum = "DELETED"
)

var mappingBdsLakeConfigurationLifecycleStateEnum = map[string]BdsLakeConfigurationLifecycleStateEnum{
	"CREATING":   BdsLakeConfigurationLifecycleStateCreating,
	"ACTIVATING": BdsLakeConfigurationLifecycleStateActivating,
	"ACTIVE":     BdsLakeConfigurationLifecycleStateActive,
	"INACTIVE":   BdsLakeConfigurationLifecycleStateInactive,
	"UPDATING":   BdsLakeConfigurationLifecycleStateUpdating,
	"FAILED":     BdsLakeConfigurationLifecycleStateFailed,
	"DELETING":   BdsLakeConfigurationLifecycleStateDeleting,
	"DELETED":    BdsLakeConfigurationLifecycleStateDeleted,
}

var mappingBdsLakeConfigurationLifecycleStateEnumLowerCase = map[string]BdsLakeConfigurationLifecycleStateEnum{
	"creating":   BdsLakeConfigurationLifecycleStateCreating,
	"activating": BdsLakeConfigurationLifecycleStateActivating,
	"active":     BdsLakeConfigurationLifecycleStateActive,
	"inactive":   BdsLakeConfigurationLifecycleStateInactive,
	"updating":   BdsLakeConfigurationLifecycleStateUpdating,
	"failed":     BdsLakeConfigurationLifecycleStateFailed,
	"deleting":   BdsLakeConfigurationLifecycleStateDeleting,
	"deleted":    BdsLakeConfigurationLifecycleStateDeleted,
}

// GetBdsLakeConfigurationLifecycleStateEnumValues Enumerates the set of values for BdsLakeConfigurationLifecycleStateEnum
func GetBdsLakeConfigurationLifecycleStateEnumValues() []BdsLakeConfigurationLifecycleStateEnum {
	values := make([]BdsLakeConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingBdsLakeConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsLakeConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for BdsLakeConfigurationLifecycleStateEnum
func GetBdsLakeConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBdsLakeConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsLakeConfigurationLifecycleStateEnum(val string) (BdsLakeConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingBdsLakeConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
