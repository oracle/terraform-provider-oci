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

// DistributedDatabaseShardDatabaseWithExadbXs Details of a distributed database shard running on ExaDB-XS.
type DistributedDatabaseShardDatabaseWithExadbXs struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	// This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// Status of distributed database shard.
	Status DistributedDatabaseShardDatabaseWithExadbXsStatusEnum `mandatory:"true" json:"status"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseShardDatabaseWithExadbXs) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseShardDatabaseWithExadbXs) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseShardDatabaseWithExadbXs) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedDatabaseShardDatabaseWithExadbXs) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

func (m DistributedDatabaseShardDatabaseWithExadbXs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseShardDatabaseWithExadbXs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseShardDatabaseWithExadbXsStatusEnumStringValues(), ",")))
	}

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
func (m DistributedDatabaseShardDatabaseWithExadbXs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseShardDatabaseWithExadbXs DistributedDatabaseShardDatabaseWithExadbXs
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseShardDatabaseWithExadbXs
	}{
		"XS_EXISTING_CLUSTER",
		(MarshalTypeDistributedDatabaseShardDatabaseWithExadbXs)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseShardDatabaseWithExadbXsStatusEnum Enum with underlying type: string
type DistributedDatabaseShardDatabaseWithExadbXsStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardDatabaseWithExadbXsStatusEnum
const (
	DistributedDatabaseShardDatabaseWithExadbXsStatusFailed                DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "FAILED"
	DistributedDatabaseShardDatabaseWithExadbXsStatusDeleting              DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "DELETING"
	DistributedDatabaseShardDatabaseWithExadbXsStatusDeleted               DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "DELETED"
	DistributedDatabaseShardDatabaseWithExadbXsStatusUpdating              DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "UPDATING"
	DistributedDatabaseShardDatabaseWithExadbXsStatusCreating              DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "CREATING"
	DistributedDatabaseShardDatabaseWithExadbXsStatusCreated               DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "CREATED"
	DistributedDatabaseShardDatabaseWithExadbXsStatusReadyForConfiguration DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseShardDatabaseWithExadbXsStatusConfigured            DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "CONFIGURED"
	DistributedDatabaseShardDatabaseWithExadbXsStatusNeedsAttention        DistributedDatabaseShardDatabaseWithExadbXsStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnum = map[string]DistributedDatabaseShardDatabaseWithExadbXsStatusEnum{
	"FAILED":                  DistributedDatabaseShardDatabaseWithExadbXsStatusFailed,
	"DELETING":                DistributedDatabaseShardDatabaseWithExadbXsStatusDeleting,
	"DELETED":                 DistributedDatabaseShardDatabaseWithExadbXsStatusDeleted,
	"UPDATING":                DistributedDatabaseShardDatabaseWithExadbXsStatusUpdating,
	"CREATING":                DistributedDatabaseShardDatabaseWithExadbXsStatusCreating,
	"CREATED":                 DistributedDatabaseShardDatabaseWithExadbXsStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseShardDatabaseWithExadbXsStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseShardDatabaseWithExadbXsStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseShardDatabaseWithExadbXsStatusNeedsAttention,
}

var mappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnumLowerCase = map[string]DistributedDatabaseShardDatabaseWithExadbXsStatusEnum{
	"failed":                  DistributedDatabaseShardDatabaseWithExadbXsStatusFailed,
	"deleting":                DistributedDatabaseShardDatabaseWithExadbXsStatusDeleting,
	"deleted":                 DistributedDatabaseShardDatabaseWithExadbXsStatusDeleted,
	"updating":                DistributedDatabaseShardDatabaseWithExadbXsStatusUpdating,
	"creating":                DistributedDatabaseShardDatabaseWithExadbXsStatusCreating,
	"created":                 DistributedDatabaseShardDatabaseWithExadbXsStatusCreated,
	"ready_for_configuration": DistributedDatabaseShardDatabaseWithExadbXsStatusReadyForConfiguration,
	"configured":              DistributedDatabaseShardDatabaseWithExadbXsStatusConfigured,
	"needs_attention":         DistributedDatabaseShardDatabaseWithExadbXsStatusNeedsAttention,
}

// GetDistributedDatabaseShardDatabaseWithExadbXsStatusEnumValues Enumerates the set of values for DistributedDatabaseShardDatabaseWithExadbXsStatusEnum
func GetDistributedDatabaseShardDatabaseWithExadbXsStatusEnumValues() []DistributedDatabaseShardDatabaseWithExadbXsStatusEnum {
	values := make([]DistributedDatabaseShardDatabaseWithExadbXsStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardDatabaseWithExadbXsStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardDatabaseWithExadbXsStatusEnum
func GetDistributedDatabaseShardDatabaseWithExadbXsStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"CREATED",
		"READY_FOR_CONFIGURATION",
		"CONFIGURED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnum(val string) (DistributedDatabaseShardDatabaseWithExadbXsStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardDatabaseWithExadbXsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
