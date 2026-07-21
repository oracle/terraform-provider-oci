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

// BdsCapacityReservation A reusable BDS capacity reservation resource.
type BdsCapacityReservation struct {

	// The OCID of the BDS capacity reservation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the BDS capacity reservation.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the BDS capacity reservation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ComputeCapacityReservations *ComputeCapacityReservations `mandatory:"true" json:"computeCapacityReservations"`

	// The lifecycle state of the BDS capacity reservation.
	LifecycleState BdsCapacityReservationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The time the BDS capacity reservation was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the BDS capacity reservation was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m BdsCapacityReservation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsCapacityReservation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBdsCapacityReservationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsCapacityReservationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BdsCapacityReservationLifecycleStateEnum Enum with underlying type: string
type BdsCapacityReservationLifecycleStateEnum string

// Set of constants representing the allowable values for BdsCapacityReservationLifecycleStateEnum
const (
	BdsCapacityReservationLifecycleStateCreating BdsCapacityReservationLifecycleStateEnum = "CREATING"
	BdsCapacityReservationLifecycleStateActive   BdsCapacityReservationLifecycleStateEnum = "ACTIVE"
	BdsCapacityReservationLifecycleStateUpdating BdsCapacityReservationLifecycleStateEnum = "UPDATING"
	BdsCapacityReservationLifecycleStateDeleting BdsCapacityReservationLifecycleStateEnum = "DELETING"
	BdsCapacityReservationLifecycleStateDeleted  BdsCapacityReservationLifecycleStateEnum = "DELETED"
)

var mappingBdsCapacityReservationLifecycleStateEnum = map[string]BdsCapacityReservationLifecycleStateEnum{
	"CREATING": BdsCapacityReservationLifecycleStateCreating,
	"ACTIVE":   BdsCapacityReservationLifecycleStateActive,
	"UPDATING": BdsCapacityReservationLifecycleStateUpdating,
	"DELETING": BdsCapacityReservationLifecycleStateDeleting,
	"DELETED":  BdsCapacityReservationLifecycleStateDeleted,
}

var mappingBdsCapacityReservationLifecycleStateEnumLowerCase = map[string]BdsCapacityReservationLifecycleStateEnum{
	"creating": BdsCapacityReservationLifecycleStateCreating,
	"active":   BdsCapacityReservationLifecycleStateActive,
	"updating": BdsCapacityReservationLifecycleStateUpdating,
	"deleting": BdsCapacityReservationLifecycleStateDeleting,
	"deleted":  BdsCapacityReservationLifecycleStateDeleted,
}

// GetBdsCapacityReservationLifecycleStateEnumValues Enumerates the set of values for BdsCapacityReservationLifecycleStateEnum
func GetBdsCapacityReservationLifecycleStateEnumValues() []BdsCapacityReservationLifecycleStateEnum {
	values := make([]BdsCapacityReservationLifecycleStateEnum, 0)
	for _, v := range mappingBdsCapacityReservationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBdsCapacityReservationLifecycleStateEnumStringValues Enumerates the set of values in String for BdsCapacityReservationLifecycleStateEnum
func GetBdsCapacityReservationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
	}
}

// GetMappingBdsCapacityReservationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBdsCapacityReservationLifecycleStateEnum(val string) (BdsCapacityReservationLifecycleStateEnum, bool) {
	enum, ok := mappingBdsCapacityReservationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
