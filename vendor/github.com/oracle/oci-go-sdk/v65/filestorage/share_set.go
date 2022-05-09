// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ShareSet The FSS SMB Share Set contains the SMB configuration data of a Mount Target.
// A set of file systems to provide SMB share through one mount target.
// Composed of zero or more share resources.
type ShareSet struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the share set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual cloud network (VCN) the share set is in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the mount target the share set bind to.
	MountTargetId *string `mandatory:"true" json:"mountTargetId"`

	// A comment of the share set. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My share set`
	Comment *string `mandatory:"true" json:"comment"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the share set.
	Id *string `mandatory:"true" json:"id"`

	// A customer-provided DNS name. This name plays a critical role in
	// establishing the server's place in the customer SMB security hierarchy.
	// For example, if an SMB server has a DNS name of
	// register5.store34.california.usa.marks-hats.com, then this particular
	// server is part of the store34.california.usa.marks-hats.com security
	// domain which in turn is part of the california.usa.marks-hats.com which
	// in turn is part of the usa.marks-hats.com
	// which in turn is part of the marks-hats.com security domain.
	// Must be unique across all FQDNs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	CustomFqdn *string `mandatory:"true" json:"customFqdn"`

	// Every SMB server (i.e. each mount target) needs a Netbios name in
	// addition to its fqdn (fully qualified domain name). Normally,
	// the Netbios name is simply the hostname portion of the fqdn.
	// This doesn't work when multiple computers have the same hostname.
	// For example, a computer called orange.colors.com and a computer
	// called orange.fruit.org can interfere with each other if they both
	// use orange as their Netbios name. To avoid problems, at least one
	// computer can be configured to have a Netbios name that is
	// not its hostname.
	NetBiosName *string `mandatory:"true" json:"netBiosName"`

	// A read-only property for the connection status between the
	// mount target and the customer-provided domain controller which
	// is the domain based on the customFQDN.
	DomainConnectionStatus *string `mandatory:"true" json:"domainConnectionStatus"`

	// The current state of the share set.
	LifecycleState ShareSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the share set was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the share set is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Turn on this flag to allow unsigned SMB traffic.
	AllowUnsignedTraffic *string `mandatory:"false" json:"allowUnsignedTraffic"`

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

func (m ShareSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShareSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShareSetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShareSetLifecycleStateEnum Enum with underlying type: string
type ShareSetLifecycleStateEnum string

// Set of constants representing the allowable values for ShareSetLifecycleStateEnum
const (
	ShareSetLifecycleStateCreating ShareSetLifecycleStateEnum = "CREATING"
	ShareSetLifecycleStateActive   ShareSetLifecycleStateEnum = "ACTIVE"
	ShareSetLifecycleStateDeleting ShareSetLifecycleStateEnum = "DELETING"
	ShareSetLifecycleStateDeleted  ShareSetLifecycleStateEnum = "DELETED"
)

var mappingShareSetLifecycleStateEnum = map[string]ShareSetLifecycleStateEnum{
	"CREATING": ShareSetLifecycleStateCreating,
	"ACTIVE":   ShareSetLifecycleStateActive,
	"DELETING": ShareSetLifecycleStateDeleting,
	"DELETED":  ShareSetLifecycleStateDeleted,
}

var mappingShareSetLifecycleStateEnumLowerCase = map[string]ShareSetLifecycleStateEnum{
	"creating": ShareSetLifecycleStateCreating,
	"active":   ShareSetLifecycleStateActive,
	"deleting": ShareSetLifecycleStateDeleting,
	"deleted":  ShareSetLifecycleStateDeleted,
}

// GetShareSetLifecycleStateEnumValues Enumerates the set of values for ShareSetLifecycleStateEnum
func GetShareSetLifecycleStateEnumValues() []ShareSetLifecycleStateEnum {
	values := make([]ShareSetLifecycleStateEnum, 0)
	for _, v := range mappingShareSetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetShareSetLifecycleStateEnumStringValues Enumerates the set of values in String for ShareSetLifecycleStateEnum
func GetShareSetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingShareSetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShareSetLifecycleStateEnum(val string) (ShareSetLifecycleStateEnum, bool) {
	enum, ok := mappingShareSetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
