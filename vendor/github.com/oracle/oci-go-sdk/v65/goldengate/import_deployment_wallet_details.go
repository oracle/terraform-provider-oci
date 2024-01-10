// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImportDeploymentWalletDetails Metadata required to import wallet to deployment
type ImportDeploymentWalletDetails struct {

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The OCID of the customer's GoldenGate Service Secret.
	// If provided, it references a key that customers will be required to ensure the policies are established
	// to permit GoldenGate to use this Secret.
	NewWalletSecretId *string `mandatory:"true" json:"newWalletSecretId"`

	// Name of the secret with which secret is shown in vault
	WalletBackupSecretName *string `mandatory:"false" json:"walletBackupSecretName"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	MasterEncryptionKeyId *string `mandatory:"false" json:"masterEncryptionKeyId"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`
}

func (m ImportDeploymentWalletDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImportDeploymentWalletDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
