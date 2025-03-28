// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMountTargetDetails Details for creating the mount target.
type CreateMountTargetDetails struct {

	// The availability domain in which to create the mount target.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the mount target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which to create the mount target.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My mount target`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname for the mount target's IP address, used for
	// DNS resolution. The value is the hostname portion of the private IP
	// address's fully qualified domain name (FQDN). For example,
	// `files-1` in the FQDN `files-1.subnet123.vcn1.oraclevcn.com`.
	// Must be unique across all VNICs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// Note: This attribute value is stored in the PrivateIp (https://docs.oracle.com/iaas/en-us/iaas/api/#/en/iaas/20160918/PrivateIp/) resource,
	// not in the `mountTarget` resource.
	// To update the `hostnameLabel`, use `GetMountTarget` to obtain the
	// OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the mount target's
	// private IPs (`privateIpIds`). Then, you can use
	// UpdatePrivateIp (https://docs.oracle.com/iaas/en-us/iaas/api/#/en/iaas/20160918/PrivateIp/UpdatePrivateIp)
	// to update the `hostnameLabel` value.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `files-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// A private IP address of your choice. Must be an available IP address within
	// the subnet's CIDR. If you don't specify a value, Oracle automatically
	// assigns a private IP address from the subnet.
	// Example: `10.0.3.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The method used to map a Unix UID to secondary groups, if any.
	IdmapType MountTargetIdmapTypeEnum `mandatory:"false" json:"idmapType,omitempty"`

	LdapIdmap *CreateLdapIdmapDetails `mandatory:"false" json:"ldapIdmap"`

	// A list of Network Security Group OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this mount target.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the mount target from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.oracle.com/iaas/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	Kerberos *CreateKerberosDetails `mandatory:"false" json:"kerberos"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Throughput for mount target in Gbps. Currently only 1 Gbps of requestedThroughput is supported during create MountTarget.
	// Available shapes and corresponding throughput are listed at Mount Target Performance (https://docs.oracle.com/iaas/Content/File/Tasks/managingmounttargets.htm#performance).
	RequestedThroughput *int64 `mandatory:"false" json:"requestedThroughput"`
}

func (m CreateMountTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMountTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMountTargetIdmapTypeEnum(string(m.IdmapType)); !ok && m.IdmapType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdmapType: %s. Supported values are: %s.", m.IdmapType, strings.Join(GetMountTargetIdmapTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
