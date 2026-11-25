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

// CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails Configuration for creating a distributed autonomous database shard using an existing ADB-D VM cluster.
type CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails struct {

	// The compute count for the shard database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the shard database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	// Admin password for shard database.
	AdminPassword *string `mandatory:"false" json:"adminPassword" sensitive:"true"`

	// The OCI vault secret [/Content/General/Concepts/identifiers.htm]OCID. This cannot be used in conjunction with adminPassword.
	AdminPasswordSecretId *string `mandatory:"false" json:"adminPasswordSecretId"`

	// The version of the vault secret. If no version is specified, the latest version will be used.
	AdminPasswordSecretVersionNumber *int `mandatory:"false" json:"adminPasswordSecretVersionNumber"`

	// The shard space name for the shard database. Shard space for an existing shard cannot be changed once the shard is created.
	// The shard space name is used when creating new shards. For user-defined sharding, every shard must have a unique
	// shard space name. For system-defined sharding, a shard space name is not required.
	ShardSpaceName *string `mandatory:"false" json:"shardSpaceName"`

	// Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count.
	// The default value is `FALSE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
	OkvKeyStoreId *string `mandatory:"false" json:"okvKeyStoreId"`

	// The OKV endpoint name.
	OkvEndPointGroupName *string `mandatory:"false" json:"okvEndPointGroupName"`

	// The lag time preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// The protection mode for the shard peer.
	ProtectionMode DistributedAutonomousDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`
}

func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedAutonomousDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedAutonomousDbProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails CreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails
	}{
		"ADBD_EXISTING_CLUSTER",
		(MarshalTypeCreateDistributedAutonomousDatabaseShardWithDedicatedInfraDetails)(m),
	}

	return json.Marshal(&s)
}
