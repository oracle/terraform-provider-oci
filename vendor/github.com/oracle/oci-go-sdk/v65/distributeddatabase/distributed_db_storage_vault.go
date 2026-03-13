// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedDbStorageVault The Storage Vault for Distributed Database Resource
type DistributedDbStorageVault struct {

	// Total storage capacity in GB for vault storage.
	HighCapacityDatabaseStorage *int `mandatory:"true" json:"highCapacityDatabaseStorage"`

	// The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Vault Storage.
	DbStorageVaultId *string `mandatory:"false" json:"dbStorageVaultId"`

	// The size of additional Flash Cache in percentage of High Capacity database storage.
	AdditionalFlashCacheInPercent *int `mandatory:"false" json:"additionalFlashCacheInPercent"`
}

func (m DistributedDbStorageVault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDbStorageVault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
