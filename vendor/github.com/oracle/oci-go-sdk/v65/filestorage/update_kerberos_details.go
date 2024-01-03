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

// UpdateKerberosDetails Kerberos details needed to update configuration.
type UpdateKerberosDetails struct {

	// Kerberos realm that this mount target will join.
	KerberosRealm *string `mandatory:"false" json:"kerberosRealm"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the keytab secret in the Vault.
	KeyTabSecretId *string `mandatory:"false" json:"keyTabSecretId"`

	// Version of the keytab secret in the Vault to use.
	CurrentKeyTabSecretVersion *int `mandatory:"false" json:"currentKeyTabSecretVersion"`

	// Version of the keytab secert in the Vault to use as a backup.
	BackupKeyTabSecretVersion *int `mandatory:"false" json:"backupKeyTabSecretVersion"`

	// Specifies whether to enable or disable Kerberos.
	IsKerberosEnabled *bool `mandatory:"false" json:"isKerberosEnabled"`
}

func (m UpdateKerberosDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateKerberosDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
