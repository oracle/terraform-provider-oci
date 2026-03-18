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

// DistributedDatabaseCatalogWithExadbXs Globally distributed database catalog based on exadbxs.
type DistributedDatabaseCatalogWithExadbXs struct {

	// The name of catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the shardGroup for the catalog.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// Peer details for the catalog.
	PeerDetails []CatalogPeerWithExadbXs `mandatory:"false" json:"peerDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Status of EXADB_XS based catalog.
	Status DistributedDatabaseCatalogWithExadbXsStatusEnum `mandatory:"true" json:"status"`
}

// GetName returns Name
func (m DistributedDatabaseCatalogWithExadbXs) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseCatalogWithExadbXs) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseCatalogWithExadbXs) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedDatabaseCatalogWithExadbXs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseCatalogWithExadbXs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseCatalogWithExadbXsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseCatalogWithExadbXsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedDatabaseCatalogWithExadbXs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseCatalogWithExadbXs DistributedDatabaseCatalogWithExadbXs
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseCatalogWithExadbXs
	}{
		"EXADB_XS",
		(MarshalTypeDistributedDatabaseCatalogWithExadbXs)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseCatalogWithExadbXsStatusEnum Enum with underlying type: string
type DistributedDatabaseCatalogWithExadbXsStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogWithExadbXsStatusEnum
const (
	DistributedDatabaseCatalogWithExadbXsStatusFailed                DistributedDatabaseCatalogWithExadbXsStatusEnum = "FAILED"
	DistributedDatabaseCatalogWithExadbXsStatusDeleting              DistributedDatabaseCatalogWithExadbXsStatusEnum = "DELETING"
	DistributedDatabaseCatalogWithExadbXsStatusDeleted               DistributedDatabaseCatalogWithExadbXsStatusEnum = "DELETED"
	DistributedDatabaseCatalogWithExadbXsStatusUpdating              DistributedDatabaseCatalogWithExadbXsStatusEnum = "UPDATING"
	DistributedDatabaseCatalogWithExadbXsStatusCreating              DistributedDatabaseCatalogWithExadbXsStatusEnum = "CREATING"
	DistributedDatabaseCatalogWithExadbXsStatusCreated               DistributedDatabaseCatalogWithExadbXsStatusEnum = "CREATED"
	DistributedDatabaseCatalogWithExadbXsStatusReadyForConfiguration DistributedDatabaseCatalogWithExadbXsStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseCatalogWithExadbXsStatusConfigured            DistributedDatabaseCatalogWithExadbXsStatusEnum = "CONFIGURED"
	DistributedDatabaseCatalogWithExadbXsStatusNeedsAttention        DistributedDatabaseCatalogWithExadbXsStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseCatalogWithExadbXsStatusEnum = map[string]DistributedDatabaseCatalogWithExadbXsStatusEnum{
	"FAILED":                  DistributedDatabaseCatalogWithExadbXsStatusFailed,
	"DELETING":                DistributedDatabaseCatalogWithExadbXsStatusDeleting,
	"DELETED":                 DistributedDatabaseCatalogWithExadbXsStatusDeleted,
	"UPDATING":                DistributedDatabaseCatalogWithExadbXsStatusUpdating,
	"CREATING":                DistributedDatabaseCatalogWithExadbXsStatusCreating,
	"CREATED":                 DistributedDatabaseCatalogWithExadbXsStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseCatalogWithExadbXsStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseCatalogWithExadbXsStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseCatalogWithExadbXsStatusNeedsAttention,
}

var mappingDistributedDatabaseCatalogWithExadbXsStatusEnumLowerCase = map[string]DistributedDatabaseCatalogWithExadbXsStatusEnum{
	"failed":                  DistributedDatabaseCatalogWithExadbXsStatusFailed,
	"deleting":                DistributedDatabaseCatalogWithExadbXsStatusDeleting,
	"deleted":                 DistributedDatabaseCatalogWithExadbXsStatusDeleted,
	"updating":                DistributedDatabaseCatalogWithExadbXsStatusUpdating,
	"creating":                DistributedDatabaseCatalogWithExadbXsStatusCreating,
	"created":                 DistributedDatabaseCatalogWithExadbXsStatusCreated,
	"ready_for_configuration": DistributedDatabaseCatalogWithExadbXsStatusReadyForConfiguration,
	"configured":              DistributedDatabaseCatalogWithExadbXsStatusConfigured,
	"needs_attention":         DistributedDatabaseCatalogWithExadbXsStatusNeedsAttention,
}

// GetDistributedDatabaseCatalogWithExadbXsStatusEnumValues Enumerates the set of values for DistributedDatabaseCatalogWithExadbXsStatusEnum
func GetDistributedDatabaseCatalogWithExadbXsStatusEnumValues() []DistributedDatabaseCatalogWithExadbXsStatusEnum {
	values := make([]DistributedDatabaseCatalogWithExadbXsStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogWithExadbXsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogWithExadbXsStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogWithExadbXsStatusEnum
func GetDistributedDatabaseCatalogWithExadbXsStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseCatalogWithExadbXsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogWithExadbXsStatusEnum(val string) (DistributedDatabaseCatalogWithExadbXsStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogWithExadbXsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
