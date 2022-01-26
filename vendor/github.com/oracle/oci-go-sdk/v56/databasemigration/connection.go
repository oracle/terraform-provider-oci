// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Connection Database Connection resource used for migrations.
type Connection struct {

	// The OCID of the resource
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Database connection type.
	DatabaseType DatabaseConnectionTypesEnum `mandatory:"true" json:"databaseType"`

	// Database Connection display name identifier.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Connection resource.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The time the Connection resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the cloud database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	ConnectDescriptor *ConnectDescriptor `mandatory:"false" json:"connectDescriptor"`

	// OCID of the Secret in the OCI vault containing the Database Connection credentials.
	CredentialsSecretId *string `mandatory:"false" json:"credentialsSecretId"`

	// This name is the distinguished name used while creating the certificate on target database.
	CertificateTdn *string `mandatory:"false" json:"certificateTdn"`

	SshDetails *SshDetails `mandatory:"false" json:"sshDetails"`

	AdminCredentials *AdminCredentials `mandatory:"false" json:"adminCredentials"`

	PrivateEndpoint *PrivateEndpointDetails `mandatory:"false" json:"privateEndpoint"`

	VaultDetails *VaultDetails `mandatory:"false" json:"vaultDetails"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information
	// for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time of the last Connection resource details update. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Connection) String() string {
	return common.PointerString(m)
}
