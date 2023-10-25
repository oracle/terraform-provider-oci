// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAmazonRedshiftConnectionDetails The information to update a the Amazon Redshift Connection.
type UpdateAmazonRedshiftConnectionDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	KeyId *string `mandatory:"false" json:"keyId"`

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Connection URL.
	// e.g.: 'jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb'
	ConnectionUrl *string `mandatory:"false" json:"connectionUrl"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`
}

// GetDisplayName returns DisplayName
func (m UpdateAmazonRedshiftConnectionDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m UpdateAmazonRedshiftConnectionDetails) GetDescription() *string {
	return m.Description
}

// GetFreeformTags returns FreeformTags
func (m UpdateAmazonRedshiftConnectionDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateAmazonRedshiftConnectionDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetVaultId returns VaultId
func (m UpdateAmazonRedshiftConnectionDetails) GetVaultId() *string {
	return m.VaultId
}

// GetKeyId returns KeyId
func (m UpdateAmazonRedshiftConnectionDetails) GetKeyId() *string {
	return m.KeyId
}

// GetNsgIds returns NsgIds
func (m UpdateAmazonRedshiftConnectionDetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m UpdateAmazonRedshiftConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAmazonRedshiftConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAmazonRedshiftConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAmazonRedshiftConnectionDetails UpdateAmazonRedshiftConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeUpdateAmazonRedshiftConnectionDetails
	}{
		"AMAZON_REDSHIFT",
		(MarshalTypeUpdateAmazonRedshiftConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
