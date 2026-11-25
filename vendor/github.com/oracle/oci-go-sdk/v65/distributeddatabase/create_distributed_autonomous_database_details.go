// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDistributedAutonomousDatabaseDetails Details required for creation of the Globally distributed autonomous database.
type CreateDistributedAutonomousDatabaseDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Globally distributed autonomous database compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Globally distributed autonomous database.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Oracle Database version for the shards and catalog used in Globally distributed autonomous database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Unique name prefix for the Globally distributed autonomous databases. Only alpha-numeric values are allowed. First character
	// has to be a letter followed by any combination of letter and number.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint associated with Globally distributed autonomous database.
	PrivateEndpointIds []string `mandatory:"true" json:"privateEndpointIds"`

	// Possible workload types. Currently only OLTP workload type is supported.
	DbWorkloadType CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum `mandatory:"true" json:"dbWorkloadType"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The listener port number for the Globally distributed autonomous database. The listener port number
	// has to be unique for a customer tenancy across all distributed autonomous databases. Same port number
	// should not be re-used for any other distributed autonomous database.
	ListenerPort *int `mandatory:"true" json:"listenerPort"`

	// Ons local port number for Globally distributed autonomous database. The onsPortLocal has to be unique for
	// a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for
	// any other distributed autonomous database.
	OnsPortLocal *int `mandatory:"true" json:"onsPortLocal"`

	// Ons remote port number for Globally distributed autonomous database. The onsPortRemote has to be unique for
	// a customer tenancy across all distributed autonomous databases. Same port number should not be re-used for
	// any other distributed autonomous database.
	OnsPortRemote *int `mandatory:"true" json:"onsPortRemote"`

	// Collection of catalog for the Globally distributed autonomous database.
	CatalogDetails []CreateAutonomousCatalogDetails `mandatory:"true" json:"catalogDetails"`

	// The collection of OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the notification topics associated with the globally distributed autonomous database.
	NotificationTopicIds []string `mandatory:"false" json:"notificationTopicIds"`

	// The TLS listener port number for Globally distributed autonomous database. The TLS listener port number
	// has to be unique for a customer tenancy across all distributed autonomous databases. Same port number
	// should not be re-used for any other distributed autonomous database. The listenerPortTls is mandatory
	// for dedicated infrastructure based distributed autonomous databases.
	ListenerPortTls *int `mandatory:"false" json:"listenerPortTls"`

	// Count of chunks associated with system raft clusters or system data guard databases.
	SystemChunkCount *int `mandatory:"false" json:"systemChunkCount"`

	// Number of replication units associated with system raft clusters.
	SystemRaftReplicationUnitCount *int `mandatory:"false" json:"systemRaftReplicationUnitCount"`

	// Collection of composite raft shards.
	CompositeRaftShardSpaces []CreateAutonomousCompositeRaftShardSpaceDetails `mandatory:"false" json:"compositeRaftShardSpaces"`

	// Collection of composite data guard shard spaces.
	CompositeDataGuardShardSpaces []CreateAutonomousCompositeDataGuardShardSpaceDetails `mandatory:"false" json:"compositeDataGuardShardSpaces"`

	// Collection of system raft clusters.
	SystemRaftClusters []CreateAutonomousSystemRaftClusterDetails `mandatory:"false" json:"systemRaftClusters"`

	SystemDataGuardDatabases *CreateAutonomousSystemDataGuardDatabaseDetails `mandatory:"false" json:"systemDataGuardDatabases"`

	// Collection of user defined shard spaces.
	UserShardSpaces []CreateAutonomousUserShardSpaceDetails `mandatory:"false" json:"userShardSpaces"`

	DbBackupConfig *DistributedAutonomousDbBackupConfig `mandatory:"false" json:"dbBackupConfig"`

	// The list of network security group (NSG) details to be associated with the distributed autonomous database.
	VcnNsgIds []VcnNsgIdsDetails `mandatory:"false" json:"vcnNsgIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDistributedAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDistributedAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum(string(m.DbWorkloadType)); !ok && m.DbWorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkloadType: %s. Supported values are: %s.", m.DbWorkloadType, strings.Join(GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum Enum with underlying type: string
type CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum string

// Set of constants representing the allowable values for CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum
const (
	CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeOltp CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum = "OLTP"
	CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeDw   CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum = "DW"
)

var mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum = map[string]CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum{
	"OLTP": CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeOltp,
	"DW":   CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeDw,
}

var mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumLowerCase = map[string]CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum{
	"oltp": CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeOltp,
	"dw":   CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeDw,
}

// GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumValues Enumerates the set of values for CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumValues() []CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum {
	values := make([]CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum, 0)
	for _, v := range mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumStringValues Enumerates the set of values in String for CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum
func GetCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum(val string) (CreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnum, bool) {
	enum, ok := mappingCreateDistributedAutonomousDatabaseDetailsDbWorkloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
