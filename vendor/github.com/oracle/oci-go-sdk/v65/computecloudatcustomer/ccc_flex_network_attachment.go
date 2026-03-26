// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccFlexNetworkAttachment The Oracle Cloud Infrastructure resource representing an association between a storage device network
// and a subnet, both of which are in the customer’s data center.
type CccFlexNetworkAttachment struct {

	// The Compute Cloud@Customer flexNetworkAttachment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This cannot be changed once created.
	Id *string `mandatory:"true" json:"id"`

	// The name that will be used to display the Compute Cloud@Customer flexNetworkAttachment
	// in the Oracle Cloud console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The flexNetworkAttachment compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Compute Cloud@Customer Infrastructure OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm),
	// which is immutable on creation.
	InfrastructureId *string `mandatory:"true" json:"infrastructureId"`

	// The network OCID assigned by the infrastructure administrative service.
	FlexNetworkId *string `mandatory:"true" json:"flexNetworkId"`

	// The subnet OCID assigned by the infrastructure administrative service.
	InfraSubnetId *string `mandatory:"true" json:"infraSubnetId"`

	// Compute Cloud@Customer flexNetworkAttachment creation date and time, using an RFC3339 formatted
	// datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Compute Cloud@Customer flexNetworkAttachment.
	LifecycleState CccFlexNetworkAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer flexNetworkAttachment.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Compute Cloud@Customer flexNetworkAttachment updated date and time, using an RFC3339 formatted
	// datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current lifecycle state in more detail.
	// For example, this can be used to provide actionable information for a resource that is in
	// a Failed state.
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

func (m CccFlexNetworkAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccFlexNetworkAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccFlexNetworkAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCccFlexNetworkAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccFlexNetworkAttachmentLifecycleStateEnum Enum with underlying type: string
type CccFlexNetworkAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for CccFlexNetworkAttachmentLifecycleStateEnum
const (
	CccFlexNetworkAttachmentLifecycleStateActive   CccFlexNetworkAttachmentLifecycleStateEnum = "ACTIVE"
	CccFlexNetworkAttachmentLifecycleStateCreating CccFlexNetworkAttachmentLifecycleStateEnum = "CREATING"
	CccFlexNetworkAttachmentLifecycleStateUpdating CccFlexNetworkAttachmentLifecycleStateEnum = "UPDATING"
	CccFlexNetworkAttachmentLifecycleStateDeleting CccFlexNetworkAttachmentLifecycleStateEnum = "DELETING"
	CccFlexNetworkAttachmentLifecycleStateDeleted  CccFlexNetworkAttachmentLifecycleStateEnum = "DELETED"
	CccFlexNetworkAttachmentLifecycleStateFailed   CccFlexNetworkAttachmentLifecycleStateEnum = "FAILED"
)

var mappingCccFlexNetworkAttachmentLifecycleStateEnum = map[string]CccFlexNetworkAttachmentLifecycleStateEnum{
	"ACTIVE":   CccFlexNetworkAttachmentLifecycleStateActive,
	"CREATING": CccFlexNetworkAttachmentLifecycleStateCreating,
	"UPDATING": CccFlexNetworkAttachmentLifecycleStateUpdating,
	"DELETING": CccFlexNetworkAttachmentLifecycleStateDeleting,
	"DELETED":  CccFlexNetworkAttachmentLifecycleStateDeleted,
	"FAILED":   CccFlexNetworkAttachmentLifecycleStateFailed,
}

var mappingCccFlexNetworkAttachmentLifecycleStateEnumLowerCase = map[string]CccFlexNetworkAttachmentLifecycleStateEnum{
	"active":   CccFlexNetworkAttachmentLifecycleStateActive,
	"creating": CccFlexNetworkAttachmentLifecycleStateCreating,
	"updating": CccFlexNetworkAttachmentLifecycleStateUpdating,
	"deleting": CccFlexNetworkAttachmentLifecycleStateDeleting,
	"deleted":  CccFlexNetworkAttachmentLifecycleStateDeleted,
	"failed":   CccFlexNetworkAttachmentLifecycleStateFailed,
}

// GetCccFlexNetworkAttachmentLifecycleStateEnumValues Enumerates the set of values for CccFlexNetworkAttachmentLifecycleStateEnum
func GetCccFlexNetworkAttachmentLifecycleStateEnumValues() []CccFlexNetworkAttachmentLifecycleStateEnum {
	values := make([]CccFlexNetworkAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingCccFlexNetworkAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCccFlexNetworkAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for CccFlexNetworkAttachmentLifecycleStateEnum
func GetCccFlexNetworkAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCccFlexNetworkAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccFlexNetworkAttachmentLifecycleStateEnum(val string) (CccFlexNetworkAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingCccFlexNetworkAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
