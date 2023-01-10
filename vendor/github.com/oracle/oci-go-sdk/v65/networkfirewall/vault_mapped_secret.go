// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VaultMappedSecret Mapped secret stored in OCI vault used in the firewall policy rules.
type VaultMappedSecret struct {

	// OCID for the Vault Secret to be used.
	VaultSecretId *string `mandatory:"true" json:"vaultSecretId"`

	// Version number of the secret to be used.
	VersionNumber *int `mandatory:"true" json:"versionNumber"`

	// Type of the secrets mapped based on the policy.
	//  * `SSL_INBOUND_INSPECTION`: For Inbound inspection of SSL traffic.
	//  * `SSL_FORWARD_PROXY`: For forward proxy certificates for SSL inspection.
	Type MappedSecretTypeEnum `mandatory:"true" json:"type"`
}

//GetType returns Type
func (m VaultMappedSecret) GetType() MappedSecretTypeEnum {
	return m.Type
}

func (m VaultMappedSecret) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultMappedSecret) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMappedSecretTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMappedSecretTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VaultMappedSecret) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVaultMappedSecret VaultMappedSecret
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeVaultMappedSecret
	}{
		"OCI_VAULT",
		(MarshalTypeVaultMappedSecret)(m),
	}

	return json.Marshal(&s)
}
