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

// UpdateShareSetDetails Details for updating the share set.
type UpdateShareSetDetails struct {

	// A comment of the share set. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My share set`
	Comment *string `mandatory:"false" json:"comment"`

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

func (m UpdateShareSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateShareSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingShareSetSmbEncryptionEnum(string(m.SmbEncryption)); !ok && m.SmbEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SmbEncryption: %s. Supported values are: %s.", m.SmbEncryption, strings.Join(GetShareSetSmbEncryptionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
