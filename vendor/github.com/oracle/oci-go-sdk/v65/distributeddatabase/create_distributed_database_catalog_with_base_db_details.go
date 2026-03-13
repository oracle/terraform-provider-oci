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

// CreateDistributedDatabaseCatalogWithBaseDbDetails Globally distributed database catalog with base database.
type CreateDistributedDatabaseCatalogWithBaseDbDetails struct {

	// The number of CPU cores to enable.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The admin password for the shard associated with Globally distributed database.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The data storage size to be allocated in GBs.
	DataStorageSizeInGbs *int `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The name of the availability domain that the base database system will be located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The identifier of the subnet for the Dbsystem instance.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the DBSystem instance. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The public key portion of the key pair to use for SSH access to the DB system.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// Fault Domain in which this base database is provisioned.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The details of the peer Dbsystem.
	PeerDetails []CreateCatalogPeerWithBaseDbDetails `mandatory:"false" json:"peerDetails"`

	DataCollectionOptions *BaseDbDistributedDatabaseDataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	// The Oracle Database edition to use for creating a Base database based catalog.
	DbEdition BaseDbEditionEnum `mandatory:"true" json:"dbEdition"`

	// The type of redundancy configured for the Base database based catalog.
	DiskRedundancy BaseDbDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The Oracle license model which applies to the Base database based catalog.
	LicenseModel BaseDbLicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// The storage option used for Base database based catalog.
	StorageManagement BaseDbStorageManagementEnum `mandatory:"true" json:"storageManagement"`

	// Block Volume Performance mode for Base database based catalog.
	StorageVolumePerformanceMode BaseDbStorageVolumePerformanceModeEnum `mandatory:"true" json:"storageVolumePerformanceMode"`
}

func (m CreateDistributedDatabaseCatalogWithBaseDbDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedDatabaseCatalogWithBaseDbDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBaseDbEditionEnum(string(m.DbEdition)); !ok && m.DbEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbEdition: %s. Supported values are: %s.", m.DbEdition, strings.Join(GetBaseDbEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbDiskRedundancyEnum(string(m.DiskRedundancy)); !ok && m.DiskRedundancy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiskRedundancy: %s. Supported values are: %s.", m.DiskRedundancy, strings.Join(GetBaseDbDiskRedundancyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetBaseDbLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbStorageManagementEnum(string(m.StorageManagement)); !ok && m.StorageManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageManagement: %s. Supported values are: %s.", m.StorageManagement, strings.Join(GetBaseDbStorageManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbStorageVolumePerformanceModeEnum(string(m.StorageVolumePerformanceMode)); !ok && m.StorageVolumePerformanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageVolumePerformanceMode: %s. Supported values are: %s.", m.StorageVolumePerformanceMode, strings.Join(GetBaseDbStorageVolumePerformanceModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedDatabaseCatalogWithBaseDbDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedDatabaseCatalogWithBaseDbDetails CreateDistributedDatabaseCatalogWithBaseDbDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedDatabaseCatalogWithBaseDbDetails
	}{
		"BASE_DB",
		(MarshalTypeCreateDistributedDatabaseCatalogWithBaseDbDetails)(m),
	}

	return json.Marshal(&s)
}
