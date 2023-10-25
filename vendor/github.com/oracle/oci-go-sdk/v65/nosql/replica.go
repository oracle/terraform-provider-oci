// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Replica Information about a MR table replica
type Replica struct {

	// A customer-facing region identifier
	Region *string `mandatory:"false" json:"region"`

	// The OCID of the replica table
	TableId *string `mandatory:"false" json:"tableId"`

	// Maximum sustained write throughput limit of the replica table.
	MaxWriteUnits *int `mandatory:"false" json:"maxWriteUnits"`

	// The capacity mode of the replica.
	CapacityMode ReplicaCapacityModeEnum `mandatory:"false" json:"capacityMode,omitempty"`

	// The state of the replica.
	LifecycleState ReplicaLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m Replica) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Replica) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReplicaCapacityModeEnum(string(m.CapacityMode)); !ok && m.CapacityMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CapacityMode: %s. Supported values are: %s.", m.CapacityMode, strings.Join(GetReplicaCapacityModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReplicaLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicaLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicaCapacityModeEnum Enum with underlying type: string
type ReplicaCapacityModeEnum string

// Set of constants representing the allowable values for ReplicaCapacityModeEnum
const (
	ReplicaCapacityModeProvisioned ReplicaCapacityModeEnum = "PROVISIONED"
	ReplicaCapacityModeOnDemand    ReplicaCapacityModeEnum = "ON_DEMAND"
)

var mappingReplicaCapacityModeEnum = map[string]ReplicaCapacityModeEnum{
	"PROVISIONED": ReplicaCapacityModeProvisioned,
	"ON_DEMAND":   ReplicaCapacityModeOnDemand,
}

var mappingReplicaCapacityModeEnumLowerCase = map[string]ReplicaCapacityModeEnum{
	"provisioned": ReplicaCapacityModeProvisioned,
	"on_demand":   ReplicaCapacityModeOnDemand,
}

// GetReplicaCapacityModeEnumValues Enumerates the set of values for ReplicaCapacityModeEnum
func GetReplicaCapacityModeEnumValues() []ReplicaCapacityModeEnum {
	values := make([]ReplicaCapacityModeEnum, 0)
	for _, v := range mappingReplicaCapacityModeEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicaCapacityModeEnumStringValues Enumerates the set of values in String for ReplicaCapacityModeEnum
func GetReplicaCapacityModeEnumStringValues() []string {
	return []string{
		"PROVISIONED",
		"ON_DEMAND",
	}
}

// GetMappingReplicaCapacityModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicaCapacityModeEnum(val string) (ReplicaCapacityModeEnum, bool) {
	enum, ok := mappingReplicaCapacityModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReplicaLifecycleStateEnum Enum with underlying type: string
type ReplicaLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicaLifecycleStateEnum
const (
	ReplicaLifecycleStateCreating ReplicaLifecycleStateEnum = "CREATING"
	ReplicaLifecycleStateUpdating ReplicaLifecycleStateEnum = "UPDATING"
	ReplicaLifecycleStateActive   ReplicaLifecycleStateEnum = "ACTIVE"
	ReplicaLifecycleStateDeleting ReplicaLifecycleStateEnum = "DELETING"
)

var mappingReplicaLifecycleStateEnum = map[string]ReplicaLifecycleStateEnum{
	"CREATING": ReplicaLifecycleStateCreating,
	"UPDATING": ReplicaLifecycleStateUpdating,
	"ACTIVE":   ReplicaLifecycleStateActive,
	"DELETING": ReplicaLifecycleStateDeleting,
}

var mappingReplicaLifecycleStateEnumLowerCase = map[string]ReplicaLifecycleStateEnum{
	"creating": ReplicaLifecycleStateCreating,
	"updating": ReplicaLifecycleStateUpdating,
	"active":   ReplicaLifecycleStateActive,
	"deleting": ReplicaLifecycleStateDeleting,
}

// GetReplicaLifecycleStateEnumValues Enumerates the set of values for ReplicaLifecycleStateEnum
func GetReplicaLifecycleStateEnumValues() []ReplicaLifecycleStateEnum {
	values := make([]ReplicaLifecycleStateEnum, 0)
	for _, v := range mappingReplicaLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicaLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicaLifecycleStateEnum
func GetReplicaLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
	}
}

// GetMappingReplicaLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicaLifecycleStateEnum(val string) (ReplicaLifecycleStateEnum, bool) {
	enum, ok := mappingReplicaLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
