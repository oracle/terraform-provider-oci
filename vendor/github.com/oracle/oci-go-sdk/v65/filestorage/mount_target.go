// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MountTarget Provides access to a collection of file systems through one or more VNICs on a
// specified subnet. The set of file systems is controlled through the
// referenced export set.
type MountTarget struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the mount target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My mount target`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the mount target.
	Id *string `mandatory:"true" json:"id"`

	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The current state of the mount target.
	LifecycleState MountTargetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCIDs of the private IP addresses associated with this mount target.
	PrivateIpIds []string `mandatory:"true" json:"privateIpIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the mount target is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the mount target was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the mount target is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated export set. Controls what file
	// systems will be exported through Network File System (NFS) protocol on this
	// mount target.
	ExportSetId *string `mandatory:"false" json:"exportSetId"`

	// The method used to map a Unix UID to secondary groups. If NONE, the mount target will not use the Unix UID for ID mapping.
	IdmapType MountTargetIdmapTypeEnum `mandatory:"false" json:"idmapType,omitempty"`

	LdapIdmap *LdapIdmap `mandatory:"false" json:"ldapIdmap"`

	// A list of Network Security Group OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with this mount target.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the mount target from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	Kerberos *Kerberos `mandatory:"false" json:"kerberos"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MountTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMountTargetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMountTargetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMountTargetIdmapTypeEnum(string(m.IdmapType)); !ok && m.IdmapType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdmapType: %s. Supported values are: %s.", m.IdmapType, strings.Join(GetMountTargetIdmapTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingMountTargetLifecycleStateEnum = map[string]MountTargetLifecycleStateEnum{
	"CREATING": MountTargetLifecycleStateCreating,
	"ACTIVE":   MountTargetLifecycleStateActive,
	"DELETING": MountTargetLifecycleStateDeleting,
	"DELETED":  MountTargetLifecycleStateDeleted,
	"FAILED":   MountTargetLifecycleStateFailed,
}

var mappingMountTargetLifecycleStateEnumLowerCase = map[string]MountTargetLifecycleStateEnum{
	"creating": MountTargetLifecycleStateCreating,
	"active":   MountTargetLifecycleStateActive,
	"deleting": MountTargetLifecycleStateDeleting,
	"deleted":  MountTargetLifecycleStateDeleted,
	"failed":   MountTargetLifecycleStateFailed,
}

// GetMountTargetLifecycleStateEnumValues Enumerates the set of values for MountTargetLifecycleStateEnum
func GetMountTargetLifecycleStateEnumValues() []MountTargetLifecycleStateEnum {
	values := make([]MountTargetLifecycleStateEnum, 0)
	for _, v := range mappingMountTargetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMountTargetLifecycleStateEnumStringValues Enumerates the set of values in String for MountTargetLifecycleStateEnum
func GetMountTargetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMountTargetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMountTargetLifecycleStateEnum(val string) (MountTargetLifecycleStateEnum, bool) {
	enum, ok := mappingMountTargetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MountTargetIdmapTypeEnum Enum with underlying type: string
type MountTargetIdmapTypeEnum string

// Set of constants representing the allowable values for MountTargetIdmapTypeEnum
const (
	MountTargetIdmapTypeLdap MountTargetIdmapTypeEnum = "LDAP"
	MountTargetIdmapTypeNone MountTargetIdmapTypeEnum = "NONE"
)

var mappingMountTargetIdmapTypeEnum = map[string]MountTargetIdmapTypeEnum{
	"LDAP": MountTargetIdmapTypeLdap,
	"NONE": MountTargetIdmapTypeNone,
}

var mappingMountTargetIdmapTypeEnumLowerCase = map[string]MountTargetIdmapTypeEnum{
	"ldap": MountTargetIdmapTypeLdap,
	"none": MountTargetIdmapTypeNone,
}

// GetMountTargetIdmapTypeEnumValues Enumerates the set of values for MountTargetIdmapTypeEnum
func GetMountTargetIdmapTypeEnumValues() []MountTargetIdmapTypeEnum {
	values := make([]MountTargetIdmapTypeEnum, 0)
	for _, v := range mappingMountTargetIdmapTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMountTargetIdmapTypeEnumStringValues Enumerates the set of values in String for MountTargetIdmapTypeEnum
func GetMountTargetIdmapTypeEnumStringValues() []string {
	return []string{
		"LDAP",
		"NONE",
	}
}

// GetMappingMountTargetIdmapTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMountTargetIdmapTypeEnum(val string) (MountTargetIdmapTypeEnum, bool) {
	enum, ok := mappingMountTargetIdmapTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
