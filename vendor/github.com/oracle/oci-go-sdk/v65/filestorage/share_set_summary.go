// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ShareSetSummary Summary information for a share set.
type ShareSetSummary struct {

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
	// domain, which in turn is part of california.usa.marks-hats.com, which
	// in turn is part of usa.marks-hats.com,
	// which in turn is part of the marks-hats.com security domain.
	// Must be unique across all FQDNs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	CustomFqdn *string `mandatory:"true" json:"customFqdn"`

	// Every SMB server (i.e. each mount target) needs a NetBIOS name in
	// addition to its FQDN (fully qualified domain name). Normally,
	// the NetBIOS name is simply the hostname portion of the FQDN.
	// This doesn't work when multiple computers have the same hostname.
	// For example, a computer called orange.colors.com and a computer
	// called orange.fruit.org can interfere with each other if they both
	// use orange as their NetBIOS name. To avoid problems, configure at least one
	// computer to have a NetBIOS name that is not its hostname.
	NetBiosName *string `mandatory:"true" json:"netBiosName"`

	// A read-only property for the connection status between the
	// mount target and the customer-provided domain controller which
	// is the domain based on the custom FQDN.
	DomainConnectionStatus ShareSetDomainConnectionStatusEnum `mandatory:"true" json:"domainConnectionStatus"`

	// The current state of the share set.
	LifecycleState ShareSetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time that the share set was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A read-only property that emits the status of the mount target share set's join domain operation
	// to a domain controller for SMB access.
	JoinDomainResult *string `mandatory:"true" json:"joinDomainResult"`

	// The availability domain that the share set is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// A read-only field that's only populated when the join domain operation FAILED.
	JoinDomainErrorMessage *string `mandatory:"false" json:"joinDomainErrorMessage"`
}

func (m ShareSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShareSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareSetDomainConnectionStatusEnum(string(m.DomainConnectionStatus)); !ok && m.DomainConnectionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DomainConnectionStatus: %s. Supported values are: %s.", m.DomainConnectionStatus, strings.Join(GetShareSetDomainConnectionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingShareSetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetShareSetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
