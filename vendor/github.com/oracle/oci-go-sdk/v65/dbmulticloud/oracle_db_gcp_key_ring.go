// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// <b>Microsoft Azure</b>:<br>
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
// <b>Google Cloud</b>:<br>
// 1. Oracle Google Cloud Connector Resource: This is for installing Google Identity in ExaCS VM Cluster.<br>
// 2. Discover Google Key-Rings and Keys Resource: This is for to discover Google Key-Rings.<br>
// 3. Google Key-Rings Resource: This is for to maintain Google Key-Rings in Oracle Cloud.<br>
// 4. Google Key Resource: This is for to maintain Google Key in Oracle Cloud for a Google Key-Ring.<br>
//

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OracleDbGcpKeyRing Oracle DB GCP Key-Ring Resource Object.
type OracleDbGcpKeyRing struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB GCP Key Ring Resource.
	Id *string `mandatory:"true" json:"id"`

	// The Compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that has this Oracle DB GCP Key-Ring Resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB Connector Resource.
	OracleDbConnectorId *string `mandatory:"false" json:"oracleDbConnectorId"`

	// Display name of DB GCP Key Ring Resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// GCP Key-Ring Id
	GcpKeyRingId *string `mandatory:"false" json:"gcpKeyRingId"`

	// Key-Ring Resource Type.
	Type *string `mandatory:"false" json:"type"`

	// GCP Key -Ring Resource Location.
	Location *string `mandatory:"false" json:"location"`

	// Resource's properties.
	Properties map[string]string `mandatory:"false" json:"properties"`

	// The lifecycle state of the Oracle DB GCP Key-Ring Resource.
	LifecycleState OracleDbGcpKeyRingLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the DB GCP Key-Ring resource was created in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the DB GCP Key-Ring resource was last modified, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-23T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OracleDbGcpKeyRing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleDbGcpKeyRing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleDbGcpKeyRingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOracleDbGcpKeyRingLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDbGcpKeyRingLifecycleStateEnum Enum with underlying type: string
type OracleDbGcpKeyRingLifecycleStateEnum string

// Set of constants representing the allowable values for OracleDbGcpKeyRingLifecycleStateEnum
const (
	OracleDbGcpKeyRingLifecycleStateCreating OracleDbGcpKeyRingLifecycleStateEnum = "CREATING"
	OracleDbGcpKeyRingLifecycleStateActive   OracleDbGcpKeyRingLifecycleStateEnum = "ACTIVE"
	OracleDbGcpKeyRingLifecycleStateUpdating OracleDbGcpKeyRingLifecycleStateEnum = "UPDATING"
	OracleDbGcpKeyRingLifecycleStateDeleting OracleDbGcpKeyRingLifecycleStateEnum = "DELETING"
	OracleDbGcpKeyRingLifecycleStateDeleted  OracleDbGcpKeyRingLifecycleStateEnum = "DELETED"
	OracleDbGcpKeyRingLifecycleStateFailed   OracleDbGcpKeyRingLifecycleStateEnum = "FAILED"
)

var mappingOracleDbGcpKeyRingLifecycleStateEnum = map[string]OracleDbGcpKeyRingLifecycleStateEnum{
	"CREATING": OracleDbGcpKeyRingLifecycleStateCreating,
	"ACTIVE":   OracleDbGcpKeyRingLifecycleStateActive,
	"UPDATING": OracleDbGcpKeyRingLifecycleStateUpdating,
	"DELETING": OracleDbGcpKeyRingLifecycleStateDeleting,
	"DELETED":  OracleDbGcpKeyRingLifecycleStateDeleted,
	"FAILED":   OracleDbGcpKeyRingLifecycleStateFailed,
}

var mappingOracleDbGcpKeyRingLifecycleStateEnumLowerCase = map[string]OracleDbGcpKeyRingLifecycleStateEnum{
	"creating": OracleDbGcpKeyRingLifecycleStateCreating,
	"active":   OracleDbGcpKeyRingLifecycleStateActive,
	"updating": OracleDbGcpKeyRingLifecycleStateUpdating,
	"deleting": OracleDbGcpKeyRingLifecycleStateDeleting,
	"deleted":  OracleDbGcpKeyRingLifecycleStateDeleted,
	"failed":   OracleDbGcpKeyRingLifecycleStateFailed,
}

// GetOracleDbGcpKeyRingLifecycleStateEnumValues Enumerates the set of values for OracleDbGcpKeyRingLifecycleStateEnum
func GetOracleDbGcpKeyRingLifecycleStateEnumValues() []OracleDbGcpKeyRingLifecycleStateEnum {
	values := make([]OracleDbGcpKeyRingLifecycleStateEnum, 0)
	for _, v := range mappingOracleDbGcpKeyRingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDbGcpKeyRingLifecycleStateEnumStringValues Enumerates the set of values in String for OracleDbGcpKeyRingLifecycleStateEnum
func GetOracleDbGcpKeyRingLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOracleDbGcpKeyRingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDbGcpKeyRingLifecycleStateEnum(val string) (OracleDbGcpKeyRingLifecycleStateEnum, bool) {
	enum, ok := mappingOracleDbGcpKeyRingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
