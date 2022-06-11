// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// BdsLakehouseConfiguration The lakehouse configuration information.
type BdsLakehouseConfiguration struct {

	// The ID of the lakehouse configuration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the lakehouse.
	LakehouseId *string `mandatory:"true" json:"lakehouseId"`

	// The current state of the lakehouse configuration lifecycle.
	LifecycleState BdsLakehouseConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the configuration was created. It is displayed in an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the lakehouse configuration
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the BDS API key used for the lakehouse configuration.
	BdsApiKeyId *string `mandatory:"false" json:"bdsApiKeyId"`

	// The time when the configuration was updated. It is displayed in an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsLakehouseConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsLakehouseConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsLakehouseConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsLakehouseConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsLakehouseConfigurationLifecycleStateEnum Enum with underlying type: string
type BdsLakehouseConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for BdsLakehouseConfigurationLifecycleStateEnum
const (
	BdsLakehouseConfigurationLifecycleStateCreating   BdsLakehouseConfigurationLifecycleStateEnum = "CREATING"
	BdsLakehouseConfigurationLifecycleStateActivating BdsLakehouseConfigurationLifecycleStateEnum = "ACTIVATING"
	BdsLakehouseConfigurationLifecycleStateActive     BdsLakehouseConfigurationLifecycleStateEnum = "ACTIVE"
	BdsLakehouseConfigurationLifecycleStateInactive   BdsLakehouseConfigurationLifecycleStateEnum = "INACTIVE"
	BdsLakehouseConfigurationLifecycleStateUpdating   BdsLakehouseConfigurationLifecycleStateEnum = "UPDATING"
	BdsLakehouseConfigurationLifecycleStateFailed     BdsLakehouseConfigurationLifecycleStateEnum = "FAILED"
	BdsLakehouseConfigurationLifecycleStateDeleting   BdsLakehouseConfigurationLifecycleStateEnum = "DELETING"
	BdsLakehouseConfigurationLifecycleStateDeleted    BdsLakehouseConfigurationLifecycleStateEnum = "DELETED"
)

var mappingBdsLakehouseConfigurationLifecycleStateEnum = map[string]BdsLakehouseConfigurationLifecycleStateEnum{
	"CREATING":   BdsLakehouseConfigurationLifecycleStateCreating,
	"ACTIVATING": BdsLakehouseConfigurationLifecycleStateActivating,
	"ACTIVE":     BdsLakehouseConfigurationLifecycleStateActive,
	"INACTIVE":   BdsLakehouseConfigurationLifecycleStateInactive,
	"UPDATING":   BdsLakehouseConfigurationLifecycleStateUpdating,
	"FAILED":     BdsLakehouseConfigurationLifecycleStateFailed,
	"DELETING":   BdsLakehouseConfigurationLifecycleStateDeleting,
	"DELETED":    BdsLakehouseConfigurationLifecycleStateDeleted,
}

var mappingBdsLakehouseConfigurationLifecycleStateEnumLowerCase = map[string]BdsLakehouseConfigurationLifecycleStateEnum{
	"creating":   BdsLakehouseConfigurationLifecycleStateCreating,
	"activating": BdsLakehouseConfigurationLifecycleStateActivating,
	"active":     BdsLakehouseConfigurationLifecycleStateActive,
	"inactive":   BdsLakehouseConfigurationLifecycleStateInactive,
	"updating":   BdsLakehouseConfigurationLifecycleStateUpdating,
	"failed":     BdsLakehouseConfigurationLifecycleStateFailed,
	"deleting":   BdsLakehouseConfigurationLifecycleStateDeleting,
	"deleted":    BdsLakehouseConfigurationLifecycleStateDeleted,
}

// GetBdsLakehouseConfigurationLifecycleStateEnumValues Enumerates the set of values for BdsLakehouseConfigurationLifecycleStateEnum
func GetBdsLakehouseConfigurationLifecycleStateEnumValues() []BdsLakehouseConfigurationLifecycleStateEnum {
	values := make([]BdsLakehouseConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingBdsLakehouseConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsLakehouseConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for BdsLakehouseConfigurationLifecycleStateEnum
func GetBdsLakehouseConfigurationLifecycleStateEnumStringValues() []string {
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

// GetMappingBdsLakehouseConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsLakehouseConfigurationLifecycleStateEnum(val string) (BdsLakehouseConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingBdsLakehouseConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
