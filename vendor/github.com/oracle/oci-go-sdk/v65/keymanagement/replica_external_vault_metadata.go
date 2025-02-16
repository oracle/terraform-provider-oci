// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicaExternalVaultMetadata Metadata of the replica region External Vault
type ReplicaExternalVaultMetadata struct {

	// OCID of the EKMS private endpoint in the replica region and must be in ACTIVE state
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// Replica region URL of the IDCS domain
	IdcsAccountNameUrl *string `mandatory:"true" json:"idcsAccountNameUrl"`
}

func (m ReplicaExternalVaultMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicaExternalVaultMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ReplicaExternalVaultMetadata) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeReplicaExternalVaultMetadata ReplicaExternalVaultMetadata
	s := struct {
		DiscriminatorParam string `json:"vaultType"`
		MarshalTypeReplicaExternalVaultMetadata
	}{
		"EXTERNAL",
		(MarshalTypeReplicaExternalVaultMetadata)(m),
	}

	return json.Marshal(&s)
}
