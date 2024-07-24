// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedCatalogDetails Details of ATP-D based catalog.
type DedicatedCatalogDetails struct {

	// Catalog name
	Name *string `mandatory:"true" json:"name"`

	// The compute amount available to the underlying autonomous database associated with shard or catalog.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Name of the shard-group to which the catalog belongs.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The time the catalog was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last created. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Status of shard or catalog or gsm for the sharded database.
	Status DedicatedCatalogDetailsStatusEnum `mandatory:"true" json:"status"`

	// Determines the auto-scaling mode.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// Identifier of the primary cloudAutonomousVmCluster for the catalog.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	EncryptionKeyDetails *DedicatedShardOrCatalogEncryptionKeyDetails `mandatory:"false" json:"encryptionKeyDetails"`

	// The time the ssl certificate associated with catalog expires. An RFC3339 formatted datetime string
	TimeSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeSslCertificateExpires"`

	// Identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// Identifier of the underlying container database.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// Identifier of the underlying container database parent.
	ContainerDatabaseParentId *string `mandatory:"false" json:"containerDatabaseParentId"`

	// Identifier of the peer cloudAutonomousVmCluster for the catalog.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// Additional metadata related to catalog's underlying supporting resource.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`
}

func (m DedicatedCatalogDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedCatalogDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedCatalogDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDedicatedCatalogDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DedicatedCatalogDetailsStatusEnum Enum with underlying type: string
type DedicatedCatalogDetailsStatusEnum string

// Set of constants representing the allowable values for DedicatedCatalogDetailsStatusEnum
const (
	DedicatedCatalogDetailsStatusFailed                DedicatedCatalogDetailsStatusEnum = "FAILED"
	DedicatedCatalogDetailsStatusDeleting              DedicatedCatalogDetailsStatusEnum = "DELETING"
	DedicatedCatalogDetailsStatusDeleted               DedicatedCatalogDetailsStatusEnum = "DELETED"
	DedicatedCatalogDetailsStatusUpdating              DedicatedCatalogDetailsStatusEnum = "UPDATING"
	DedicatedCatalogDetailsStatusCreating              DedicatedCatalogDetailsStatusEnum = "CREATING"
	DedicatedCatalogDetailsStatusCreated               DedicatedCatalogDetailsStatusEnum = "CREATED"
	DedicatedCatalogDetailsStatusReadyForConfiguration DedicatedCatalogDetailsStatusEnum = "READY_FOR_CONFIGURATION"
	DedicatedCatalogDetailsStatusConfigured            DedicatedCatalogDetailsStatusEnum = "CONFIGURED"
	DedicatedCatalogDetailsStatusNeedsAttention        DedicatedCatalogDetailsStatusEnum = "NEEDS_ATTENTION"
)

var mappingDedicatedCatalogDetailsStatusEnum = map[string]DedicatedCatalogDetailsStatusEnum{
	"FAILED":                  DedicatedCatalogDetailsStatusFailed,
	"DELETING":                DedicatedCatalogDetailsStatusDeleting,
	"DELETED":                 DedicatedCatalogDetailsStatusDeleted,
	"UPDATING":                DedicatedCatalogDetailsStatusUpdating,
	"CREATING":                DedicatedCatalogDetailsStatusCreating,
	"CREATED":                 DedicatedCatalogDetailsStatusCreated,
	"READY_FOR_CONFIGURATION": DedicatedCatalogDetailsStatusReadyForConfiguration,
	"CONFIGURED":              DedicatedCatalogDetailsStatusConfigured,
	"NEEDS_ATTENTION":         DedicatedCatalogDetailsStatusNeedsAttention,
}

var mappingDedicatedCatalogDetailsStatusEnumLowerCase = map[string]DedicatedCatalogDetailsStatusEnum{
	"failed":                  DedicatedCatalogDetailsStatusFailed,
	"deleting":                DedicatedCatalogDetailsStatusDeleting,
	"deleted":                 DedicatedCatalogDetailsStatusDeleted,
	"updating":                DedicatedCatalogDetailsStatusUpdating,
	"creating":                DedicatedCatalogDetailsStatusCreating,
	"created":                 DedicatedCatalogDetailsStatusCreated,
	"ready_for_configuration": DedicatedCatalogDetailsStatusReadyForConfiguration,
	"configured":              DedicatedCatalogDetailsStatusConfigured,
	"needs_attention":         DedicatedCatalogDetailsStatusNeedsAttention,
}

// GetDedicatedCatalogDetailsStatusEnumValues Enumerates the set of values for DedicatedCatalogDetailsStatusEnum
func GetDedicatedCatalogDetailsStatusEnumValues() []DedicatedCatalogDetailsStatusEnum {
	values := make([]DedicatedCatalogDetailsStatusEnum, 0)
	for _, v := range mappingDedicatedCatalogDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedCatalogDetailsStatusEnumStringValues Enumerates the set of values in String for DedicatedCatalogDetailsStatusEnum
func GetDedicatedCatalogDetailsStatusEnumStringValues() []string {
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

// GetMappingDedicatedCatalogDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedCatalogDetailsStatusEnum(val string) (DedicatedCatalogDetailsStatusEnum, bool) {
	enum, ok := mappingDedicatedCatalogDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
