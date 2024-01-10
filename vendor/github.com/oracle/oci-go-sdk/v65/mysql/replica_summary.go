// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicaSummary Summary of the read replica.
type ReplicaSummary struct {

	// The OCID of the read replica.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DB System the read replica is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID of the compartment that contains the read replica.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the read replica. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the read replica.
	LifecycleState ReplicaSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the read replica was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The MySQL version currently in use by the read replica.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The IP address the read replica is configured to listen on.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The port the read replica is configured to listen on.
	Port *int `mandatory:"true" json:"port"`

	// The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port.
	PortX *int `mandatory:"true" json:"portX"`

	// User provided description of the read replica.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the state of the read replica.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time the read replica was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The name of the Availability Domain the read replica is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The name of the Fault Domain the read replica is located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Specifies whether the read replica can be deleted. Set to true to prevent deletion, false (default) to allow.
	// Note that if a read replica is delete protected it also prevents the entire DB System from being deleted. If
	// the DB System is delete protected, read replicas can still be deleted individually if they are not delete
	// protected themselves.
	IsDeleteProtected *bool `mandatory:"false" json:"isDeleteProtected"`

	// The shape currently in use by the read replica. The shape determines the resources allocated:
	// CPU cores and memory for VM shapes, CPU cores, memory and storage for non-VM (bare metal) shapes.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The OCID of the Configuration currently in use by the read replica.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	ReplicaOverrides *ReplicaOverrides `mandatory:"false" json:"replicaOverrides"`
}

func (m ReplicaSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicaSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicaSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicaSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicaSummaryLifecycleStateEnum Enum with underlying type: string
type ReplicaSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicaSummaryLifecycleStateEnum
const (
	ReplicaSummaryLifecycleStateCreating       ReplicaSummaryLifecycleStateEnum = "CREATING"
	ReplicaSummaryLifecycleStateActive         ReplicaSummaryLifecycleStateEnum = "ACTIVE"
	ReplicaSummaryLifecycleStateInactive       ReplicaSummaryLifecycleStateEnum = "INACTIVE"
	ReplicaSummaryLifecycleStateUpdating       ReplicaSummaryLifecycleStateEnum = "UPDATING"
	ReplicaSummaryLifecycleStateDeleting       ReplicaSummaryLifecycleStateEnum = "DELETING"
	ReplicaSummaryLifecycleStateDeleted        ReplicaSummaryLifecycleStateEnum = "DELETED"
	ReplicaSummaryLifecycleStateNeedsAttention ReplicaSummaryLifecycleStateEnum = "NEEDS_ATTENTION"
	ReplicaSummaryLifecycleStateFailed         ReplicaSummaryLifecycleStateEnum = "FAILED"
)

var mappingReplicaSummaryLifecycleStateEnum = map[string]ReplicaSummaryLifecycleStateEnum{
	"CREATING":        ReplicaSummaryLifecycleStateCreating,
	"ACTIVE":          ReplicaSummaryLifecycleStateActive,
	"INACTIVE":        ReplicaSummaryLifecycleStateInactive,
	"UPDATING":        ReplicaSummaryLifecycleStateUpdating,
	"DELETING":        ReplicaSummaryLifecycleStateDeleting,
	"DELETED":         ReplicaSummaryLifecycleStateDeleted,
	"NEEDS_ATTENTION": ReplicaSummaryLifecycleStateNeedsAttention,
	"FAILED":          ReplicaSummaryLifecycleStateFailed,
}

var mappingReplicaSummaryLifecycleStateEnumLowerCase = map[string]ReplicaSummaryLifecycleStateEnum{
	"creating":        ReplicaSummaryLifecycleStateCreating,
	"active":          ReplicaSummaryLifecycleStateActive,
	"inactive":        ReplicaSummaryLifecycleStateInactive,
	"updating":        ReplicaSummaryLifecycleStateUpdating,
	"deleting":        ReplicaSummaryLifecycleStateDeleting,
	"deleted":         ReplicaSummaryLifecycleStateDeleted,
	"needs_attention": ReplicaSummaryLifecycleStateNeedsAttention,
	"failed":          ReplicaSummaryLifecycleStateFailed,
}

// GetReplicaSummaryLifecycleStateEnumValues Enumerates the set of values for ReplicaSummaryLifecycleStateEnum
func GetReplicaSummaryLifecycleStateEnumValues() []ReplicaSummaryLifecycleStateEnum {
	values := make([]ReplicaSummaryLifecycleStateEnum, 0)
	for _, v := range mappingReplicaSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicaSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicaSummaryLifecycleStateEnum
func GetReplicaSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingReplicaSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicaSummaryLifecycleStateEnum(val string) (ReplicaSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingReplicaSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
