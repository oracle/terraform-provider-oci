// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccCapacityRequest A single request of some quantity of a specific server type, in a specific location and expected delivery date. The maximum amount possible to request is the smallest number between the number of servers available for purchase and the number of servers allowed by the constraints (For example, power, network, physical space, and so on).
type OccCapacityRequest struct {

	// The OCID of the capacity request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the request was made.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the availability catalog against which the capacity request was placed.
	OccAvailabilityCatalogId *string `mandatory:"true" json:"occAvailabilityCatalogId"`

	// The display name of the capacity request.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The OCID of the customer group to which this customer belongs to.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The name of the region for which the capacity request was made.
	Region *string `mandatory:"true" json:"region"`

	// The availability domain (AD) for which the capacity request was made.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	DateExpectedCapacityHandover *common.SDKTime `mandatory:"true" json:"dateExpectedCapacityHandover"`

	// The different states the capacity request goes through.
	RequestState OccCapacityRequestRequestStateEnum `mandatory:"true" json:"requestState"`

	// The time when the capacity request was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the capacity request was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the resource.
	LifecycleState OccCapacityRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A list of resources requested as part of this request
	Details []OccCapacityRequestBaseDetails `mandatory:"true" json:"details"`

	// Meaningful text about the capacity request.
	Description *string `mandatory:"false" json:"description"`

	// Type of Capacity Request(New or Transfer)
	RequestType OccCapacityRequestRequestTypeEnum `mandatory:"false" json:"requestType,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccCapacityRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCapacityRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCapacityRequestRequestStateEnum(string(m.RequestState)); !ok && m.RequestState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestState: %s. Supported values are: %s.", m.RequestState, strings.Join(GetOccCapacityRequestRequestStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccCapacityRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccCapacityRequestLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOccCapacityRequestRequestTypeEnum(string(m.RequestType)); !ok && m.RequestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestType: %s. Supported values are: %s.", m.RequestType, strings.Join(GetOccCapacityRequestRequestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccCapacityRequestRequestTypeEnum Enum with underlying type: string
type OccCapacityRequestRequestTypeEnum string

// Set of constants representing the allowable values for OccCapacityRequestRequestTypeEnum
const (
	OccCapacityRequestRequestTypeNew      OccCapacityRequestRequestTypeEnum = "NEW"
	OccCapacityRequestRequestTypeTransfer OccCapacityRequestRequestTypeEnum = "TRANSFER"
)

var mappingOccCapacityRequestRequestTypeEnum = map[string]OccCapacityRequestRequestTypeEnum{
	"NEW":      OccCapacityRequestRequestTypeNew,
	"TRANSFER": OccCapacityRequestRequestTypeTransfer,
}

var mappingOccCapacityRequestRequestTypeEnumLowerCase = map[string]OccCapacityRequestRequestTypeEnum{
	"new":      OccCapacityRequestRequestTypeNew,
	"transfer": OccCapacityRequestRequestTypeTransfer,
}

// GetOccCapacityRequestRequestTypeEnumValues Enumerates the set of values for OccCapacityRequestRequestTypeEnum
func GetOccCapacityRequestRequestTypeEnumValues() []OccCapacityRequestRequestTypeEnum {
	values := make([]OccCapacityRequestRequestTypeEnum, 0)
	for _, v := range mappingOccCapacityRequestRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCapacityRequestRequestTypeEnumStringValues Enumerates the set of values in String for OccCapacityRequestRequestTypeEnum
func GetOccCapacityRequestRequestTypeEnumStringValues() []string {
	return []string{
		"NEW",
		"TRANSFER",
	}
}

// GetMappingOccCapacityRequestRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCapacityRequestRequestTypeEnum(val string) (OccCapacityRequestRequestTypeEnum, bool) {
	enum, ok := mappingOccCapacityRequestRequestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccCapacityRequestRequestStateEnum Enum with underlying type: string
type OccCapacityRequestRequestStateEnum string

// Set of constants representing the allowable values for OccCapacityRequestRequestStateEnum
const (
	OccCapacityRequestRequestStateCreated            OccCapacityRequestRequestStateEnum = "CREATED"
	OccCapacityRequestRequestStateSubmitted          OccCapacityRequestRequestStateEnum = "SUBMITTED"
	OccCapacityRequestRequestStateRejected           OccCapacityRequestRequestStateEnum = "REJECTED"
	OccCapacityRequestRequestStateInProgress         OccCapacityRequestRequestStateEnum = "IN_PROGRESS"
	OccCapacityRequestRequestStateCompleted          OccCapacityRequestRequestStateEnum = "COMPLETED"
	OccCapacityRequestRequestStatePartiallyCompleted OccCapacityRequestRequestStateEnum = "PARTIALLY_COMPLETED"
	OccCapacityRequestRequestStateCancelled          OccCapacityRequestRequestStateEnum = "CANCELLED"
	OccCapacityRequestRequestStateDeleted            OccCapacityRequestRequestStateEnum = "DELETED"
)

var mappingOccCapacityRequestRequestStateEnum = map[string]OccCapacityRequestRequestStateEnum{
	"CREATED":             OccCapacityRequestRequestStateCreated,
	"SUBMITTED":           OccCapacityRequestRequestStateSubmitted,
	"REJECTED":            OccCapacityRequestRequestStateRejected,
	"IN_PROGRESS":         OccCapacityRequestRequestStateInProgress,
	"COMPLETED":           OccCapacityRequestRequestStateCompleted,
	"PARTIALLY_COMPLETED": OccCapacityRequestRequestStatePartiallyCompleted,
	"CANCELLED":           OccCapacityRequestRequestStateCancelled,
	"DELETED":             OccCapacityRequestRequestStateDeleted,
}

var mappingOccCapacityRequestRequestStateEnumLowerCase = map[string]OccCapacityRequestRequestStateEnum{
	"created":             OccCapacityRequestRequestStateCreated,
	"submitted":           OccCapacityRequestRequestStateSubmitted,
	"rejected":            OccCapacityRequestRequestStateRejected,
	"in_progress":         OccCapacityRequestRequestStateInProgress,
	"completed":           OccCapacityRequestRequestStateCompleted,
	"partially_completed": OccCapacityRequestRequestStatePartiallyCompleted,
	"cancelled":           OccCapacityRequestRequestStateCancelled,
	"deleted":             OccCapacityRequestRequestStateDeleted,
}

// GetOccCapacityRequestRequestStateEnumValues Enumerates the set of values for OccCapacityRequestRequestStateEnum
func GetOccCapacityRequestRequestStateEnumValues() []OccCapacityRequestRequestStateEnum {
	values := make([]OccCapacityRequestRequestStateEnum, 0)
	for _, v := range mappingOccCapacityRequestRequestStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCapacityRequestRequestStateEnumStringValues Enumerates the set of values in String for OccCapacityRequestRequestStateEnum
func GetOccCapacityRequestRequestStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"SUBMITTED",
		"REJECTED",
		"IN_PROGRESS",
		"COMPLETED",
		"PARTIALLY_COMPLETED",
		"CANCELLED",
		"DELETED",
	}
}

// GetMappingOccCapacityRequestRequestStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCapacityRequestRequestStateEnum(val string) (OccCapacityRequestRequestStateEnum, bool) {
	enum, ok := mappingOccCapacityRequestRequestStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccCapacityRequestLifecycleStateEnum Enum with underlying type: string
type OccCapacityRequestLifecycleStateEnum string

// Set of constants representing the allowable values for OccCapacityRequestLifecycleStateEnum
const (
	OccCapacityRequestLifecycleStateCreating OccCapacityRequestLifecycleStateEnum = "CREATING"
	OccCapacityRequestLifecycleStateUpdating OccCapacityRequestLifecycleStateEnum = "UPDATING"
	OccCapacityRequestLifecycleStateActive   OccCapacityRequestLifecycleStateEnum = "ACTIVE"
	OccCapacityRequestLifecycleStateDeleting OccCapacityRequestLifecycleStateEnum = "DELETING"
	OccCapacityRequestLifecycleStateDeleted  OccCapacityRequestLifecycleStateEnum = "DELETED"
	OccCapacityRequestLifecycleStateFailed   OccCapacityRequestLifecycleStateEnum = "FAILED"
)

var mappingOccCapacityRequestLifecycleStateEnum = map[string]OccCapacityRequestLifecycleStateEnum{
	"CREATING": OccCapacityRequestLifecycleStateCreating,
	"UPDATING": OccCapacityRequestLifecycleStateUpdating,
	"ACTIVE":   OccCapacityRequestLifecycleStateActive,
	"DELETING": OccCapacityRequestLifecycleStateDeleting,
	"DELETED":  OccCapacityRequestLifecycleStateDeleted,
	"FAILED":   OccCapacityRequestLifecycleStateFailed,
}

var mappingOccCapacityRequestLifecycleStateEnumLowerCase = map[string]OccCapacityRequestLifecycleStateEnum{
	"creating": OccCapacityRequestLifecycleStateCreating,
	"updating": OccCapacityRequestLifecycleStateUpdating,
	"active":   OccCapacityRequestLifecycleStateActive,
	"deleting": OccCapacityRequestLifecycleStateDeleting,
	"deleted":  OccCapacityRequestLifecycleStateDeleted,
	"failed":   OccCapacityRequestLifecycleStateFailed,
}

// GetOccCapacityRequestLifecycleStateEnumValues Enumerates the set of values for OccCapacityRequestLifecycleStateEnum
func GetOccCapacityRequestLifecycleStateEnumValues() []OccCapacityRequestLifecycleStateEnum {
	values := make([]OccCapacityRequestLifecycleStateEnum, 0)
	for _, v := range mappingOccCapacityRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCapacityRequestLifecycleStateEnumStringValues Enumerates the set of values in String for OccCapacityRequestLifecycleStateEnum
func GetOccCapacityRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccCapacityRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCapacityRequestLifecycleStateEnum(val string) (OccCapacityRequestLifecycleStateEnum, bool) {
	enum, ok := mappingOccCapacityRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
