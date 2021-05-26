// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// KeyStoreTypeFromOracleKeyVaultDetails Details for Oracle Key Vault
type KeyStoreTypeFromOracleKeyVaultDetails struct {

	// The list of Oracle Key Vault connection IP addresses.
	ConnectionIps []string `mandatory:"true" json:"connectionIps"`

	// The administrator username to connect to Oracle Key Vault
	AdminUsername *string `mandatory:"true" json:"adminUsername"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure secret (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	SecretId *string `mandatory:"true" json:"secretId"`
}

func (m KeyStoreTypeFromOracleKeyVaultDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m KeyStoreTypeFromOracleKeyVaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKeyStoreTypeFromOracleKeyVaultDetails KeyStoreTypeFromOracleKeyVaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKeyStoreTypeFromOracleKeyVaultDetails
	}{
		"ORACLE_KEY_VAULT",
		(MarshalTypeKeyStoreTypeFromOracleKeyVaultDetails)(m),
	}

	return json.Marshal(&s)
}
