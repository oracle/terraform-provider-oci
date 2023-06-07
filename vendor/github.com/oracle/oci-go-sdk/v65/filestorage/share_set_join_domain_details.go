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

// ShareSetJoinDomainDetails Details for the mount target share set to join a domain controller.
type ShareSetJoinDomainDetails struct {

	// The account name of a sufficiently powerful administrator
	// account. The default domain is the left-most domain of the
	// customer-provided DNS name.
	// This user name can be specified in several different ways:
	//   * As an ordinary name (e.g. joe).
	//   * As a name at some domain (e.g. sally@some.domain.com).
	//   * As a name under some domain (e.g. some.domain.com\administrator).
	Account *string `mandatory:"true" json:"account"`

	// The credential password of the account performing the join domain.
	Password *string `mandatory:"true" json:"password"`

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

	// The organizational unit (OU) is a container in an Active Directory (AD) that can
	// hold user accounts, service accounts, computer accounts, and other OUs.
	// This parameter specifies the OU that the mount target will join within the
	// AD domain. You can then assign administrators to specific OUs, and apply
	// group policy to enforce targeted configuration settings.
	OrganizationalUnit *string `mandatory:"false" json:"organizationalUnit"`

	// Every SMB server (i.e. each mount target) needs a NetBIOS name in
	// addition to its FQDN (fully qualified domain name). Normally,
	// the NetBIOS name is simply the hostname portion of the FQDN.
	// This doesn't work when multiple computers have the same hostname.
	// For example, a computer called orange.colors.com and a computer
	// called orange.fruit.org can interfere with each other if they both
	// use orange as their NetBIOS name. To avoid problems, configure at least one
	// computer to have a NetBIOS name that is not its hostname.
	NetBiosName *string `mandatory:"false" json:"netBiosName"`

	// Enable this flag to allow unsigned SMB traffic.
	IsUnsignedTrafficAllowed *bool `mandatory:"false" json:"isUnsignedTrafficAllowed"`

	// Describes the mount target's policy on SMB encryption.
	SmbEncryption ShareSetSmbEncryptionEnum `mandatory:"false" json:"smbEncryption,omitempty"`
}

func (m ShareSetJoinDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShareSetJoinDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingShareSetSmbEncryptionEnum(string(m.SmbEncryption)); !ok && m.SmbEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SmbEncryption: %s. Supported values are: %s.", m.SmbEncryption, strings.Join(GetShareSetSmbEncryptionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
