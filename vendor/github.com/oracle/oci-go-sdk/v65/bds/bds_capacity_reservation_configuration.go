// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// BdsCapacityReservationConfiguration A configuration between a BDS cluster and a BDS capacity reservation.
type BdsCapacityReservationConfiguration struct {

	// The OCID of the BDS capacity reservation configuration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the BDS cluster associated with the BDS capacity reservation.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// The OCID of the BDS capacity reservation associated with the BDS cluster.
	BdsCapacityReservationId *string `mandatory:"true" json:"bdsCapacityReservationId"`

	// The display name of the BDS capacity reservation configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The lifecycle state of the BDS capacity reservation configuration.
	LifecycleState BdsCapacityReservationConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the BDS capacity reservation configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the BDS capacity reservation configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsCapacityReservationConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsCapacityReservationConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsCapacityReservationConfigurationLifecycleStateEnum Enum with underlying type: string
type BdsCapacityReservationConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for BdsCapacityReservationConfigurationLifecycleStateEnum
const (
	BdsCapacityReservationConfigurationLifecycleStateCreating BdsCapacityReservationConfigurationLifecycleStateEnum = "CREATING"
	BdsCapacityReservationConfigurationLifecycleStateActive   BdsCapacityReservationConfigurationLifecycleStateEnum = "ACTIVE"
	BdsCapacityReservationConfigurationLifecycleStateInactive BdsCapacityReservationConfigurationLifecycleStateEnum = "INACTIVE"
	BdsCapacityReservationConfigurationLifecycleStateUpdating BdsCapacityReservationConfigurationLifecycleStateEnum = "UPDATING"
	BdsCapacityReservationConfigurationLifecycleStateDeleting BdsCapacityReservationConfigurationLifecycleStateEnum = "DELETING"
	BdsCapacityReservationConfigurationLifecycleStateDeleted  BdsCapacityReservationConfigurationLifecycleStateEnum = "DELETED"
)

var mappingBdsCapacityReservationConfigurationLifecycleStateEnum = map[string]BdsCapacityReservationConfigurationLifecycleStateEnum{
	"CREATING": BdsCapacityReservationConfigurationLifecycleStateCreating,
	"ACTIVE":   BdsCapacityReservationConfigurationLifecycleStateActive,
	"INACTIVE": BdsCapacityReservationConfigurationLifecycleStateInactive,
	"UPDATING": BdsCapacityReservationConfigurationLifecycleStateUpdating,
	"DELETING": BdsCapacityReservationConfigurationLifecycleStateDeleting,
	"DELETED":  BdsCapacityReservationConfigurationLifecycleStateDeleted,
}

var mappingBdsCapacityReservationConfigurationLifecycleStateEnumLowerCase = map[string]BdsCapacityReservationConfigurationLifecycleStateEnum{
	"creating": BdsCapacityReservationConfigurationLifecycleStateCreating,
	"active":   BdsCapacityReservationConfigurationLifecycleStateActive,
	"inactive": BdsCapacityReservationConfigurationLifecycleStateInactive,
	"updating": BdsCapacityReservationConfigurationLifecycleStateUpdating,
	"deleting": BdsCapacityReservationConfigurationLifecycleStateDeleting,
	"deleted":  BdsCapacityReservationConfigurationLifecycleStateDeleted,
}

// GetBdsCapacityReservationConfigurationLifecycleStateEnumValues Enumerates the set of values for BdsCapacityReservationConfigurationLifecycleStateEnum
func GetBdsCapacityReservationConfigurationLifecycleStateEnumValues() []BdsCapacityReservationConfigurationLifecycleStateEnum {
	values := make([]BdsCapacityReservationConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingBdsCapacityReservationConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for BdsCapacityReservationConfigurationLifecycleStateEnum
func GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum(val string) (BdsCapacityReservationConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingBdsCapacityReservationConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
