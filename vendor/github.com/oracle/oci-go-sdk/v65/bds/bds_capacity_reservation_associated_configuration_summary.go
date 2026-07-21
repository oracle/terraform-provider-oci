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

// BdsCapacityReservationAssociatedConfigurationSummary Summary of a BDS capacity reservation configuration associated with a BDS capacity reservation.
type BdsCapacityReservationAssociatedConfigurationSummary struct {

	// The OCID of the BDS capacity reservation configuration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the BDS cluster linked through the BDS capacity reservation configuration.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// The display name of the BDS capacity reservation configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The lifecycle state of the BDS capacity reservation configuration.
	LifecycleState BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the compartment that contains the BDS cluster.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time the configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsCapacityReservationAssociatedConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsCapacityReservationAssociatedConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum Enum with underlying type: string
type BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum
const (
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateCreating BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "CREATING"
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateActive   BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "ACTIVE"
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateInactive BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "INACTIVE"
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateUpdating BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "UPDATING"
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleting BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "DELETING"
	BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleted  BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = "DELETED"
)

var mappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum = map[string]BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum{
	"CREATING": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateCreating,
	"ACTIVE":   BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateActive,
	"INACTIVE": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateInactive,
	"UPDATING": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateUpdating,
	"DELETING": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleting,
	"DELETED":  BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleted,
}

var mappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumLowerCase = map[string]BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum{
	"creating": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateCreating,
	"active":   BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateActive,
	"inactive": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateInactive,
	"updating": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateUpdating,
	"deleting": BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleting,
	"deleted":  BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateDeleted,
}

// GetBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumValues Enumerates the set of values for BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum
func GetBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumValues() []BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum {
	values := make([]BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum
func GetBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum(val string) (BdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBdsCapacityReservationAssociatedConfigurationSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
