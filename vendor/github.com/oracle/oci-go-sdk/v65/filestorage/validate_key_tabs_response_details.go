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

// ValidateKeyTabsResponseDetails Validate keytabs response details.
type ValidateKeyTabsResponseDetails struct {

	// An array of keytab entries (principal, encryptionType, keyVersionNumber).
	CurrentKerberosKeytabEntries []KerberosKeytabEntry `mandatory:"true" json:"currentKerberosKeytabEntries"`

	// An array of keytab entries (principal, encryptionType, keyVersionNumber).
	BackupKerberosKeytabEntries []KerberosKeytabEntry `mandatory:"false" json:"backupKerberosKeytabEntries"`
}

func (m ValidateKeyTabsResponseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidateKeyTabsResponseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
