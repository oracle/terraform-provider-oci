// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MountTarget Provides access to a collection of file systems through one or more VNICs on a
// specified subnet. The set of file systems is controlled through the
// referenced export set.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type MountTarget struct {

	// The OCID of the compartment that contains the mount target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My mount target`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the mount target.
	Id *string `mandatory:"true" json:"id"`

	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The current state of the mount target.
	LifecycleState MountTargetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCIDs of the private IP addresses associated with this mount target.
	PrivateIpIds []string `mandatory:"true" json:"privateIpIds"`

	// The OCID of the subnet the mount target is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the mount target was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the mount target is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the associated export set. Controls what file
	// systems will be exported through Network File System (NFS) protocol on this
	// mount target.
	ExportSetId *string `mandatory:"false" json:"exportSetId"`

	// A list of Network Security Group OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with this mount target.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the mount target from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m MountTarget) String() string {
	return common.PointerString(m)
}

// MountTargetLifecycleStateEnum Enum with underlying type: string
type MountTargetLifecycleStateEnum string

// Set of constants representing the allowable values for MountTargetLifecycleStateEnum
const (
	MountTargetLifecycleStateCreating MountTargetLifecycleStateEnum = "CREATING"
	MountTargetLifecycleStateActive   MountTargetLifecycleStateEnum = "ACTIVE"
	MountTargetLifecycleStateDeleting MountTargetLifecycleStateEnum = "DELETING"
	MountTargetLifecycleStateDeleted  MountTargetLifecycleStateEnum = "DELETED"
	MountTargetLifecycleStateFailed   MountTargetLifecycleStateEnum = "FAILED"
)

var mappingMountTargetLifecycleState = map[string]MountTargetLifecycleStateEnum{
	"CREATING": MountTargetLifecycleStateCreating,
	"ACTIVE":   MountTargetLifecycleStateActive,
	"DELETING": MountTargetLifecycleStateDeleting,
	"DELETED":  MountTargetLifecycleStateDeleted,
	"FAILED":   MountTargetLifecycleStateFailed,
}

// GetMountTargetLifecycleStateEnumValues Enumerates the set of values for MountTargetLifecycleStateEnum
func GetMountTargetLifecycleStateEnumValues() []MountTargetLifecycleStateEnum {
	values := make([]MountTargetLifecycleStateEnum, 0)
	for _, v := range mappingMountTargetLifecycleState {
		values = append(values, v)
	}
	return values
}
