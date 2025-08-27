// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVaultAndEncryptionKeyDetails Update properties for a customer-managed vault and encryption key in the destination region.
// The customer-managed encryption key in this will be used to encrypt the resource or containing resources after they
// move to the destination region.
type UpdateVaultAndEncryptionKeyDetails struct {

	// The OCID of the destination region vault for the customer-managed encryption key.
	// Example: `ocid1.vault.oc1..uniqueID`
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The OCID of the customer-managed encryption key in the destination region vault.
	// Example: `ocid1.key.oc1..uniqueID`
	EncryptionKeyId *string `mandatory:"true" json:"encryptionKeyId"`
}

func (m UpdateVaultAndEncryptionKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVaultAndEncryptionKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
