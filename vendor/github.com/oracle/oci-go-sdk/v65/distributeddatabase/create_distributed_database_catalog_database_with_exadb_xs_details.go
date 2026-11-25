// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails Configuration for creating a distributed database catalog using an existing ExaDB-XS VM cluster.
type CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The password to open the TDE wallet.
	// The password must be 9 to 255 characters and contain at least two uppercase letters, two lowercase letters, two numeric characters, and two special characters.
	// The allowed special characters are _, \#, and -.
	TdeWalletPassword *string `mandatory:"false" json:"tdeWalletPassword" sensitive:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OCI vault secret. This cannot be used in conjunction with tdeWalletPassword.
	TdeWalletPasswordSecretId *string `mandatory:"false" json:"tdeWalletPasswordSecretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	TdeWalletPasswordSecretVersionNumber *int `mandatory:"false" json:"tdeWalletPasswordSecretVersionNumber"`

	// The admin password for the catalog associated with the distributed database.
	AdminPassword *string `mandatory:"false" json:"adminPassword" sensitive:"true"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
	AdminPasswordSecretId *string `mandatory:"false" json:"adminPasswordSecretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	AdminPasswordSecretVersionNumber *int `mandatory:"false" json:"adminPasswordSecretVersionNumber"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	// This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetTdeWalletPassword returns TdeWalletPassword
func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) GetTdeWalletPassword() *string {
	return m.TdeWalletPassword
}

// GetTdeWalletPasswordSecretId returns TdeWalletPasswordSecretId
func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) GetTdeWalletPasswordSecretId() *string {
	return m.TdeWalletPasswordSecretId
}

// GetTdeWalletPasswordSecretVersionNumber returns TdeWalletPasswordSecretVersionNumber
func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) GetTdeWalletPasswordSecretVersionNumber() *int {
	return m.TdeWalletPasswordSecretVersionNumber
}

func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedDbProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDistributedDbTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails CreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails
	}{
		"XS_EXISTING_CLUSTER",
		(MarshalTypeCreateDistributedDatabaseCatalogDatabaseWithExadbXsDetails)(m),
	}

	return json.Marshal(&s)
}
