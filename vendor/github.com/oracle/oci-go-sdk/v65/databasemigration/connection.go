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

	// Database manual connection subtype. This value can only be specified for manual connections.
	ManualDatabaseSubType DatabaseManualConnectionSubTypesEnum `mandatory:"false" json:"manualDatabaseSubType,omitempty"`

	// True if the Autonomous Connection is dedicated. Not provided for Non-Autonomous Connections.
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The OCID of the cloud database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	ConnectDescriptor *ConnectDescriptor `mandatory:"false" json:"connectDescriptor"`

	// OCID of the Secret in the OCI vault containing the Database Connection credentials.
	CredentialsSecretId *string `mandatory:"false" json:"credentialsSecretId"`

	// This name is the distinguished name used while creating the certificate on target database.
	CertificateTdn *string `mandatory:"false" json:"certificateTdn"`

	SshDetails *SshDetails `mandatory:"false" json:"sshDetails"`

	AdminCredentials *AdminCredentials `mandatory:"false" json:"adminCredentials"`

	ReplicationCredentials *AdminCredentials `mandatory:"false" json:"replicationCredentials"`

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

	// An array of Network Security Group OCIDs used to define network access for Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m Connection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Connection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseConnectionTypesEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseConnectionTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseManualConnectionSubTypesEnum(string(m.ManualDatabaseSubType)); !ok && m.ManualDatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManualDatabaseSubType: %s. Supported values are: %s.", m.ManualDatabaseSubType, strings.Join(GetDatabaseManualConnectionSubTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
