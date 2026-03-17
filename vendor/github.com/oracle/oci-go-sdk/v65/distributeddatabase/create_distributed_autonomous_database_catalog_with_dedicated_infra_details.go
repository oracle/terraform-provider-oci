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

// CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails Globally distributed autonomous database catalog based on Dedicated infrastructure.
type CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails struct {

	// Admin password for catalog database.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The compute count for the catalog database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the catalog database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Determines the auto-scaling mode for the catalog database.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	// This field is deprecated. This should not be used while creation of new distributed autonomous database. To set the peers
	// on catalog of distributed autonomous database please use peerDetails.
	PeerCloudAutonomousVmClusterIds []string `mandatory:"false" json:"peerCloudAutonomousVmClusterIds"`

	// The details required for creation of the peer for the autonomous dedicated infrastructure based catalog.
	PeerDetails []CreateCatalogPeerWithDedicatedInfraDetails `mandatory:"false" json:"peerDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
	OkvKeyStoreId *string `mandatory:"false" json:"okvKeyStoreId"`

	// The OKV endpoint name.
	OkvEndPointGroup *string `mandatory:"false" json:"okvEndPointGroup"`
}

func (m CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails CreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails
	}{
		"ADB_D",
		(MarshalTypeCreateDistributedAutonomousDatabaseCatalogWithDedicatedInfraDetails)(m),
	}

	return json.Marshal(&s)
}
