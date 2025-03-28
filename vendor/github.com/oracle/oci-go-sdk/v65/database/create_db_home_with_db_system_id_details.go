// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbHomeWithDbSystemIdDetails Note that a valid `dbSystemId` value must be supplied for the `CreateDbHomeWithDbSystemId` API operation to successfully complete.
type CreateDbHomeWithDbSystemIdDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The user-provided name of the Database Home.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The database software image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId *string `mandatory:"false" json:"databaseSoftwareImageId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// If true, the customer acknowledges that the specified Oracle Database software is an older release that is not currently supported by OCI.
	IsDesupportedVersion *bool `mandatory:"false" json:"isDesupportedVersion"`

	// Indicates whether unified autiding is enabled or not. Set to True to enable unified auditing on respective DBHome.
	IsUnifiedAuditingEnabled *bool `mandatory:"false" json:"isUnifiedAuditingEnabled"`

	// A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	Database *CreateDatabaseDetails `mandatory:"false" json:"database"`
}

// GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetKmsKeyId returns KmsKeyId
func (m CreateDbHomeWithDbSystemIdDetails) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m CreateDbHomeWithDbSystemIdDetails) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m CreateDbHomeWithDbSystemIdDetails) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

// GetFreeformTags returns FreeformTags
func (m CreateDbHomeWithDbSystemIdDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateDbHomeWithDbSystemIdDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetIsDesupportedVersion returns IsDesupportedVersion
func (m CreateDbHomeWithDbSystemIdDetails) GetIsDesupportedVersion() *bool {
	return m.IsDesupportedVersion
}

// GetIsUnifiedAuditingEnabled returns IsUnifiedAuditingEnabled
func (m CreateDbHomeWithDbSystemIdDetails) GetIsUnifiedAuditingEnabled() *bool {
	return m.IsUnifiedAuditingEnabled
}

func (m CreateDbHomeWithDbSystemIdDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbHomeWithDbSystemIdDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdDetails CreateDbHomeWithDbSystemIdDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdDetails
	}{
		"NONE",
		(MarshalTypeCreateDbHomeWithDbSystemIdDetails)(m),
	}

	return json.Marshal(&s)
}
