// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalKeyManagerMetadataSummary Summary about metadata of external key manager to be returned to the customer as a response.
type ExternalKeyManagerMetadataSummary struct {

	// URL of the vault on external key manager.
	ExternalVaultEndpointUrl *string `mandatory:"true" json:"externalVaultEndpointUrl"`

	// OCID of the private endpoint.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// Vendor of the external key manager.
	Vendor *string `mandatory:"false" json:"vendor"`

	OauthMetadataSummary *OauthMetadataSummary `mandatory:"false" json:"oauthMetadataSummary"`
}

func (m ExternalKeyManagerMetadataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalKeyManagerMetadataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
