// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateConnectionDetails Details to update in a Database Connection resource.
type UpdateConnectionDetails struct {

	// Database Connection display name identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the cloud database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	ConnectDescriptor *UpdateConnectDescriptor `mandatory:"false" json:"connectDescriptor"`

	// This name is the distinguished name used while creating the certificate on target database. Not required for source container database connections.
	CertificateTdn *string `mandatory:"false" json:"certificateTdn"`

	// cwallet.sso containing containing the TCPS/SSL certificate; base64 encoded String. Not required for source container database connections.
	TlsWallet *string `mandatory:"false" json:"tlsWallet"`

	// keystore.jks file contents; base64 encoded String. Not required for source container database connections.
	TlsKeystore *string `mandatory:"false" json:"tlsKeystore"`

	SshDetails *UpdateSshDetails `mandatory:"false" json:"sshDetails"`

	AdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"adminCredentials"`

	PrivateEndpoint *UpdatePrivateEndpoint `mandatory:"false" json:"privateEndpoint"`

	VaultDetails *UpdateVaultDetails `mandatory:"false" json:"vaultDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
