// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateConnectionDetails Details to create a Database Connection resource.
type CreateConnectionDetails struct {

	// OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Database connection type.
	DatabaseType DatabaseConnectionTypesEnum `mandatory:"true" json:"databaseType"`

	AdminCredentials *CreateAdminCredentials `mandatory:"true" json:"adminCredentials"`

	VaultDetails *CreateVaultDetails `mandatory:"true" json:"vaultDetails"`

	// Database Connection display name identifier.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Database manual connection subtype. This value can only be specified for manual connections.
	ManualDatabaseSubType DatabaseManualConnectionSubTypesEnum `mandatory:"false" json:"manualDatabaseSubType,omitempty"`

	// The OCID of the cloud database. Required if the database connection type is Autonomous.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	ConnectDescriptor *CreateConnectDescriptor `mandatory:"false" json:"connectDescriptor"`

	// This name is the distinguished name used while creating the certificate on target database. Requires a TLS wallet to be specified.
	// Not required for source container database connections.
	CertificateTdn *string `mandatory:"false" json:"certificateTdn"`

	// cwallet.sso containing containing the TCPS/SSL certificate; base64 encoded String. Not required for source container database connections.
	TlsWallet *string `mandatory:"false" json:"tlsWallet"`

	// keystore.jks file contents; base64 encoded String. Requires a TLS wallet to be specified. Not required for source container database connections.
	TlsKeystore *string `mandatory:"false" json:"tlsKeystore"`

	SshDetails *CreateSshDetails `mandatory:"false" json:"sshDetails"`

	ReplicationCredentials *CreateAdminCredentials `mandatory:"false" json:"replicationCredentials"`

	PrivateEndpoint *CreatePrivateEndpoint `mandatory:"false" json:"privateEndpoint"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m CreateConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseConnectionTypesEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseConnectionTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseManualConnectionSubTypesEnum(string(m.ManualDatabaseSubType)); !ok && m.ManualDatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManualDatabaseSubType: %s. Supported values are: %s.", m.ManualDatabaseSubType, strings.Join(GetDatabaseManualConnectionSubTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
