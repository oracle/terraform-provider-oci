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

// DistributedDatabaseCatalogDatabaseWithDbHome Details of a distributed database catalog using a pre-existing Database Home.
type DistributedDatabaseCatalogDatabaseWithDbHome struct {

	// The name of the catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"true" json:"dbHomeId"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

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

	// Status of the distributed database catalog.
	Status DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum `mandatory:"true" json:"status"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseCatalogDatabaseWithDbHome) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseCatalogDatabaseWithDbHome) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseCatalogDatabaseWithDbHome) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedDatabaseCatalogDatabaseWithDbHome) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

func (m DistributedDatabaseCatalogDatabaseWithDbHome) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseCatalogDatabaseWithDbHome) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumStringValues(), ",")))
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
func (m DistributedDatabaseCatalogDatabaseWithDbHome) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseCatalogDatabaseWithDbHome DistributedDatabaseCatalogDatabaseWithDbHome
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseCatalogDatabaseWithDbHome
	}{
		"EXISTING_DB_HOME",
		(MarshalTypeDistributedDatabaseCatalogDatabaseWithDbHome)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum Enum with underlying type: string
type DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum
const (
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusFailed                DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "FAILED"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleting              DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "DELETING"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleted               DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "DELETED"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusUpdating              DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "UPDATING"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreating              DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "CREATING"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreated               DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "CREATED"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusReadyForConfiguration DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusConfigured            DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "CONFIGURED"
	DistributedDatabaseCatalogDatabaseWithDbHomeStatusNeedsAttention        DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum = map[string]DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum{
	"FAILED":                  DistributedDatabaseCatalogDatabaseWithDbHomeStatusFailed,
	"DELETING":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleting,
	"DELETED":                 DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleted,
	"UPDATING":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusUpdating,
	"CREATING":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreating,
	"CREATED":                 DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseCatalogDatabaseWithDbHomeStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseCatalogDatabaseWithDbHomeStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseCatalogDatabaseWithDbHomeStatusNeedsAttention,
}

var mappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumLowerCase = map[string]DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum{
	"failed":                  DistributedDatabaseCatalogDatabaseWithDbHomeStatusFailed,
	"deleting":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleting,
	"deleted":                 DistributedDatabaseCatalogDatabaseWithDbHomeStatusDeleted,
	"updating":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusUpdating,
	"creating":                DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreating,
	"created":                 DistributedDatabaseCatalogDatabaseWithDbHomeStatusCreated,
	"ready_for_configuration": DistributedDatabaseCatalogDatabaseWithDbHomeStatusReadyForConfiguration,
	"configured":              DistributedDatabaseCatalogDatabaseWithDbHomeStatusConfigured,
	"needs_attention":         DistributedDatabaseCatalogDatabaseWithDbHomeStatusNeedsAttention,
}

// GetDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumValues Enumerates the set of values for DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumValues() []DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum {
	values := make([]DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum(val string) (DistributedDatabaseCatalogDatabaseWithDbHomeStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogDatabaseWithDbHomeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
