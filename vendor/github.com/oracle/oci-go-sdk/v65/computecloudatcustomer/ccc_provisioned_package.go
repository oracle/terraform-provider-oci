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

// CccProvisionedPackage Compute Cloud@Customer provisioned package information. This resource tracks
// Compute Cloud@Customer marketplace packages that
// are provisioned to a Compute Cloud@Customer infrastructure.
type CccProvisionedPackage struct {

	// Compute Cloud@Customer marketplace provisioned package OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This cannot be changed once created.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer marketplace provisioned package display name.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the
	// Compute Cloud@Customer infrastructure where the marketplace provisioned package will reside.
	CccInfrastructureId *string `mandatory:"true" json:"cccInfrastructureId"`

	// Compute Cloud@Customer marketplace listing identifier
	// that this provisioned package is associated with.
	// This cannot be updated once provisioned.
	CccListingId *string `mandatory:"true" json:"cccListingId"`

	// Compute Cloud@Customer package
	// identifier that relates to a version of the package that will be provisioned.
	CccPackageId *string `mandatory:"true" json:"cccPackageId"`

	// The time the marketplace provisioned package was created, using an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Lifecycle state of the resource.
	LifecycleState CccProvisionedPackageLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The total number of OCPUs that can be used for the compute instances that use this
	// image on the Compute Cloud@Customer infrastructure. This limit can be changed after
	// provisioning. If the value is higher then the total number of OCPUs available,
	// the value will accepted but the maximum will be the total number of OCPUs on the
	// Compute Cloud@Customer infrastructure.
	TotalOcpuLimit *int `mandatory:"true" json:"totalOcpuLimit"`

	// An optional description of the Compute Cloud@Customer marketplace provisioned package.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time the marketplace provisioned package was updated, using an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, the message can be used to provide actionable information for a resource in
	// a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Compute Cloud@Customer infrastructure boot image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This image OCID will be used on the rack to launch a compute instance.
	CccImageId *string `mandatory:"false" json:"cccImageId"`

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

func (m CccProvisionedPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccProvisionedPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccProvisionedPackageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCccProvisionedPackageLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccProvisionedPackageLifecycleStateEnum Enum with underlying type: string
type CccProvisionedPackageLifecycleStateEnum string

// Set of constants representing the allowable values for CccProvisionedPackageLifecycleStateEnum
const (
	CccProvisionedPackageLifecycleStateCreating CccProvisionedPackageLifecycleStateEnum = "CREATING"
	CccProvisionedPackageLifecycleStateActive   CccProvisionedPackageLifecycleStateEnum = "ACTIVE"
	CccProvisionedPackageLifecycleStateInactive CccProvisionedPackageLifecycleStateEnum = "INACTIVE"
	CccProvisionedPackageLifecycleStateUpdating CccProvisionedPackageLifecycleStateEnum = "UPDATING"
	CccProvisionedPackageLifecycleStateDeleting CccProvisionedPackageLifecycleStateEnum = "DELETING"
	CccProvisionedPackageLifecycleStateDeleted  CccProvisionedPackageLifecycleStateEnum = "DELETED"
	CccProvisionedPackageLifecycleStateFailed   CccProvisionedPackageLifecycleStateEnum = "FAILED"
)

var mappingCccProvisionedPackageLifecycleStateEnum = map[string]CccProvisionedPackageLifecycleStateEnum{
	"CREATING": CccProvisionedPackageLifecycleStateCreating,
	"ACTIVE":   CccProvisionedPackageLifecycleStateActive,
	"INACTIVE": CccProvisionedPackageLifecycleStateInactive,
	"UPDATING": CccProvisionedPackageLifecycleStateUpdating,
	"DELETING": CccProvisionedPackageLifecycleStateDeleting,
	"DELETED":  CccProvisionedPackageLifecycleStateDeleted,
	"FAILED":   CccProvisionedPackageLifecycleStateFailed,
}

var mappingCccProvisionedPackageLifecycleStateEnumLowerCase = map[string]CccProvisionedPackageLifecycleStateEnum{
	"creating": CccProvisionedPackageLifecycleStateCreating,
	"active":   CccProvisionedPackageLifecycleStateActive,
	"inactive": CccProvisionedPackageLifecycleStateInactive,
	"updating": CccProvisionedPackageLifecycleStateUpdating,
	"deleting": CccProvisionedPackageLifecycleStateDeleting,
	"deleted":  CccProvisionedPackageLifecycleStateDeleted,
	"failed":   CccProvisionedPackageLifecycleStateFailed,
}

// GetCccProvisionedPackageLifecycleStateEnumValues Enumerates the set of values for CccProvisionedPackageLifecycleStateEnum
func GetCccProvisionedPackageLifecycleStateEnumValues() []CccProvisionedPackageLifecycleStateEnum {
	values := make([]CccProvisionedPackageLifecycleStateEnum, 0)
	for _, v := range mappingCccProvisionedPackageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCccProvisionedPackageLifecycleStateEnumStringValues Enumerates the set of values in String for CccProvisionedPackageLifecycleStateEnum
func GetCccProvisionedPackageLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCccProvisionedPackageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccProvisionedPackageLifecycleStateEnum(val string) (CccProvisionedPackageLifecycleStateEnum, bool) {
	enum, ok := mappingCccProvisionedPackageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
